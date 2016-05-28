/**
 * Chen - My Network protocol tests
 */

#include "boost/make_shared.hpp"
#include "boost/utility/enable_if.hpp"
#include "chen_log.hxx"

_CHEN_BEGIN_

class LogImpl : public Log
{
public:
    LogImpl(void)
        : Log()
    {
        /* NOP */
    }


    virtual void write_nofmt(const char *src, const size_t srcLine,
                             Flags flags, const char *log)
    {

    }

    virtual void write_vfmt(const char *src, const size_t srcLine,
                            Flags flags, const char *fmt, va_list args)
    {

    }

    virtual void write(const char *src, const size_t srcLine,
                       Flags flags, const char *fmt, ...)
    {
        if (is_skip(flags))
        {
            return;
        }
    }

    virtual bool is_skip(Flags flags)
    {
        return false;
    }
};

boost::shared_ptr<Log> Log::get_instance(void)
{
    static boost::shared_ptr<Log> inst;
    if (inst == nullptr)
    {
        inst = boost::make_shared<LogImpl>();
    }
    return inst;
}

Log::Log(void)
{
    /* NOP */
}

_CHEN_END_
