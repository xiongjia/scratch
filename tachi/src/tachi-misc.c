/**
 */

#include "tachi-misc.h"


SOCKET create_listening(const char *addr,
                        const unsigned short port,
                        int backlog) {
  struct sockaddr_in sock_addr;
  SOCKET sock;
  u_long non_blocking;
  int opt_val;
  int rs;

  sock = socket(PF_INET, SOCK_STREAM, IPPROTO_TCP);
  if (INVALID_SOCKET == sock) {
    // xxx add logs
    return INVALID_SOCKET;
  }

  opt_val = 1;
  rs = setsockopt(sock, SOL_SOCKET, SO_REUSEADDR,
    (const char *)&opt_val, sizeof(opt_val));
  if (SOCKET_ERROR == rs) {
    printf("setsockopt error\n");
    closesocket(sock);
    return INVALID_SOCKET;
  }

  memset(&sock_addr, 0, sizeof(sock_addr));
  sock_addr.sin_family = PF_INET;
  sock_addr.sin_addr.s_addr = inet_addr(addr);
  sock_addr.sin_port = htons(port);

  rs = bind(sock, &sock_addr, sizeof(struct sockaddr_in));
  if (SOCKET_ERROR == rs) {
    printf("bind error");
    closesocket(sock);
    return INVALID_SOCKET;
  }

  rs = listen(sock, backlog);
  if (SOCKET_ERROR == rs) {
    printf("listen error");
    closesocket(sock);
    return INVALID_SOCKET;
  }

#if 0
  non_blocking = 1;
  rs = ioctlsocket(sock, FIONBIO, &non_blocking);
  if (SOCKET_ERROR == rs) {
    printf("ioctlsocket error");
    closesocket(sock);
    return INVALID_SOCKET;
  }
#endif

  return sock;
}
