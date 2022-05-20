/*
 */

#include <stdio.h>
#include "ash_tests.h"
#include "ash_time.h"

void test_simple_time(void) {
  time_t now = ash_time_now();
  char buf[512];

  printf("%s\n", ash_time_asc(now, buf, sizeof(buf)));
}


ash_unit_test_t unittest_simple_time = {
  "simple-time",
  test_simple_time
};
