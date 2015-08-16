/**
 * Zeratul - A simple socks proxy. (Boost ASIO)
 */

#include "boost/make_shared.hpp"
#include "boost/log/core.hpp"
#include "boost/log/sources/logger.hpp"
#include "boost/log/utility/setup/common_attributes.hpp"
#include "boost/log/utility/setup/file.hpp"
#include "boost/log/utility/setup/console.hpp"
#include "boost/log/expressions.hpp"
#include "boost/log/support/date_time.hpp"
#include "boost/date_time/posix_time/posix_time_types.hpp"
#include "boost/log/sinks/sync_frontend.hpp"
#include "boost/log/sinks/text_file_backend.hpp"
#include "boost/log/sinks/text_ostream_backend.hpp"

#include "z_misc.hxx"

namespace logger      = boost::log;
namespace log_attr    = boost::log::attributes;
namespace log_trivial = boost::log::trivial;
namespace log_expr    = boost::log::expressions;
namespace log_kwords  = boost::log::keywords;
namespace posix_tm    = boost::posix_time;

_ZERATUL_BEGIN_

class LogMgrImpl : public LogMgr
{
private:
    bool m_init;

public:
    LogMgrImpl(void);

    virtual void Initialize(void);
};

LogMgrImpl::LogMgrImpl(void)
    : LogMgr()
    , m_init(false)
{
    /* NOP */
}

void LogMgrImpl::Initialize(void)
{
    if (m_init)
    {
        return;
    }
    m_init = true;

    /* init boost log
     * 1. Add common attributes
     * 2. set log filter to trace
     */
    logger::add_common_attributes();
    auto logCore = logger::core::get();
    logCore->add_global_attribute("Scope", log_attr::named_scope());
    logCore->set_filter(log_trivial::severity >= log_trivial::trace);

    /* log formatter (console) :
     * [Severity Level] (ThreadId) [Scope] Log message
     *
     * log formatter (file) :
     * [TimeStamp] [ThreadId] [Severity Level] [Scope] Log message
     */
    auto fmtTM = log_expr::format_date_time<posix_tm::ptime>("TimeStamp",
        "%Y-%m-%d %H:%M:%S.%f");
    auto fmtThreadId = log_expr::
        attr<log_attr::current_thread_id::value_type>("ThreadID");
    auto fmtSeverity = log_expr::attr<log_trivial::severity_level>("Severity");
    auto fmtScope = log_expr::format_named_scope("Scope",
        log_kwords::format = "%n(%F:%l)",
        log_kwords::iteration = log_expr::reverse,
        log_kwords::depth = 2);

    auto consoleLogFmt = log_expr::format("[%1%] (%2%) [%3%] %4%")
        % fmtSeverity % fmtThreadId % fmtScope
        % boost::log::expressions::smessage;

    /* console sink */
    auto consoleSink = boost::log::add_console_log(std::clog);
    consoleSink->set_formatter(consoleLogFmt);

    /* TODO Adding FS Sink */
}

LogMgr::LogMgr(void)
{
    /* NOP */
}

boost::shared_ptr<LogMgr> LogMgr::Instance(void)
{
    static boost::shared_ptr<LogMgr> inst = boost::make_shared<LogMgrImpl>();
    return inst;
}

_ZERATUL_END_
