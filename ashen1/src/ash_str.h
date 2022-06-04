/**
 */

#ifndef _ASH_STR_H_
#define _ASH_STR_H_ 1

#include <stdarg.h>

#include "ash_types.h"
#include "ash_mm.h"
#include "ash_list.h"

char *ash_pstrdup(ash_pool_t *pool, const char *src);

char *ash_pstrcat(ash_pool_t *pool, ...);

ash_list_t *ash_slist_create(ash_pool_t *pool, ...);

ash_list_t *ash_slist_vcreate(ash_pool_t *pool, va_list ap);

boolean_t ash_slist_compare(ash_list_t *l1, ash_list_t *l2);

int32_t ash_vsnprintf(char *buf, size_t buf_sz, const char *fmt, va_list ap);

int32_t ash_snprintf(char *buf, size_t buf_sz, const char *fmt, ...);

ash_list_t *ash_parse_lines(ash_pool_t *pool, const char *src,
                            uint32_t max_line_sz);

#endif /* !defined(_ASH_STR_H_) */
