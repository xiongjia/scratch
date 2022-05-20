/*
 */

#include "ash_mm.h"

#define ASH_BUCKET_SIZE  (1024 * 4)

typedef struct _ash_bucket_s ash_bucket_t;

struct _ash_bucket_s {
  uchar_t *current;
  uchar_t *last;
  ash_bucket_t *next;
};

struct _ash_pool_s {
  ash_bucket_t *init;
  ash_bucket_t *buckets;
};

ash_pool_t* ash_pool_create(size_t size) {
  const size_t req_sz = size + sizeof(ash_pool_t) + sizeof(ash_bucket_t);
  size_t allcate_sz; 
  ash_pool_t *pool;

  allcate_sz = req_sz < ASH_BUCKET_SIZE ? ASH_BUCKET_SIZE : req_sz;
  pool = ash_malloc(allcate_sz);
  if (pool == NULL) {
    return NULL;
  }
  pool->buckets = NULL;
  pool->init = (ash_bucket_t*)((uchar_t*)pool + sizeof(ash_pool_t));
  pool->init->next = NULL;
  pool->init->last = (uchar_t*)pool + allcate_sz;
  pool->init->current = (uchar_t*)pool + sizeof(ash_pool_t) + sizeof(ash_bucket_t);
  return  pool;
}

void ash_pool_destroy(ash_pool_t *pool) {
  ash_bucket_t *bucket;

  if (pool == NULL) {
    return;
  }
  for (bucket = pool->buckets; bucket != NULL;) {
    pool->buckets = bucket->next;
    ash_free(bucket);
    bucket = pool->buckets;
  }
  ash_free(pool);
}

void* ash_pool_alloc(ash_pool_t* pool, size_t size) {
  ash_bucket_t *bucket;
  size_t free_sz;
  size_t req_sz;
  size_t allcate_sz;
  uchar_t* mm;

  if (pool == NULL || pool->init == NULL) {
    return NULL;
  }

  free_sz = pool->init->last - pool->init->current;
  if (free_sz >= size) {
    mm = pool->init->current;
    pool->init->current += size;
    return mm;
  }
  
  for (bucket = pool->buckets; bucket != NULL; bucket = bucket->next) {
    free_sz = bucket->last - bucket->current;
    if (free_sz < size) {
      continue;
    }
    mm = bucket->current;
    bucket->current += size;
    return mm;
  }

  req_sz = size + sizeof(ash_bucket_t);
  allcate_sz = req_sz < ASH_BUCKET_SIZE ? ASH_BUCKET_SIZE : req_sz;
  bucket = ash_malloc(allcate_sz);
  if (bucket == NULL) {
    return NULL;
  }
  mm = (uchar_t*)bucket + sizeof(ash_bucket_t);
  bucket->current = (uchar_t*)bucket + sizeof(ash_bucket_t) + size;
  bucket->last = (uchar_t*)bucket + allcate_sz;
  bucket->next = pool->buckets;
  pool->buckets = bucket;
  return mm;
}
