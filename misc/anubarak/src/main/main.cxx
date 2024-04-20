/**
 */

#include <iostream>

#include <boost/make_shared.hpp>
#include <boost/program_options.hpp>
namespace po = boost::program_options;

#include "anubarak.h"
#include "ab_logger.h"

class ConsoleLogger : public ab::LoggerHandler
{
public:
    virtual void OnLog(const char *src,
                       const int srcLine,
                       ab::LoggerMask mask,
                       const char *log)
    {
        /* [<src>:<srcLine>] [0x<mask>]  <log> */
        std::cout << "[" << src << ":" << std::oct << srcLine << "] "
                  << "[0x" << std::hex << mask << "] "
                  << (log == NULL ? "" : log) << std::endl;
    }
};

int main(int argc, char **argv)
{
    /* check program options */
    po::options_description mainOptDesc("Anubarak");
    mainOptDesc.add_options()
        ("luaFile,l",
            po::value<std::string>()->default_value("core/anubarak.lua"),
            "The test LUA script filename.")
        ("help,h", "Print help messages");

    po::variables_map varMap;
    try
    {
        po::store(po::parse_command_line(argc, argv, mainOptDesc), varMap);
        po::notify(varMap);
    }
    catch (po::error &optErr)
    {
        /* Invalid options */
        std::cerr << "ERROR: " << optErr.what() << std::endl << std::endl;
        std::cout << mainOptDesc << std::endl;
        return 0;
    }

    if (varMap.count("help"))
    {
        /* print usage */
        std::cout << mainOptDesc << std::endl;
        return 0;
    }

    /* initialize logger */
    auto logger = ab::Logger::instance();
    logger->RegisterHandler(boost::make_shared<ConsoleLogger>());
    logger->SetLogLevel(ab::Logger::Level::All);


    /* create Anubarak */
    AB_LOG(AB_LOG_DBG, "Creating Anubarak");
    auto anubarak = ab::Anubarak::Create();

    const std::string &testLuaFile = varMap["luaFile"].as<std::string>();
    AB_LOG(AB_LOG_DBG, "The testing lua file %s", testLuaFile.c_str());
    anubarak->Run(testLuaFile.c_str());
    return 0;
}
