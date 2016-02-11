/**
* Chen - My Network protocol tests
*/

#define _CRT_SECURE_NO_WARNINGS 1
#include <stdarg.h>
#include <stdio.h>
#include <vector>
#include "boost/make_shared.hpp"
#include "boost/thread/mutex.hpp"
#include "boost/foreach.hpp"
#include "chen_logger.hxx"

_CHEN_BEGIN_

class LoggerImpl : public Logger
{
private:
    boost::mutex m_mutex;
    Level m_logLevel;
    std::vector<boost::shared_ptr<LoggerHandler>> m_handler;

public:
    LoggerImpl(void);

    virtual void reg_handler(boost::shared_ptr<LoggerHandler> handler);

    virtual void set_log_level(Level logLevel)
    {
        m_logLevel = logLevel;
    }

    virtual const Level get_log_level(void) const
    {
        return m_logLevel;
    }

    virtual bool is_ignore(Log::Flags flags)
    {
        switch (m_logLevel)
        {
        case LevelInf:
            return (flags & (Log::Err | Log::War | Log::Inf)) == 0;
        case LevelWar:
            return (flags & (Log::Err | Log::War)) == 0;
        case LevelDbg:
            return (flags & (Log::Err | Log::War | Log::Inf | Log::Dbg)) == 0;
        case LevelErr:
            return (flags & Log::Err) == 0;
        case LevelAll:
            return false;
        default:
        case LevelNone:
            return true;
        }
    }

    virtual void log(const char *src,
                     size_t srcLine,
                     Log::Flags flags,
                     const char *fmt, ...);

    virtual void log_nofmt(const char *src,
                           size_t srcLine,
                           Log::Flags flags,
                           const char *log);

private:
    void log_vfmt(const char *src,
                  size_t srcLine,
                  Log::Flags flags,
                  const char *fmt,
                  va_list args);
};

LoggerImpl::LoggerImpl(void)
    : m_logLevel(LevelNone)
{
    /* NOP */
}

void LoggerImpl::reg_handler(boost::shared_ptr<LoggerHandler> handler)
{
    boost::mutex::scoped_lock lock(m_mutex);
    m_handler.push_back(handler);
}

void LoggerImpl::log_nofmt(const char *src,
                           size_t srcLine,
                           Log::Flags flags,
                           const char *log)
{
    if (is_ignore(flags) || nullptr == log)
    {
        return;
    }

    Log logItem(src, srcLine, flags, log);
    boost::mutex::scoped_lock lock(m_mutex);
    BOOST_FOREACH(auto handler, m_handler)
    {
        handler->append(logItem);
    }
}

void LoggerImpl::log(const char *src,
                     size_t srcLine,
                     Log::Flags flags,
                     const char *fmt, ...)
{
    if (is_ignore(flags) || nullptr == fmt)
    {
        return;
    }
    va_list args;
    va_start(args, fmt);
    log_vfmt(src, srcLine, flags, fmt, args);
    va_end(args);
}

void LoggerImpl::log_vfmt(const char *src,
                          size_t srcLine,
                          Log::Flags flags,
                          const char *fmt,
                          va_list args)
{
    if (is_ignore(flags) || nullptr == fmt)
    {
        return;
    }
    size_t needed = std::vsnprintf(nullptr, 0, fmt, args);
    std::string logMsg;
    logMsg.resize(needed + sizeof(char));
    std::vsnprintf(&logMsg[0], logMsg.size(), fmt, args);
    log_nofmt(src, srcLine, flags, logMsg.c_str());
}

boost::shared_ptr<Logger> Logger::create(void)
{
    static boost::shared_ptr<Logger> inst = boost::make_shared<LoggerImpl>();
    return inst;
}

_CHEN_END_
