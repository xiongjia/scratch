/*
 */

#ifndef _TACHI_STR_H_
#define _TACHI_STR_H_ 1

#include "tachi-general.h"
#include "tachi-mm.h"

typedef struct _tachi_str {
  uint32_t len;
  uchar_t *data;
} tachi_str_t;

#define tachi_string(_str)  { sizeof(_str) - 1, (uchar_t*) _str }

#define tachi_str_is_empty(_str) (_str == NULL || _str[0] = '\0')

char* tachi_pstrdup(tachi_pool_t *pool, const char *src);

#endif /* !defined(_TACHI_STR_H_) */
