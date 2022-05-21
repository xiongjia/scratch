/*
 */

#ifndef _ASH_TESTS_H_
#define _ASH_TESTS_H_ 1

typedef struct _ash_unit_test_s {
  const char *test_name;
  void (*invoke_test)(void);
} ash_unit_test_t;

extern ash_unit_test_t unittest_simple_str;
extern ash_unit_test_t unittest_simple_time;
extern ash_unit_test_t unittest_filename;

#endif /* !defined(_ASH_TESTS_H_) */
