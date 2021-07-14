/*
 */

#ifndef _ASH_ALLOC_H_
#define _ASH_ALLOC_H_ 1

#include "ash_general.h"

#ifdef __cplusplus
extern "C" {
#endif /* defined(__cplusplus) */

void * ash_alloc(size_t size);

void * ash_calloc(size_t size);

void ash_free(void *mem);

#ifdef __cplusplus
}
#endif /* defined(__cplusplus) */

#endif /* !defined(_ASH_ALLOC_H_) */

