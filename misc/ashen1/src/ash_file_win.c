/**
 */

#include <Windows.h>

#include "ash_file.h"
#include "ash_vfmt.h"

typedef struct _ash_vbuf_write_ctx_s {
  HANDLE fd;
} ash_vbuf_write_ctx_t;

static boolean_t ash_vfmt_flush(void *context,
                                char *buf, size_t count,
                                boolean_t finish) {
  ash_vbuf_write_ctx_t *ctx = (ash_vbuf_write_ctx_t*)context;
  if (NULL == ctx || NULL == buf || 0 == count) {
    return ASH_TRUE;
  }
  WriteFile(ctx->fd, buf,
    (DWORD)count, NULL, NULL);
  return ASH_TRUE;
}

void ash_file_vwrite_stdout(const char *fmt, va_list ap) {
  ash_vbuf_write_ctx_t ctx;
  ash_vfmt_buff_t vbuf;
  char buf[128];
  ctx.fd = GetStdHandle(STD_OUTPUT_HANDLE);
  ASH_VFMT_BUFF_INIT_WITH_FLUSH(&vbuf, buf, sizeof(buf), &ctx, ash_vfmt_flush);
  ash_vformatter(&vbuf, fmt, ap);
}

void ash_file_vwrite_stderr(const char *fmt, va_list ap)  {
  ash_vbuf_write_ctx_t ctx;
  ash_vfmt_buff_t vbuf;
  char buf[128];
  ctx.fd = GetStdHandle(STD_ERROR_HANDLE);
  ASH_VFMT_BUFF_INIT_WITH_FLUSH(&vbuf, buf, sizeof(buf), &ctx, ash_vfmt_flush);
  ash_vformatter(&vbuf, fmt, ap);
}

ash_fd_t ash_open_file(const char_t *name, uint32_t mode,
                       uint32_t create, uint32_t access) {
  return CreateFileA(name, mode,
    FILE_SHARE_READ|FILE_SHARE_WRITE|FILE_SHARE_DELETE,
    NULL, create,
    FILE_FLAG_BACKUP_SEMANTICS, NULL);
}

int32_t ash_read_fd(ash_fd_t fd, uchar_t *buf, int32_t buf_size) {
  DWORD rd_sz = 0;
  BOOL rt = ReadFile(fd, buf, buf_size, &rd_sz, NULL);
  return rt ? rd_sz : -1;
}

int32_t ash_write_fd(ash_fd_t fd, uchar_t *buf, int32_t buf_size) {
  DWORD wr_sz = 0;
  BOOL rt = WriteFile(fd, buf, buf_size, &wr_sz, NULL);
  return rt ? wr_sz : -1;
}
