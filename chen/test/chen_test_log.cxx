/**
 * Chen - My Network protocol tests
 */

#include "boost/test/unit_test.hpp"
#pragma warning(push)
#   pragma warning(disable:4819)
#   include "boost/format.hpp"
#pragma warning(pop) 
#include "chen_log.hxx"

BOOST_AUTO_TEST_SUITE(chen)
BOOST_AUTO_TEST_CASE(log)
{
    BOOST_TEST_MESSAGE("Che - Log");
    auto log = chen::Log::get_instance();
    auto oldHandler = log->get_handler();
    auto oldLevel = log->get_level();

    std::string lastLog;
    chen::Log::Flags lastFlag;
    log->set_handler([&](const chen::Log::LogItem &logItem) {
        lastLog = logItem.log;
        lastFlag = logItem.flags;
    });

    log->set_level(chen::Log::LEVEL_ERR);
    const char *logFmt = "Test Log %d, %s";
    CHEN_LOG(CHEN_LOG_ERR, logFmt, 1, "string");
    BOOST_REQUIRE_EQUAL(lastFlag, chen::Log::FErr);
    boost::format fmt(logFmt);
    const std::string name = boost::str(fmt % 1 % "string");
    BOOST_REQUIRE_EQUAL(name, lastLog);

    lastLog.clear();
    log->set_level(chen::Log::LEVEL_ERR);
    const char *infLog = "info log";
    CHEN_LOG(CHEN_LOG_INF, infLog);
    BOOST_REQUIRE(lastLog.empty());

    log->set_level(chen::Log::LEVEL_ALL);
    CHEN_LOG(CHEN_LOG_INF, infLog);
    BOOST_REQUIRE_EQUAL(lastLog, infLog);

    /* restore log leve & handler */
    log->set_level(oldLevel);
    log->set_handler(oldHandler);
}
BOOST_AUTO_TEST_SUITE_END()
