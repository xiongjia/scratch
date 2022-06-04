/*
 */

#include "ash_list.h"


ash_list_t* ash_list_create(ash_pool_t *pool) {
  ash_list_t *list;

  list = (ash_list_t *)ash_pool_alloc(pool, sizeof(ash_list_t));
  if (NULL == list) {
    return NULL;
  }
  list->pool = pool;
  list->count = 0;
  list->head = NULL;
  list->tail = NULL;
  return list;
}

boolean_t ash_list_push(ash_list_t *list, void *data) {
  ash_list_node_t *node;

  if (NULL == list || NULL == list->pool) {
    return ASH_FALSE;
  }

  node = (ash_list_node_t*)ash_pool_alloc(list->pool, sizeof(ash_list_node_t));
  if (NULL == node) {
    return ASH_FALSE;
  }

  node->element = data;
  node->next = NULL;
  if (NULL == list->head) {
    list->head = node;
    list->tail = node;
    list->count++;
    return ASH_TRUE;
  }

  if (NULL == list->tail) {
    return ASH_FALSE;
  }

  list->tail->next = node;
  list->tail = node;
  list->count++;
  return ASH_TRUE;
}

boolean_t ash_list_compare(ash_list_t *l1, ash_list_t *l2,
    boolean_t (*compare)(ash_list_node_t *n1, ash_list_node_t *n2)) {
  ash_list_node_t *n1;
  ash_list_node_t *n2;

  if (l1 == l2) {
    return ASH_TRUE;
  }
  if (NULL == l1 && NULL == l2) {
    return ASH_TRUE;
  }
  if (NULL == l1 || NULL == l2) {
    return ASH_FALSE;
  }
  if (l1->count != l2->count) {
    return ASH_FALSE;
  }

  for (n1 = l1->head, n2 = l2->head; NULL != n1;
      n1 = n1->next, n2 = n2->next) {
    if (!compare(n1, n2)) {
      return ASH_FALSE;
    }
  }
  return ASH_TRUE;
}
