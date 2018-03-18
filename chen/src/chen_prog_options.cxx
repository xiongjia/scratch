/**
* Chen - My Network protocol tests
*/

#include <string>

#include "boost/make_shared.hpp"
#include "boost/program_options.hpp"

#include "chen_prog_options.hxx"

namespace po = boost::program_options;

class OptionsDesc {
private:
  po::options_description options_desc_{ "MainOptions" };

public:
  static boost::shared_ptr<OptionsDesc> GetInstance(void);

  OptionsDesc(void) {
    options_desc_.add_options()
      ("help,h", "Print help messages")
      ("logLevel,l", po::value<std::string>()->default_value("all"), "log level");
  }
};

class OptionsImpl : public chen::Options {
private:
  
  chen::Log::Level log_level_ = chen::Log::LevelAll;

public:
  OptionsImpl(void);

  virtual const chen::Log::Level GetLogLevel(void) const {
    return log_level_;
  }
};

boost::shared_ptr<OptionsDesc> OptionsDesc::GetInstance(void) {
  static boost::shared_ptr<OptionsDesc> instance = boost::make_shared<OptionsDesc>();
  return instance;
}

OptionsImpl::OptionsImpl(void) {

}

_CHEN_BEGIN_

boost::shared_ptr<Options> Options::Create(int, const char **) {
  return boost::make_shared<OptionsImpl>();
}

_CHEN_END_
