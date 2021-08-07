/**
 *
 */

#include <gtest/gtest.h>
#include "ash_palloc.h"
#include "ash_hash.h"

TEST(AshHash, SimpleHash) {
  ash_pool_t *pool = ash_pool_create(512);

  unsigned long hash = ash_hash_djb("data", 4);
  printf("hash1 = %lu\n", hash);
  hash = ash_hash_djb("dat1", 4);
  printf("hash2 = %lu\n", hash);

  ash_pool_destroy(pool);
}

