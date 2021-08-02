/**
 *
 */

#include "ash_palloc.h"

typedef struct _ash_pool_data_t ash_pool_data_t;
typedef struct _ash_pool_data_large_t ash_pool_data_large_t;
typedef struct _ash_pool_t ash_pool_t;

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
  ash_pool_t *current;
};

