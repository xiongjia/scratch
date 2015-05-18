/**
 */

#include <boost/make_shared.hpp>

#include "anubarak.h"
#include "ab_logger.h"

class ConsoleLogger : public ab::LoggerHandler
{
public:
    virtual void OnLog(const char *src,
                       const int srcLine,
                       ab::LoggerMask mask,
                       const char *log)
    {
        fprintf(stdout, "[%s:%d] [0x%x] %s\n", src, srcLine, mask, log);
    }
};


int main(int argc, char **argv)
{
    /* init logger */
    boost::shared_ptr<ab::Logger> logger = ab::Logger::instance();
    logger->RegisterHandler(boost::make_shared<ConsoleLogger>());
    logger->SetLogLevel(ab::Logger::Level::All);

    /* create Anubarak */
    boost::shared_ptr<ab::Anubarak> anubarak = ab::Anubarak::Create();
    AB_LOG(ab::LoggerMask::Dbg, "Creating Anubarak");
    return 0;
}
