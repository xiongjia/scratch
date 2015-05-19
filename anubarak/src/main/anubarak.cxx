/**
 *
 */

#include <boost/make_shared.hpp>
#include <boost/filesystem.hpp>

extern "C"
{
    #include <lua.h>
    #include <lauxlib.h>
    #include <lualib.h>
}

#include "anubarak.h"
#include "ab_logger.h"

_ANUBARAK_BEGIN

class AnubarakImpl : public Anubarak
{
private:
    lua_State *m_lua;

public:
    AnubarakImpl(void)
        : Anubarak()
        , m_lua(NULL)
    {
        /* TODO
         * 1. update lua panic
         * 2. load built-in LUA modules
         */
        m_lua = luaL_newstate();
        luaL_openlibs(m_lua);
    }

    virtual ~AnubarakImpl(void)
    {
        if (NULL != m_lua)
        {
            lua_close(m_lua);
            m_lua = NULL;
        }
    }

    virtual void Run(const char *testLuaFile)
    {
        auto srcPath = boost::filesystem::system_complete(testLuaFile);
        AB_LOG(AB_LOG_DBG, "Test Lua file: %s", srcPath.string().c_str());
        int error = luaL_loadfile(m_lua, srcPath.string().c_str()) ||
                    lua_pcall(m_lua, 0, 0, 0);
        AB_LOG(AB_LOG_DBG, "Result: %d", error);
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
