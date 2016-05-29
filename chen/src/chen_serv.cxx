/**
 * Chen - My Network protocol tests
 */

#include "boost/make_shared.hpp"
#include "boost/bind.hpp"
#include "boost/asio.hpp"
#include "chen_serv.hxx"
#include "chen_log.hxx"

_CHEN_BEGIN_

class ServerImpl : public Server
{
private:
    boost::asio::io_service m_ioSvc;


public:
    ServerImpl(void)
        : Server()
    {
        /* NOP */

    }

    virtual bool run(void)
    {
        boost::asio::signal_set signals(m_ioSvc, SIGINT, SIGTERM);
        signals.async_wait(boost::bind(&ServerImpl::on_signals, this, _1, _2));
        m_ioSvc.run();
        return true;
    }

public:
    void on_signals(const boost::system::error_code &errCode,
                    const int signalNum)
    {
        CHEN_LOG(CHEN_LOG_DBG, "Recive signale: %d, ErrCode: %d",
                 signalNum, errCode.value());
    }
};

boost::shared_ptr<Server> Server::create(void)
{
    return boost::make_shared<ServerImpl>();
}

Server::Server(void)
{
    /* NOP */
}

_CHEN_END_
