/**
 */

#include "tachi-fs.h"

tachi_fd_t tachi_file_open(uchar_t *name, ulong_t mode, ulong_t create) {
  HANDLE fd;

  if (tachi_str_is_empty(name)) {
    return TACHI_INVALID_FILE;
  }
  fd = CreateFileA(name, mode,
    FILE_SHARE_READ | FILE_SHARE_WRITE | FILE_SHARE_DELETE,
    NULL, create, FILE_FLAG_BACKUP_SEMANTICS, NULL);
  return fd;
}
