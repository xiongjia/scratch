/**
 */

#include "ab_mod_logger.h"
#include "ab_logger.h"

_ANUBARAK_BEGIN

namespace
{
    int ABLLogger(lua_State *lua, ab::LoggerMask mask)
    {
        const char *logStr = luaL_optlstring(lua, 1, NULL, NULL);
        if (NULL != logStr)
        {
            ab::Logger::instance()->LogNoFmt(NULL, 0, mask, logStr);
        }
        return 0;
    }

    int ABLLoggerDbg(lua_State *lua)
    {
        return ABLLogger(lua, AB_LOG_DBG);
    }

    int ABLLoggerInf(lua_State *lua)
    {
        return ABLLogger(lua, AB_LOG_INF);
    }

    int ABLLoggerErr(lua_State *lua)
    {
        return ABLLogger(lua, AB_LOG_ERR);
    }

    int ABLLoggerWar(lua_State *lua)
    {
        return ABLLogger(lua, AB_LOG_WAR);
    }

    const struct luaL_Reg LOGGER_LIB[] =
    {
        { "debug", ABLLoggerDbg },
        { "info",  ABLLoggerInf },
        { "err",   ABLLoggerErr },
        { "warn",  ABLLoggerWar },
        { NULL, NULL }
    };
}

ModLogger::ModLogger(void)
{
    /* NOP */
}

const char* ModLogger::GetModuleName(void) const
{
    return "logger";
}

const luaL_Reg* ModLogger::GetLUAReg(void)
{
    return LOGGER_LIB;
}

_ANUBARAK_END
