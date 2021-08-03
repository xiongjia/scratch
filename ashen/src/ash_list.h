/**
 *
 */

#ifndef _ASH_LIST_H_
#define _ASH_LIST_H_ 1

#include "ash_general.h"
#include "ash_palloc.h"

typedef struct _ash_list_t ash_list_t;
typedef struct _ash_list_element_t ash_list_element_t;

struct _ash_list_element_t {
  void *data;

  ash_list_element_t *next;
  ash_list_element_t *prev;
};

struct _ash_list_t {
  ash_pool_t *pool;

  ash_list_element_t *head;
  ash_list_element_t *tail;
  size_t count;
};

#define ASH_LIST_FOR_EACH(_list, _pos) \
  for (_pos = _list->head; _pos != NULL; _pos = _pos->next) 

#define ASH_LIST_COUNT(_list)  (_list->count)

ASH_DECLARE(ash_list_t*) ash_list_create(ash_pool_t *pool);

ASH_DECLARE(void) ash_list_push(ash_list_t *list, void *data);

#endif /* !defined(_ASH_LIST_H_) */

