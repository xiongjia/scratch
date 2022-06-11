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

boolean_t ash_vbuf_rdline_rd_from_fd(void *context,
                                     char *buf, uint32_t buf_size,
                                     uint32_t *read_sz) {
  ash_vbuf_fd_ctx_t *buf_ctx = (ash_vbuf_fd_ctx_t *)context;
  int32_t rd_sz;

  if (0 >= buf_size) {
    return ASH_FALSE;
  }

  rd_sz = ash_read_fd(buf_ctx->fd, buf, buf_size);
  if (0 > rd_sz) {
    return ASH_FALSE;
  }
  *read_sz = rd_sz;
  return ASH_TRUE;
}

boolean_t ash_vbuf_rdline_wr_from_fd(void *context, char *buf,
                                     uint32_t buf_size) {
  ash_vbuf_fd_ctx_t *buf_ctx = (ash_vbuf_fd_ctx_t *)context;
  ash_vbuf_dist_t dist_buf;
  dist_buf.dist_buf = buf_ctx->dist_buf;
  dist_buf.dist_buf_size = buf_ctx->dist_buf_size;
  dist_buf.dist_buf_pos = buf_ctx->dist_buf_pos;
  if (!ash_vbuf_writ_dist(&dist_buf, buf, buf_size)) {
    return ASH_FALSE;
  }
  buf_ctx->dist_buf_pos = dist_buf.dist_buf_pos;
  return ASH_TRUE;
}

void ash_vbuf_rdline_lineend_from_fd(void *context) {
  ash_vbuf_fd_ctx_t *buf_ctx = (ash_vbuf_fd_ctx_t *)context;
  buf_ctx->dist_buf_pos = 0;
}
