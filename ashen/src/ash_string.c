/**
 */

#include "ash_string.h"


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

