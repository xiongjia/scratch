/**
 * 
 */


#include <stdio.h>
#include <sys/socket.h>

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


  return 0;
}
