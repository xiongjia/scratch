/*
 */

#include "ash_config.h"
#include "ash_general.h"
#include "ash_alloc.h"

#include <stdlib.h>
#include <string.h>

ASH_DECLARE(void *) ash_alloc(size_t size) {
  return malloc(size);
}

ASH_DECLARE(void *) ash_calloc(size_t size) {
  void *mem = ash_alloc(size);
  if (NULL == mem) {
    return NULL;
  }
  memset(mem, 0, size);
  return mem;
}

ASH_DECLARE(void) ash_free(void *mem) {
  if (NULL != mem) {
    free(mem);
  }
}

