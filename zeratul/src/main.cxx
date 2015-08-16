/**
 * Zeratul - A simple socks proxy. (Boost ASIO Example)
 */

#include "zeratul.hxx"

int main(int argc, char **argv)
{
    zeratul::LogMgr::Instance()->Initialize();
    BOOST_LOG_NAMED_SCOPE("main");
    BOOST_LOG_TRIVIAL(info) << "Info Log";

    try
    {
        static boost::asio::io_service ioSvc;

        /* create server */
        auto serv = zeratul::Server::Create(ioSvc, 9090);

        /* handle signals */
        boost::asio::signal_set signals(ioSvc, SIGINT, SIGTERM);
        signals.async_wait([](const boost::system::error_code& errCode,
                              int signalNum)
        {
            BOOST_LOG_TRIVIAL(info) << "Terminate signals: " << signalNum;
            if (!errCode)
            {
                ioSvc.stop();
            }
            else
            {
                BOOST_LOG_TRIVIAL(error) << "Signals Error: " << errCode;
            }
        });
        ioSvc.run();
    }
    catch (std::exception &e)
    {
        BOOST_LOG_TRIVIAL(fatal) << "Exception: " << e.what();
    }
    return 0;
}
