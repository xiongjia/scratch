/**
* Chen - My Network protocol tests
*/

#ifndef _LIBCHEN_CHEN_PROGO_OPTIONS_HXX_
#define _LIBCHEN_CHEN_PROGO_OPTIONS_HXX_ 1

#include "boost/shared_ptr.hpp"
#include "boost/utility.hpp"

#include "chen_types.hxx"
#include "chen_log.hxx"

_CHEN_BEGIN_

class Options : boost::noncopyable {
public:
  static boost::shared_ptr<Options> Create(int argc, const char **argv);

public:
  virtual const Log::Level GetLogLevel(void) const = 0;

  virtual const char *GetHelpMessage(void) const = 0;
  virtual const char *GetParseErr(void) const = 0;
  virtual bool IsParseErr(void) const = 0;
  virtual bool NeedShowHelp(void) const = 0;
};

_CHEN_END_

#endif /* !defined(_LIBCHEN_CHEN_PROGO_OPTIONS_HXX_) */
