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
    static int OnMsgBeginFunc(http_parser *parse)
    {
        HTTPParserImpl *self = static_cast<HTTPParserImpl*>(parse->data);
        self->OnMsgBegin(parse);
        return 0;
    }

    void OnMsgBegin(http_parser * /* parse */)
    {
        m_msgCompleted = false;
        m_url.clear();
        m_hdrCompleted = false;
        m_lastHdrField.clear();
        m_hdr->clear();
    }

    static int OnUrlFunc(http_parser *parse, const char *at, size_t len)
    {
        HTTPParserImpl *self = static_cast<HTTPParserImpl*>(parse->data);
        self->OnUrl(parse, at, len);
        return 0;
    }

    void OnUrl(http_parser * /* parse */, const char *at, size_t len)
    {
        m_url = std::string(at, len);
    }

    static int OnHdrFieldFunc(http_parser *parse, const char *at, size_t len)
    {
        HTTPParserImpl *self = static_cast<HTTPParserImpl*>(parse->data);
        self->OnHdrField(parse, at, len);
        return 0;
    }

    void OnHdrField(http_parser * /* parse */, const char *at, size_t len)
    {
        m_lastHdrField = std::string(at, len);
    }

    static int OnHdrValueFunc(http_parser *parse, const char *at, size_t len)
    {
        HTTPParserImpl *self = static_cast<HTTPParserImpl*>(parse->data);
        self->OnHdrValue(parse, at, len);
        return 0;
    };

    void OnHdrValue(http_parser * /* parse */, const char *at, size_t len)
    {
        if (m_lastHdrField.empty())
        {
            return;
        }
        std::string hdrVal = std::string(at, len);
        (*m_hdr)[m_lastHdrField] = hdrVal;
    }

    static int OnHdrCompleteFunc(http_parser *parse)
    {
        HTTPParserImpl *self = static_cast<HTTPParserImpl*>(parse->data);
        self->OnHdrComplete(parse);
        return 0;
    }

    void OnHdrComplete(http_parser * /* parse */)
    {
        m_hdrCompleted = true;
    }

    static int OnBodyFunc(http_parser *parse, const char *at, size_t len)
    {
        return 0;
    }

    static int OnMsgCompleteFunc(http_parser *parse)
    {
        HTTPParserImpl *self = static_cast<HTTPParserImpl*>(parse->data);
        self->OnMsgComplete(parse);
        return 1;
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
    , m_msgCompleted(false)
    , m_hdrCompleted(false)
    , m_hdr(boost::make_shared<HttpHeader>())
    , m_method(HTTP_GET)
{
    memset(&m_setting, 0, sizeof(m_setting));
    m_setting.on_message_begin = HTTPParserImpl::OnMsgBeginFunc;
    m_setting.on_url = HTTPParserImpl::OnUrlFunc;
    m_setting.on_header_field = HTTPParserImpl::OnHdrFieldFunc;
    m_setting.on_header_value = HTTPParserImpl::OnHdrValueFunc;
    m_setting.on_headers_complete = HTTPParserImpl::OnHdrCompleteFunc;
    m_setting.on_body = HTTPParserImpl::OnBodyFunc;
    m_setting.on_message_complete = HTTPParserImpl::OnMsgCompleteFunc;

    m_paser.data = this;
    http_parser_init(&m_paser, HTTP_REQUEST);
}

} /* namespace roach */
