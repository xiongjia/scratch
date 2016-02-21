/**
 * Chen - My Network protocol tests
 */

#include <string>
#include <ctime>
#include "boost/test/unit_test.hpp"
#include "boost/make_shared.hpp"
#include "chen_logger.hxx"
#include "chen_log_handler.hxx"

class LogTester : public chen::LoggerHandler
{
public:
    chen::Log::Flags m_lastFlags;
    std::time_t      m_lastTM;
    std::string      m_lastMsg;

public:
    static boost::shared_ptr<LogTester> create(void)
    {
        return boost::make_shared<LogTester>();
    }

public:
    LogTester(void)
        : chen::LoggerHandler()
        , m_lastFlags(chen::Log::None)
        , m_lastTM(0)
    {
        /* NOP */
    }

    virtual void append(const chen::Log &log)
    {
        m_lastFlags = log.get_flags();
        m_lastMsg = log.get_log();
        m_lastTM = log.get_tm();
    }
};

BOOST_AUTO_TEST_SUITE(chen)
BOOST_AUTO_TEST_CASE(log)
{
    BOOST_TEST_MESSAGE("Chen - Log");
    auto logger = chen::Logger::create();
    logger->set_log_level(chen::Logger::LevelAll);
    auto testHandler = LogTester::create();
    logger->reg_handler(testHandler);

    logger->log(__FILE__, __LINE__, chen::Log::Err | chen::Log::Dbg,
                "Test Log %s", "my test1");
    BOOST_CHECK_EQUAL(testHandler->m_lastMsg, "Test Log my test1");
    BOOST_CHECK_EQUAL(testHandler->m_lastFlags,
                      chen::Log::Err | chen::Log::Dbg);
}
BOOST_AUTO_TEST_SUITE_END()
