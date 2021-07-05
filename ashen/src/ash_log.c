/**
 *
 */

#include "ash_config.h"
#include "ash_general.h"
#include "ash_log.h"
#include <stdlib.h>
#include <stdnoreturn.h>

struct _ash_log_context {
  ash_uint_t log_level;
};

noreturn void exit_test();

noreturn void exit_test() {
  exit(1);
}

ash_log_context_t* ash_log_create(void) {
  return ASH_NULL;
}

void ash_log_destroy(ash_log_context_t *context) {
  exit(1);
}

