/**
 */

#include "ash_hash.h"

typedef struct _ash_hash_entry_t ash_hash_entry_t;

typedef struct _ash_hash_t ash_hash_t;

struct _ash_hash_entry_t {
  const void *key;
  const size_t klen;

  const void *val;
};

struct _ash_hash_t {
  ash_pool_t *pool;

};

ASH_DECLARE(unsigned long) ash_hash_djb(const char *key, size_t klen) {
  unsigned int hash = 5381;
  size_t i = 0;

  /* An algorithm produced by Professor Daniel J. Bernstein and shown 
   * first to the world on the usenet newsgroup comp.lang.c. It is one of 
   * the most efficient hash functions ever published.
   */
  for (i = 0; i < klen; ++i) {
    hash = ((hash << 5) + hash) + key[i];
  }
  return hash;
}

