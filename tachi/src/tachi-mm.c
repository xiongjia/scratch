/**
 * 
 */

#include "tachi-mm.h"


typedef struct _tachi_pool_data {
  uchar_t *hdr;
  uchar_t *last;
  uchar_t* current;

  tachi_pool *next;
} tachi_pool_data;

struct _tachi_pool {
  tachi_pool_data *data;
};


tachi_pool* tachi_create_pool(void) {

  return NULL;
}
