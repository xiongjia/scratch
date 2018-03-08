/**
 * Chen - My Network protocol tests
 */

#include "boost/thread/mutex.hpp"
#include "boost/make_shared.hpp"
#include "chen_log.hxx"

class LogImpl;

_CHEN_BEGIN_

boost::shared_ptr<Log> Log::GetInstance(void) {
  static boost::shared_ptr<Log> inst = boost::make_shared<LogImpl>();
  return inst;
}

_CHEN_END_

class LogImpl : public chen::Log {
private:
    boost::mutex m_mutex;
    Level        m_level;
    Logger       m_handler;

public:
    LogImpl(void)
        : Log()
        , m_level(Level::LEVEL_NONE)
    {
        /* NOP */
    }

    virtual void set_level(Level level = Level::LEVEL_NONE) { m_level = level; }
    virtual const Level get_level(void) const { return m_level; }

    virtual void write_nofmt(const char *src, const size_t srcLine,
                             Flags flags, const char *log)
    {
        if (is_skip(flags) || nullptr == log || '\0' == *log)
        {
            return;
        }
        src = get_src_filename(src);
        boost::mutex::scoped_lock scopedLock(m_mutex);
        if (!m_handler)
        {
            return;
        }
        LogItem logItem(src, srcLine, flags, log);
        m_handler(logItem);
    }

    virtual void write_vfmt(const char *src, const size_t srcLine,
                            Flags flags, const char *fmt, va_list args)
    {
        if (is_skip(flags) || nullptr == fmt || '\0' == *fmt)
        {
            return;
        }
        const size_t needed = _vsnprintf_c(nullptr, 0, fmt, args);
        std::string msg;
        msg.resize(needed + 1);
        _vsnprintf_c(&msg[0], msg.size(), fmt, args);
        write_nofmt(src, srcLine, flags, msg.c_str());
    }

    virtual void write(const char *src, const size_t srcLine,
                       Flags flags, const char *fmt, ...)
    {
        if (is_skip(flags) || nullptr == fmt || '\0' == *fmt)
        {
            return;
        }
        va_list args;
        va_start(args, fmt);
        write_vfmt(src, srcLine, flags, fmt, args);
        va_end(args);
    }

    virtual bool is_skip(Flags flags)
    {
        if (!get_handler())
        {
            return true;
        }
        switch (m_level)
        {
        case Level::LEVEL_NONE:
            return true;
        case Level::LEVEL_ERR:
            return !(flags & Flags::FErr);
        case Level::LEVEL_WARN:
            return !(flags & Flags::FErr ||
                     flags & Flags::FWar);
        case Level::LEVEL_INF:
            return !(flags & Flags::FErr ||
                     flags & Flags::FWar ||
                     flags & Flags::FInf);
        default:
        case Level::LEVEL_DBG:
        case Level::LEVEL_ALL:
            return false;
        }
    }

    virtual void set_handler(const Logger &handler)
    {
        boost::mutex::scoped_lock scopedLock(m_mutex);
        m_handler = handler;
    }

    virtual const Logger& get_handler(void) const
    {
        return m_handler;
    }

private:
    const char *get_src_filename(const char *src)
    {
        const char *srcRelFName = (nullptr == src ? "" : src);
        const char *slash = strrchr(srcRelFName, '/');
        if (nullptr == slash)
        {
            slash = strrchr(srcRelFName, '\\');
        }
        if (nullptr != slash)
        {
            srcRelFName = slash + 1;
        }
        return srcRelFName;
    }
};

