/**
 */

#include "ash_tests.h"
#include "ash_file.h"

void test_filename(void) {
  const char *src = __FILE__;
  char *name;
   
  name = ash_file_get_filename(src);
  printf("file name %s\n", name);
}

ash_unit_test_t unittest_filename = {
  "simple-filename",
  test_filename
};
