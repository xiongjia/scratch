/**
 * Chen - My Network protocol tests
 */

#ifndef _LIBCHEN_CHEN_LOG_HXX_
#define _LIBCHEN_CHEN_LOG_HXX_ 1

#include <stdarg.h>

#include "boost/shared_ptr.hpp"
#include "boost/utility.hpp"
#include "boost/function.hpp"
#include "boost/thread.hpp"
#include "boost/date_time/posix_time/posix_time.hpp"
#include "boost/date_time/c_local_time_adjustor.hpp"

#include "chen_types.hxx"

_CHEN_BEGIN_

class LogItem;

class Log : boost::noncopyable {
public:
  typedef enum {
    Error  = (1 << 0),
    Warn   = (1 << 1),
    Info   = (1 << 2),
    Debug  = (1 << 3)
  } Flags;

  typedef enum {
    LevelNone  = 0,
    LevelError = 1,
    LevelInfo  = 2,
    LevelWarn  = 3,
    LevelDebug = 4,
    LevelAll   = 100
  } Level;

  typedef boost::function<void(const LogItem &)> Logger;

public:
  static boost::shared_ptr<Log> GetInstance(void);

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
};

class LogItem : boost::noncopyable
{
private:
    typedef boost::date_time::c_local_adjustor<boost::posix_time::ptime>
        boost_tm_local_adjustor;

private:
    const char                     *m_src;
    const size_t                    m_srcLine;
    const Log::Flags                m_flags;
    const char                     *m_log;
    const boost::thread::id         m_threadId;
    const boost::posix_time::ptime  m_tm;

public:
    const char *get_src(void) const { return m_src; }
    const size_t get_srcline(void) const { return m_srcLine; }
    const Log::Flags get_flags(void) const { return m_flags; }
    const char *get_log(void) const { return m_log; }
    const boost::thread::id &get_threadid(void) const { return m_threadId; }
    const boost::posix_time::ptime &get_tm(void) const { return m_tm; }
    const boost::posix_time::ptime &get_tm_local(void) const
    {
        return boost_tm_local_adjustor::utc_to_local(m_tm);
    }

public:
    LogItem(const char *src, const size_t srcLine, 
             const Log::Flags flags, const char *log)
        : m_src(src)
        , m_srcLine(srcLine)
        , m_log(log)
        , m_flags(flags)
        , m_threadId(boost::this_thread::get_id())
        , m_tm(boost::posix_time::microsec_clock::universal_time())
    {
        /* NOP */
    }
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

#endif /* !defined(_LIBCHEN_CHEN_LOG_HXX_) */
