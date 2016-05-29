/**
 * Chen - My Network protocol tests
 */

#include <iostream>
#include "boost/utility.hpp"
#include "chen_serv.hxx"
#include "chen_log.hxx"

class Chen : boost::noncopyable
{
private:
    void log(const chen::LogItem &logItem)
    {
        std::cout
            << boost::posix_time::to_simple_string(logItem.get_tm_local())
            << " [" << logItem.get_src() << ":"
            << logItem.get_srcline() << "] "
            << logItem.get_log()
            << std::endl;
    }

public:
    Chen(void)
    {
        auto log = chen::Log::get_instance();
        log->set_level(chen::Log::LEVEL_ALL);
        log->set_handler(boost::bind(&Chen::log, this, _1));
    }

public:
    int run(void)
    {
        CHEN_LOG(CHEN_LOG_INF, "creating server ...");
        auto serv = chen::Server::create();
        CHEN_LOG(CHEN_LOG_INF, "start server");
        const int retVal = serv->run();
        CHEN_LOG(CHEN_LOG_INF, "server exit (%d)", retVal);
        return retVal;
    }
};

int main(int /* argc */, char ** /* argv */)
{
    Chen chen;
    return chen.run();
}
