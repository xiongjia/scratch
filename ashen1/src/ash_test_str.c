/*
 */

#include "ash_mm.h"
#include "ash_str.h"
#include "ash_tests.h"

void test_str(void) {
  ash_pool_t *pool;
  char *dup;

  pool = ash_pool_create(0);
  dup = ash_pstrdup(pool, "12345678");
  printf("dup = %s\n", dup);
  ash_pool_destroy(pool);
}
