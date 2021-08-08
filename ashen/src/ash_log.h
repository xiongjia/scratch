/*
 */

#ifndef _ASH_LOG_H_
#define _ASH_LOG_H_ 1

#include "ash_general.h"

ASH_DECLARE_CDECL(void) ash_log_append(const char *src, const int src_line,
                                       int level, int flags, const char *fmt, ...);

#endif /* !defined(_ASH_LOG_H_) */

