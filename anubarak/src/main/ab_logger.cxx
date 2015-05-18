/**
 */

#include <boost/make_shared.hpp>
#include <boost/thread/mutex.hpp>

#include "ab_logger.h"

_ANUBARAK_BEGIN

class LoggerImpl : public Logger
{
public:
    static boost::shared_ptr<Logger> m_logger;

private:
    Level m_level;

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

    virtual void Log(const char *src,
                     const int srcLine,
                     LoggerMask mask,
                     const char *fmt, ...)
    {
        /* XXX */
    }

    virtual void LogNoFmt(const char *src,
                          const int srcLine,
                          LoggerMask mask,
                          const char *log)
    {
        /* XXX */
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
