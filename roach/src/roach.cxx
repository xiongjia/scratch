/**
 *
 */

#include "roach.hxx"

namespace roach {

class RoachImpl : public Roach
{
private:
    boost::shared_ptr<Logger> m_logger;

public:
    RoachImpl(Logger::Level level, Logger::Handler handler)
        : Roach()
        , m_logger(Logger::Create(level, handler))
    {
        /* NOP */
    }

    virtual void SetLogger(Logger::Level level, Logger::Handler handler)
    {
        m_logger->SetLevel(level);
        m_logger->SetHandler(handler);
    }

    virtual boost::shared_ptr<Server> CreateServ(void)
    {
        m_logger->Log(Logger::Dbg, "Creating server");
        return Server::Create(m_logger);
    }
};

Roach::Roach(void)
{
    /* NOP */
}

boost::shared_ptr<Roach> Roach::Create(Logger::Level level,
                                       Logger::Handler handler)
{
    return boost::shared_ptr<Roach>(new RoachImpl(level, handler));
}

} /* namespace roach */
