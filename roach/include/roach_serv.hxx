/**
*
*/

#ifndef _ROACH_SERV_HXX_
#define _ROACH_SERV_HXX_ 1

#include "boost/utility.hpp"
#include "boost/shared_ptr.hpp"
#include "roach_context.hxx"

namespace roach {

class Server : boost::noncopyable
{
public:
    static boost::shared_ptr<Server> Create(boost::shared_ptr<Context> ctx);

    virtual bool Run(const char *ip, const int port) = 0;
    virtual void Stop(void) = 0;

    virtual bool IsRunning(void) const = 0;

protected:
    Server(void);
};

} /* namespace roach */

#endif /* !defined(_ROACH_SERV_HXX_) */
