/**
 * Chen - My Network protocol tests
 */

#include <iostream>
#include "boost/utility.hpp"
#include "chen_serv.hxx"
#include "chen_log.hxx"

static void Logger(const char *src, size_t line,
                   chen::Log::Flags /* flags */, const char *msg) {
  std::cout << "[" << src << ":" << line << "] " << msg << std::endl;
}

class Chen : boost::noncopyable {
public:
  int Run(void) {
    CHEN_LOG_INFO("creating server ...");
    return 0;
  }
};

int main(int /* argc */, char ** /* argv */) {
  chen::Log::GetInstance()->SetHandler(Logger);
  chen::Log::GetInstance()->SetLevel(chen::Log::LevelAll);
  Chen chen;
  return chen.Run();
}
