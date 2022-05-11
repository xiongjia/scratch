/**
 * 
 */


#include <winsock2.h>
#include <ws2def.h>

#include "tachi-proc.h"

void tachi_proc_startup(void) {
  WSADATA wsa_data;
  WSAStartup(MAKEWORD(2, 2), &wsa_data);
}

void tachi_proc_cleanup(void) {
  WSACleanup();
}
