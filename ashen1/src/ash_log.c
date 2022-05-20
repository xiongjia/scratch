/**
 */

#include "ash_log.h"

void ash_log_write(ash_log_t *log, const char* msg) {
  time_t now;

  if (log == NULL || log->append == NULL) {
    return;
  }

  now = time(&now);
  log->append(log, now, "test");
}
