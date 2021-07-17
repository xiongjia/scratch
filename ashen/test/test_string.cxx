/**
 */

#include <gtest/gtest.h>
#include "ash_string.h"

TEST(AshenString, Basic) {
  ash_str_t test1 = ash_string("abc");
  EXPECT_EQ(test1.len, 3);
  EXPECT_EQ(test1.data, "abc");
}
