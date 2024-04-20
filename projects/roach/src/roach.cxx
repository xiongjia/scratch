/**
 * Roach - A simple HTTP Server. (libuv & http-parser & boost & C++11 & CMake)
 */

#include "roach.hxx"
#include "boost/make_shared.hpp"

namespace roach {

class RoachImpl : public Roach
{
private:
    boost::shared_ptr<Logger> m_logger;
    boost::shared_ptr<Context> m_ctx;

public:
    RoachImpl()
        : Roach()
        , m_logger(Logger::Create())
    {
        m_ctx = Context::Create(m_logger);
    }

    virtual boost::shared_ptr<Server> CreateServ(void)
    {
        m_logger->Log(Logger::Dbg, "Creating server");
        return Server::Create(m_ctx);
    }

    virtual boost::shared_ptr<Context> GetContext(void)
    {
        return m_ctx;
    }
};

Roach::Roach(void)
{
    /* NOP */
}

boost::shared_ptr<Roach> Roach::Create(void)
{
    return boost::make_shared<RoachImpl>();
}

} /* namespace roach */
