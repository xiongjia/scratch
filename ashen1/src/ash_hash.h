/**
 */

#ifndef _ASH_HASH_H_
#define _ASH_HASH_H_ 1

#include "ash_types.h"
#include "ash_mm.h"

uint32_t ash_hash_times33(const char *key, uint32_t klen, uint32_t hash);

#endif /* !defined(_ASH_HASH_H_) */
