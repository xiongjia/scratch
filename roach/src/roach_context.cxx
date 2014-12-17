/**
*
*/

#include "roach_context.hxx"
#include "boost/make_shared.hpp"

namespace roach {

class ContextImpl : public Context
{
private:
    boost::shared_ptr<Logger> m_logger;

public:
    ContextImpl(boost::shared_ptr<Logger> logger)
        : Context()
        , m_logger(logger)
    {
        /* NOP */
    }

    virtual boost::shared_ptr<Logger> GetLogger(void)
    {
        return m_logger;
    }
};

Context::Context(void)
{
    /* NOP */
}

boost::shared_ptr<Context> Context::Create(boost::shared_ptr<Logger> logger)
{
    return boost::make_shared<ContextImpl>(logger);
}

} /* namespace roach */
