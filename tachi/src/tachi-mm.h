/**
 */

#ifndef _TACHI_MM_H_
#define _TACHI_MM_H_ 1

#include "tachi-general.h"

typedef struct _tachi_pool tachi_pool_t;

#define tachi_malloc(_sz) malloc(_sz);
#define tachi_free(_ptr)  free(_ptr);


tachi_pool_t* tachi_pool_create(size_t size);
void tachi_pool_destroy(tachi_pool_t *pool);

void* tachi_pool_alloc(tachi_pool_t *pool, size_t size);


#endif /* !defined(_TACHI_MM_H_) */
