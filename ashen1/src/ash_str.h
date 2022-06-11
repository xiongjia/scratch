/**
 */

#ifndef _ASH_STR_H_
#define _ASH_STR_H_ 1

#include <stdarg.h>

#include "ash_types.h"
#include "ash_mm.h"
#include "ash_list.h"
#include "ash_vbuf.h"

typedef struct _ash_vbuf_str_ctx_s {
  const char *src;
  uint32_t pos;

  char *dist_buf;
  uint32_t dist_buf_size;
  uint32_t dist_buf_pos;
} ash_vbuf_str_ctx_t;

char *ash_pstrdup(ash_pool_t *pool, const char *src);

char *ash_pstrcat(ash_pool_t *pool, ...);

ash_list_t *ash_slist_create(ash_pool_t *pool, ...);

ash_list_t *ash_slist_vcreate(ash_pool_t *pool, va_list ap);

boolean_t ash_slist_compare(const ash_list_t *l1, const ash_list_t *l2);

int32_t ash_vsnprintf(char *buf, size_t buf_sz, const char *fmt, va_list ap);

int32_t ash_snprintf(char *buf, size_t buf_sz, const char *fmt, ...);

ash_list_t *ash_parse_lines(ash_pool_t *pool, const char *src,
                            uint32_t max_line_sz);

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

#endif /* !defined(_ASH_STR_H_) */
