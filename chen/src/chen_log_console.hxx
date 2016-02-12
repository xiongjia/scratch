/**
 * Chen - My Network protocol tests
 */

#ifndef _CHEN_LOG_CONSOLE_HXX_
#define _CHEN_LOG_CONSOLE_HXX_ 1

#include <iostream>
#include "boost/shared_ptr.hpp"
#include "chen_log_handler.hxx"

_CHEN_BEGIN_

class LogConsole : public chen::LoggerHandler
{
private:
    const std::ostream &m_output;

public:
    static boost::shared_ptr<LoggerHandler> create(
        const std::ostream &output = std::cout);

public:
    virtual void append(const chen::Log &log);

    LogConsole(const std::ostream &output)
        : LoggerHandler()
        , m_output(output)
    {
        /* NOP */
    }
};

_CHEN_END_
#endif /* !defined(_CHEN_LOG_CONSOLE_HXX_) */
