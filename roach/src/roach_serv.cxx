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

public:
    ServerImpl(boost::shared_ptr<Context> ctx)
        : Server()
        , m_ctx(ctx)
        , m_logger(ctx->GetLogger())
        , m_uvLoop(uv_default_loop())
    {
        /* NOP */
    }

    static void OnConnect(uv_stream_t* server_handle, int status)
    {

    }

    virtual bool Listen(const char *ip, const int port)
    {
        int result;

        struct sockaddr_in addr;
        result = uv_ip4_addr(ip, port, &addr);
        if (0 != result)
        {
            /* xxx LOG */
            return false;
        }

        result = uv_tcp_init(m_uvLoop, &m_listener);
        if (0 != result)
        {
            /* xxx LOG */
            return false;
        }

        result = uv_tcp_bind(&m_listener, (const struct sockaddr*)&addr, 0);
        if (0 != result)
        {
            /* xxx LOG */
            return false;
        }

        result = uv_listen((uv_stream_t*)&m_listener, 128, OnConnect);
        if (0 != result)
        {
            /* xxx LOG */
            return false;
        }
        return true;
    }

    virtual bool Run(void)
    {
        uv_run(m_uvLoop, UV_RUN_DEFAULT);
        return true;
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
