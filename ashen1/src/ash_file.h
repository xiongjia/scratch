/*
 */

#ifndef _ASH_FILE_H_
#define _ASH_FILE_H_ 1

#if defined(WIN32)
#include <Windows.h>
#endif /* defined(WIN32) */

#include "ash_types.h"

typedef HANDLE  ash_fd_t;

#define ash_close_file      CloseHandle
#define ASH_INVALID_FILE    INVALID_HANDLE_VALUE

#define ASH_FILE_RDONLY             GENERIC_READ
#define ASH_FILE_WRONLY             GENERIC_WRITE
#define ASH_FILE_RDWR               GENERIC_READ|GENERIC_WRITE
#define ASH_FILE_APPEND             FILE_APPEND_DATA|SYNCHRONIZE
#define ASH_FILE_NONBLOCK           0

#define ASH_FILE_CREATE_OR_OPEN     OPEN_ALWAYS
#define ASH_FILE_OPEN               OPEN_EXISTING
#define ASH_FILE_TRUNCATE           CREATE_ALWAYS

#define ASH_FILE_DEFAULT_ACCESS     0
#define ASH_FILE_OWNER_ACCESS       0

ash_fd_t ash_open_file(char_t *name, uint32_t mode,
                       uint32_t create, uint32_t access);

int32_t ash_read_fd(ash_fd_t fd, uchar_t *buf, int32_t buf_size);

const char *ash_file_get_filename(const char *full_filename);

void ash_file_write_stdout(const char *fmt, ...);

void ash_file_vwrite_stdout(const char *fmt, va_list ap);

void ash_file_write_stderr(const char *fmt, ...);

void ash_file_vwrite_stderr(const char *fmt, va_list ap);

#endif /* !defined(_ASH_FILE_H_) */
