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

static void test_file_io(ash_unit_test_context_t* ctx,
                         ash_unit_test_case_context_t *tc) {
  const char *test_file = "c:/wrk/tmp/test1.prop";
  ash_fd_t fd;
  uchar_t buf[1024] = {0};
  int32_t rd_sz;

  fd = ash_open_file(test_file,
    ASH_FILE_RDONLY, ASH_FILE_OPEN, ASH_FILE_DEFAULT_ACCESS);
  if (ASH_INVALID_FILE == fd) {
    return;
  }

   rd_sz = ash_read_fd(fd, buf, sizeof(buf));
}

ash_unit_test_case_t unittest_filename = {
  "simple-filename",
  test_filename
};

ash_unit_test_case_t unittest_file_io = {
  "simple-fileio",
  test_file_io
};
