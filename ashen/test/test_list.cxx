/**
 *
 */

#include <gtest/gtest.h>
#include "ash_palloc.h"
#include "ash_list.h"

TEST(AshList, SimpleList) {
  ash_pool_t *pool = ash_pool_create(512);
  ash_list_t *list = ash_list_create(pool);

  ash_list_push(list, (void *)"1");
  ash_list_push(list, (void *)"2");

  ash_list_element_t *node = NULL;
  ASH_LIST_FOR_EACH(list, node) {
    printf("Item: %s\n", (const char *)node->data);
  }

  EXPECT_EQ(2, ASH_LIST_COUNT(list));
  ash_pool_destroy(pool);
}

