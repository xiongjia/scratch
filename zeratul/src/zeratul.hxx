/**
 * Zeratul - A simple socks proxy. (Boost ASIO)
 */

#ifndef _ZERATUL_HXX_
#define _ZERATUL_HXX_ 1

#include "boost/asio.hpp"
#include "boost/utility.hpp"
#include "boost/shared_ptr.hpp"

#include "z_types.hxx"
#include "z_misc.hxx"

_ZERATUL_BEGIN_

class Server : boost::noncopyable
{
public:
    static boost::shared_ptr<Server> Create(boost::asio::io_service &ioSvc,
                                            unsigned short port);

protected:
    Server(void);
};

_ZERATUL_END_

#endif /* !defined(_ZERATUL_HXX_) */

