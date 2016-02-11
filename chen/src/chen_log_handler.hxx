/**
 * Chen - My Network protocol tests
 */

#ifndef _CHEN_LOG_HANDLER_HXX_
#define _CHEN_LOG_HANDLER_HXX_ 1

#include "chen_log.hxx"

_CHEN_BEGIN_

class LoggerHandler : boost::noncopyable
{
public:
    virtual void append(const Log &log) = 0;
};

_CHEN_END_
#endif /* !defined(_CHEN_LOG_HANDLER_HXX_) */
