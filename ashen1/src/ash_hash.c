/**
 */

#include <memory.h>
#include "ash_hash.h"

#define ASH_HASH_MAX_BUCKET_SIZE 64

typedef struct _ash_hash_entry_s ash_hash_entry_t;

struct _ash_hash_entry_s {
  uchar_t *key;
  size_t klen;
  const void *val;

  ash_hash_entry_t *next;
};

struct _ash_hash_s {
  ash_pool_t *pool;
  uint32_t buckets_len;
  ash_hash_entry_t **bucket;
};

static uint32_t ash_hash_make_key(ash_hash_t *hash,
                                  const char *key, uint32_t klen) {
  return ash_hash_dbj_default(key, klen) % hash->buckets_len;
}

static ash_hash_entry_t* ash_hash_find_entry(ash_hash_t *hash,
                                             const unsigned long khash,
                                             const char *key, size_t klen) {
  ash_hash_entry_t *bucket;
  for (bucket = hash->bucket[khash]; bucket != NULL; bucket = bucket->next) {
    if (bucket->klen == klen && (memcmp(key, bucket->key, klen) == 0)) {
      return bucket;
    }
  }
  return NULL;
}

uint32_t ash_hash_dbj(const char *key, uint32_t klen, uint32_t hash) {
  uint32_t hash_val = hash;
  uint32_t i;

  /* An algorithm produced by Professor Daniel J. Bernstein and shown 
   * first to the world on the usenet newsgroup comp.lang.c. It is one of 
   * the most efficient hash functions ever published.
   */
  for (i = 0; i < klen; i++) {
    hash = ((hash << 5) + hash) + (uint32_t)key[i];
  }
  return hash; 
}

uint32_t ash_hash_dbj_default(const char *key, uint32_t klen) {
  return ash_hash_dbj(key, klen, 5381);
}

ash_hash_t *ash_hash_create(ash_pool_t *pool) {
  ash_hash_t *hash;

  if (NULL == pool) {
    return NULL;
  }
  hash = ash_pool_alloc(pool, sizeof(ash_hash_t));
  hash->pool = pool;
  hash->buckets_len = ASH_HASH_MAX_BUCKET_SIZE;
  hash->bucket = ash_pool_calloc(pool,
    sizeof(ash_hash_entry_t *) * hash->buckets_len);
  return hash;
}

void ash_hash_set(ash_hash_t *hash,
                  const uchar_t *key, uint32_t klen,
                  const void *val) {
  ash_hash_entry_t *bucket;
  uint32_t khash;
  uchar_t *bucket_key;

  if (hash == NULL || key == NULL || klen == 0) {
    return;
  }

  khash = ash_hash_make_key(hash, key, klen);
  bucket = ash_hash_find_entry(hash, khash, key, klen);
  if (bucket != NULL) {
    bucket->val = val;
    return;
  }

  bucket = ash_pool_alloc(hash->pool, sizeof(ash_hash_entry_t));
  bucket_key = ash_pool_calloc(hash->pool, klen);
  if (bucket == NULL || bucket_key == NULL) {
    return;
  }
  memcpy(bucket_key, key, klen);
  bucket->key = bucket_key;
  bucket->klen = klen;
  bucket->val = val;
  bucket->next = hash->bucket[khash];
  hash->bucket[khash] = bucket;
}

const void *ash_hash_get(ash_hash_t *hash,
                         const uchar_t *key, uint32_t klen) {
  ash_hash_entry_t *bucket;
  uint32_t khash;

  if (hash == NULL || key == NULL || klen == 0) {
    return NULL;
  }
  khash = ash_hash_make_key(hash, key, klen);
  bucket = ash_hash_find_entry(hash, khash, key, klen);
  if (bucket == NULL) {
    return NULL;
  }
  return bucket->val;
}
