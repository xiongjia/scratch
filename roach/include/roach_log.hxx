/**
 *
 */

#ifndef _ROACH_LOG_HXX_
#define _ROACH_LOG_HXX_ 1

#include "boost/utility.hpp"
#include "boost/shared_ptr.hpp"
#include <functional>

namespace roach {

class Logger : boost::noncopyable
{
public:
    typedef enum _Mask
    {
        Err = (1 << 0),
        War = (1 << 1),
        Inf = (1 << 2),
        Dbg = (1 << 3),
    } Mask;

    typedef enum _Level
    {
        LevelNone = 0, LevelErr, LevelInf, LevelWar, LevelDbg, LevelAll
    } Level;

    typedef std::function<void(const char *log)> Handler;

public:
    static boost::shared_ptr<Logger> Create(Level logLevel = LevelNone,
                                            Handler handler = nullptr);

    virtual void SetHandler(Handler handler) = 0;
    virtual void SetLevel(Level logLevel) = 0;

    virtual void Log(Mask mask, const char *fmt, ...) = 0;
    virtual void LogNoFmt(Mask mask, const char *log) = 0;

protected:
    Logger(void);
};

} /* namespace roach */

#endif /* !defined(_ROACH_LOG_HXX_) */
