/**
 *
 */

#include <gtest/gtest.h>
#include "ash_alloc.h"

TEST(HelloTest, BasicAssertions) {
  EXPECT_STRNE("hello", "world");
  EXPECT_EQ(7 * 6, 42);
}

