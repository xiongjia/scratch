/**
* Roach - A simple HTTP Server. (libuv & http-parser & boost & C++11 & CMake)
*/

#ifndef _ROACH_SERV_HANDLER_HXX_
#define _ROACH_SERV_HANDLER_HXX_ 1

#include <functional>
#include <string>
#include <map>

#include "boost/utility.hpp"
#include "boost/shared_ptr.hpp"

namespace roach {

/* Http Header */
typedef std::map<std::string, std::string> HttpHeader;

class HttpRequest : boost::noncopyable
{
public:
    static boost::shared_ptr<HttpRequest> Create(void);

    virtual void SetHeader(boost::shared_ptr<HttpHeader> hdr) = 0;
    virtual boost::shared_ptr<HttpHeader> GetHeader(void) = 0;

    virtual void SetUrl(const char *url) = 0;
    virtual const char* GetUrl(void) const = 0;

    virtual void SetMethod(const char *method) = 0;
    virtual const char* GetMethod(void) const = 0;

protected:
    HttpRequest(void);
};

class HttpResponse : boost::noncopyable
{
public:
    virtual void WritePlainText(const unsigned short statusCode,
                                const char *status,
                                const char *content) = 0;

protected:
    HttpResponse(void);
};

typedef std::function<bool(boost::shared_ptr<HttpRequest>,
                           boost::shared_ptr<HttpResponse>)> ServHandler;


} /* namespace roach */

#endif /* !defined(_ROACH_SERV_HANDLER_HXX_) */
