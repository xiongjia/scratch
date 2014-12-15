/**
 *
 */

#include <stdarg.h>
#include "boost/thread/thread.hpp"
#include "boost/thread/mutex.hpp"
#include "boost/utility.hpp"

#include "roach_log.hxx"

namespace roach {

class LoggerImpl : public Logger
{
private:
    Level        m_level;
    Handler      m_handler;
    boost::mutex m_mutex;

public:
    LoggerImpl(Level logLevel, Handler handler)
        : Logger()
        , m_level(logLevel)
        , m_handler(handler)
    {
        /* NOP */
    }

    bool IsIgnore(Mask mask)
    {
        if (m_handler == nullptr)
        {
            return true;
        }
        switch (m_level)
        {
        case LevelInf:
            return (mask & (Err|War|Inf)) == 0;
        case LevelWar:
            return (mask & (Err|War)) == 0;
        case LevelDbg:
            return (mask & (Err|War|Inf|Dbg)) == 0;
        case LevelErr:
            return (mask & Err) == 0;
        case LevelAll:
            return false;
        default:
        case LevelNone:
            return true;
        }
    }

    void LogVFmt(Mask mask, const char *fmt, va_list args)
    {
        if (IsIgnore(mask))
        {
            return;
        }
        size_t needed = _vsnprintf_c(NULL, 0, fmt, args);
        std::string msg;
        msg.resize(needed + 1);
        _vsnprintf_c(&msg[0], msg.size(), fmt, args);
        LogNoFmt(mask, msg.c_str());
    }

    virtual void Log(Mask mask, const char *fmt, ...)
    {
        if (IsIgnore(mask))
        {
            return;
        }
        va_list args;
        va_start(args, fmt);
        LogVFmt(mask, fmt, args);
        va_end(args);
    }

    virtual void LogNoFmt(Mask mask, const char *log)
    {
        boost::mutex::scoped_lock scopedLock(m_mutex);
        if (IsIgnore(mask) || NULL == log)
        {
            return;
        }
        m_handler(log);
    }

    virtual void SetHandler(Handler handler)
    {
        boost::mutex::scoped_lock scopedLock(m_mutex);
        m_handler = handler;
    }

    virtual void SetLevel(Level logLevel)
    {
        boost::mutex::scoped_lock scopedLock(m_mutex);
        m_level = logLevel;
    }
};

Logger::Logger(void)
{
    /* NOP */
}

boost::shared_ptr<Logger> Logger::Create(Level logLevel,
                                         Handler handler)
{
    return boost::shared_ptr<Logger>(new LoggerImpl(logLevel, handler));
}

} /* namespace roach */
