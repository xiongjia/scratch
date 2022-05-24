/*
 */

#include <stdio.h>
#include "ash_mm.h"
#include "ash_str.h"
#include "ash_tests.h"

static void test_simple_str(ash_unit_test_context_t *ctx) {
  ash_pool_t *pool;
  size_t sz;
  char buf[512];
  char *dup;

  pool = ash_pool_create(0);
  dup = ash_pstrdup(pool, "12345678");
  printf("dup = %s\n", dup);
  ash_pool_destroy(pool);

  sz = ash_snprintf(buf, sizeof(buf), "test");
  printf("d: %s, sz %d\n", buf, sz);

  sz = ash_snprintf(buf, sizeof(buf), "str: %% %s", "abc");
  printf("d: %s, sz %d\n", buf, sz);
}

ash_unit_test_t unittest_simple_str = {
  .test_name = "simple-string",
  .invoke = test_simple_str
};
