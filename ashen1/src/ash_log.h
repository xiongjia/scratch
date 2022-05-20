/**
 */

#ifndef _ASH_LOG_H_
#define _ASH_LOG_H_ 1

#include <time.h>
#include <stdarg.h>
#include "ash_types.h"

typedef struct _ash_log_s ash_log_t;

typedef uint32_t ash_log_opt_t;
#define ASH_LOG_ERR    ((ash_log_opt_t)(1 << 0))
#define ASH_LOG_WARN   ((ash_log_opt_t)(1 << 1))
#define ASH_LOG_INFO   ((ash_log_opt_t)(1 << 2))
#define ASH_LOG_DBG    ((ash_log_opt_t)(1 << 3))
#define ASH_LOG_ALL    (ASH_LOG_ERR|ASH_LOG_WARN|ASH_LOG_INFO|ASH_LOG_DBG)

struct _ash_log_s {
  void *context;
  const char *name;
  ash_log_opt_t level_filter;
  void (*append)(ash_log_t *log, time_t ts, const char *msg);
};

void ash_log_write_nofmt(const char *src, const size_t line,
                         ash_log_t *log, ash_log_opt_t mask,
                         const char *msg);

void ash_log_write_vfmt(const char *src, const size_t line,
                        ash_log_t *log, ash_log_opt_t mask,
                        const char *fmt, va_list args);

void ash_log_write(const char* src, const size_t line,
                   ash_log_t* log, ash_log_opt_t mask,
                   const char* fmt, ...);

#define ASH_LOG_WRITE_NOFMT(_log, _mask, msg) \
  ash_log_write_nofmt(__FILE__, __LINE__, _log, _mask, _msg);

#define ASH_LOG_WRITE(_log, _mask, _fmt, ...) \
  ash_log_write(__FILE__, __LINE__, _log, _mask, _fmt, ##__VA_ARGS__);

#endif /* !defined(_ASH_LOG_H_) */

