/**
 *
 */

#include <gtest/gtest.h>
#include "ash_alloc.h"

TEST(HelloTest, BasicAssertions) {
  const char *result = test_gtest();
  EXPECT_STREQ("test1", result);
  EXPECT_STRNE("hello", "world");
  EXPECT_EQ(7 * 6, 42);
}

