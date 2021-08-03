/**
 *
 */

#include "ash_list.h"

ASH_DECLARE(ash_list_t*) ash_list_create(ash_pool_t *pool) {
  ash_list_t *list;
  if (pool == NULL) {
    return NULL;
  }

  list = ash_pool_alloc(pool, sizeof(ash_list_t));
  if (list == NULL) {
    return NULL;
  }
  list->pool = pool;
  list->count = 0;
  list->head = list->tail = NULL;
  return list;
}

ASH_DECLARE(void) ash_list_push(ash_list_t *list, void *data) {
  ash_list_element_t *element;

  if (list == NULL) {
    return;
  }

  element = ash_pool_alloc(list->pool, sizeof(ash_list_element_t));
  if (element == NULL) {
    return;
  }
  element->data = data;
  element->next = NULL;

  if (list->head == NULL) {
    element->prev = NULL;
    list->head = element;
    list->tail = element;
  } else {
    element->prev = list->tail;
    list->tail->next = element;
    list->tail = element;
  }
  list->count++;
}


