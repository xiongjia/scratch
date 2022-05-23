/*
 */

#include <string.h>
#include <memory.h>
#include "ash_str.h"

typedef struct _ash_vformatter_buff_s {
  size_t count;
  char *curpos;
  char *endpos;
} ash_vformatter_buff_t;

static void ash_vformatter_inst_char(ash_vformatter_buff_t *vbuf, char ch) {
  if (vbuf->endpos <= vbuf->curpos) {
    return;
  }
  *(vbuf->curpos) = ch;
  vbuf->curpos++;
  vbuf->count++;
}

static size_t ash_vformatter(ash_vformatter_buff_t *vbuf, const char *fmt, va_list ap) {
  while (*fmt) {
    if (*fmt != '%') {
      ash_vformatter_inst_char(vbuf, *fmt);
      fmt++;
      continue;
    }
    // TODO
  }
  ash_vformatter_inst_char(vbuf, '\0');
  return vbuf->count;
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
  ash_vformatter_buff_t vbuf;

  vbuf.count = 0;
  vbuf.curpos = buf;
  vbuf.endpos = buf + buf_size;
  return ash_vformatter(&vbuf, fmt, ap);
}

size_t ash_snprintf(char *buf, size_t buf_size, const char *fmt, ...) {
  va_list args;
  size_t res;

  va_start(args, fmt);
  res = ash_vsnprintf(buf, buf_size, fmt, args);
  va_end(args);
  return res;
}
