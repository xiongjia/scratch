/**
 * Roach - A simple HTTP Server. (libuv & http-parser & boost & C++11 & CMake)
 */


#include "boost/make_shared.hpp"
#include "roach_serv_handler.hxx"

namespace roach {

class HttpRequestImpl : public HttpRequest
{
public:
    boost::shared_ptr<HttpHeader> m_hdr;
    std::string m_url;
    std::string m_method;

public:
    HttpRequestImpl(void)
        : HttpRequest()
    {
        /* NOP */
    }

    virtual void SetHeader(boost::shared_ptr<HttpHeader> hdr)
    {
        m_hdr = hdr;
    }

    virtual boost::shared_ptr<HttpHeader> GetHeader(void)
    {
        return m_hdr;
    }

    virtual void SetUrl(const char *url)
    {
        m_url = url;
    }

    virtual const char* GetUrl(void) const
    {
        return m_url.c_str();
    }

    virtual void SetMethod(const char *method)
    {
        m_method = method;
    }

    virtual const char* GetMethod(void) const
    {
        return m_method.c_str();
    }
};

boost::shared_ptr<HttpRequest> HttpRequest::Create(void)
{
    return boost::make_shared<HttpRequestImpl>();
}

HttpRequest::HttpRequest(void)
{
    /* NOP */
}

HttpResponse::HttpResponse(void)
{
    /* NOP */
}

} /* namespace roach */
