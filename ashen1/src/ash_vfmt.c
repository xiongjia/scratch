/*
 */

#include "ash_vfmt.h"

#define _VFMT_INT32_LEN   (sizeof("-2147483648") - 1)
#define _VFMT_INT64_LEN   (sizeof("-9223372036854775808") - 1)
#define _VFMT_MAX_UINT32_VALUE  (uint32_t) 0xffffffff
#define _VFMT_MAX_INT32_VALUE   (uint32_t) 0x7fffffff

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

static void ash_vfmt_append_num(ash_vfmt_buff_t *vbuf, uint64_t num,
                                uint32_t hexadecimal, uchar_t zero, uint32_t width) {
  static char_t hex_l[] = "0123456789abcdef";
  static char_t hex_u[] = "0123456789ABCDEF";
  uint32_t ui32;
  uint32_t len;
  char tmp_buf[_VFMT_INT64_LEN + 1] = {0};
  char *pos;

  pos = tmp_buf + _VFMT_INT64_LEN;
  if (0 == hexadecimal) {
    if (num <= (uint64_t)_VFMT_MAX_UINT32_VALUE) {
      ui32 = (uint32_t)num;
      do {
        *--pos = (uchar_t) (ui32 % 10 + '0');
      } while (ui32 /= 10);
    } else {
      do {
        *--pos = (uchar_t) (num % 10 + '0');
      } while (num /= 10);
    }
  } else if (hexadecimal == 1) {
    do {
      *--pos = hex_l[(uint32_t) (num & 0xf)];
    } while (num >>= 4);
  } else if (hexadecimal == 2) {
    do {
      *--pos = hex_u[(uint32_t) (num & 0xf)];
    } while (num >>= 4);
  }
  /* zero or space padding */
  len = tmp_buf + _VFMT_INT64_LEN - pos;
  while (len++ < width) {
    ash_vfmt_append_char(vbuf, zero);
  }

  for (; pos != tmp_buf + _VFMT_INT64_LEN; ++pos) {
    ash_vfmt_append_char(vbuf, *pos);
  }
}

size_t ash_vformatter(ash_vfmt_buff_t *vbuf, const char *fmt, va_list ap) {
  uchar_t zero;
  uint32_t width;
  char ch;

  while (ch = *fmt++) {
    if ('%' != ch) {
      ash_vfmt_append_char(vbuf, ch);
      continue;
    }

    ch = *fmt++;
    _VFMT_IS_FMT_END(ch);

    zero = ch == '0' ? '0' : ' ';
    width = 0;
    for (; ch >= '0' && ch <= '9'; ch = *fmt++) {
      width = width * 10 + (ch - '0');
    }
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
      ash_vfmt_append_num(vbuf, va_arg(ap, uint64_t), 0, zero, width);
      break;
    case 'h':
      ash_vfmt_append_num(vbuf, va_arg(ap, uint64_t), 1, zero, width);
      break;
    case 'H':
      ash_vfmt_append_num(vbuf, va_arg(ap, uint64_t), 2, zero, width);
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
