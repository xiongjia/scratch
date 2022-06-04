/**
 */

#ifndef _ASH_LIST_H_
#define _ASH_LIST_H_ 1

#include "ash_mm.h"

typedef struct _ash_list_node_s ash_list_node_t;

struct _ash_list_node_s {
  void *element;
  ash_list_node_t *next;
};

typedef struct _ash_list_s {
  ash_pool_t *pool;
  ash_list_node_t *head;
  ash_list_node_t *tail;
  size_t count;
} ash_list_t;

ash_list_t* ash_list_create(ash_pool_t *pool);

boolean_t ash_list_push(ash_list_t *list, void *data);

boolean_t ash_list_compare(ash_list_t *l1, ash_list_t *l2,
      boolean_t (*compare)(ash_list_node_t *n1, ash_list_node_t *n2));

#endif /* !defined(_ASH_LIST_H_) */
