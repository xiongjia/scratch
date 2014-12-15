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

    uv_loop_t *m_uvLoop;
    uv_tcp_t  m_listener;

public:
    ServerImpl(boost::shared_ptr<Logger> logger)
        : Server()
        , m_logger(logger)
        , m_uvLoop(uv_default_loop())
    {
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

boost::shared_ptr<Server> Server::Create(boost::shared_ptr<Logger> logger)
{
    return boost::shared_ptr<Server>(new ServerImpl(logger));
}

} /* namespace roach */
