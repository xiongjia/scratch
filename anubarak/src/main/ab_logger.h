/**
 */

#ifndef _AB_LOGGER_H_
#define _AB_LOGGER_H_ 1

#include <boost/utility.hpp>
#include <boost/shared_ptr.hpp>

#include "ab_types.h"

_ANUBARAK_BEGIN

typedef enum _LoggerMask
{
    Err = (1 << 0), War = (1 << 1), Inf = (1 << 2), Dbg = (1 << 3),
} LoggerMask;

class LoggerHandler : boost::noncopyable{public:    virtual void OnLog(const char *src, const int srcLine,
                       LoggerMask mask, const char *log) = 0;
};

class Logger : boost::noncopyable
{
public:
    typedef enum _Level
    {
        None = 0, Err, Inf, War, Dbg, All
    } Level;

public:
    static boost::shared_ptr<Logger> instance(void);

    virtual void SetLogLevel(Level level) = 0;
    virtual const Level GetLogLevel(void) const = 0;

    virtual void Log(const char *src, const int srcLine,
                     LoggerMask mask, const char *fmt, ...) = 0;

    virtual void LogNoFmt(const char *src, const int srcLine,
                          LoggerMask mask, const char *log) = 0;

protected:
    Logger(void);
};

_ANUBARAK_END

#define AB_LOG_NOFMT(_mask, _log) \
    ab::Logger::instance()->LogNoFmt(__FILE__, __LINE__, _mask, _log);

#if defined(WIN32)
#   define AB_LOG(_mask, _log, ...) \
        ab::Logger::instance()->Log(__FILE__, __LINE__, _mask, _log, __VA_ARGS__);
#else
#   define AB_LOG(_mask, _log, ...) \
        ab::Logger::instance()->Log(__FILE__, __LINE__, _mask, _log, ##__VA_ARGS__);
#endif /* defined(WIN32) */

#endif /* !defined(_AB_LOGGER_H_) */
