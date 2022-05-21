/**
 */

#include <string.h>
#include <wtypes.h>
#include "ash_log.h"
#include "ash_file.h"

#define ASH_LOG_BUF_SIZE   (1024 * 2)

boolean_t ash_log_skip(ash_log_t* log, ash_log_opt_t mask) {
  if (log == NULL || log->append == NULL) {
    return ASH_TRUE;
  }
  if (log->level_filter == 0) {
    return ASH_FALSE;
  }
  return (log->level_filter & mask) ? ASH_FALSE : ASH_TRUE;
}

void ash_log_write_nofmt(const char* src, const size_t line,
                         ash_log_t* log, ash_log_opt_t mask,
                         const char* msg) {
  time_t now;
  const char *src_filename;
  if (ash_log_skip(log, mask)) {
    return;
  }
  src_filename = ash_file_get_filename(src);
  now = time(&now);
  log->append(log, src_filename, line, now, msg == NULL ? "" : msg);
}

void ash_log_write(const char* src, const size_t line,
                   ash_log_t* log, ash_log_opt_t mask,
                   const char* fmt, ...) {
  va_list args;
  if (ash_log_skip(log, mask)) {
    return;
  }
  va_start(args, fmt);
  ash_log_write_vfmt(src, line, log, mask, fmt, args);
  va_end(args);
}

void ash_log_write_vfmt(const char* src, const size_t line,
                        ash_log_t* log, ash_log_opt_t mask,
                        const char* fmt, va_list args) {
  char buf[ASH_LOG_BUF_SIZE];
  if (ash_log_skip(log, mask)) {
    return;
  }
  vsnprintf_s(buf, _countof(buf), _TRUNCATE, fmt, args);
  ash_log_write_nofmt(src, line, log, mask, buf);
}
