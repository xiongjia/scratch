/**
 */

#include "ash_time.h"

time_t ash_time_now(void) {
  time_t now;
  return time(&now);
}

char* ash_time_asc(time_t tm, char* buf, size_t buf_size) {
  struct tm tm_data;
  errno_t err;
  _localtime64_s(&tm_data , &tm);
  err = asctime_s(buf, buf_size, &tm_data);
  if (err != 0) {
    return NULL;
  }
  return buf;
}
