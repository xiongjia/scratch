/**
 * Roach - A simple HTTP Server. (libuv & http-parser & boost & C++11 & CMake)
 */

#include <functional>
#include <string>

#include "boost/make_shared.hpp"
#if defined(_MSC_VER)
#   pragma warning (push)
#   pragma warning (disable : 4819)
#   include "boost/format.hpp"
#   pragma warning (push)
#   pragma warning (default : 4819)
#else
#   include "boost/format.hpp"
#endif /* defined(_MSC_VER) */

#include "uv.h"

#include "roach.hxx"
#include "roach_serv.hxx"
#include "roach_uvbufpool.hxx"
#include "roach_misc.hxx"
#include "roach_http_parser.hxx"

namespace roach {

class Connection;
class ShutdownReq;

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
        if (!uv_is_active(conn->GetHandle()))
        {
            if (afterShutdown)
            {
                afterShutdown(conn, 0);
            }
            return;
        }

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

class UVSender : boost::noncopyable
{
private:
    boost::shared_ptr<Logger> m_logger;
    uv_write_t m_wr;
    uv_buf_t m_buf;
    Connection *m_conn;

public:
    UVSender(boost::shared_ptr<Logger> logger,
             Connection *conn, const char *buf, size_t len)
        : m_logger(logger)
        , m_conn(conn)
    {
        m_wr.data = this;
        m_buf.base = (char*)malloc(len);
        if (NULL != m_buf.base)
        {
            memcpy(m_buf.base, buf, len);
            m_buf.len = len;
        }
        else
        {
            m_buf.len = 0;
        }
    }

    ~UVSender(void)
    {
        if (NULL != m_buf.base)
        {
            free(m_buf.base);
        }
    }

    uv_write_t* GetWrite(void)
    {
        return &m_wr;
    }

    uv_buf_t* GetBuf(void)
    {
        return &m_buf;
    }

    Connection* GetConn(void)
    {
        return m_conn;
    }

    boost::shared_ptr<Logger> GetLogger(void)
    {
        return m_logger;
    }
};

class HttpResponseImpl : public HttpResponse
{
private:
    boost::shared_ptr<Logger> m_logger;
    Connection *m_conn;

public:
    HttpResponseImpl(Connection *conn, boost::shared_ptr<Logger> logger)
        : HttpResponse()
        , m_logger(logger)
        , m_conn(conn)
    {
        /* NOP */
    }

    virtual void WritePlainText(const unsigned short statusCode,
        const char *status,
        const char *content)
    {
        auto repFmt = boost::format("HTTP/1.1 %d %s\r\n"
            "Content-Type: text/plain\r\n"
            "Content-Length: %d\r\n"
            "\r\n"
            "%s");
        std::string contentStr(content);
        auto rep = boost::str(repFmt %
            statusCode % status % contentStr.size() % contentStr);
        UVSender *snd = new UVSender(m_logger, m_conn, &rep[0], rep.size());
        uv_write(snd->GetWrite(), m_conn->GetStream(), snd->GetBuf(), 1, 
                [](uv_write_t *req, int status) {
            UVSender *snd = static_cast<UVSender*>(req->data);
            Connection *conn = snd->GetConn();
            boost::shared_ptr<Logger> logger = snd->GetLogger();
            logger->Log(Logger::Dbg, "write response status %d", status);
            uv_close(conn->GetHandle(), NULL);
            delete snd;
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
    volatile bool m_isRunning;

    uv_loop_t* GetUVLoop(void)
    {
        return &m_uvLoop;
    }

public:
    ServerImpl(boost::shared_ptr<Context> ctx);

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
    , m_logger(ctx->GetLogger())
    , m_ctx(ctx)
    , m_uvBufPool(UVBufPool::Create(m_ctx))
    , m_isRunning(false)
{
    uv_loop_init(&m_uvLoop);
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
    uv_tcp_nodelay(listener.GetTCP(), 1);
    uv_tcp_keepalive(listener.GetTCP(), 1, 60);
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
            /* TODO: write an error status back */
            m_logger->LogNoFmt(Logger::Err,
                "Invalid HTTP request Closing connection");
            Connection::Shutdown(conn, [](Connection *conn, int /* status */) {
                UVAsyncDelete<Connection>(conn->GetLoop(), conn);
            });
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
    boost::shared_ptr<HttpResponse> rep = boost::make_shared<HttpResponseImpl>(conn, m_logger);
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

    /* default handler */
    m_logger->Log(Logger::Dbg, "Calling default handler");
    rep->WritePlainText(404, "Not Found", "404 - Not Found");
}

} /* namespace roach */
