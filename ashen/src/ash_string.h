/**
 */

#ifndef _ASH_STRING_H_
#define _ASH_STRING_H_ 1

#include "ash_general.h"
#include "ash_palloc.h"

typedef struct _ash_str {
  size_t len;
  char *data;
} ash_str_t;

#define ash_string(str)     { sizeof(str) - 1, (char *) str }
#define ash_null_string     { 0, NULL }

ASH_DECLARE(char *) ash_str_duplicate(ash_pool_t *pool, const char *src);

#endif /* !defined(_ASH_STRING_H_) */

