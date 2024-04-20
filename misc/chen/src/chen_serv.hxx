/**
 * Chen - My Network protocol tests
 */

#ifndef _CHEN_SERV_HXX_
#define _CHEN_SERV_HXX_ 1

#include "boost/shared_ptr.hpp"
#include "boost/utility.hpp"
#include "chen_types.hxx"

_CHEN_BEGIN_

#if 0
class Server : boost::noncopyable
{
public:
    static boost::shared_ptr<Server> create(void);

public:
    virtual bool run(void) = 0;

protected:
    Server(void);
};
#endif

_CHEN_END_

#endif /* !defined(_CHEN_SERV_HXX_) */
