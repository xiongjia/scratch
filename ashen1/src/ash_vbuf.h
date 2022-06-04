/**
 */

#ifndef _ASH_VBUF_H_
#define _ASH_VBUF_H_ 1

#include "ash_types.h"

#define _ASH_VBUF_CACHE_SIZE 128

typedef struct _ash_vbuf_s {
  boolean_t failed;

  void *context;
  char cache[_ASH_VBUF_CACHE_SIZE];

  uint32_t read_pos;
  uint32_t write_pos;

  boolean_t read_eof;
  boolean_t (*read)(void *context,
                    char *buf, uint32_t buf_size, uint32_t *read_sz);
  boolean_t (*write)(void *context, char *buf, uint32_t buf_size);
  void (*line_end)(void *context);
} ash_vbuf_t;

typedef struct _ash_vbuf_str_ctx_s {
  const char *src;
  uint32_t pos;

  char *dist_buf;
  uint32_t dist_buf_size;
  uint32_t dist_buf_pos;
} ash_vbuf_str_ctx_t;

#define ASH_VBUF_INIT(_vbuf, _ctx, _read, _write, _line_end) \
    ((ash_vbuf_t*)_vbuf)->failed = ASH_FALSE; \
    ((ash_vbuf_t*)_vbuf)->context = _ctx; \
    ((ash_vbuf_t*)_vbuf)->read = _read; \
    ((ash_vbuf_t*)_vbuf)->read_eof = ASH_FALSE; \
    ((ash_vbuf_t*)_vbuf)->write = _write; \
    ((ash_vbuf_t*)_vbuf)->line_end = _line_end; \
    ((ash_vbuf_t*)_vbuf)->read_pos = 0; \
    ((ash_vbuf_t*)_vbuf)->write_pos = 0;

boolean_t ash_vbuf_rdline_rd_from_buf(void *context,
                                      char *buf, uint32_t buf_size,
                                      uint32_t *read_sz);

boolean_t ash_vbuf_rdline_wr_from_buf(void *context, char *buf,
                                      uint32_t buf_size);

void ash_vbuf_rdline_lineend_from_buf(void *context);

#define ASH_VBUF_STR_RD_LINE_INIT(_vbuf, _ctx, _src, _buf, _buf_sz) \
  ((ash_vbuf_str_ctx_t*)_ctx)->src = (NULL == _src) ? "" : _src; \
  ((ash_vbuf_str_ctx_t*)_ctx)->pos = 0; \
  ((ash_vbuf_str_ctx_t*)_ctx)->dist_buf = _buf; \
  ((ash_vbuf_str_ctx_t*)_ctx)->dist_buf_size = _buf_sz; \
  ((ash_vbuf_str_ctx_t*)_ctx)->dist_buf_pos = 0; \
  ASH_VBUF_INIT(_vbuf, _ctx, \
    ash_vbuf_rdline_rd_from_buf, \
    ash_vbuf_rdline_wr_from_buf, \
    ash_vbuf_rdline_lineend_from_buf);

boolean_t ash_vbuf_rdline(ash_vbuf_t *vbuf, boolean_t *finish);

#endif /* !defined(_ASH_VBUF_H_) */
