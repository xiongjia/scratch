/**
 */

#ifndef _ASH_LOG_H_
#define _ASH_LOG_H_ 1

#include <time.h>
#include <stdarg.h>

typedef struct _ash_log_s ash_log_t;

struct _ash_log_s {
  void *context;
  const char *name;

  void (*append)(ash_log_t *log, time_t ts, const char *msg);
};

void ash_log_write(ash_log_t *log, const char *msg);

#endif /* !defined(_ASH_LOG_H_) */

