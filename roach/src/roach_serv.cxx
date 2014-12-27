/**
 * Roach - A simple HTTP Server. (libuv & http-parser & boost & C++11 & CMake)
 */

#include <functional>
#include "uv.h"
#include "http_parser.h"

#include "roach.hxx"
#include "roach_serv.hxx"
#include "roach_uvbufpool.hxx"
#include "roach_misc.hxx"

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
    http_parser_settings m_parserCfg;
    http_parser m_paser;

public:
    Connection(uv_loop_t *uvLoop, void *uvCtx = nullptr)
        : UVTCPWrap(uvLoop, uvCtx)
    {
        memset(&m_parserCfg, 0, sizeof(m_parserCfg));
        m_parserCfg.on_message_begin = [](http_parser *parse)  -> int {
                return 0; 
            };
        m_parserCfg.on_url = [](http_parser *parse, const char *at, size_t length) -> int {
                return 0;
            };
        m_parserCfg.on_header_field = [](http_parser *parse, const char *at, size_t length) -> int {
                return 0;
            };
        m_parserCfg.on_header_value = [](http_parser *parse, const char *at, size_t length) -> int {
                return 0;
            };
        m_parserCfg.on_headers_complete = [](http_parser *parse)  -> int {
                return 1;
            };
        m_parserCfg.on_body = [](http_parser *parse, const char *at, size_t length) -> int {
                return 1;
            };
        m_parserCfg.on_message_complete = [](http_parser *parse)  -> int {
                return 1;
            };
        m_paser.data = this;
        http_parser_init(&m_paser, HTTP_REQUEST);
    }

    bool Parse(char *buf, size_t len)
    {
        size_t  parsed = http_parser_execute(&m_paser, &m_parserCfg, buf, len);
        return parsed < len ? false : true;
    }

    static void Shutdown(Connection *conn, 
                         ShutdownReq::AfterShutdown afterShutdown)
    {
        ShutdownReq *shutdown = new ShutdownReq(conn, afterShutdown);
        uv_shutdown(shutdown->GetShutdownReq(), conn->GetStream(), 
            [](uv_shutdown_t *req, int status) {
                ShutdownReq *shutdown = static_cast<ShutdownReq*>(req->data);
                ShutdownReq::AfterShutdown afterShutdown(shutdown->GetAfterShutdown());
                if (afterShutdown)
                {
                    afterShutdown(shutdown->GetConn(), status);
                }
                delete shutdown;
            });
    }
};

class ServerImpl : public Server
{
private:
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

    void OnConnect(UVTCPWrap *listener, int status);
    void OnConnRead(Connection *conn, ssize_t nread, const uv_buf_t* buf);

    void OnAllocBuf(size_t suggestedSize, uv_buf_t *buf);
    void OnFreeBuf(const uv_buf_t *buf);
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
    ShutdownLoop::Shutdown(GetUVLoop());
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
            delete conn;
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
        conn->Parse(buf->base, nread);
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

} /* namespace roach */
