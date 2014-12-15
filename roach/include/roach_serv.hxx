/**
*
*/

#ifndef _ROACH_SERV_HXX_
#define _ROACH_SERV_HXX_ 1

#include "boost/utility.hpp"
#include "boost/shared_ptr.hpp"
#include "roach_log.hxx"

namespace roach {

class Server : boost::noncopyable
{
public:
    static boost::shared_ptr<Server> Create(boost::shared_ptr<Logger> logger);

    virtual bool Listen(const char *ip, const int port) = 0;
    virtual bool Run(void) = 0;

protected:
    Server(void);
};

} /* namespace roach */

#endif /* !defined(_ROACH_SERV_HXX_) */
