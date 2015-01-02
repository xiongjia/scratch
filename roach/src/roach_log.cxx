/**
 * Roach - A simple HTTP Server. (libuv & http-parser & boost & C++11 & CMake)
 */

#include <stdarg.h>
#include <vector>
#include "boost/thread/thread.hpp"
#include "boost/thread/mutex.hpp"
#include "boost/utility.hpp"
#include "boost/make_shared.hpp"
#include "boost/foreach.hpp"
#include "roach_log.hxx"

namespace roach {

class LoggerImpl : public Logger
{
private:
    boost::mutex m_mutex;
    Level        m_level;
    std::vector<boost::shared_ptr<LoggerHandler>> m_handler;

public:
    LoggerImpl(void);

    bool IsIgnore(Mask mask)
    {
        if (m_handler.empty())
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
#if defined(_MSC_VER)
        size_t needed = _vsnprintf_c(NULL, 0, fmt, args);
        std::string msg;
        msg.resize(needed + 1);
        _vsnprintf_c(&msg[0], msg.size(), fmt, args);
        LogNoFmt(mask, msg.c_str());
#else
        size_t needed = vsnprintf(NULL, 0, fmt, args);
        std::string msg;
        msg.resize(needed + 1);
        vsnprintf(&msg[0], msg.size(), fmt, args);
        LogNoFmt(mask, msg.c_str());
#endif /**/
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
        BOOST_FOREACH(auto handler, m_handler)
        {
            handler->OnLog(log);
        }
    }

    virtual void RegisterHandler(boost::shared_ptr<LoggerHandler> handler)
    {
        if (!handler)
        {
            return;
        }
        boost::mutex::scoped_lock scopedLock(m_mutex);
        m_handler.push_back(handler);
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

boost::shared_ptr<Logger> Logger::Create(void)
{
    return boost::make_shared<LoggerImpl>();
}

LoggerImpl::LoggerImpl(void)
    : Logger()
    , m_level(LevelErr)
{
    /* NOP */
}

} /* namespace roach */
