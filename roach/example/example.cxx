/**
 *
 */

#include "roach.hxx"

static void Logger(const char *msg)
{
    fprintf(stdout, "%s\n", msg);
}

int main(int argc, char **argv)
{
    auto roach = roach::Roach::Create(roach::Logger::LevelDbg, Logger);
    auto serv = roach->CreateServ();

    serv->Listen("0.0.0.0", 8999);
    serv->Run();

    return 0;
}
