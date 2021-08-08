/**
 */

#include <stdarg.h>
#include <stdio.h>
#include "ash_string.h"

static size_t ash_str_vsprintf_len(const char *fmt, va_list vargs) {
  return vsnprintf(NULL, 0, fmt, vargs);
}

static char * ash_str_vsprintf(ash_pool_t *pool, size_t buf_len,
                               const char *fmt, va_list vargs) {
  char *buf = ash_pool_calloc(pool, buf_len);
  if (buf == NULL) {
    return NULL;
  }
  vsnprintf(buf, buf_len, fmt, vargs);
  return buf;
}

ASH_DECLARE(char *) ash_str_duplicate(ash_pool_t *pool, const char *src) {
  size_t len;
  char *dest;

  if (src == NULL) {
    return NULL;
  }

  len = strlen(src) + sizeof(char);
  dest = ash_pool_alloc(pool, len);
  if (dest == NULL) {
    return NULL;
  }
  memcpy(dest, src, len);
  return dest;
}

ASH_DECLARE_CDECL(char *) ash_str_sprintf(ash_pool_t *pool,
                                          const char *fmt, ...) {
  va_list vargs;
  size_t buf_len;
  char *result;

  va_start(vargs, fmt);
  buf_len = ash_str_vsprintf_len(fmt, vargs);
  va_end(vargs);

  va_start(vargs, fmt);
  result = ash_str_vsprintf(pool, buf_len + sizeof(char), fmt, vargs);
  va_end(vargs);
  return result;
}

