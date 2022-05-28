/*
 */

#include <stdio.h>
#include "ash_mm.h"
#include "ash_str.h"
#include "ash_tests.h"

static void test_simple_str(ash_unit_test_context_t *ctx,
                            ash_unit_test_case_context_t *tc) {
  size_t sz;
  char buf[1024];
  char *dup;

  dup = ash_pstrdup(tc->pool, "123");
  ASHTU_STR_EQU(tc, "123", dup);
  dup = ash_pstrcat(tc->pool, "123", "456", NULL);
  ASHTU_STR_EQU(tc, "123456", dup);

  sz = ash_snprintf(buf, sizeof(buf), "test");
  ASHTU_STR_EQU(tc, "test", buf);
  ASHTU_SIZE_EQU(tc, 5, sz);
  sz = ash_snprintf(NULL, 0, "test");
  ASHTU_SIZE_EQU(tc, 5, sz);
  sz = ash_snprintf(buf, 1, "test");
  ASHTU_SIZE_EQU(tc, (size_t)-1, sz);

  ash_snprintf(buf, sizeof(buf), "str: %% %s%", "abc");
  ASHTU_STR_EQU(tc, "str: % abc", buf);
}

ash_unit_test_case_t unittest_simple_str = {
  .test_name = "simple-string",
  .invoke = test_simple_str
};
