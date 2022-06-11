/*
 */

#include <string.h>
#include "ash_str.h"
#include "ash_vfmt.h"
#include "ash_vbuf.h"
#include "ash_mm.h"

static boolean_t ash_str_node_compare(ash_list_node_t *n1,
                                      ash_list_node_t *n2) {
  const char *s1;
  const char *s2;

  if (NULL == n1 && NULL == n2) {
    return ASH_TRUE;
  }
  if (NULL == n1 || NULL == n2) {
    return ASH_FALSE;
  }
  s1 = (const char*)(NULL == n1->element ? "" : n1->element);
  s2 = (const char*)(NULL == n2->element ? "" : n2->element);
  if (strcmp(s1, s2) == 0) {
    return ASH_TRUE;
  } else {
    return ASH_FALSE;
  }
}

char* ash_pstrdup(ash_pool_t *pool, const char *src) {
  size_t alloc_sz;
  char *dup;

  if (pool == NULL || src == NULL) {
    return NULL;
  }
  alloc_sz = strlen(src) + 1;
  dup = ash_pool_alloc(pool, alloc_sz);
  memcpy(dup, src, alloc_sz);
  return dup;
}

char *ash_pstrcat(ash_pool_t *pool, ...) {
  va_list args;
  size_t total_len = 0;
  size_t len;
  char *val;
  char *rt;
  char *cur;

  va_start(args, pool);
  while ((val = va_arg(args, char *)) != NULL) {
    total_len += strlen(val);
  }
  va_end(args);

  rt = (char*)ash_pool_alloc(pool, total_len + 1);
  if (NULL == rt) {
    return NULL;
  }

  cur = rt;
  va_start(args, pool);
  while ((val = va_arg(args, char *)) != NULL) {
    len = strlen(val);
    memcpy(cur, val, len);
    cur += len;
  }
  va_end(args);
  *cur = '\0';
  return rt;
}

ash_list_t *ash_slist_create(ash_pool_t *pool, ...) {
  ash_list_t *list;
  va_list args;

  va_start(args, pool);
  list = ash_slist_vcreate(pool, args);
  va_end(args);
  return list;
}

ash_list_t *ash_slist_vcreate(ash_pool_t *pool, va_list ap) {
  ash_list_t *list;
  char *val;

  list = ash_list_create(pool);
  if (NULL == list) {
    return NULL;
  }

  while ((val = va_arg(ap, char *)) != NULL) {
    if (!ash_list_push(list, ash_pstrdup(pool, val))) {
      return NULL;
    }
  }
  return list;
}

int32_t ash_vsnprintf(char *buf, size_t buf_sz,
                      const char *fmt, va_list ap) {
  ash_vfmt_buff_t vbuf;
  ASH_VFMT_BUFF_INIT(&vbuf, buf, buf_sz)
  return ash_vformatter(&vbuf, fmt, ap);
}

int32_t ash_snprintf(char *buf, size_t buf_sz, const char *fmt, ...) {
  va_list args;
  int32_t res;

  va_start(args, fmt);
  res = ash_vsnprintf(buf, buf_sz, fmt, args);
  va_end(args);
  return res;
}

ash_list_t *ash_parse_lines(ash_pool_t *pool, const char *src,
                            uint32_t max_line_sz) {
  ash_vbuf_str_ctx_t ctx;
  ash_list_t *lines;
  ash_vbuf_t vbuf;
  boolean_t rs;
  boolean_t finish = ASH_FALSE;
  char *buf;


  if (NULL == pool || NULL == src || 0 >= max_line_sz) {
    return NULL;
  }
  buf = (char*)ash_pool_alloc(pool, max_line_sz + 1);
  if (NULL == buf) {
    return NULL;
  }
  lines = ash_list_create(pool);
  if (NULL == lines) {
    return NULL;
  }
  ASH_VBUF_STR_RD_LINE_INIT(&vbuf, &ctx, src, buf, max_line_sz);
  for (;;) {
    rs = ash_vbuf_rdline(&vbuf, &finish);
    if (!rs) {
      return NULL;
    }
    if (!ash_list_push(lines, ash_pstrdup(pool, buf))) {
      return NULL;
    }
    if (finish) {
      break;
    }
  }
  return lines;
}

boolean_t ash_slist_compare(const ash_list_t *l1, const ash_list_t *l2) {
  return ash_list_compare(l1, l2, ash_str_node_compare);
}

void ash_vbuf_rdline_lineend_from_buf(void *context) {
  ash_vbuf_str_ctx_t *buf_ctx = (ash_vbuf_str_ctx_t *)context;
  buf_ctx->dist_buf_pos = 0;
}

boolean_t ash_vbuf_rdline_wr_from_buf(void *context, char *buf,
                                      uint32_t buf_size) {
  ash_vbuf_str_ctx_t *buf_ctx = (ash_vbuf_str_ctx_t *)context;
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

boolean_t ash_vbuf_rdline_rd_from_buf(void *context,
                                      char *buf, uint32_t buf_size,
                                      uint32_t *read_sz) {
  ash_vbuf_str_ctx_t *buf_ctx = (ash_vbuf_str_ctx_t *)context;
  uint32_t rd_sz = 0;

  if (0 >= buf_size) {
    return ASH_FALSE;
  }

  for (;;) {
    if ('\0' == buf_ctx->src[buf_ctx->pos]) {
      break;
    }

    buf[rd_sz++] = buf_ctx->src[buf_ctx->pos++];
    if (rd_sz >= buf_size) {
      break;
    }
  }
  *read_sz = rd_sz;
  return ASH_TRUE;
}
