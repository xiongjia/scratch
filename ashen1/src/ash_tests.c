/**
 */

#include <stdio.h>
#include <string.h>
#include "ash_types.h"
#include "ash_tests.h"
#include "ash_log.h"

static ash_unit_test_t *ALL_TESTS[] = {
  &unittest_simple_str,
  &unittest_simple_time,
  &unittest_filename
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

static void log_write(ash_log_t* log, time_t ts, const char* msg) {
  printf("%s\n", msg);
}

ash_log_t LOG = {
  NULL, "unit-tests", ASH_LOG_ALL, log_write
};

int main(const int argc, const char **argv) {

  ASH_LOG_WRITE(&LOG, ASH_LOG_DBG, "test %s", "data");

  run_tests(argc <= 1 ? NULL : argv[1]);
  return 0;
}
