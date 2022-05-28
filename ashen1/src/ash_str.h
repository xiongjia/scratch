/**
 */

#ifndef _ASH_STR_H_
#define _ASH_STR_H_ 1

#include <stdarg.h>
#include "ash_types.h"
#include "ash_mm.h"

char *ash_pstrdup(ash_pool_t *pool, const char *src);

char *ash_pstrcat(ash_pool_t *pool, ...);

int32_t ash_vsnprintf(char *buf, size_t buf_sz, const char *fmt, va_list ap);

int32_t ash_snprintf(char *buf, size_t buf_sz, const char *fmt, ...);

#endif /* !defined(_ASH_STR_H_) */
