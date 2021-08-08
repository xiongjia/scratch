/**
 */

#ifndef _ASH_HASH_H_
#define _ASH_HASH_H_ 1

#include "ash_general.h"
#include "ash_palloc.h"

typedef struct _ash_hash_t ash_hash_t;

ASH_DECLARE(unsigned long) ash_hash_djb(const char *key, size_t klen);

ASH_DECLARE(ash_hash_t *) ash_hash_create(ash_pool_t *pool);

ASH_DECLARE(void) ash_hash_set(ash_hash_t *hash,
  const char *key, size_t klen, const void *val);

ASH_DECLARE(const void *) ash_hash_get(ash_hash_t *hash,
  const char *key, size_t klen);

#endif /* !defined(_ASH_HASH_H_) */

