/**
 * Chen - My Network protocol tests
 */

#include <iostream>

#include "boost/thread.hpp"
#include "boost/thread/mutex.hpp"
#include "boost/make_shared.hpp"
#include "boost/date_time/posix_time/posix_time.hpp"

#include "chen_log.hxx"

namespace pt = boost::posix_time;

class LogImpl : public chen::Log {
private:
  Level level_;
  boost::mutex mutex_;
  Logger handler_;

public:
  LogImpl(void);

  virtual void Append(const char *src, size_t line,
                      Flags flags, const char *fmt, ...);

  virtual void AppendVFmt(const char *src, size_t line,
                          Flags flags, const char *fmt, va_list args);

  virtual void AppendNoFmt(const char *src, size_t line,
                           Flags flags, const char *log);

  virtual void SetLevel(Level level) { level_ = level; }
  virtual const Level GetLevel(void) const { return level_; }

  virtual void SetHandler(const Logger &handler);
  virtual const Logger &GetHandler(void) const { return handler_; }

  virtual bool NeedAppend(Flags flags);

private:
  static bool IsEmptyString(const char *str);

  static const char *GetSrcFilename(const char *src);

  static void ConsoleLogHandle(const char *src, size_t line,
                               chen::Log::Flags flags, const char *msg);
};

void LogImpl::ConsoleLogHandle(const char *src, size_t line,
                               chen::Log::Flags flags, const char *msg) {
  auto now = pt::microsec_clock::local_time();
  std::cout
    << "[" << src << ":" << line << ":" << std::hex << flags << "]("
    << boost::this_thread::get_id() << ")] "
    << "[" << pt::to_iso_string(now) << "] "
    << msg << std::endl;
}

bool LogImpl::IsEmptyString(const char *str) {
  return (nullptr == str || '\0' == *str);
}

const char *LogImpl::GetSrcFilename(const char *src) {
  const char *srcRelFName = (nullptr == src ? "" : src);
  const char *slash = strrchr(srcRelFName, '/');
  if (nullptr == slash) {
    slash = strrchr(srcRelFName, '\\');
  }
  if (nullptr != slash) {
    srcRelFName = slash + 1;
  }
  return srcRelFName;
}

LogImpl::LogImpl(void)
  : Log()
  , handler_(LogImpl::ConsoleLogHandle)
  , level_(Level::LevelNone) {
  /* NOP */
}

bool LogImpl::NeedAppend(Flags flags) {
  switch (level_) {
    case Level::LevelError:
      return !(flags & Flags::Error);
    case Level::LevelWarn:
      return !(flags & (Flags::Error | Flags::Warn));
    case Level::LevelInfo:
      return !(flags & (Flags::Error | Flags::Warn | Flags::Info));
    case Level::LevelDebug:
    case Level::LevelAll:
      return false;
    default:
    case Level::LevelNone:
      return true;
  }
}

void LogImpl::SetHandler(const Logger &handler) {
  boost::mutex::scoped_lock scopedLock(mutex_);
  handler_ = handler;
}

void LogImpl::Append(const char *src, size_t line,
                     Flags flags, const char *fmt, ...) {
  if (IsEmptyString(fmt) || NeedAppend(flags)) {
    return;
  }
  va_list args;
  va_start(args, fmt);
  AppendVFmt(src, line, flags, fmt, args);
  va_end(args);
}

void LogImpl::AppendVFmt(const char *src, size_t line,
                         Flags flags, const char *fmt, va_list args) {
  if (IsEmptyString(fmt) || NeedAppend(flags)) {
    return;
  }

  const size_t needed = _vsnprintf_c(nullptr, 0, fmt, args);
  std::string msg;
  msg.resize(needed + 1);
  _vsnprintf_c(&msg[0], msg.size(), fmt, args);
  AppendNoFmt(src, line, flags, msg.c_str());
}

void LogImpl::AppendNoFmt(const char *src, size_t line,
                          Flags flags, const char *log) {
  if (IsEmptyString(log) || NeedAppend(flags)) {
    return;
  }
  src = GetSrcFilename(src);
  boost::mutex::scoped_lock scopedLock(mutex_);
  if (!handler_) {
    return;
  }
  handler_(src, line, flags, log);
}

_CHEN_BEGIN_

boost::shared_ptr<Log> Log::GetInstance(void) {
  static boost::shared_ptr<Log> inst = boost::make_shared<LogImpl>();
  return inst;
}

_CHEN_END_
