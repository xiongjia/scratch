/**
 */

extern "C"
{
    #include <lua.h>
    #include <lauxlib.h>
    #include <lualib.h>
}

#include "ab_module.h"
#include "ab_logger.h"

_ANUBARAK_BEGIN

Module::Module(void)
{
    /* NOP */
}

void Module::Register(lua_State *lua)
{
    const char *modName = GetModuleName();
    AB_LOG(AB_LOG_INF, "Load LUA module (%s)", modName);
    luaL_register(lua, modName, GetLUAReg());
}

_ANUBARAK_END
