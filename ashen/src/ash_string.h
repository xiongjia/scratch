/**
 */

#ifndef _ASH_STRING_H_
#define _ASH_STRING_H_ 1

#include "ash_general.h"
#include "ash_palloc.h"

#define ASH_STR_IS_NULL_OR_EMPTY(_str) \
  ((_str) == NULL || (_str)[0] == '\0')

ASH_DECLARE(char *) ash_str_duplicate(ash_pool_t *pool, const char *src);

ASH_DECLARE_CDECL(char *) ash_str_sprintf(ash_pool_t *pool,
    const char *fmt, ...);

#endif /* !defined(_ASH_STRING_H_) */

