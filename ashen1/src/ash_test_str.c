/*
 */

#include <stdio.h>
#include "ash_mm.h"
#include "ash_str.h"
#include "ash_tests.h"

static void test_simple_str(ash_unit_test_context_t* ctx) {
  ash_pool_t *pool;
  char *dup;

  pool = ash_pool_create(0);
  dup = ash_pstrdup(pool, "12345678");
  printf("dup = %s\n", dup);
  ash_pool_destroy(pool);
}

ash_unit_test_t unittest_simple_str = {
  "simple-string",
  test_simple_str
};
