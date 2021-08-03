/**
 */

#ifndef _ASH_PALLOC_H_
#define _ASH_PALLOC_H_ 1

#include "ash_general.h"

typedef struct _ash_pool_t ash_pool_t;


ASH_DECLARE(ash_pool_t *) ash_pool_create(size_t size);

ASH_DECLARE(void) ash_pool_destroy(ash_pool_t *pool);


ASH_DECLARE(void *) ash_pool_alloc(ash_pool_t *pool, size_t size);

ASH_DECLARE(void *) ash_pool_calloc(ash_pool_t *pool, size_t size);

ASH_DECLARE(void) ash_pool_free(ash_pool_t *pool, void *mem);


#endif /* !defined(_ASH_PALLOC_H_) */

