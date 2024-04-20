/**
 * Roach - A simple HTTP Server. (libuv & http-parser & boost & C++11 & CMake)
 */

#ifndef _ROACH_HTTP_PARSER_HXX_
#define _ROACH_HTTP_PARSER_HXX_ 1

#include <string>
#include <map>

#include "boost/utility.hpp"
#include "boost/shared_ptr.hpp"
#include "roach_serv_handler.hxx"

namespace roach {

class HTTPParser : boost::noncopyable
{
public:
    static boost::shared_ptr<HTTPParser> Create(void);

    virtual bool Parse(const char *buf, const size_t bufLen) = 0;

    virtual bool IsCompleted(void) const = 0;

    virtual const char* GetMethod(void) const = 0;
    virtual const char* GetUrl(void) const = 0;
    virtual boost::shared_ptr<HttpHeader> GetHeader(void) = 0;

protected:
    HTTPParser(void);
};



} /* namespace roach */

#endif /* !defined(_ROACH_HTTP_PARSER_HXX_) */
