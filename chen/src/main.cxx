/**
 * Chen - My Network protocol tests
 */

#include <stdio.h>
#include "boost/make_shared.hpp"
#include "chen_log.hxx"

class TestLogger : public chen::LoggerHandler
{
public:
    virtual void append(const char *src,
                        size_t srcLine,
                        chen::Logger::Flags flags,
                        const char *log)
    {
        printf("%s\n", log);
    }
};

int main(int argc, char **argv)
{
    auto logger = chen::Logger::create();
    logger->reg_handler(boost::make_shared<TestLogger>());
    logger->set_log_level(chen::Logger::LevelAll);
    logger->log(__FILE__, __LINE__, chen::Logger::Err, "test %s", "my test");
    return 0;
}

