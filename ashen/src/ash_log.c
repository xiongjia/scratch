/**
 *
 */

#include "ash_config.h"
#include "ash_general.h"
#include "ash_log.h"

struct _ash_log_context {
  ash_uint_t log_level;
};

ash_log_context_t* ash_log_create(void) {
  return ASH_NULL;
}

void ash_log_destroy(ash_log_context_t *context) {

}

