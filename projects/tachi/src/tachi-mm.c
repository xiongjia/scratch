/**
 * 
 */

#include "tachi-mm.h"

#define TACHI_BUCKET_SIZE  (1024 * 4)

typedef struct _tachi_bucket tachi_bucket_t;

typedef struct _tachi_bucket {
  uchar_t *hdr;
  uchar_t *last;
  uchar_t* current;
  tachi_bucket_t *next;
};

struct _tachi_pool {
  tachi_bucket_t *bucket;
};

tachi_pool_t* tachi_pool_create(size_t size) {
  const size_t req_sz = size + sizeof(tachi_pool_t) + sizeof(tachi_bucket_t);
  tachi_pool_t *pool;
  size_t allcate_sz;
  uchar_t *mm;

  allcate_sz = req_sz < TACHI_BUCKET_SIZE ? TACHI_BUCKET_SIZE : req_sz;
  mm = tachi_malloc(allcate_sz);
  if (mm == NULL) {
    return NULL;
  }
  pool = mm;
  pool->bucket = (tachi_bucket_t*)(mm + sizeof(tachi_pool_t));
  pool->bucket->current = mm + sizeof(tachi_pool_t) + sizeof(tachi_bucket_t);
  pool->bucket->hdr = pool->bucket->current;
  pool->bucket->last = mm + allcate_sz;
  pool->bucket->next = NULL;
  return pool;
}

void tachi_pool_destroy(tachi_pool_t *pool) {
  tachi_bucket_t *bucket;
  if (pool == NULL || pool->bucket == NULL) {
    return;
  }
  bucket = pool->bucket->next;
  for (; bucket != NULL;) {
    pool->bucket->next = bucket->next;
    tachi_free(bucket);
    bucket = pool->bucket->next;
  }
  tachi_free(pool);
}

void* tachi_pool_alloc(tachi_pool_t *pool, size_t size) {
  tachi_bucket_t *bucket;
  size_t free_sz;
  size_t req_sz;
  size_t allcate_sz;
  uchar_t *mm;

  if (pool == NULL || pool->bucket == NULL) {
    return NULL;
  }

  bucket = pool->bucket;
  for (; bucket != NULL; bucket = bucket->next) {
    free_sz = bucket->last - bucket->current;
    if (free_sz < size) {
      continue;
    }
    mm = bucket->current;
    bucket->current += size;
    return mm;
  }

  req_sz = size + sizeof(tachi_bucket_t);
  allcate_sz = req_sz < TACHI_BUCKET_SIZE ? TACHI_BUCKET_SIZE : req_sz;
  mm = tachi_malloc(allcate_sz);
  if (mm == NULL) {
    return NULL;
  }
  bucket = mm;
  bucket->current = mm + sizeof(tachi_bucket_t) + size;
  bucket->hdr = mm + sizeof(tachi_bucket_t);
  bucket->last = mm + allcate_sz;
  bucket->next = pool->bucket->next;
  pool->bucket->next = bucket;
  return bucket->hdr;
}
