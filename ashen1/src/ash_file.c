/*
 */

#include <stdarg.h>
#include <string.h>

#include "ash_file.h"

const char* ash_file_get_filename(const char *full_filename) {
  char *rs;

  if (full_filename == NULL || full_filename[0] == '\0') {
    return full_filename;
  }
  rs = strrchr(full_filename, '/');
  if (rs == NULL) {
    rs = strrchr(full_filename, '\\');
  }
  return rs == NULL ? full_filename : rs + 1;
}

void ash_file_write_stdout(const char *fmt, ...) {
  va_list args;
  va_start(args, fmt);
  ash_file_vwrite_stdout(fmt, args);
  va_end(args);
}

void ash_file_write_stderr(const char *fmt, ...) {
  va_list args;
  va_start(args, fmt);
  ash_file_vwrite_stderr(fmt, args);
  va_end(args);
}
