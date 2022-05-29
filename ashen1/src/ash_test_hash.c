/**
 */

#include "ash_hash.h"
#include "ash_tests.h"


static void test_simple_hash(ash_unit_test_context_t *ctx,
                             ash_unit_test_case_context_t *tc) {
  ash_hash_t *hash;
  const char *data;

  hash = ash_hash_create(tc->pool);
  ash_hash_set(hash, "test", sizeof("test"), "data");
  data = (const char *)ash_hash_get(hash, "test", sizeof("test"));
  ASHTU_STR_EQU(tc, "data", data);
}

ash_unit_test_case_t unittest_simple_hash = {
  .test_name = "simple-hash",
  .invoke = test_simple_hash
};
