/**
 * Chen - My Network protocol tests
 */

#ifndef _CHEN_LOGGER_HXX_
#define _CHEN_LOGGER_HXX_ 1

#include "chen_types.hxx"
#include "boost/utility.hpp"
#include "boost/shared_ptr.hpp"
#include <thread>

#include "chen_log_handler.hxx"

_CHEN_BEGIN_

class Logger : boost::noncopyable
{
public:
    typedef enum {
        LevelNone = 0, LevelErr, LevelInf, LevelWar, LevelDbg, LevelAll,
    } Level;

public:
    static boost::shared_ptr<Logger> create(void);

public:
    virtual void set_log_level(Level logLevel) = 0;
    virtual const Level get_log_level(void) const = 0;

    virtual void reg_handler(boost::shared_ptr<LoggerHandler> handler) = 0;

    virtual void log(const char *src,
                     size_t srcLine,
                     Log::Flags flags,
                     const char *fmt, ...) = 0;

    virtual void log_nofmt(const char *src,
                           size_t srcLine,
                           Log::Flags flags,
                           const char *log) = 0;

    virtual bool is_ignore(Log::Flags flags) = 0;

protected:
    Logger(void)
    {
        /* NOP */
    }
};

_CHEN_END_
#endif /* !defined(_CHEN_LOGGER_HXX_) */
