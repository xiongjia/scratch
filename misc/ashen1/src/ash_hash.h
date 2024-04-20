/**
 */

#ifndef _ASH_HASH_H_
#define _ASH_HASH_H_ 1

#include "ash_types.h"
#include "ash_mm.h"

typedef struct _ash_hash_s ash_hash_t;

uint32_t ash_hash_dbj(const char *key, uint32_t klen, uint32_t hash);

uint32_t ash_hash_dbj_default(const char *key, uint32_t klen);

ash_hash_t *ash_hash_create(ash_pool_t *pool);

void ash_hash_set(ash_hash_t *hash,
                  const uchar_t *key, uint32_t klen,
                  const void *val);

const void *ash_hash_get(ash_hash_t *hash,
                         const uchar_t *key, uint32_t klen);

#endif /* !defined(_ASH_HASH_H_) */
