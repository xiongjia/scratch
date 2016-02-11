/**
 * Chen - My Network protocol tests
 */

#include "chen_log.hxx"

_CHEN_BEGIN_

Log::Log(const char *src,
         size_t srcLine,
         Flags flags,
         const char *log)
    : m_src(src)
    , m_srcFilename(m_src.filename().string())
    , m_srcLine(srcLine)
    , m_flags(flags)
    , m_log(log)
    , m_tm(std::time(nullptr))
    , m_threadId(std::this_thread::get_id())
{
    /* NOP */
}

_CHEN_END_
