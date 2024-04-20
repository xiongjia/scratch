/**
 * Chen - My Network protocol tests
 */

#ifndef _LIBCHEN_CHEN_LOG_HXX_
#define _LIBCHEN_CHEN_LOG_HXX_ 1

#include <stdarg.h>

#include "boost/shared_ptr.hpp"
#include "boost/utility.hpp"
#include "boost/function.hpp"
#include "chen_types.hxx"

#define LOG(_level, _fmt, ...) \
  chen::Log::GetInstance()->Append(__FILE__, __LINE__, _level, \
                                   _fmt, ##__VA_ARGS__)

#define LOG_ERR(_fmt, ...) LOG(chen::Log::Error, _fmt, ##__VA_ARGS__)
#define LOG_WAR(_fmt, ...) LOG(chen::Log::Warn, _fmt, ##__VA_ARGS__)
#define LOG_INF(_fmt, ...) LOG(chen::Log::Info, _fmt, ##__VA_ARGS__)
#define LOG_DBG(_fmt, ...) LOG(chen::Log::Debug, _fmt, ##__VA_ARGS__)


_CHEN_BEGIN_

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

  typedef boost::function<void(const char *src,
                               size_t line,
                               Flags flags,
                               const char *msg)> Logger;

public:
  static boost::shared_ptr<Log> GetInstance(void);

  virtual void Append(const char *src, size_t line,
                      Flags flags, const char *fmt, ...) = 0;

  virtual void AppendNoFmt(const char *src, size_t line,
                           Flags flags, const char *log) = 0;

  virtual void AppendVFmt(const char *src, size_t line,
                          Flags flags, const char *fmt, va_list args) = 0;

  virtual bool NeedAppend(Flags flags) = 0;

  virtual void SetLevel(Level level) = 0;
  virtual const Level GetLevel(void) const = 0;

  virtual void SetHandler(const Logger &handler) = 0;
  virtual const Logger &GetHandler(void) const = 0;
};

_CHEN_END_

#endif /* !defined(_LIBCHEN_CHEN_LOG_HXX_) */
