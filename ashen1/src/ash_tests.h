/*
 */

#ifndef _ASH_TESTS_H_
#define _ASH_TESTS_H_ 1

#include "ash_types.h"
#include "ash_log.h"

typedef struct _ash_unit_test_context_s {
  time_t tm_start;
  time_t tm_end;
  const char *test_filter;
  ash_log_t *log;
} ash_unit_test_context_t;

typedef struct _ash_unit_test_s {
  const char *test_name;
  void (*invoke)(ash_unit_test_context_t *ctx);
} ash_unit_test_t;

extern ash_unit_test_t unittest_simple_str;
extern ash_unit_test_t unittest_simple_time;
extern ash_unit_test_t unittest_filename;

#endif /* !defined(_ASH_TESTS_H_) */
