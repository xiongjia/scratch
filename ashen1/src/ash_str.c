/*
 */

#include <string.h>
#include <memory.h>
#include "ash_str.h"


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
