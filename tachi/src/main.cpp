/**
 * 
 */


#include <stdio.h>

#include "tachi-log.hxx"

#include <boost/log/core.hpp>
#include <boost/log/expressions.hpp>
#include <boost/log/sources/logger.hpp>
#include <boost/log/utility/setup/file.hpp>
#include <boost/log/utility/setup/console.hpp>
#include <boost/log/utility/setup/common_attributes.hpp>
#include <boost/log/support/date_time.hpp>
#include <boost/log/sinks/sync_frontend.hpp>
#include <boost/log/sinks/text_file_backend.hpp>
#include <boost/log/sinks/text_ostream_backend.hpp>


static void init_log(void) {
  boost::log::add_common_attributes();
  boost::log::core::get()->add_global_attribute("Scope", boost::log::attributes::named_scope());
  boost::log::core::get()->set_filter(boost::log::trivial::severity >= boost::log::trivial::trace);

  auto fmtTimeStamp = boost::log::expressions::format_date_time<boost::posix_time::ptime>("TimeStamp", "%Y-%m-%d %H:%M:%S.%f");
  auto fmtThreadId = boost::log::expressions::attr<boost::log::attributes::current_thread_id::value_type>("ThreadID");
  auto fmtSeverity = boost::log::expressions:: attr<boost::log::trivial::severity_level>("Severity");
  auto fmtScope = boost::log::expressions::format_named_scope("Scope",
    boost::log::keywords::format = "%n(%f:%l)",
    boost::log::keywords::iteration = boost::log::expressions::reverse,
    boost::log::keywords::depth = 2);
  boost::log::formatter logFmt = boost::log::expressions::format("[%1%] (%2%) [%3%] [%4%] %5%")
    % fmtTimeStamp % fmtThreadId % fmtSeverity % fmtScope % boost::log::expressions::smessage;

  auto consoleSink = boost::log::add_console_log(std::clog);
  consoleSink->set_formatter(logFmt);
}

static void Test(void)
{
  TACHI_LOG_FUNCTION();
  TACHI_LOG(info) << "Info Log in Test()";
}

int main(const int argc, const char **argv) {
  TACHI_LOG_NAMED_SCOPE("main");
  TACHI_LOG(info) << "Info Log";
  TACHI_LOG(error) << "Info Log";

  Test();
  return 0;
}
