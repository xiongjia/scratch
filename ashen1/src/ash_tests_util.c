/**
 */

#include <stdio.h>
#include <string.h>
#include <stdarg.h>

#include "ash_file.h"
#include "ash_str.h"
#include "ash_tests_util.h"

void ash_testutil_str_equal(ash_unit_test_case_context_t *tc,
                            const char *expected, const char *actual,
                            const char *src, const int line) {
  if (expected == NULL && actual == NULL) {
    return;
  }

  if (expected != NULL && actual != NULL && strcmp(expected, actual) == 0) {
    return;
  }

  tc->failed = ASH_TRUE;
  if (!tc->verbose) {
    return;
  }
  ash_file_write_stderr("[%s:%d]: expected <%s>, but saw <%s>\n",
    ash_file_get_filename(src), line,  expected, actual);
}

void ash_testutil_size_equal(ash_unit_test_case_context_t *tc,
                             const size_t expected, const size_t actual,
                             const char *src, const int line) {
  if (expected == actual) {
    return;
  }

  tc->failed = ASH_TRUE;
  if (!tc->verbose) {
    return;
  }
  ash_file_write_stderr("[%s:%d]: expected <%llu>, but saw <%llu>\n",
    ash_file_get_filename(src), line, expected, actual);
}

void ash_testutil_int32_equal(ash_unit_test_case_context_t *tc,
                              const int32_t expected, const int32_t actual,
                              const char *src, const int line) {
  if (expected == actual) {
    return;
  }

  tc->failed = ASH_TRUE;
  if (!tc->verbose) {
    return;
  }
  ash_file_write_stderr("[%s:%d]: expected <%d>, but saw <%d>\n",
    ash_file_get_filename(src), line, expected, actual);
}

void ash_testutil_slist_equal(ash_unit_test_case_context_t *tc,
                              const ash_list_t *expected,
                              const ash_list_t *actual,
                              const char *src, const int line) {
  if (ash_slist_compare(actual, expected)) {
    return;
  }
  tc->failed = ASH_TRUE;
  ash_file_write_stderr("[%s:%d]: string lists are not equal\n",
    ash_file_get_filename(src), line);
}
