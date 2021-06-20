/**
 * 
 */


#include <stdio.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netdb.h>

#include "tachi-log.hxx"


int main(const int argc, const char **argv) {
  TachiLog::getInstance().init();
  TACHI_LOG(info) << "Starting ...";


  // local tests 
  
  int sockfd = socket(AF_INET, SOCK_STREAM, 0);
  if (sockfd < 0) {
    TACHI_LOG(error) << "  create socket error ";
    return 1;
  }


	struct sockaddr_in serv_addr; 
  memset(&serv_addr, 0, sizeof(serv_addr));
 	serv_addr.sin_family = AF_INET;
  serv_addr.sin_addr.s_addr = INADDR_ANY;
  serv_addr.sin_port = htons(2333); 

  int binded = bind(sockfd, (struct sockaddr *)&serv_addr, sizeof(serv_addr));
 	if (binded < 0) {
 		TACHI_LOG(error) << "Error on binding";
    return 1;
 	}
 
  listen(sockfd, 100);
  for (;;) {
	  struct sockaddr client_addr;
    socklen_t client_addr_len = sizeof(client_addr);

    // TODO Select or other API
    int new_sockfd = accept(sockfd, &client_addr, &client_addr_len); 
    if (new_sockfd < 0) {
      TACHI_LOG(error) << " invalid socket";
      continue;
    }
   
    TACHI_LOG(debug) << " new connection"; 
    close(new_sockfd);
  }




  return 0;
}
