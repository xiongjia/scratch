/**
 * 
 */

#ifndef _TACHI_MISC_H_
#define _TACHI_MISC_H_ 1

#include <winsock2.h>


SOCKET create_listening(const char* addr,
                        const unsigned short port,
                        int backlog);


#endif /* !defined(_TACHI_MISC_H_) */
