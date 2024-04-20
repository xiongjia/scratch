/**
 * Roach - A simple HTTP Server. (libuv & http-parser & boost & C++11 & CMake)
 */

#ifndef _ROACH_HXX_
#define _ROACH_HXX_ 1

#include "boost/utility.hpp"
#include "boost/shared_ptr.hpp"

#include "roach_log.hxx"
#include "roach_context.hxx"
#include "roach_serv.hxx"

namespace roach {

class Roach : boost::noncopyable
{
public:
    static boost::shared_ptr<Roach> Create(void);

public:
    virtual boost::shared_ptr<Context> GetContext(void) = 0;

    virtual boost::shared_ptr<Server> CreateServ(void) = 0;
protected:
    Roach(void);
};


} /* namespace roach */

#endif /* !defined(_ROACH_HXX_) */
