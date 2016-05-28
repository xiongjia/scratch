/**
 * Chen - My Network protocol tests
 */

#ifndef _CHEN_LOG_HXX_
#define _CHEN_LOG_HXX_ 1

#include <stdarg.h>
#include "boost/shared_ptr.hpp"
#include "boost/utility.hpp"

#include "chen_types.hxx"

_CHEN_BEGIN_

class Log : boost::noncopyable
{
public:
    typedef enum {
        none = 0,
        err  = (1 << 0),
        war  = (1 << 1),
        inf  = (1 << 2),
        dbg  = (1 << 3)
    } Flags;

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

protected:
    Log(void);
};

_CHEN_END_

#endif /* !defined(_CHEN_LOG_HXX_) */
