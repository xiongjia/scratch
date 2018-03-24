/**
* Chen - My Network protocol tests
*/

#include <string>
#include <sstream>
#include <map>

#include "boost/algorithm/string/trim.hpp"
#include "boost/algorithm/string/case_conv.hpp"
#include "boost/make_shared.hpp"
#include "boost/program_options.hpp"
#include "boost/assign/list_of.hpp"

#include "chen_prog_options.hxx"

namespace po = boost::program_options;

static const std::map<std::string, chen::Log::Level> LEVELS = boost::assign::map_list_of
  ("none",  chen::Log::LevelNone)
  ("error", chen::Log::LevelError)
  ("info",  chen::Log::LevelInfo)
  ("warn",  chen::Log::LevelWarn)
  ("debug", chen::Log::LevelDebug)
  ("all",   chen::Log::LevelAll);

class OptionsImpl : public chen::Options {
private:
  chen::Log::Level log_level_ = chen::Log::LevelAll;
  std::string parse_err_;
  std::string help_msg_;
  bool need_show_help_ = false;
  bool is_parse_err_ = true;

  static chen::Log::Level ParseLogLevel(const std::string &level);

public:
  OptionsImpl(int argc, const char **argv);

  virtual const char *GetParseErr(void) const { return parse_err_.c_str(); }
  virtual bool IsParseErr(void) const { return is_parse_err_; }
  virtual bool NeedShowHelp(void) const { return need_show_help_; }

  virtual const chen::Log::Level GetLogLevel(void) const { return log_level_; }
  virtual const char *GetHelpMessage(void) const { return help_msg_.c_str(); };
};

OptionsImpl::OptionsImpl(int argc, const char **argv) {
  po::options_description options_desc{ "MainOptions" };
  options_desc.add_options()
    ("help,h", "Print help messages")
    ("logLevel,l",
     po::value<std::string>()->default_value("all"),
     "log level");

  std::stringstream ss;
  ss << options_desc;
  help_msg_ = ss.str();

  po::variables_map vm;
  try {
    po::store(po::parse_command_line(argc, argv, options_desc), vm);
  } catch (po::error &e) {
    parse_err_ = e.what();
    need_show_help_ = true;
    is_parse_err_ = true;
    return;
  }
  if (vm.count("help")) {
    need_show_help_ = true;
  }

  std::string log_level = vm["logLevel"].as<std::string>();
  boost::trim(log_level);
  boost::to_lower(log_level);
  log_level_ = ParseLogLevel(log_level);
}

chen::Log::Level OptionsImpl::ParseLogLevel(const std::string &level) {
  auto item = LEVELS.find(level);
  if (item == LEVELS.end()) {
    return chen::Log::LevelNone;
  }
  return item->second;
}

_CHEN_BEGIN_

boost::shared_ptr<Options> Options::Create(int argc, const char **argv) {
  return boost::make_shared<OptionsImpl>(argc, argv);
}

_CHEN_END_
