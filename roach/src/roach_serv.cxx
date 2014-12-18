/**
*
*/

#include "uv.h"

#include "roach.hxx"
#include "roach_serv.hxx"

namespace roach {

class ServerImpl : public Server
{
private:
    boost::shared_ptr<Logger> m_logger;
    boost::shared_ptr<Context> m_ctx;

    uv_loop_t *m_uvLoop;
    uv_tcp_t  m_listener;
    volatile bool m_isRunning;

    static const int LISTEN_BACKLOG = 128;

public:
    ServerImpl(boost::shared_ptr<Context> ctx)
        : Server()
        , m_ctx(ctx)
        , m_logger(ctx->GetLogger())
        , m_uvLoop(uv_default_loop())
        , m_isRunning(false)
    {
        /* NOP */
    }

    static void OnConnect(uv_stream_t* server_handle, int status)
    {

    }

    virtual bool Run(const char *ip, const int port)
    {
        if (m_isRunning)
        {
            m_logger->Log(Logger::Err, "Server is already running");
            return false;
        }

        struct sockaddr_in addr;
        int result = uv_ip4_addr(ip, port, &addr);
        if (0 != result)
        {
            m_logger->Log(Logger::Err,
                "Invalid IP or Port. (%s %d)", ip, port);
            return false;
        }

        uv_tcp_t listener;
        result = uv_tcp_init(m_uvLoop, &listener);
        if (0 != result)
        {
            m_logger->Log(Logger::Err, "uv_tcp_init() error");
            return false;
        }

        result = uv_tcp_bind(&listener,
            static_cast<const struct sockaddr*>(static_cast<void*>(&addr)), 0);
        if (0 != result)
        {
            m_logger->Log(Logger::Err, "Cannot bind to %s %d", ip, port);
            return false;
        }

        result = uv_listen((uv_stream_t*)&listener, LISTEN_BACKLOG, OnConnect);
        if (0 != result)
        {
            uv_close(static_cast<uv_handle_t*>(static_cast<void*>(&listener)), NULL);
            m_logger->Log(Logger::Err, "Cannot listen on %s %d", ip, port);
            return false;
        }

        m_isRunning = true;
        result = uv_run(m_uvLoop, UV_RUN_DEFAULT);
        if (0 != result)
        {
            m_isRunning = false;
            uv_close(static_cast<uv_handle_t*>(static_cast<void*>(&listener)), NULL);
            m_logger->Log(Logger::Err, "Cannot start UV Loop");
            return false;
        }

        uv_close(static_cast<uv_handle_t*>(static_cast<void*>(&listener)), NULL);
        m_isRunning = false;
        m_logger->Log(Logger::Inf, "UV Loop stopped");
        return true;
    }

    virtual bool IsRunning(void) const
    {
        return m_isRunning;
    }

    virtual void Stop(void)
    {
        if (!IsRunning())
        {
            m_logger->Log(Logger::War, "Cannot stop serevr (It's not running)");
            return;
        }
        uv_stop(m_uvLoop);
        m_logger->Log(Logger::Inf, "uv_stop() called");
    }
};

Server::Server(void)
{
    /* NOP */
}

boost::shared_ptr<Server> Server::Create(boost::shared_ptr<Context> ctx)
{
    return boost::shared_ptr<Server>(new ServerImpl(ctx));
}

} /* namespace roach */
