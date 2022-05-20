/**
 */

#ifndef _ASH_TIME_H_
#define _ASH_TIME_H_ 1

#include <time.h>

time_t ash_time_now(void);

char *ash_time_asc(time_t tm, char *buf, size_t buf_size);

#endif /* !defined(_ASH_TIME_H_) */
