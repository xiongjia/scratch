/**
 */

#ifndef _ASH_STRING_H_
#define _ASH_STRING_H_ 1

#include "ash_general.h"
#include "ash_palloc.h"

ASH_DECLARE(char *) ash_str_duplicate(ash_pool_t *pool, const char *src);

#endif /* !defined(_ASH_STRING_H_) */

