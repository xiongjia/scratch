/**
 * local tests
 */


#include "tachi-proc.h"
#include "tachi-config.h"
#include "tachi-misc.h"

#include <stdio.h>
#include <winsock2.h>
#include <ws2def.h>

static int connection_count = 0;

static void do_accept(tachi_config* config, SOCKET serv_sock) {
  SOCKADDR_IN client_addr;
  int addr_len = sizeof(SOCKADDR);
  SOCKET sock_client;

  sock_client = accept(serv_sock, (SOCKADDR*)&client_addr, &addr_len);
  if (INVALID_SOCKET  == sock_client) {
    printf("[%d] Invalid input connection\n", config->pid);
    return;
  }

  printf("[%d] new connection (%u)\n", config->pid, ++connection_count);

  // Sleep(1000 * 30);
  shutdown(sock_client, SD_BOTH);
}

static void main_loop(tachi_config *config, SOCKET serv_sock) {
  TIMEVAL tv;
  fd_set rd_set;
  int rs;

  for (;;) {
    FD_ZERO(&rd_set);
    FD_SET(serv_sock, &rd_set);
    tv.tv_sec = 30;
    tv.tv_usec = 0;

    printf("[%d] waiting ...\n", config->pid);
    rs = select(0, &rd_set, NULL, NULL, &tv);
    if (SOCKET_ERROR == rs) {
      continue;
    }

    if (FD_ISSET(serv_sock, &rd_set)) {
      do_accept(config, serv_sock);
    }
  }
}

static void work(tachi_config *config) {
  SOCKET serv_sock;

  serv_sock = create_listening(config->addr, config->port, 32);
  if (INVALID_SOCKET == serv_sock) {
    return;
  }
  main_loop(config, serv_sock);
  closesocket(serv_sock);
}

int main(const int argc, const char **argv) {
  tachi_proc_startup();

  tachi_config config;
  config.addr = "127.0.0.1";
  config.port = 8893;
  config.pid = GetCurrentProcessId();
  work(&config);

  tachi_proc_cleanup();
  return 0;
}
