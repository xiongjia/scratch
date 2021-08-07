/**
 */

#ifndef _ASH_HASH_H_
#define _ASH_HASH_H_ 1

#include "ash_general.h"
#include "ash_palloc.h"

ASH_DECLARE(unsigned long) ash_hash_djb(const char *key, size_t klen);

#endif /* !defined(_ASH_HASH_H_) */

