/**
*/

#ifndef _AB_MODULE_H_
#define _AB_MODULE_H_ 1

#include <boost/utility.hpp>

extern "C"
{
    #include <lua.h>
    #include <lauxlib.h>
}

#include "ab_types.h"

_ANUBARAK_BEGIN

class Module : boost::noncopyable{public:
    virtual const char* GetModuleName(void) const = 0;
    virtual void Register(lua_State *lua);

protected:
    Module(void);
    virtual const luaL_Reg* GetLUAReg(void) = 0;
};

_ANUBARAK_END

#endif /* !defined(_AB_MODULE_H_) */
