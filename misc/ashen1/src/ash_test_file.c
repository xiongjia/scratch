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
  ash_vbuf_fd_ctx_t vbuf_ctx;
  ash_vbuf_t vbuf;
  boolean_t rt;
  boolean_t finish = ASH_FALSE;
  ash_fd_t fd;
  uchar_t buf[1024] = {0};

  fd = ash_open_file(test_file,
    ASH_FILE_RDONLY, ASH_FILE_OPEN, ASH_FILE_DEFAULT_ACCESS);
  if (ASH_INVALID_FILE == fd) {
    return;
  }

  ASH_VBUF_FD_RD_LINE_INIT(&vbuf, &vbuf_ctx, fd, buf, sizeof(buf));
  for (;;) {
    rt = ash_vbuf_rdline(&vbuf, &finish);
    if (!rt) {
      break;
    }
    ash_file_write_stdout("%s\n", buf);
    if (finish) {
      break;
    }
  }
  ash_close_file(fd);
}

ash_unit_test_case_t unittest_filename = {
  "simple-filename",
  test_filename
};

ash_unit_test_case_t unittest_file_io = {
  "simple-fileio",
  test_file_io
};
