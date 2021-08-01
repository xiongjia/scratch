/*
 */

#ifndef _ASH_ALLOC_H_
#define _ASH_ALLOC_H_ 1

#include "ash_general.h"

ASH_DECLARE(void *) ash_alloc(size_t size);

ASH_DECLARE(void *) ash_calloc(size_t size);

ASH_DECLARE(void) ash_free(void *mem);

#endif /* !defined(_ASH_ALLOC_H_) */

