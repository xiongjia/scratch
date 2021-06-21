/**
 * gnu autotools test
 *
 */

#include "config.h"
#include <stdio.h>

int main(const int argc, const char **argv) {
  printf("Package =  %s (%s)\n", PACKAGE_NAME, PACKAGE_STRING);
  return 0;
}

