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
#if 0
    auto log = chen::Log::get_instance();
    auto oldHandler = log->get_handler();
    auto oldLevel = log->get_level();

    const boost::thread::id curThreadId = boost::this_thread::get_id();
    std::string lastLog;
    chen::Log::Flags lastFlags;
    boost::thread::id lastThreadId;
    boost::posix_time::ptime lastLogTM;
    log->set_handler([&](const chen::LogItem &logItem)
    {
        lastLog = logItem.get_log();
        lastFlags = logItem.get_flags();
        lastThreadId = logItem.get_threadid();
        lastLogTM = logItem.get_tm();
    });

    auto now = boost::posix_time::microsec_clock::universal_time();
    log->set_level(chen::Log::LEVEL_ERR);
    const char *logFmt = "Test Log %d, %s";
    CHEN_LOG(CHEN_LOG_ERR, logFmt, 1, "string");
    BOOST_REQUIRE(lastFlags == chen::Log::FErr &&
                  lastThreadId == curThreadId &&
                  now <= lastLogTM);
    boost::format fmt(logFmt);
    const std::string logMsg = boost::str(fmt % 1 % "string");
    BOOST_REQUIRE_EQUAL(logMsg, lastLog);

    lastLog.clear();
    log->set_level(chen::Log::LEVEL_ERR);
    const char *infLog = "info log";
    CHEN_LOG(CHEN_LOG_INF, infLog);
    BOOST_REQUIRE(lastLog.empty());

    log->set_level(chen::Log::LEVEL_ALL);
    CHEN_LOG(CHEN_LOG_INF, infLog);
    BOOST_REQUIRE_EQUAL(lastLog, infLog);
    BOOST_REQUIRE(lastFlags == chen::Log::FInf &&
                  lastThreadId == curThreadId);

    /* restore log leve & handler */
    log->set_level(oldLevel);
    log->set_handler(oldHandler);
#endif
}
BOOST_AUTO_TEST_SUITE_END()
