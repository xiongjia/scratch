/**
 * Chen - My Network protocol tests
 */

#ifndef _CHEN_LOG_HXX_
#define _CHEN_LOG_HXX_ 1

#include "boost/utility.hpp"
#include "boost/filesystem.hpp"
#include <thread>
#include <ctime>

#include "chen_types.hxx"

_CHEN_BEGIN_

class Log : boost::noncopyable
{
public:
    typedef enum {
        Err = (1 << 0),
        War = (1 << 1),
        Inf = (1 << 2),
        Dbg = (1 << 3)
    } Flags;

private:
    boost::filesystem::path  m_src;
    std::string              m_srcFilename;
    const size_t             m_srcLine;
    Flags                    m_flags;
    const char              *m_log;
    std::time_t              m_tm;
    std::thread::id          m_threadId;

public:
    Log(const char *src,
        size_t srcLine,
        Flags flags,
        const char *log);

    const char *get_src(void) const
    {
        return m_srcFilename.c_str();
    }

    const size_t get_src_line(void) const
    {
        return m_srcLine;
    }

    const Flags get_flags(void) const
    {
        return m_flags;
    }

    const char *get_log(void) const
    {
        return m_log;
    }

    const std::time_t get_tm(void) const
    {
        return m_tm;
    }

    const std::thread::id& get_thread_id(void) const
    {
        return m_threadId;
    }
};

_CHEN_END_
#endif /* !defined(_CHEN_LOG_HXX_) */
