/**
 */

#include "ash_hash.h"

typedef struct _ash_hash_entry_t ash_hash_entry_t;

#define ASH_HASH_MAX_BUCKET_SIZE 64

struct _ash_hash_entry_t {
  char *key;
  size_t klen;

  const void *val;
  ash_hash_entry_t *next;
};

struct _ash_hash_t {
  ash_pool_t *pool;

  unsigned long buckets_len;
  ash_hash_entry_t **bucket;
};

static unsigned long ash_hash_make_key(ash_hash_t *hash,
                                       const char *key, size_t klen) {
  return ash_hash_djb(key, klen) % hash->buckets_len;
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

ASH_DECLARE(unsigned long) ash_hash_djb(const char *key, size_t klen) {
  unsigned int hash = 5381;
  size_t i = 0;

  /* An algorithm produced by Professor Daniel J. Bernstein and shown 
   * first to the world on the usenet newsgroup comp.lang.c. It is one of 
   * the most efficient hash functions ever published.
   */
  for (i = 0; i < klen; ++i) {
    hash = ((hash << 5) + hash) + key[i];
  }
  return hash;
}

ASH_DECLARE(ash_hash_t *) ash_hash_create(ash_pool_t *pool) {
  ash_hash_t *hash;

  if (pool == NULL) {
    return NULL;
  }
  hash = ash_pool_alloc(pool, sizeof(ash_hash_t));
  hash->pool = pool;
  hash->buckets_len = ASH_HASH_MAX_BUCKET_SIZE;
  hash->bucket = ash_pool_calloc(pool,
    sizeof(ash_hash_entry_t *) * hash->buckets_len);
  return hash;
}

ASH_DECLARE(void) ash_hash_set(ash_hash_t *hash,
                               const char *key, size_t klen,
                               const void *val) {
  ash_hash_entry_t *bucket;
  unsigned long khash;
  char *bucket_key;

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

ASH_DECLARE(const void *) ash_hash_get(ash_hash_t *hash,
                                       const char *key, size_t klen) {
  ash_hash_entry_t *bucket;
  unsigned long khash;

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


