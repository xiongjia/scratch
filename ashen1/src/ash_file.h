/*
 */

#ifndef _ASH_FILE_H_
#define _ASH_FILE_H_ 1

#if defined(WIN32)
#include <Windows.h>
#endif /* defined(WIN32) */

#include "ash_types.h"
#include "ash_vbuf.h"

typedef HANDLE ash_fd_t;

#define ASH_FILE_BUFF_SIZE  512

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

ash_fd_t ash_open_file(const char_t *name, uint32_t mode,
                       uint32_t create, uint32_t access);

int32_t ash_read_fd(ash_fd_t fd, uchar_t *buf, int32_t buf_size);

int32_t ash_write_fd(ash_fd_t fd, uchar_t *buf, int32_t buf_size);

const char *ash_file_get_filename(const char *full_filename);

void ash_file_write_stdout(const char *fmt, ...);

void ash_file_vwrite_stdout(const char *fmt, va_list ap);

void ash_file_write_stderr(const char *fmt, ...);

void ash_file_vwrite_stderr(const char *fmt, va_list ap);

typedef struct _ash_vbuf_fd_ctx_s {
  ash_fd_t fd;

  char *dist_buf;
  uint32_t dist_buf_size;
  uint32_t dist_buf_pos;
} ash_vbuf_fd_ctx_t;

boolean_t ash_vbuf_rdline_rd_from_fd(void *context,
                                     char *buf, uint32_t buf_size,
                                     uint32_t *read_sz);

boolean_t ash_vbuf_rdline_wr_from_fd(void *context, char *buf,
                                     uint32_t buf_size);

void ash_vbuf_rdline_lineend_from_fd(void *context);

#define ASH_VBUF_FD_RD_LINE_INIT(_vbuf, _ctx, _fd, _buf, _buf_sz) \
  ((ash_vbuf_fd_ctx_t*)_ctx)->fd = _fd; \
  ((ash_vbuf_fd_ctx_t*)_ctx)->dist_buf = _buf; \
  ((ash_vbuf_fd_ctx_t*)_ctx)->dist_buf_size = _buf_sz; \
  ((ash_vbuf_fd_ctx_t*)_ctx)->dist_buf_pos = 0; \
  ASH_VBUF_INIT(_vbuf, _ctx, \
    ash_vbuf_rdline_rd_from_fd, \
    ash_vbuf_rdline_wr_from_fd, \
    ash_vbuf_rdline_lineend_from_fd);


#endif /* !defined(_ASH_FILE_H_) */
