/**
 */

#include "tachi-str.h"

char* tachi_pstrdup(tachi_pool_t* pool, const char* src) {
  int src_len;
  char *dup;

  src_len = src == NULL ? 1 : strlen(src);
  dup = tachi_pool_alloc(pool, src_len);
  if (dup == NULL) {
    return NULL;
  }

  if (src == NULL) {
    dup[0] = '\0';
  }
  else {
    memcpy(dup, src, src_len + 1);
  }
  return dup;
}
