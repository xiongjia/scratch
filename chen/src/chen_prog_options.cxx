/**
* Chen - My Network protocol tests
*/

#include <string>
#include <sstream>

#include "boost/make_shared.hpp"
#include "boost/program_options.hpp"

#include "chen_prog_options.hxx"

namespace po = boost::program_options;

class OptionsDesc {
private:
  po::options_description options_desc_{ "MainOptions" };
  std::string parse_err_;

public:
  static boost::shared_ptr<OptionsDesc> GetInstance(void);

  OptionsDesc(void);

  bool Parse(int argc, const char **argv);
  const char *GetParseErr(void) const { return parse_err_.c_str(); }
};

boost::shared_ptr<OptionsDesc> OptionsDesc::GetInstance(void) {
  static boost::shared_ptr<OptionsDesc> instance = boost::make_shared<OptionsDesc>();
  return instance;
}

OptionsDesc::OptionsDesc(void) {
  options_desc_.add_options()
    ("help,h", "Print help messages")
    ("logLevel,l",
     po::value<std::string>()->default_value("all"),
     "log level");
}

bool OptionsDesc::Parse(int argc, const char **argv) {
  try {
    po::variables_map vm;
    po::store(po::parse_command_line(argc, argv, options_desc_), vm);
  } catch (po::error &e) {
    parse_err_ = e.what();
    return false;
  }
  return true;
}

class OptionsImpl : public chen::Options {
private:
  po::options_description options_desc_{ "MainOptions" };

  chen::Log::Level log_level_ = chen::Log::LevelAll;
  std::string parse_err_;
  std::string help_msg_;
  bool need_show_help_ = false;
  bool is_parse_err_ = true;

public:
  OptionsImpl(int argc, const char **argv);

  virtual const char *GetParseErr(void) const { return parse_err_.c_str(); }
  virtual bool IsParseErr(void) const { return is_parse_err_; }
  virtual bool NeedShowHelp(void) const { return need_show_help_; }

  virtual const chen::Log::Level GetLogLevel(void) const {
    return log_level_;
  }
  virtual const char *GetHelpMessage(void);
};
\
OptionsImpl::OptionsImpl(int argc, const char **argv) {
  options_desc_.add_options()
    ("help,h", "Print help messages")
    ("logLevel,l",
     po::value<std::string>()->default_value("all"),
     "log level");

  po::variables_map vm;
  try {
    po::store(po::parse_command_line(argc, argv, options_desc_), vm);
  } catch (po::error &e) {
    parse_err_ = e.what();
    need_show_help_ = true;
    is_parse_err_ = true;
    return;
  }
  if (vm.count("help")) {
    need_show_help_ = true;
  }
}

const char *OptionsImpl::GetHelpMessage(void) {
  if (!help_msg_.empty()) {
    return help_msg_.c_str();
  }
  std::stringstream ss;
  ss << options_desc_;
  help_msg_ = ss.str();
  return help_msg_.c_str();
}

_CHEN_BEGIN_

boost::shared_ptr<Options> Options::Create(int argc, const char **argv) {
  return boost::make_shared<OptionsImpl>(argc, argv);
}

_CHEN_END_
