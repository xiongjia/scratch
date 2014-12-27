/**
 * Roach - A simple HTTP Server. (libuv & http-parser & boost & C++11 & CMake)
 */

#ifndef _ROACH_HTTP_PARSER_HXX_
#define _ROACH_HTTP_PARSER_HXX_ 1

#include "boost/utility.hpp"
#include "boost/shared_ptr.hpp"

namespace roach {

class HTTPParser : boost::noncopyable
{
public:
    static boost::shared_ptr<HTTPParser> Create(void);

    virtual bool Parse(const char *buf, const size_t bufLen) = 0;

protected:
    HTTPParser(void);
};



} /* namespace roach */

#endif /* !defined(_ROACH_HTTP_PARSER_HXX_) */
