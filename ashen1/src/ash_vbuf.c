/**
 */

#include "ash_vbuf.h"

static boolean_t ash_vbuf_reload(ash_vbuf_t *vbuf) {
  boolean_t read_rs;
  uint32_t reload_size;
  uint32_t read_size;

  if (vbuf->read_eof) {
    return ASH_TRUE;
  }

  reload_size = sizeof(vbuf->cache) - vbuf->read_pos;
  if (0 >= reload_size) {
    return ASH_TRUE;
  }

  read_size = 0;
  read_rs = vbuf->read(vbuf->context,
    vbuf->cache + vbuf->read_pos, reload_size, &read_size);
  if (!read_rs) {
    return ASH_FALSE;
  }
  if (read_size == 0) {
    vbuf->read_eof = ASH_TRUE;
  } else {
    vbuf->read_pos += read_size;
  }
  return ASH_TRUE;
}

static void ash_vbuf_dump(ash_vbuf_t *vbuf) {
  uint32_t remain_sz;
  uint32_t i;

  if (vbuf->write_pos >= vbuf->read_pos) {
    vbuf->write_pos = 0;
    vbuf->read_pos = 0;
  } else {
    remain_sz =  vbuf->read_pos - vbuf->write_pos;
    for (i = 0; i < remain_sz; ++i) {
      vbuf->cache[i] = vbuf->cache[vbuf->write_pos + i];
    }
    vbuf->write_pos = 0;
    vbuf->read_pos = remain_sz;
  }
}

static boolean_t ash_vbuf_readline_flush(ash_vbuf_t *vbuf, boolean_t *line_end) {
  uint32_t write_size = 0;

  *line_end = ASH_FALSE;
  for (;;) {
    if (vbuf->write_pos + write_size >= vbuf->read_pos) {
      break;
    }

    if ('\n' == vbuf->cache[vbuf->write_pos + write_size++]) {
      *line_end = ASH_TRUE;
      break;
    }
  }

  if (0 >= write_size) {
    ash_vbuf_dump(vbuf);
    return ASH_TRUE;
  }

  if (!vbuf->write(vbuf->context,
        vbuf->cache + vbuf->write_pos,
        *line_end ? write_size - 1 : write_size)) {
    return ASH_FALSE;
  }
  vbuf->write_pos += write_size;
  ash_vbuf_dump(vbuf);
  return ASH_TRUE;
}

boolean_t ash_vbuf_rdline(ash_vbuf_t *vbuf, boolean_t *finish) {
  boolean_t line_end;
  boolean_t rs;

  if (NULL == vbuf || NULL == finish
      || NULL == vbuf->read || NULL == vbuf->write) {
    return ASH_FALSE;
  }

  for (;;) {
    rs = ash_vbuf_reload(vbuf);
    if (!rs) {
      return ASH_FALSE;
    }

    rs = ash_vbuf_readline_flush(vbuf, &line_end);
    if (!rs) {
      return ASH_FALSE;
    }
    if (line_end) {
      break;
    }

    if (0 >= vbuf->read_pos && vbuf->read_eof) {
      break;
    }
  }
  if (NULL != vbuf->line_end) {
    vbuf->line_end(vbuf->context);
  }
  *finish = vbuf->read_eof && 0 >= vbuf->read_pos;
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

boolean_t ash_vbuf_rdline_wr_from_buf(void *context, char *buf,
                                      uint32_t buf_size) {
  ash_vbuf_str_ctx_t *buf_ctx = (ash_vbuf_str_ctx_t *)context;
  uint32_t remain_sz = buf_ctx->dist_buf_size - buf_ctx->dist_buf_pos - 1;
  uint32_t idx;
  uint32_t write_sz;
  char *dist;

  if (buf_size > remain_sz) {
    return ASH_FALSE;
  }
  dist = buf_ctx->dist_buf + buf_ctx->dist_buf_pos;
  for (idx = 0, write_sz = 0; idx < buf_size; ++idx) {
    if ('\r' == buf[idx]) {
      continue;
    }
    dist[write_sz++] = buf[idx];
  }
  buf_ctx->dist_buf_pos += write_sz;
  buf_ctx->dist_buf[buf_ctx->dist_buf_pos] = '\0';
  return ASH_TRUE;
}

void ash_vbuf_rdline_lineend_from_buf(void *context) {
  ash_vbuf_str_ctx_t *buf_ctx = (ash_vbuf_str_ctx_t *)context;
  buf_ctx->dist_buf_pos = 0;
}
