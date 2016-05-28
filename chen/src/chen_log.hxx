/**
 * Chen - My Network protocol tests
 */

#ifndef _CHEN_LOG_HXX_
#define _CHEN_LOG_HXX_ 1

#include <stdarg.h>
#include "boost/shared_ptr.hpp"
#include "boost/utility.hpp"
#include "boost/function.hpp"

#include "chen_types.hxx"

_CHEN_BEGIN_

class Log : boost::noncopyable
{
public:
    typedef enum
    {
        FErr  = (1 << 0),
        FWar  = (1 << 1),
        FInf  = (1 << 2),
        FDbg  = (1 << 3)
    } Flags;

    typedef enum
    {
        LEVEL_NONE = 0,
        LEVEL_ERR  = 1,
        LEVEL_INF  = 2,
        LEVEL_WARN = 3,
        LEVEL_DBG  = 4,
        LEVEL_ALL  = 100
    } Level;

    typedef struct
    {
        const char *src;
        size_t      srcLine;
        Flags       flags;
        const char *log;
    } LogItem;
    typedef boost::function<void(const LogItem &)> Logger;

public:
    static boost::shared_ptr<Log> get_instance(void);

public:
    virtual bool is_skip(Flags flags) = 0;

    virtual void write_nofmt(const char *src, const size_t srcLine,
                             Flags flags, const char *log) = 0;

    virtual void write_vfmt(const char *src, const size_t srcLine,
                            Flags flags, const char *fmt, va_list args) = 0;

    virtual void write(const char *src, const size_t srcLine,
                       Flags flags, const char *fmt, ...) = 0;

    virtual void set_level(Level level = Level::LEVEL_NONE) = 0;
    virtual const Level get_level(void) const = 0;

    virtual void set_handler(const Logger &handler) = 0;
    virtual const Logger& get_handler(void) const = 0;
protected:
    Log(void);
};

_CHEN_END_

#define CHEN_LOG_ERR    (chen::Log::FErr)
#define CHEN_LOG_WAR    (chen::Log::FWar)
#define CHEN_LOG_INF    (chen::Log::FInf)
#define CHEN_LOG_DBG    (chen::Log::FDbg)

#define CHEN_LOG(_flags, _fmt, ...) \
    chen::Log::get_instance()->write(__FILE__, __LINE__, _flags, _fmt, __VA_ARGS__)

#define CHEN_LOG_NOFMT(_flags, _msg) \
    chen::Log::get_instance()->write_nofmt(__FILE__, __LINE__, _flags, _msg)

#endif /* !defined(_CHEN_LOG_HXX_) */
