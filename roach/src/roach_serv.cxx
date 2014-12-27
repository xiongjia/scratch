/**
 * Roach - A simple HTTP Server. (libuv & http-parser & boost & C++11 & CMake)
 */

#include <functional>
#include "boost/make_shared.hpp"
#include "uv.h"

#include "roach.hxx"
#include "roach_serv.hxx"
#include "roach_uvbufpool.hxx"
#include "roach_misc.hxx"
#include "roach_http_parser.hxx"

namespace roach {

class Connection;
class ShutdownReq;

class HttpResponseImpl : public HttpResponse
{
private:
    Connection *m_conn;

public:
    HttpResponseImpl(Connection *conn)
        : HttpResponse()
        , m_conn(conn)
    {
        /* NOP */
    }

    virtual void Write(const char *buf, const size_t len)
    {

    }
};

class ShutdownReq : boost::noncopyable
{
public:
    typedef std::function<void(Connection*, int)> AfterShutdown;

private:
    uv_shutdown_t m_shutdown;
    AfterShutdown m_afterShutdown;
    Connection *m_conn;

public:
    ShutdownReq(Connection *conn = nullptr,
                AfterShutdown afterShtdown = nullptr)
        : m_afterShutdown(afterShtdown)
        , m_conn(conn)
    {
        m_shutdown.data = this;
    }

    void SetConn(Connection *conn)
    {
        m_conn = conn;
    }

    Connection *GetConn(void)
    {
        return m_conn;
    }

    void SetAfterShutdown(AfterShutdown afterShutdown)
    {
        m_afterShutdown = afterShutdown;
    }

    AfterShutdown GetAfterShutdown(void)
    {
        return m_afterShutdown;
    }

    uv_shutdown_t* GetShutdownReq(void)
    {
        return &m_shutdown;
    }
};

class Connection : public UVTCPWrap
{
private:
    boost::shared_ptr<HTTPParser> m_parser;

public:
    Connection(uv_loop_t *uvLoop, void *uvCtx = nullptr)
        : UVTCPWrap(uvLoop, uvCtx)
        , m_parser(HTTPParser::Create())
    {
        /* NOP */
    }

    boost::shared_ptr<HTTPParser> GetParser(void)
    {
        return m_parser;
    }

    static void Shutdown(Connection *conn, 
                         ShutdownReq::AfterShutdown afterShutdown)
    {
        ShutdownReq *shutdown = new ShutdownReq(conn, afterShutdown);
        uv_shutdown(shutdown->GetShutdownReq(), conn->GetStream(), 
            [](uv_shutdown_t *req, int status) {
                ShutdownReq *shutdown = static_cast<ShutdownReq*>(req->data);
                Connection *conn = shutdown->GetConn();
                uv_loop_t *loop = conn->GetLoop();
                ShutdownReq::AfterShutdown afterShutdown(shutdown->GetAfterShutdown());
                if (afterShutdown)
                {
                    afterShutdown(conn, status);
                }
                UVAsyncDelete<ShutdownReq>(loop, shutdown);
            });
    }
};

class ServerImpl : public Server
{
private:
    ServHandler m_srvHandler;
    boost::shared_ptr<Logger> m_logger;
    boost::shared_ptr<Context> m_ctx;
    boost::shared_ptr<UVBufPool> m_uvBufPool;

    uv_loop_t m_uvLoop;
    uv_tcp_t  m_listener;
    volatile bool m_isRunning;

    uv_loop_t* GetUVLoop(void)
    {
        return &m_uvLoop;
    }

public:
    ServerImpl(boost::shared_ptr<Context> ctx);
    ~ServerImpl(void);

    virtual bool Run(const char *ip, const int port);
    virtual void Stop(void);

    virtual bool IsRunning(void) const
    {
        return m_isRunning;
    }

    virtual void SetHandler(ServHandler handler)
    {
        m_srvHandler = handler;
    }

    void OnConnect(UVTCPWrap *listener, int status);
    void OnConnRead(Connection *conn, ssize_t nread, const uv_buf_t* buf);

    void OnAllocBuf(size_t suggestedSize, uv_buf_t *buf);
    void OnFreeBuf(const uv_buf_t *buf);

    void OnHttpRequest(Connection *conn);

