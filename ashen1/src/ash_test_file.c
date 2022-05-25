/**
 */

#include "ash_tests.h"
#include "ash_file.h"

#include <stdio.h>

static void test_filename(ash_unit_test_context_t* ctx) {
  const char *src = __FILE__;
  const char *name;
   
  name = ash_file_get_filename(src);
  printf("file name %s\n", name);
}

ash_unit_test_case_t unittest_filename = {
  "simple-filename",
  test_filename
};
