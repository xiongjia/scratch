/*
 */

#ifndef _TACHI_FS_H_
#define _TACHI_FS_H_ 1

#include <Windows.h>
#include "tachi-general.h"

#define TACHI_INVALID_FILE  INVALID_HANDLE_VALUE
typedef HANDLE tachi_fd_t;

#define tachi_file_close(_fd)  CloseHandle(_fd)

tachi_fd_t tachi_file_open(uchar_t *name, ulong_t mode, ulong_t create);


#endif /* !defined(_TACHI_FS_H_) */
