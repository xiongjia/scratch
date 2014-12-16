/**
 *
 */

#ifndef _ROACH_HXX_
#define _ROACH_HXX_ 1

#include "boost/utility.hpp"
#include "boost/shared_ptr.hpp"

#include "roach_log.hxx"
#include "roach_serv.hxx"

namespace roach {

class Roach : boost::noncopyable
{
public:
    static boost::shared_ptr<Roach> Create(void);

public:
    virtual void SetLevel(Logger::Level logLevel) = 0;
    virtual void RegisterHandler(boost::shared_ptr<LoggerHandler> handler) = 0;

    virtual boost::shared_ptr<Server> CreateServ(void) = 0;

protected:
    Roach(void);
};


} /* namespace roach */

#endif /* !defined(_ROACH_HXX_) */
