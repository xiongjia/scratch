/**
 */

#include <gtest/gtest.h>
#include "ash_string.h"

TEST(AshenString, Basic) {
  ash_str_t test1 = ash_string("abc");
  EXPECT_EQ(test1.len, 3);
  EXPECT_STREQ(test1.data, "abc");

  ash_pool_t *pool = ash_pool_create(1024);
  char *dup = ash_str_duplicate(pool, test1.data);
  printf("dup = %s\n", dup);
  EXPECT_STREQ(dup, "abc");
  ash_pool_destroy(pool);
}
