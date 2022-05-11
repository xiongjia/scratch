/**
 * local tests
 */


#include "tachi-proc.h"
#include "tachi-config.h"

#include <stdio.h>
#include <winsock2.h>
#include <ws2def.h>

void wrk() {
  FD_SET fd_read;

  TIMEVAL tv;
  tv.tv_sec = 10;
  tv.tv_usec = 0;

  for (;;) {
    FD_ZERO(&fd_read);

    // nRet = select(0, &fdRead, NULL, NULL, &tv);

  }
}

int main(const int argc, const char **argv) {
  tachi_proc_startup();

  tachi_config config;
  config.port = 8893;

  // create socket 
  SOCKET servSock = socket(PF_INET, SOCK_STREAM, IPPROTO_TCP);

  // bind
  struct sockaddr_in sockAddr;
  memset(&sockAddr, 0, sizeof(sockAddr));
  sockAddr.sin_family = PF_INET;
  sockAddr.sin_addr.s_addr = inet_addr("127.0.0.1");
  sockAddr.sin_port = htons(config.port);
  bind(servSock, &sockAddr, sizeof(struct sockaddr_in));

  listen(servSock, 32);

  int nonBlock = 1;
  ioctlsocket(servSock, FIONBIO, &nonBlock);

  FD_SET writeSet;
  FD_SET readSet;

  FD_ZERO(&readSet);
  FD_ZERO(&writeSet);

  FD_SET(servSock, &readSet);
  int total = select(0, &readSet, &writeSet, NULL, NULL);
  if (SOCKET_ERROR == total) {
    printf("Select error\n");
    return 1;
  }

  if (FD_ISSET(servSock, &readSet)) {
    SOCKADDR_IN addrClient;
    int nLenAddrClient = sizeof(SOCKADDR);
    SOCKET sockClient = accept(servSock, (SOCKADDR*)&addrClient, &nLenAddrClient);

    int nonBlock = 1;
    ioctlsocket(sockClient, FIONBIO, &nonBlock);

  }

  closesocket(servSock);

  tachi_proc_cleanup();
  return 0;
}
