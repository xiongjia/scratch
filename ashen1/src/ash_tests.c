/**
 */

#include <stdio.h>
#include <string.h>

#include "ash_tests.h"
#include "ash_time.h"
#include "ash_file.h"

static ash_unit_test_case_t *ALL_TESTS[] = {
  &unittest_simple_str,
  &unittest_simple_time,
  &unittest_filename
};

static void run_test(ash_unit_test_context_t *ctx,
                     ash_unit_test_case_t *unit_test) {
  ash_unit_test_case_context_t test_case = {
    .failed = ASH_FALSE,
    .verbose = ASH_TRUE,
    .pool = ash_pool_create_default()
  };

  ASH_LOG_WRITE(ctx->log, ASH_LOG_INFO,
    "Test case = %s", unit_test->test_name);
  unit_test->invoke(ctx, &test_case);
  ash_pool_destroy(test_case.pool);
}

static void run_tests(ash_unit_test_context_t *ctx) {
  ash_unit_test_case_t *unit_test;
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

static void log_write(ash_log_t *log, const char *src, size_t line,
                      time_t ts, const char *msg) {
  ash_file_write_stdout("%s\n", msg);
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
