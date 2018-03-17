/**
 * Chen - My Network protocol tests
 */

#include "boost/test/unit_test.hpp"
#include "boost/format.hpp"

#include "chen_log.hxx"

BOOST_AUTO_TEST_SUITE(chen)
BOOST_AUTO_TEST_CASE(log)
{
    BOOST_TEST_MESSAGE("Che - Log");
    auto log = chen::Log::GetInstance();
    auto oldHandler = log->GetHandler();
    std::string lastLog;
    chen::Log::Flags lastFlags = chen::Log::Flags::Debug;
    log->SetHandler([&](const char *, size_t,
                        chen::Log::Flags flags, const char *msg) {
      lastLog = msg;
      lastFlags = flags;
    });
    const char *fmt = "Test Log %d %s";
    log->SetLevel(chen::Log::LevelAll);
    CHEN_LOG_DEBUG(fmt, 1, "string");
    BOOST_REQUIRE(lastFlags == chen::Log::Debug);
    BOOST_REQUIRE_EQUAL(boost::str(boost::format(fmt) % 1 % "string"),
                        lastLog);
}
BOOST_AUTO_TEST_SUITE_END()
