/**
 */

#ifndef _ASH_VFMT_H_
#define _ASH_VFMT_H_ 1

#include <stdarg.h>
#include "ash_types.h"

typedef struct _ash_vfmt_buff_s {
  size_t count;
  size_t buf_sz;
  char *buf;
  char *cur;
} ash_vfmt_buff_t;

#define ASH_VFMT_BUFF_INIT(_vbuf, _buf, _buf_sz) \
  ((ash_vfmt_buff_t*)_vbuf)->count = 0; \
  ((ash_vfmt_buff_t*)_vbuf)->buf_sz = _buf_sz; \
  ((ash_vfmt_buff_t*)_vbuf)->buf = _buf; \
  ((ash_vfmt_buff_t*)_vbuf)->cur = _buf;

size_t ash_vformatter(ash_vfmt_buff_t *vbuf,
                      const char *fmt, va_list ap);

#endif /* !defined(_ASH_VFMT_H_) */
