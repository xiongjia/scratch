/**
 */

#ifndef _ASH_MM_H_
#define _ASH_MM_H_ 1

#include <stdlib.h>
#include "ash_types.h"

typedef struct _ash_pool_s ash_pool_t;

#define ash_malloc(_sz) malloc(_sz)
#define ash_free(_ptr)  free(_ptr)

ash_pool_t* ash_pool_create_default(void);

ash_pool_t* ash_pool_create(size_t size);

void ash_pool_destroy(ash_pool_t *pool);

void* ash_pool_alloc(ash_pool_t *pool, size_t size);

#endif /* !defined(_ASH_MM_H_) */
