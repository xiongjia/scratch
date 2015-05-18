/**
 */

#ifndef _ANUBARAK_H_
#define _ANUBARAK_H_ 1

#include <boost/utility.hpp>
#include <boost/shared_ptr.hpp>

#include "ab_types.h"

_ANUBARAK_BEGIN

class Anubarak : boost::noncopyable
{
public:
    static boost::shared_ptr<Anubarak> Create(void);

protected:
    Anubarak(void);
};

_ANUBARAK_END

#endif /* !defined(_ANUBARAK_H_) */
