/*
 */

#ifndef _TACHI_STR_H_
#define _TACHI_STR_H_ 1

#include "tachi-general.h"

typedef struct _tachi_str {
  uint32_t len;
  uchar_t *data;
} tachi_str_t;

#define tachi_string(_str)  { sizeof(_str) - 1, (uchar_t*) _str }



#endif /* !defined(_TACHI_STR_H_) */