    void OnServHandler(Connection *conn,
                       boost::shared_ptr<HttpRequest> req,
                       boost::shared_ptr<HttpResponse> rep);
};

Server::Server(void)
{
    /* NOP */
}

boost::shared_ptr<Server> Server::Create(boost::shared_ptr<Context> ctx)
{
    return boost::shared_ptr<Server>(new ServerImpl(ctx));
}

ServerImpl::ServerImpl(boost::shared_ptr<Context> ctx)
    : Server()
    , m_ctx(ctx)
    , m_logger(ctx->GetLogger())
    , m_uvBufPool(UVBufPool::Create(m_ctx))
    , m_isRunning(false)
{
    uv_loop_init(&m_uvLoop);
}

ServerImpl::~ServerImpl(void)
{
    uv_loop_close(&m_uvLoop);
}

void ServerImpl::Stop(void)
{
    if (!IsRunning())
    {
        m_logger->Log(Logger::War, "Cannot stop serevr (It's not running)");
        return;
    }

    UVShutdownLoop(GetUVLoop());
    m_logger->Log(Logger::Inf, "uv_stop() request has been sent");
}

bool ServerImpl::Run(const char *ip, const int port)
{
    if (m_isRunning)
    {
        m_logger->Log(Logger::Err, "Server is already running");
        return false;
    }

    auto uvAddrCreator = UVAddr::CreateIP4(ip, port);
    if (uvAddrCreator.first != 0)
    {
        m_logger->Log(Logger::Err,
            "Invalid IP or Port. (%s %d). Error: %s",
            ip, port, uv_strerror(uvAddrCreator.first));
        return false;
    }
    auto uvAddr = uvAddrCreator.second;

    UVTCPWrap listener(GetUVLoop(), this);
    int uvErr = uv_tcp_bind(listener.GetTCP(), uvAddr->GetSockAddr(), 0);
    if (0 != uvErr)
    {
        m_logger->Log(Logger::Err, "Cannot bind to %s %d: %s",
            ip, port, uv_strerror(uvErr));
        return false;
    }

    uvErr = uv_listen(listener.GetStream(), SOMAXCONN,
        [](uv_stream_t *handle, int status) {
            UVTCPWrap *listener = static_cast<UVTCPWrap*>(handle->data);
            ServerImpl *serv = static_cast<ServerImpl*>(listener->GetUVCtx());
            serv->OnConnect(listener, status);
        });
    if (0 != uvErr)
    {
        m_logger->Log(Logger::Err, "Cannot listen on %s %d: %s",
            ip, port, uv_strerror(uvErr));
        return false;
    }

    m_isRunning = true;
    uv_run(GetUVLoop(), UV_RUN_DEFAULT);
    m_isRunning = false;
    m_logger->Log(Logger::Inf, "UV Loop stopped");
    return true;
}

void ServerImpl::OnConnect(UVTCPWrap *listener, int status)
{
    if (0 != status)
    {
        m_logger->Log(Logger::Err, "Connection error: %s",
            uv_strerror(status));
        return;
    }

    Connection *conn = new Connection(GetUVLoop(), this);
    int uvErr = uv_accept(listener->GetStream(), conn->GetStream());
    if (0 != uvErr)
    {
        m_logger->Log(Logger::Err, "Connection error: %s",
            uv_strerror(uvErr));
        delete conn;
        return;
    }

    uv_read_start(conn->GetStream(),
        [](uv_handle_t *handle, size_t suggestedSize, uv_buf_t *buf) {
            Connection *conn = static_cast<Connection*>(handle->data);
            ServerImpl *serv = static_cast<ServerImpl*>(conn->GetUVCtx());
            serv->OnAllocBuf(suggestedSize, buf);
        },
        [](uv_stream_t *handle, ssize_t nread, const uv_buf_t *buf) {
            Connection *conn = static_cast<Connection*>(handle->data);
            ServerImpl *serv = static_cast<ServerImpl*>(conn->GetUVCtx());
            serv->OnConnRead(conn, nread, buf);
        });
}

void ServerImpl::OnConnRead(Connection *conn,
                            ssize_t nread,
                            const uv_buf_t *buf)
{
    if (0 > nread)
    {
        if (UV_EOF != nread)
        {
            m_logger->Log(Logger::Err, "Read error: nread = %d", nread);
        }
        else
        {
            m_logger->LogNoFmt(Logger::Dbg, "Read UV_EOF");
        }

        /* shutdown the connection */
        m_logger->LogNoFmt(Logger::Dbg, "Closing connection");
        Connection::Shutdown(conn, [](Connection *conn, int /* status */) {
            UVAsyncDelete<Connection>(conn->GetLoop(), conn);
        });        
    }
    else if (0 == nread)
    {
        m_logger->LogNoFmt(Logger::Dbg, "Read 0 bytes");
    }
    else
    {
        /* parse HTTP request */
        m_logger->Log(Logger::Dbg, "Read %d bytes", nread);
        boost::shared_ptr<HTTPParser> parser = conn->GetParser();
        bool parse = parser->Parse(buf->base, nread);
        if (!parse)
        {
            /* TODO: write a client/server error back */
        }
        else
        {
            if (parser->IsCompleted())
            {
                OnHttpRequest(conn);
            }
        }
    }

    /* free the buf */
    OnFreeBuf(buf);
}

void ServerImpl::OnAllocBuf(size_t suggestedSize, uv_buf_t *buf)
{
    m_uvBufPool->Alloc(suggestedSize, buf);
}

void ServerImpl::OnFreeBuf(const uv_buf_t *buf)
{
    m_uvBufPool->Free(buf);
}

void ServerImpl::OnHttpRequest(Connection *conn)
{
    boost::shared_ptr<HTTPParser> parser = conn->GetParser();
    boost::shared_ptr<HttpRequest> req = HttpRequest::Create();
    req->SetMethod(parser->GetMethod());
    req->SetUrl(parser->GetUrl());
    req->SetHeader(parser->GetHeader());

    boost::shared_ptr<HttpResponse> rep = boost::make_shared<HttpResponseImpl>(conn);

    OnServHandler(conn, req, rep);
}

void ServerImpl::OnServHandler(Connection *conn,
                               boost::shared_ptr<HttpRequest> req,
                               boost::shared_ptr<HttpResponse> rep)
{
    if (m_srvHandler)
    {
        m_logger->Log(Logger::Dbg, "Calling customer handler");
        if (m_srvHandler(req, rep))
        {
            return;
        }
    }

    m_logger->Log(Logger::Dbg, "Calling default handler");    
    /* TODO default handler */
}

} /* namespace roach */
