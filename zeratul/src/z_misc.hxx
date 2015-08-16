/**
 * Zeratul - A simple socks proxy. (Boost ASIO Example)
 */

#ifndef _Z_MISC_HXX_
#define _Z_MISC_HXX_ 1

#include "z_types.hxx"
#include "boost/utility.hpp"
#include "boost/shared_ptr.hpp"
#include "boost/log/trivial.hpp"
#include "boost/log/attributes/named_scope.hpp"

_ZERATUL_BEGIN_

class LogMgr : boost::noncopyable
{
public:
    static boost::shared_ptr<LogMgr> Instance(void);

    virtual void Initialize(void) = 0;

protected:
    LogMgr(void);
};


_ZERATUL_END_

#endif /* !defined(_Z_MISC_HXX_) */
