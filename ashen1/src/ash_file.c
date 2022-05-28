/*
 */

#include <Windows.h>
#include <stdarg.h>
#include <string.h>

#include "ash_vfmt.h"
#include "ash_file.h"

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
  WriteFile(ctx->fd, buf, (DWORD)count, NULL, NULL);
  return ASH_TRUE;
}

const char* ash_file_get_filename(const char* full_filename) {
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

void ash_file_vwrite_stdout(const char *fmt, va_list ap) {
  ash_vbuf_write_ctx_t ctx;
  ash_vfmt_buff_t vbuf;
  char buf[128];

  ctx.fd = GetStdHandle(STD_OUTPUT_HANDLE);
  ASH_VFMT_BUFF_INIT_WITH_FLUSH(&vbuf, buf, sizeof(buf), &ctx, ash_vfmt_flush);
  ash_vformatter(&vbuf, fmt, ap);
}

void ash_file_write_stderr(const char *fmt, ...) {
  va_list args;
  va_start(args, fmt);
  ash_file_vwrite_stderr(fmt, args);
  va_end(args);
}

void ash_file_vwrite_stderr(const char *fmt, va_list ap)  {
  ash_vbuf_write_ctx_t ctx;
  ash_vfmt_buff_t vbuf;
  char buf[128];
  ctx.fd = GetStdHandle(STD_ERROR_HANDLE);
  ASH_VFMT_BUFF_INIT_WITH_FLUSH(&vbuf, buf, sizeof(buf), &ctx, ash_vfmt_flush);
  ash_vformatter(&vbuf, fmt, ap);
}
