/*
 */

#include <stdio.h>
#include "ash_tests.h"
#include "ash_time.h"

static void test_simple_time(ash_unit_test_context_t* ctx,
                             ash_unit_test_case_context_t *test_case) {
  time_t now = ash_time_now();
  char buf[512];

  printf("%s\n", ash_time_asc(now, buf, sizeof(buf)));
}

ash_unit_test_case_t unittest_simple_time = {
  "simple-time",
  test_simple_time
};
