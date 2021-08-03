/**
 *
 */

#include "ash_alloc.h"
#include "ash_palloc.h"

typedef struct _ash_pool_data_t ash_pool_data_t;
typedef struct _ash_pool_data_large_t ash_pool_data_large_t;

#define ASH_POOL_MIN_SIZE  (1024 * 8)

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
  const size_t alloc_size = data_max_size + sizeof(ash_pool_t);
  ash_pool_t *pool = NULL;
  const unsigned char *mem = ash_alloc(alloc_size);
  if (mem == NULL) {
    return NULL;
  }

  pool = (ash_pool_t *)mem;
  pool->large = NULL;
  pool->data_max_size = data_max_size;

  pool->data.next = NULL;
  pool->data.last = (unsigned char *)(mem + sizeof(ash_pool_t));
  pool->data.end = (unsigned char *)(mem + alloc_size);
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
  void *mem;
  const size_t alloc_size = ASH_MAX(size, 1);

  if (pool == NULL) {
    return NULL;
  }

  if (size > pool->data_max_size) {
    /* alloc large mem */
    large = ash_alloc(alloc_size + sizeof(ash_pool_data_large_t));
    if (large == NULL) {
      return NULL;
    }
    large->alloc = (large + sizeof(ash_pool_data_large_t));
    large->next = pool->large;
    pool->large = large;
    return large->alloc;
  }

  for (data = &(pool->data); data != NULL; data = data->next) {
    if ((data->end - data->last) > alloc_size) {
      mem = data->last;
      data->last += alloc_size;
      return mem;
    }
  }

  data = ash_alloc(sizeof(ash_pool_data_t) + pool->data_max_size);
  mem = data + sizeof(ash_pool_data_t);
  data->last = (unsigned char *)(data + sizeof(ash_pool_data_t) + alloc_size);
  data->end = (unsigned char *)(data + sizeof(ash_pool_data_t) + pool->data_max_size);
  data->next = pool->data.next;
  pool->data.next = data;
  return mem;
}

ASH_DECLARE(void *) ash_pool_calloc(ash_pool_t *pool, size_t size) {
  void *mem = ash_pool_alloc(pool, size);
  if (mem != NULL) {
    memset(mem, 0, size);
  }
  return mem;
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

