/**
 */

#include <stdio.h>

int main(const int argc, const char **argv) {
  int i;

  for (i = 0; i < argc; ++i) {
    printf("argv[%d] = %s", i, argv[i]);
  }

  printf("PASS\n");
  return 0;
}
