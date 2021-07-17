/**
 */

#ifndef _ASH_STRING_H_
#define _ASH_STRING_H_ 1

#include "ash_general.h"

#ifdef __cplusplus
extern "C" {
#endif /* defined(__cplusplus) */

typedef struct _ash_str {
  size_t len;
  char *data;
} ash_str_t;

#define ash_string(str)     { sizeof(str) - 1, (char *) str }
#define ash_null_string     { 0, NULL }

#ifdef __cplusplus
}
#endif /* defined(__cplusplus) */

#endif /* !defined(_ASH_STRING_H_) */

