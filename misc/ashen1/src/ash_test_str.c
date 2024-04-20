/*
 */

#include <string.h>
#include <stdio.h>
#include "ash_mm.h"
#include "ash_str.h"
#include "ash_tests.h"

static void test_simple_str(ash_unit_test_context_t *ctx,
                            ash_unit_test_case_context_t *tc) {
  ash_list_t *expected_lines;
  ash_list_t *lines;
  int32_t sz;
  char buf[1024];
  char *dup;

  dup = ash_pstrdup(tc->pool, "123");
  ASHTU_STR_EQU(tc, "123", dup);
  dup = ash_pstrcat(tc->pool, "123", "456", NULL);
  ASHTU_STR_EQU(tc, "123456", dup);

  sz = ash_snprintf(buf, sizeof(buf), "test");
  ASHTU_STR_EQU(tc, "test", buf);
  ASHTU_INT32_EQU(tc, 5, sz);
  sz = ash_snprintf(NULL, 0, "test");
  ASHTU_INT32_EQU(tc, 5, sz);
  sz = ash_snprintf(buf, 1, "test");
  ASHTU_INT32_EQU(tc, -1, sz);

  ash_snprintf(buf, sizeof(buf), "str: %% %c %s%", '1', "abc");
  ASHTU_STR_EQU(tc, "str: % 1 abc", buf);
  ash_snprintf(buf, sizeof(buf), "test\\ - \n");
  ASHTU_STR_EQU(tc, "test\\ - \n", buf);

  ash_snprintf(buf, sizeof(buf), "%d", 123);
  ASHTU_STR_EQU(tc, "123", buf);
  ash_snprintf(buf, sizeof(buf), "%h", 0x1A2b);
  ASHTU_STR_EQU(tc, "1a2b", buf);
  ash_snprintf(buf, sizeof(buf), "%H", 0x1A2b);
  ASHTU_STR_EQU(tc, "1A2B", buf);

  ash_snprintf(buf, sizeof(buf), "%5d", 123);
  ASHTU_STR_EQU(tc, "  123", buf);
  ash_snprintf(buf, sizeof(buf), "%05d", 123);
  ASHTU_STR_EQU(tc, "00123", buf);
  ash_snprintf(buf, sizeof(buf), "%03d", 123);
  ASHTU_STR_EQU(tc, "123", buf);
  ash_snprintf(buf, sizeof(buf), "%03u", 123);
  ASHTU_STR_EQU(tc, "123", buf);
  ash_snprintf(buf, sizeof(buf), "%d", -123);
  ASHTU_STR_EQU(tc, "-123", buf);

  ash_snprintf(buf, sizeof(buf), "%lld", (int64_t)123);
  ASHTU_STR_EQU(tc, "123", buf);

  expected_lines = ash_slist_create(tc->pool, "123", "456", "789", NULL);
  lines = ash_parse_lines(tc->pool, "123\n456\r\n789", 512);
  ASHTU_SLIST_EQU(tc, expected_lines, lines);

  dup = ash_pstrdup(tc->pool, "  123  \t  \v ");
  ash_str_trim(dup, NULL);
  ASHTU_STR_EQU(tc, "123", dup);
}

ash_unit_test_case_t unittest_simple_str = {
  .test_name = "simple-string",
  .invoke = test_simple_str
};
