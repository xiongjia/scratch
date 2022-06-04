/**
 */

#ifndef _ASH_CFG_H_
#define _ASH_CFG_H_ 1

#include "ash_mm.h"
#include "ash_hash.h"

typedef struct _ash_cfg_s {
  ash_pool_t *pool;
  ash_hash_t *content;
} ash_cfg_t;

ash_cfg_t *ash_cfg_load_from_buff(ash_pool_t *pool, const char *buf);

#endif /* !defined(_ASH_CFG_H_) */
