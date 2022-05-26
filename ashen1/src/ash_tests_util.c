/**
 */

#include <stdio.h>
#include <string.h>
#include "ash_file.h"
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
  fprintf(stderr, "[%s:%d]: expected <%s>, but saw <%s>\n",
    ash_file_get_filename(src),  line,
    expected == NULL ? "<null>" : expected,
    actual == NULL ? "<null>" : actual);
  fflush(stderr);
}
