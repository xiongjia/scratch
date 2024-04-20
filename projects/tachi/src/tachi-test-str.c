/**
 */

#include "tachi-str.h"

void tachi_test_str(void) {
  tachi_pool_t *pool;
  char *dup;
  tachi_str_t test = tachi_string("PASS");
  printf("%s\n", test.data);

  pool = tachi_pool_create(100);
  dup = tachi_pstrdup(pool, "1234567890");
  printf("%s\n", dup);

  tachi_pool_destroy(pool);
}
