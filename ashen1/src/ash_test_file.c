/**
 */

#include "ash_tests.h"
#include "ash_file.h"

#include <stdio.h>

static void test_filename(ash_unit_test_context_t* ctx,
                          ash_unit_test_case_context_t *tc) {
  const char *name;
   
  name = ash_file_get_filename("c:\\test\\test1.c");
  ASHTU_STR_EQU(tc, "test1.c", name);
  name = ash_file_get_filename("/test/test1.c");
  ASHTU_STR_EQU(tc, "test1.c", name);
  name = ash_file_get_filename("test1.c");
  ASHTU_STR_EQU(tc, "test1.c", name);
}

ash_unit_test_case_t unittest_filename = {
  "simple-filename",
  test_filename
};
