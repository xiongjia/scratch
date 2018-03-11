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

protected:
  Log(void);
};

_CHEN_END_

#define CHEN_LOG(_flags, _fmt, ...) \
  chen::Log::GetInstance()->Append(__FILE__, __LINE__, \
                                   _flags, _fmt, __VA_ARGS__)

#define CHEN_LOG_ERROR(_fmt, ...) \
  chen::Log::GetInstance()->Append(__FILE__, __LINE__, \
                                   chen::Log::Error, _fmt, __VA_ARGS__)

#define CHEN_LOG_WARN(_fmt, ...) \
  chen::Log::GetInstance()->Append(__FILE__, __LINE__, \
                                   chen::Log::Warn, _fmt, __VA_ARGS__)

#define CHEN_LOG_DEBUG(_fmt, ...) \
  chen::Log::GetInstance()->Append(__FILE__, __LINE__, \
                                   chen::Log::Debug, _fmt, __VA_ARGS__)

#define CHEN_LOG_INFO(_fmt, ...) \
  chen::Log::GetInstance()->Append(__FILE__, __LINE__, \
                                   chen::Log::Info, _fmt, __VA_ARGS__)

#define CHEN_LOG_NOFMT(_flags, _msg) \
  chen::Log::GetInstance()->AppendNoFmt(__FILE__, __LINE__, _flags, _msg)

#endif /* !defined(_LIBCHEN_CHEN_LOG_HXX_) */
