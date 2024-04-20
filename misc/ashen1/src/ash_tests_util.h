/**
 */

#ifndef _ASH_TESTS_UTIL_H_
#define _ASH_TESTS_UTIL_H_ 1

#include "ash_types.h"
#include "ash_log.h"
#include "ash_mm.h"
#include "ash_list.h"

typedef struct _ash_unit_test_context_s {
  time_t tm_start;
  time_t tm_end;
  const char *test_filter;
  ash_log_t *log;
} ash_unit_test_context_t;

typedef struct _ash_unit_test_case_context_s {
  boolean_t verbose;
  boolean_t failed;
  ash_pool_t *pool;
} ash_unit_test_case_context_t;

typedef struct _ash_unit_test_case_s {
  const char *test_name;
  void (*invoke)(ash_unit_test_context_t *ctx,
                 ash_unit_test_case_context_t *test_case);
} ash_unit_test_case_t;

void ash_testutil_str_equal(ash_unit_test_case_context_t *tc,
                            const char *expected, const char *actual,
                            const char *src, const int line);

void ash_testutil_size_equal(ash_unit_test_case_context_t *tc,
                             const size_t expected, const size_t actual,
                             const char *src, const int line);

void ash_testutil_int32_equal(ash_unit_test_case_context_t *tc,
                              const int32_t expected, const int32_t actual,
                              const char *src, const int line);

void ash_testutil_slist_equal(ash_unit_test_case_context_t *tc,
                              const ash_list_t *expected,
                              const ash_list_t *actual,
                              const char *src, const int line);

#define ASHTU_STR_EQU(_tc, _expected, _actual) \
  ash_testutil_str_equal(_tc, _expected, _actual, __FILE__, __LINE__)

#define ASHTU_SIZE_EQU(_tc, _expected, _actual) \
  ash_testutil_size_equal(_tc, _expected, _actual, __FILE__, __LINE__)

#define ASHTU_INT32_EQU(_tc, _expected, _actual) \
  ash_testutil_int32_equal(_tc, _expected, _actual, __FILE__, __LINE__)

#define ASHTU_SLIST_EQU(_tc, _expected, _actual) \
  ash_testutil_slist_equal(_tc, _expected, _actual, __FILE__, __LINE__)

#endif /* !defined(_ASH_TESTS_UTIL_H_) */
