/**
 */

#ifndef _ASH_VBUF_H_
#define _ASH_VBUF_H_ 1

#include "ash_types.h"

#define _ASH_VBUF_CACHE_SIZE 128

typedef struct _ash_vbuf_dist_s {
  char *dist_buf;
  uint32_t dist_buf_size;
  uint32_t dist_buf_pos;
} ash_vbuf_dist_t;

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

#define ASH_VBUF_INIT(_vbuf, _ctx, _read, _write, _line_end) \
    ((ash_vbuf_t*)_vbuf)->failed = ASH_FALSE; \
    ((ash_vbuf_t*)_vbuf)->context = _ctx; \
    ((ash_vbuf_t*)_vbuf)->read = _read; \
    ((ash_vbuf_t*)_vbuf)->read_eof = ASH_FALSE; \
    ((ash_vbuf_t*)_vbuf)->write = _write; \
    ((ash_vbuf_t*)_vbuf)->line_end = _line_end; \
    ((ash_vbuf_t*)_vbuf)->read_pos = 0; \
    ((ash_vbuf_t*)_vbuf)->write_pos = 0;

boolean_t ash_vbuf_rdline(ash_vbuf_t *vbuf, boolean_t *finish);

boolean_t ash_vbuf_writ_dist(ash_vbuf_dist_t *dist_buf, char *buf,
                             uint32_t buf_size);

#endif /* !defined(_ASH_VBUF_H_) */
