/**
* Roach - A simple HTTP Server. (libuv & http-parser & boost & C++11 & CMake)
*/

#include <string>

#include "boost/make_shared.hpp"
#include "http_parser.h"

#include "roach_http_parser.hxx"

namespace roach {

class HTTPParserImpl : public HTTPParser
{
private:
    http_parser_settings m_setting;
    http_parser m_paser;

public:
    HTTPParserImpl(void)
        : HTTPParser()
    {
        memset(&m_setting, 0, sizeof(m_setting));
        m_setting.on_url = [](http_parser *parse, const char *at, size_t length) -> int {
            std::string url;
            url.resize(length + 1);
            strncpy_s(&url[0], url.size(), at, length);
            url[length] = '\0';
            return 0;
        };
        http_parser_init(&m_paser, HTTP_REQUEST);
    }

    virtual bool Parse(const char *buf, const size_t bufLen)
    {
        size_t  parsed = http_parser_execute(&m_paser, &m_setting, buf, bufLen);
        return parsed < bufLen ? false : true;
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

} /* namespace roach */
