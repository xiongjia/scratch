/*
 */

#include "ash_vfmt.h"

#define _VFMT_IS_FMT_END(_c) if ('\0' == _c) { break; }

static void ash_vfmt_append_char(ash_vfmt_buff_t *vbuf, char ch) {
  if (NULL == vbuf) {
    return;
  }
  vbuf->count++;
  if (NULL == vbuf->buf || 0 == vbuf->buf_sz || vbuf->cur == NULL) {
    return;
  }
  if ((size_t)(vbuf->cur - vbuf->buf) > vbuf->buf_sz) {
    return;
  }
  *(vbuf->cur) = ch;
  vbuf->cur++;
}

static void ash_vfmt_append_str(ash_vfmt_buff_t *vbuf, char *val) {
  char *pos;
  if (val == NULL || *val == '\0') {
    return;
  }
  for (pos = val; *pos != '\0'; pos++) {
    ash_vfmt_append_char(vbuf, *pos);
  }
}

size_t ash_vformatter(ash_vfmt_buff_t *vbuf, const char *fmt, va_list ap) {
  char ch;

  while (ch = *fmt++) {
    if ('%' != ch) {
      ash_vfmt_append_char(vbuf, ch);
      continue;
    }

    ch = *fmt++;
    _VFMT_IS_FMT_END(ch);

    switch (ch) {
    case '%':
      ash_vfmt_append_char(vbuf, '%');
      break;
    case 'c':
      ash_vfmt_append_char(vbuf, (char)va_arg(ap, int));
      break;
    case 's':
      ash_vfmt_append_str(vbuf, va_arg(ap, uchar_t*));
      break;
    case 'd':

      break;
    default:
      break;
    }
  }

  ash_vfmt_append_char(vbuf, '\0');
  if (NULL == vbuf->buf || 0 == vbuf->buf_sz) {
    return vbuf->count;
  } else {
    return vbuf->count > vbuf->buf_sz ? (size_t)-1 : vbuf->count;
  }
}
