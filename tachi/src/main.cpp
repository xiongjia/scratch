/**
 * 
 */

#include <stdio.h>
#include <sqlite3.h>

int main(const int argc, const char **argv) {
  printf("test = %s\n", sqlite3_libversion());
  return 0;
}
