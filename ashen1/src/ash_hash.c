/**
 */

#include "ash_hash.h"

uint32_t ash_hash_dbj(const char *key, uint32_t klen, uint32_t hash) {
  uint32_t hash_val = hash;
  uint32_t i;

  /* An algorithm produced by Professor Daniel J. Bernstein and shown 
   * first to the world on the usenet newsgroup comp.lang.c. It is one of 
   * the most efficient hash functions ever published.
   */
  for (i = 0; i < klen; i++) {
    hash = ((hash << 5) + hash) + (uint32_t)key[i];
  }
  return hash; 
}

uint32_t ash_hash_dbj_default(const char *key, uint32_t klen) {
  return ash_hash_dbj(key, klen, 5381);
}
