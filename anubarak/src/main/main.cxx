/**
 */

#include "anubarak.h"
#include "ab_logger.h"

int main(int argc, char **argv)
{
    boost::shared_ptr<ab::Anubarak> anubarak = ab::Anubarak::Create();

    AB_LOG(ab::LoggerMask::Dbg, "start");
    return 0;
}
