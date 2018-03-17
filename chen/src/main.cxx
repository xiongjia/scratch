/**
 * Chen - My Network protocol tests
 */

#include "boost/asio.hpp"
#include "boost/bind.hpp"
#include "boost/utility.hpp"
#include "boost/asio/thread_pool.hpp"
#include "boost/thread.hpp"
#include "chen_serv.hxx"
#include "chen_log.hxx"

class Chen : boost::noncopyable {
private:
  void OnSignals(const boost::system::error_code &err,
                 const int signal_num) {
    if (!err) {
      CHEN_LOG_DEBUG("Recive signale: %d (NO ERR)", signal_num);
    } else {
      CHEN_LOG_DEBUG("Recive signale: %d, ErrCode: %d, %s",
                     signal_num, err.value(), err.message().c_str());
    }
  }

public:
  Chen(void) {
    // boost::asio::signal_set signals(io_svc_, SIGINT, SIGTERM);
  }

  int Run(void) {
    CHEN_LOG_INFO("creating server ...");

    boost::asio::io_service io_svc_;

    boost::asio::signal_set signals(io_svc_, SIGINT);
    signals.async_wait(boost::bind(&Chen::OnSignals, this, _1, _2));

    std::thread work_thread([&io_svc_]() { io_svc_.run(); });
    work_thread.join();

    /// io_svc_.run();
    return 0;
  }
};

int main(int /* argc */, char ** /* argv */) {
  chen::Log::GetInstance()->SetLevel(chen::Log::LevelAll);
  Chen chen;
  return chen.Run();
}
