/**
 *
 */

#include <stdio.h>
#include <boost/make_shared.hpp>
#include "anubarak.h"

_ANUBARAK_BEGIN

class AnubarakImpl : public Anubarak
{
public:
    AnubarakImpl(void)
        : Anubarak()
    {
        /* NOP */
    }
};

Anubarak::Anubarak(void)
{
    /* NOP */
}

boost::shared_ptr<Anubarak> Anubarak::Create(void)
{
    return boost::make_shared<AnubarakImpl>();
}

_ANUBARAK_END
