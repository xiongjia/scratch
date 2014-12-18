/**
 *
 */

#include "roach.hxx"
#include "boost/make_shared.hpp"

class ConsoleLogger : public roach::LoggerHandler
{
public:
    virtual void OnLog(const char *log)
    {
       fprintf(stdout, "%s\n", log);
    }
};

int main(int argc, char **argv)
{
    auto roach = roach::Roach::Create();
    auto logger = roach->GetContext()->GetLogger();
    logger->SetLevel(roach::Logger::LevelDbg);
    logger->RegisterHandler(boost::make_shared<ConsoleLogger>());

    auto serv = roach->CreateServ();
    serv->Run("0.0.0.0", 8999);
    return 0;
}
