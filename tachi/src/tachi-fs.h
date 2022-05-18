/*
 */

#ifndef _TACHI_FS_H_
#define _TACHI_FS_H_ 1

#include <Windows.h>
#include "tachi-general.h"

typedef HANDLE tachi_fd_t;

tachi_fd_t tachi_file_open(uchar_t *name);


#endif /* !defined(_TACHI_FS_H_) */
