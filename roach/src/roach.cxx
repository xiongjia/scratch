/**
 *
 */

#include "roach.hxx"
#include "boost/make_shared.hpp"

namespace roach {

class RoachImpl : public Roach
{
private:
    boost::shared_ptr<Logger> m_logger;

public:
    RoachImpl()
        : Roach()
        , m_logger(Logger::Create())
    {
        /* NOP */
    }

    virtual boost::shared_ptr<Server> CreateServ(void)
    {
        m_logger->Log(Logger::Dbg, "Creating server");
        return Server::Create(m_logger);
    }

    virtual void SetLevel(Logger::Level logLevel)
    {
        m_logger->SetLevel(logLevel);
    }

    virtual void RegisterHandler(boost::shared_ptr<LoggerHandler> handler)
    {
        m_logger->RegisterHandler(handler);
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
