/**
 * gnu autotools test
 *
 */

#include "tautotools_config.h"
#include <stdio.h>

int main(const int argc, const char **argv) {
  printf("Package =  %s (%s)\n", PACKAGE_NAME, PACKAGE_STRING);
#ifdef HAVE_EXAMPLES 
  printf("Examples are enabled.\n");
#else
  printf("Examples are disabled.\n");
#endif
  return 0;
}

