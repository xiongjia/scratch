/**
 * Chen - My Network protocol tests
 */

#include <iostream>
#include "boost/make_shared.hpp"
#include "boost/date_time/posix_time/posix_time.hpp"
#include "chen_log_console.hxx"

_CHEN_BEGIN_

boost::shared_ptr<LoggerHandler> LogConsole::create(const std::ostream &output)
{
    return boost::make_shared<LogConsole>(output);
}

void LogConsole::append(const chen::Log &log)
{
    /* [date_time] [src:line] [flags] [thread id] logs */
    auto logTM = boost::posix_time::from_time_t(log.get_tm());
    std::cout
        << "[" << boost::posix_time::to_iso_extended_string(logTM) << "] "
        << "[" << log.get_src() << ":" << log.get_src_line() << "] "
        << "[" << log.get_flags() << "] "
        << "[" << log.get_thread_id() << "] "
        << log.get_log()
        << std::endl;
}

_CHEN_END_
