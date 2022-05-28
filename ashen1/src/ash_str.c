/*
 */

#include <string.h>
#include "ash_str.h"
#include "ash_vfmt.h"

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

char *ash_pstrcat(ash_pool_t *pool, ...) {
  va_list args;
  size_t total_len = 0;
  size_t len;
  char *val;
  char *rt;
  char *cur;

  va_start(args, pool);
  while ((val = va_arg(args, char *)) != NULL) {
    total_len += strlen(val);
  }
  va_end(args);

  rt = (char*)ash_pool_alloc(pool, total_len + 1);
  if (NULL == rt) {
    return NULL;
  }

  cur = rt;
  va_start(args, pool);
  while ((val = va_arg(args, char *)) != NULL) {
    len = strlen(val);
    memcpy(cur, val, len);
    cur += len;
  }
  va_end(args);
  *cur = '\0';
  return rt;
}

int32_t ash_vsnprintf(char *buf, size_t buf_sz, const char *fmt, va_list ap) {
  ash_vfmt_buff_t vbuf;
  ASH_VFMT_BUFF_INIT(&vbuf, buf, buf_sz)
  return ash_vformatter(&vbuf, fmt, ap);
}

int32_t ash_snprintf(char *buf, size_t buf_sz, const char *fmt, ...) {
  va_list args;
  int32_t res;

  va_start(args, fmt);
  res = ash_vsnprintf(buf, buf_sz, fmt, args);
  va_end(args);
  return res;
}
