/**
 */

#include "tachi-str.h"

void tachi_test_str(void) {
  tachi_str_t test1 = tachi_string("PASS");
  printf("%s\n", test1.data);
}
