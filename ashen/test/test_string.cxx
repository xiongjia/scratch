/**
 */

#include <gtest/gtest.h>
#include "ash_string.h"

TEST(AshenString, Basic) {
  ash_pool_t *pool = ash_pool_create(1024);

  char *dup = ash_str_duplicate(pool, "test data");
  EXPECT_STREQ(dup, "test data");
  ash_pool_destroy(pool);
}

