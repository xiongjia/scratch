/**
 * Chen - My Network protocol tests
 */

#ifndef _CHEN_LOG_HXX_
#define _CHEN_LOG_HXX_ 1

#include "chen_types.hxx"
#include "boost/utility.hpp"
#include "boost/shared_ptr.hpp"

_CHEN_BEGIN_

class LoggerHandler;

class Logger : boost::noncopyable
{
public:
    typedef enum {
        LevelNone = 0, LevelErr, LevelInf, LevelWar, LevelDbg, LevelAll,
    } Level;

    typedef enum {
        Err = (1 << 0),
        War = (1 << 1),
        Inf = (1 << 2),
        Dbg = (1 << 3)
    } Flags;

public:
    static boost::shared_ptr<Logger> create(void);

public:
    virtual void set_log_level(Level logLevel) = 0;

    virtual const Level get_log_level(void) const = 0;

    virtual void reg_handler(boost::shared_ptr<LoggerHandler> handler) = 0;

    virtual void log(const char *src,
                     size_t srcLine,
                     Flags flags,
                     const char *fmt, ...) = 0;

    virtual void log_nofmt(const char *src,
                           size_t srcLine,
                           Flags flags,
                           const char *log) = 0;

    virtual bool is_ignore(Flags flags) = 0;

protected:
    Logger(void)
    {
        /* NOP */
    }
};

class LoggerHandler : boost::noncopyable
{
public:
    virtual void append(const char *src,
                        size_t srcLine,
                        Logger::Flags flags,
                        const char *log) = 0;
};

_CHEN_END_
#endif /* !defined(_CHEN_LOG_HXX_) */
