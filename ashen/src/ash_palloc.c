/**
 *
 */

#include "ash_alloc.h"
#include "ash_palloc.h"

typedef struct _ash_pool_data_t ash_pool_data_t;
typedef struct _ash_pool_data_large_t ash_pool_data_large_t;

#define ASH_POOL_MIN_SIZE  (1024 * 4)

struct _ash_pool_data_t {
  unsigned char *last;
  unsigned char *end;
  ash_pool_data_t *next;
};

struct _ash_pool_data_large_t {
  void *alloc;
  ash_pool_data_large_t *next;
};

struct _ash_pool_t {
  ash_pool_data_t data;
  ash_pool_data_large_t *large;
  size_t data_max_size;
};

ASH_DECLARE(ash_pool_t *) ash_pool_create(size_t size) {
  const size_t data_max_size = ASH_MIN(ASH_POOL_MIN_SIZE, size);
  const size_t alloc_sz = data_max_size + sizeof(ash_pool_t);
  ash_pool_t *pool = NULL;
  const unsigned char *mem = ash_alloc(alloc_sz);
  if (mem == NULL) {
    return NULL;
  }

  pool = (ash_pool_t *)mem;
  pool->large = NULL;
  pool->data_max_size = data_max_size;

  pool->data.next = NULL;
  pool->data.last = (unsigned char *)(mem + sizeof(ash_pool_t));
  pool->data.end = (unsigned char *)(mem + alloc_sz);
  return pool;
}

ASH_DECLARE(void) ash_pool_destroy(ash_pool_t *pool) {
  ash_pool_data_large_t *large;
  ash_pool_data_t *data;

  if (pool == NULL) {
    return;
  }

  /* free large data */
  for (large = pool->large; large != NULL; ) {
    large = large->next;
    ash_free(pool->large);
    pool->large = large; 
  }

  /* free data */
  for (data = pool->data.next; data != NULL; ) {
    data = data->next;
    ash_free(data);
    pool->data.next = data;
  }

  /* free pool */
  ash_free(pool);
}

ASH_DECLARE(void *) ash_pool_alloc(ash_pool_t *pool, size_t size) {
  ash_pool_data_large_t *large; 
  ash_pool_data_t *data;

  if (pool == NULL) {
    return NULL;
  }

  if (size > pool->data_max_size) {
    large = ash_alloc(size + sizeof(ash_pool_data_large_t));
    if (large == NULL) {
      return NULL;
    }
    large->alloc = (large + sizeof(ash_pool_data_large_t));
    large->next = pool->large;
    pool->large = large;
    return large->alloc;
  }

  /* XXX  TODO ALLOC from data
   * for (data = pool->data; 
   */ 

 
  return NULL;
}


ASH_DECLARE(void) ash_pool_free(ash_pool_t *pool, void *mem) {
  ash_pool_data_large_t *current;
  ash_pool_data_large_t *prev;

  if (pool == NULL || mem == NULL) {
    return;
  }
  
  for (current = pool->large, prev = NULL; current != NULL;
      prev = current, current = current->next) {
    if (current->alloc != mem) {
      continue;
    }

    if (prev == NULL) {
      pool->large = current->next;
    } else {
      prev->next = current->next;
    }
    ash_free(current);
    break;
  }
}

