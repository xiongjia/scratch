/**
 */

#include <gtest/gtest.h>
#include "ash_string.h"

TEST(AshenString, Basic) {
  ash_pool_t *pool = ash_pool_create(1024);

  char *dup = ash_str_duplicate(pool, "test data");
  EXPECT_STREQ(dup, "test data");

  char *data = ash_str_sprintf(pool, "data %s", "abc");
  EXPECT_STREQ(data, "data abc");
  ash_pool_destroy(pool);
}

