/**
 * Roach - A simple HTTP Server. (libuv & http-parser & boost & C++11 & CMake)
 */

#include <stdio.h>
#include <signal.h>

#include <string>
#include "boost/make_shared.hpp"
#if defined(_MSC_VER)
#   pragma warning (push)
#   pragma warning (disable : 4819)
#   include "boost/format.hpp"
#   pragma warning (push)
#   pragma warning (default : 4819)
#else
#   include "boost/format.hpp"
#endif /* defined(_MSC_VER) */

#include "roach.hxx"

class ConsoleLogger : public roach::LoggerHandler
{
public:
    virtual void OnLog(const char *log)
    {
       fprintf(stdout, "%s\n", log);
    }
};

class Example : boost::noncopyable
{
private:
    boost::shared_ptr<roach::Roach>  m_roach;
    boost::shared_ptr<roach::Logger> m_logger;
    boost::shared_ptr<roach::Server> m_srv;

public:
    Example(void)
        : m_roach(roach::Roach::Create())
        , m_logger(m_roach->GetContext()->GetLogger())
    {
        m_logger->SetLevel(roach::Logger::LevelDbg);
        m_logger->RegisterHandler(boost::make_shared<ConsoleLogger>());
    }

    void Run(void)
    {
        m_srv = m_roach->CreateServ();
        m_srv->SetHandler([](boost::shared_ptr<roach::HttpRequest> req,
                             boost::shared_ptr<roach::HttpResponse> rep)
                             -> bool {
            /* the response content */
            auto fmt = boost::format("Method = %s; Url = %s\r\n");
            auto content = boost::str(fmt % req->GetMethod() % req->GetUrl());
            /* write the 200 response to client */
            rep->WritePlainText(200, "OK", content.c_str());
            return true;
        });
        m_srv->Run("0.0.0.0", 8999);
    }

    void Stop(void)
    {
        if (m_srv)
        {
            m_srv->Stop();
        }
    }
};

int main(int argc, char **argv)
{
    static boost::shared_ptr<Example> EXAMPLE;

    signal(SIGINT, [](int sig) {
        if (EXAMPLE)
        {
            EXAMPLE->Stop();
        }
    });

    EXAMPLE = boost::make_shared<Example>();
    EXAMPLE->Run();
    return 0;
}
