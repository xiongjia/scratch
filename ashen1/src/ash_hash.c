/**
 */

#include "ash_hash.h"

uint32_t ash_hash_times33(const char *key, uint32_t klen, uint32_t hash) {
  uint32_t hash_val = hash;
  for (uint32_t i = 0; i < klen; i++) {
    hash = ((hash << 5) + hash) + (uint32_t)key[i];
  }
  return hash; 
}
