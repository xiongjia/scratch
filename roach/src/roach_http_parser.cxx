/**
 * Roach - A simple HTTP Server. (libuv & http-parser & boost & C++11 & CMake)
 */

#include "boost/make_shared.hpp"
#include "http_parser.h"

#include "roach_http_parser.hxx"
#include "roach_misc.hxx"

namespace roach {

class HTTPParserImpl : public HTTPParser
{
private:
    http_parser_settings m_setting;
    http_parser m_paser;

    bool m_msgCompleted;
    std::string m_url;

    bool m_hdrCompleted;
    std::string m_lastHdrField;
    boost::shared_ptr<HttpHeader> m_hdr;
    http_method m_method;

public:
    void OnMsgBegin(http_parser * /* parse */)
    {
        m_msgCompleted = false;
        m_url.clear();
        m_hdrCompleted = false;
        m_lastHdrField.clear();
        m_hdr->clear();
    }

    void OnUrl(http_parser * /* parse */, const char *at, size_t len)
    {
        m_url = std::string(at, len);
    }

    void OnHdrField(http_parser * /* parse */, const char *at, size_t len)
    {
        m_lastHdrField = std::string(at, len);
    }

    void OnHdrValue(http_parser * /* parse */, const char *at, size_t len)
    {
        if (m_lastHdrField.empty())
        {
            return;
        }
        std::string hdrVal = std::string(at, len);
        (*m_hdr)[m_lastHdrField] = hdrVal;
    }

    void OnHdrComplete(http_parser * /* parse */)
    {
        m_hdrCompleted = true;
    }

    void OnMsgComplete(http_parser *parse)
    {
        m_method = static_cast<http_method>(parse->method);
        m_msgCompleted = true;
    }

public:
    HTTPParserImpl(void);

    virtual bool Parse(const char *buf, const size_t bufLen)
    {
        size_t parsed = http_parser_execute(&m_paser, &m_setting, buf, bufLen);
        return parsed < bufLen ? false : true;
    }

    virtual bool IsCompleted(void) const
    {
        return m_hdrCompleted && m_msgCompleted;
    }

    virtual const char* GetMethod(void) const
    {
        return http_method_str(m_method);
    }

    virtual const char* GetUrl(void) const
    {
        return m_url.c_str();
    }

    virtual boost::shared_ptr<HttpHeader> GetHeader(void)
    {
        return m_hdr;
    }
};

boost::shared_ptr<HTTPParser> HTTPParser::Create(void)
{
    return boost::make_shared<HTTPParserImpl>();
}

HTTPParser::HTTPParser(void)
{
    /* NOP */
}

HTTPParserImpl::HTTPParserImpl(void)
    : HTTPParser()
    , m_hdrCompleted(false)
    , m_msgCompleted(false)
    , m_method(HTTP_GET)
    , m_hdr(boost::make_shared<HttpHeader>())
{
    memset(&m_setting, 0, sizeof(m_setting));
    m_setting.on_message_begin = [](http_parser *parse) -> int {
        HTTPParserImpl *self = static_cast<HTTPParserImpl*>(parse->data);
        self->OnMsgBegin(parse);
        return 0;
    };
    m_setting.on_url = [](http_parser *parse,
                          const char *at, size_t len) -> int {
        HTTPParserImpl *self = static_cast<HTTPParserImpl*>(parse->data);
        self->OnUrl(parse, at, len);
        return 0;
    };
    m_setting.on_header_field = [](http_parser *parse,
                                   const char *at, size_t len) -> int {
        HTTPParserImpl *self = static_cast<HTTPParserImpl*>(parse->data);
        self->OnHdrField(parse, at, len);
        return 0;
    };
    m_setting.on_header_value = [](http_parser *parse,
                                   const char *at, size_t len) -> int {
        HTTPParserImpl *self = static_cast<HTTPParserImpl*>(parse->data);
        self->OnHdrValue(parse, at, len);
        return 0;
    };
    m_setting.on_headers_complete = [](http_parser *parse) -> int {
        HTTPParserImpl *self = static_cast<HTTPParserImpl*>(parse->data);
        self->OnHdrComplete(parse);
        return 0;
    };
    m_setting.on_body = [](http_parser *parse,
                           const char *at, size_t len) -> int {
        /* TODO: on_body recive the data of post request */
        return 0;
    };
    m_setting.on_message_complete = [](http_parser *parse) -> int {
        HTTPParserImpl *self = static_cast<HTTPParserImpl*>(parse->data);
        self->OnMsgComplete(parse);
        return 1;
    };

    m_paser.data = this;
    http_parser_init(&m_paser, HTTP_REQUEST);
}

} /* namespace roach */
