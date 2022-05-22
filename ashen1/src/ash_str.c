/*
 */

#include <string.h>
#include <memory.h>
#include "ash_str.h"

typedef struct _ash_vformatter_buff_s {
  char *curpos;
  char *endpos;
} ash_vformatter_buff_t;

static size_t ash_vformatter(ash_vformatter_buff_t vbuf, const char *fmt, va_list ap) {

   while (*fmt) {
      if (*fmt != '%') {

      }
   }

  return 0;
}

char* ash_pstrdup(ash_pool_t *pool, const char *src) {
  size_t alloc_sz;
  char *dup;

  if (pool == NULL || src == NULL) {
    return NULL;
  }
  alloc_sz = strlen(src) + 1;
  dup = ash_pool_alloc(pool, alloc_sz);
  memcpy(dup, src, alloc_sz);
  return dup;
}

size_t ash_vsnprintf(char *buf, size_t buf_size, const char *fmt, va_list ap) {
  /* TODO */
  return 0;
}
