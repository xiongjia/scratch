/**
 */

#ifndef _TACHI_MM_H_
#define _TACHI_MM_H_ 1

#include "tachi-general.h"

typedef struct _tachi_pool tachi_pool;

#define tachi_malloc(_sz) malloc(_size);

#define tachi_free(_ptr)  free(_ptr);


tachi_pool* tachi_create_pool(void);

#endif /* !defined(_TACHI_MM_H_) */
