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
    roach->SetLevel(roach::Logger::LevelDbg);
    roach->RegisterHandler(boost::make_shared<ConsoleLogger>());

    auto serv = roach->CreateServ();
    serv->Listen("0.0.0.0", 8999);
    serv->Run();
    return 0;
}
