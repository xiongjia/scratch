/**
 * Chen - My Network protocol tests
 */

#include <stdio.h>
#include "boost/make_shared.hpp"
#include "boost/date_time.hpp"
#include "chen_logger.hxx"
#include "chen_log_console.hxx"

int main(int argc, char **argv)
{
    auto logger = chen::Logger::create();
    logger->reg_handler(chen::LogConsole::create());
    logger->set_log_level(chen::Logger::LevelAll);
    logger->log(__FILE__, __LINE__, chen::Log::Err, "test %s", "my test1");
    logger->log(__FILE__, __LINE__, chen::Log::Err, "test %s", "my test2");
    return 0;
}

