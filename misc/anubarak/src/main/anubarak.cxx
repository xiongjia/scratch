/**
 *
 */

#include <boost/make_shared.hpp>
#include <boost/filesystem.hpp>
#include <boost/foreach.hpp>

extern "C"
{
    #include <lua.h>
    #include <lauxlib.h>
    #include <lualib.h>
}

#include "anubarak.h"
#include "ab_logger.h"

#include "ab_mod_logger.h"

_ANUBARAK_BEGIN

class AnubarakImpl : public Anubarak
{
private:
    lua_State *m_lua;
    std::vector<boost::shared_ptr<Module>> m_modules;

private:
    static int LuaPanic(lua_State *lua)
    {
        /* lua panic */
        if (LUA_TSTRING == lua_type(lua, -1))
        {
            AB_LOG(AB_LOG_ERR, "A Lua error has occured: %s",
                lua_tostring(lua, -1));
        }
        else
        {
            AB_LOG(AB_LOG_ERR, "A Lua error has occured: %p",
                lua_topointer(lua, -1));
        }
        return 0;
    }

public:
    AnubarakImpl(void)
        : Anubarak()
        , m_lua(NULL)
    {
        m_lua = luaL_newstate();
        lua_atpanic(m_lua, AnubarakImpl::LuaPanic);

        /* load build-in modules */
        m_modules.push_back(boost::make_shared<ModLogger>());
        BOOST_FOREACH(auto mod, m_modules)
        {
            mod->Register(m_lua);
        }
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
        AB_LOG(AB_LOG_INF, "Result: %d", error);
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
