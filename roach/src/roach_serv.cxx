/**
 *
 */

#include "uv.h"

#include "roach.hxx"
#include "roach_serv.hxx"
#include "roach_misc.hxx"

namespace roach {

class ServerImpl;

typedef struct _UVListener
{
    uv_tcp_t listener;
    ServerImpl *srvImpl;
} UVListener;

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

    static void OnConnect(uv_stream_t *srvHandle, int status)
    {
        if (0 != status)
        {
            return;
        }
    }

    virtual bool Run(const char *ip, const int port)
    {
        if (m_isRunning)
        {
            m_logger->Log(Logger::Err, "Server is already running");
            return false;
        }

        struct sockaddr_in addr;
        int uvErr = uv_ip4_addr(ip, port, &addr);
        if (0 != uvErr)
        {
            m_logger->Log(Logger::Err,
                "Invalid IP or Port. (%s %d). Error: %s",
                ip, port, uv_strerror(uvErr));
            return false;
        }

        UVTCPWrap listener(m_uvLoop);
        uvErr = uv_tcp_bind(listener.GetTCP(),
            static_cast<const struct sockaddr*>(static_cast<void*>(&addr)), 0);
        if (0 != uvErr)
        {
            m_logger->Log(Logger::Err, "Cannot bind to %s %d: %s",
                ip, port, uv_strerror(uvErr));
            return false;
        }

        uvErr = uv_listen(listener.GetStream(), LISTEN_BACKLOG, OnConnect);
        if (0 != uvErr)
        {
            m_logger->Log(Logger::Err, "Cannot listen on %s %d: %s",
                ip, port, uv_strerror(uvErr));
            return false;
        }

        m_isRunning = true;
        uvErr = uv_run(m_uvLoop, UV_RUN_DEFAULT);
        if (0 != uvErr)
        {
            m_isRunning = false;
            m_logger->Log(Logger::Err, "Cannot start UV Loop: %s",
                uv_strerror(uvErr));
            return false;
        }
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
