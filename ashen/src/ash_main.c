/*
 */

#include "ash_config.h"
#include "ash_general.h"
#include "ash_log.h"
#include <stdio.h>

int main(const int argc, const char **argv) {
  ash_log_context_t *log_context = ASH_NULL;

  printf("test, os type %s\n", ASH_OS_TYPE);

  log_context = ash_log_create();
  ash_log_destroy(log_context);
  return 0;
}

