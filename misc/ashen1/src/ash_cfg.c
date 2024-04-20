/*
 */

#include "ash_cfg.h"

#define ASH_CFG_MAX_LINE  1024

static void ash_cfg_parse(ash_cfg_t *cfg, const char *buf) {
  const char *pos;
  const char *hdr;

  for (pos = hdr = buf; '\0' != *pos; ++pos) {
    if ('\n' != *pos) {
      continue;
    }
  }

}

ash_cfg_t *ash_cfg_load_from_buff(ash_pool_t *pool, const char *buf) {
  ash_cfg_t *cfg = ash_pool_alloc(pool, sizeof(ash_cfg_t));
  cfg->pool = pool;
  cfg->content = ash_hash_create(pool);
  ash_cfg_parse(cfg, buf);
  return cfg;
}
