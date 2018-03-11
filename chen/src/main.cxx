/**
 * Chen - My Network protocol tests
 */

#include <iostream>

#include "boost/asio.hpp"
#include "boost/bind.hpp"
#include "boost/utility.hpp"

#include "chen_serv.hxx"
#include "chen_log.hxx"

static void Logger(const char *src, size_t line,
                   chen::Log::Flags /* flags */, const char *msg) {
  std::cout << "[" << src << ":" << line << "] " << msg << std::endl;
}

class Chen : boost::noncopyable {
private:
  boost::asio::io_service io_svc_;

  void OnSignals(const boost::system::error_code &err_code,
                 const int signal_num) {
    CHEN_LOG_DEBUG("Recive signale: %d, ErrCode: %d",
                   signal_num, err_code.value());
  }

public:
  Chen(void) {
    boost::asio::signal_set signals(io_svc_, SIGINT, SIGTERM);
    signals.async_wait(boost::bind(&Chen::OnSignals, this, _1, _2));
  }

  int Run(void) {
    CHEN_LOG_INFO("creating server ...");
    io_svc_.run();
    return 0;
  }
};

int main(int /* argc */, char ** /* argv */) {

  chen::Log::GetInstance()->SetHandler(Logger);
  chen::Log::GetInstance()->SetLevel(chen::Log::LevelAll);
  Chen chen;
  return chen.Run();
}
