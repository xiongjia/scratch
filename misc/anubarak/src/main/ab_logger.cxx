/**
 */

#include <stdarg.h>
#include <vector>

#include <boost/make_shared.hpp>
#include <boost/thread/mutex.hpp>
#include <boost/foreach.hpp>
#include <boost/filesystem.hpp>

#include "ab_logger.h"

_ANUBARAK_BEGIN

class LoggerImpl : public Logger
{
public:
    static boost::shared_ptr<Logger> m_logger;

private:
    boost::mutex m_mutex;
    Level m_level;
    std::vector<boost::shared_ptr<LoggerHandler>> m_handler;

public:
    LoggerImpl(void)
        : Logger()
        , m_level(Level::None)
    {
        /* NOP */
    }

    virtual const Level GetLogLevel(void) const
    {
        return m_level;
    }

    virtual void SetLogLevel(Level level)
    {
        m_level = level;
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

    virtual void Log(const char *src,
                     const int srcLine,
                     LoggerMask mask,
                     const char *fmt, ...)
    {
        va_list args;
        va_start(args, fmt);
        LogVFmt(src, srcLine, mask, fmt, args);
        va_end(args);
    }

    virtual void LogNoFmt(const char *src,
                          const int srcLine,
                          LoggerMask mask,
                          const char *log)
    {
        WriteLog(src, srcLine, mask, log);
    }

private:
    void WriteLog(const char *src,
                  const int srcLine,
                  LoggerMask mask,
                  const char *log)
    {
        boost::mutex::scoped_lock scopedLock(m_mutex);
        if (IsIgnore(mask))
        {
            return;
        }

        if (NULL != src)
        {
            boost::filesystem::path srcPath(src);
            const std::string &srcFilename = srcPath.filename().string();
            BOOST_FOREACH(auto handler, m_handler)
            {
                handler->OnLog(srcFilename.c_str(), srcLine, mask, log);
            }
        }
        else
        {
            BOOST_FOREACH(auto handler, m_handler)
            {
                handler->OnLog("", 0, mask, log);
            }
        }
    }

    void LogVFmt(const char *src,
                 const int srcLine,
                 LoggerMask mask,
                 const char *fmt,
                 va_list args)
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
        WriteLog(src, srcLine, mask, msg.c_str());
#else
        size_t needed = vsnprintf(NULL, 0, fmt, args);
        std::string msg;
        msg.resize(needed + 1);
        vsnprintf(&msg[0], msg.size(), fmt, args);
        WriteLog(src, srcLine, mask, msg.c_str());
#endif /* defined(_MSC_VER) */
    }

    bool IsIgnore(LoggerMask mask)
    {
        if (m_handler.empty())
        {
            return true;
        }
        switch (m_level)
        {
        case Level::Dbg:
            return (mask & (LoggerMask::Err|
                            LoggerMask::War|
                            LoggerMask::Inf|
                            LoggerMask::Dbg)) == 0;
        case Level::Inf:
            return (mask & (LoggerMask::Err|
                            LoggerMask::War|
                            LoggerMask::Inf)) == 0;
        case Level::War:
            return (mask & (LoggerMask::Err|
                            LoggerMask::War)) == 0;
        case Level::Err:
            return (mask & LoggerMask::Err) == 0;
        case Level::All:
            return false;
        default:
        case Level::None:
            return true;
        }
    }
};

boost::shared_ptr<Logger> LoggerImpl::m_logger;

Logger::Logger(void)
{
    /* NOP */
}

boost::shared_ptr<Logger> Logger::instance(void)
{
    if (!LoggerImpl::m_logger)
    {
        LoggerImpl::m_logger = boost::make_shared<LoggerImpl>();
    }
    return LoggerImpl::m_logger;
}

_ANUBARAK_END
