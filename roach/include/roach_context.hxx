/**
 *
 */

#ifndef _ROACH_CONTEXT_HXX_
#define _ROACH_CONTEXT_HXX_ 1

#include "boost/utility.hpp"
#include "boost/shared_ptr.hpp"

#include "roach_log.hxx"

namespace roach {

class Context : boost::noncopyable
{
public:
    static boost::shared_ptr<Context> Create(boost::shared_ptr<Logger> logger);

    virtual boost::shared_ptr<Logger> GetLogger(void) = 0;

protected:
    Context(void);
};

} /* namespace roach */

#endif /* !defined(_ROACH_CONTEXT_HXX_) */
