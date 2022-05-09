/**
 * local tests
 */


#include <stdio.h>
#include <winsock2.h>
#include <ws2def.h>


static void proc_init() {
  WSADATA wsaData;
  WSAStartup(MAKEWORD(2, 2), &wsaData);
}

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
  proc_init();

  const u_short port = 8893;

  // create socket 
  SOCKET servSock = socket(PF_INET, SOCK_STREAM, IPPROTO_TCP);

  // bind
  struct sockaddr_in sockAddr;
  memset(&sockAddr, 0, sizeof(sockAddr));
  sockAddr.sin_family = PF_INET;
  sockAddr.sin_addr.s_addr = inet_addr("127.0.0.1");
  sockAddr.sin_port = htons(port);
  bind(servSock, &sockAddr, sizeof(struct sockaddr_in));

  listen(servSock, 32);

  for (;;) {
    SOCKADDR_IN addrClient;
    int nLenAddrClient = sizeof(SOCKADDR);
    SOCKET sockClient = accept(servSock, (SOCKADDR*)&addrClient, &nLenAddrClient);
    if (sockClient == INVALID_SOCKET) {
      printf("invalid socket\n");
      continue;
    }


  }
  closesocket(servSock);
  WSACleanup();
  return 0;
}
