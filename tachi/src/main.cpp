/**
 * 
 */


#include <stdio.h>

#include "tachi-log.hxx"

static void Test(void)
{
  TACHI_LOG_FUNCTION();
  TACHI_LOG(info) << "Info Log in Test()";
}

int main(const int argc, const char **argv) {
  TachiLog::getInstance().init();

  TACHI_LOG_NAMED_SCOPE("main");
  TACHI_LOG(info) << "Info Log";
  TACHI_LOG(error) << "Info Log";

  Test();
  return 0;
}
