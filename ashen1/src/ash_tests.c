/**
 */

#include <stdio.h>
#include <string.h>
#include "ash_types.h"
#include "ash_tests.h"

static ash_unit_test_t *ALL_TESTS[] = {
  &unittest_simple_str,
  &unittest_simple_time
};

static void run_test(ash_unit_test_t* unit_test) {
  printf("Running unit test %s\n", unit_test->test_name);
  unit_test->invoke_test();
}

static void run_tests(const char *match) {
  ash_unit_test_t* unit_test;
  int32_t i;

  for (i = 0; i < sizeof(ALL_TESTS) / sizeof(ALL_TESTS[0]); ++i) {
    unit_test = ALL_TESTS[i];
    if (match == NULL) {
      run_test(unit_test);
      continue;
    }
    if (strstr(unit_test->test_name, match)) {
      run_test(unit_test);
    }
  }
}

int main(const int argc, const char **argv) {
  run_tests(argc <= 1 ? NULL : argv[1]);
  return 0;
}
