/**
 * 
 */


#include <boost/log/core.hpp>
#include <boost/log/expressions.hpp>
#include <boost/log/utility/setup/common_attributes.hpp>
#include <boost/log/support/date_time.hpp>
#include <boost/log/utility/setup/console.hpp>
#include "tachi-log.hxx"

TachiLog::TachiLog() {
  // NOP
}

void TachiLog::init() {
  boost::log::add_common_attributes();
  auto core = boost::log::core::get();
  core->add_global_attribute("Scope", boost::log::attributes::named_scope());
  core->set_filter(boost::log::trivial::severity >= boost::log::trivial::trace);

  auto fmtTimeStamp = boost::log::expressions::format_date_time<boost::posix_time::ptime>("TimeStamp", "%Y-%m-%d %H:%M:%S.%f");
  auto fmtThreadId = boost::log::expressions::attr<boost::log::attributes::current_thread_id::value_type>("ThreadID");
  auto fmtSeverity = boost::log::expressions::attr<boost::log::trivial::severity_level>("Severity");
  auto fmtScope = boost::log::expressions::format_named_scope("Scope",
    boost::log::keywords::format = "%n(%f:%l)",
    boost::log::keywords::iteration = boost::log::expressions::reverse,
    boost::log::keywords::depth = 2);

  boost::log::formatter logFmt = boost::log::expressions::format("[%1%] (%2%) [%3%] [%4%] %5%")
    % fmtTimeStamp % fmtThreadId % fmtSeverity % fmtScope % boost::log::expressions::smessage;
  auto consoleSink = boost::log::add_console_log(std::clog);
  consoleSink->set_formatter(logFmt);
}