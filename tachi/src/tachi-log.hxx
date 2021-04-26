/**
 *  
 */

#ifndef _TACHI_LOG_HXX_
#define _TACHI_LOG_HXX_ 1

#include <boost/log/trivial.hpp>
#include <boost/log/attributes/named_scope.hpp>
#include <boost/utility.hpp>

#define TACHI_LOG(_level) BOOST_LOG_TRIVIAL(_level)
#define TACHI_LOG_NAMED_SCOPE(_name) BOOST_LOG_NAMED_SCOPE(_name)
#define TACHI_LOG_FUNCTION BOOST_LOG_FUNCTION

class TachiLog : boost::noncopyable {
private:
  TachiLog();

public:
  static TachiLog& getInstance() {
    static TachiLog instance;
    return instance;
  }

  void init();
};

#endif
