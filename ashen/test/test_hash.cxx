/**
 *
 */

#include <gtest/gtest.h>
#include "ash_palloc.h"
#include "ash_hash.h"

TEST(AshHash, SimpleHash) {
  ash_pool_t *pool = ash_pool_create(512);
  ash_hash_t *hash = ash_hash_create(pool);

  ash_hash_set(hash, "key1", sizeof("key1"), "data1");
  ash_hash_set(hash, "key2", sizeof("key2"), "data2");
  const char *data1 = static_cast<const char *>(ash_hash_get(hash,
        "key1", sizeof("key1")));
  EXPECT_STREQ("data1", data1);
  const char *data2 = static_cast<const char *>(ash_hash_get(hash,
        "key2", sizeof("key2")));
  EXPECT_STREQ("data2", data2);
  ash_pool_destroy(pool);
}

