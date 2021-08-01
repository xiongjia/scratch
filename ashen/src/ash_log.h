/*
 */

#ifndef _ASH_LOG_H_
#define _ASH_LOG_H_ 1

ASH_DECLARE(void) ash_log_append(const char *src, const int src_line,
                                 int level, const char *fmt, ...);

#endif /* !defined(_ASH_LOG_H_) */

