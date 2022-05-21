/**
 */

#include <stdio.h>
#include <string.h>
#include "ash_tests.h"
#include "ash_time.h"

static ash_unit_test_t *ALL_TESTS[] = {
  &unittest_simple_str,
  &unittest_simple_time,
  &unittest_filename
};

static void run_test(ash_unit_test_context_t* ctx, ash_unit_test_t* unit_test) {
  printf("Running unit test %s\n", unit_test->test_name);
  unit_test->invoke_test(ctx);
}

static void run_tests(ash_unit_test_context_t *ctx) {
  ash_unit_test_t* unit_test;
  int32_t i;

  for (i = 0; i < ASH_ARRAY_COUNT(ALL_TESTS); ++i) {
    unit_test = ALL_TESTS[i];
    if (ctx->test_filter == NULL) {
      run_test(ctx, unit_test);
      continue;
    }
    if (strstr(unit_test->test_name, ctx->test_filter)) {
      run_test(ctx, unit_test);
    }
  }
}

static void log_write(ash_log_t* log, const char* src, size_t line, time_t ts,
                      const char* msg) {
  printf("%s\n", msg);
}

int main(const int argc, const char **argv) {
  ash_unit_test_context_t ctx;
  ash_log_t log = {
    NULL, "unit-tests", ASH_LOG_ALL, log_write
  };
  ctx.log = &log;
  ctx.tm_start = ctx.tm_end = ash_time_now();
  ctx.test_filter = argc <= 1 ? NULL : argv[1];
  run_tests(&ctx);
  ctx.tm_end = ash_time_now();
  return 0;
}
