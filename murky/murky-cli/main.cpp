/**
 *
 */

#include <stdio.h>
#include "zlib.h"
#include "boost/make_shared.hpp"

struct test_data {
    int x;
};

int main(const int /* argc */, const char ** /* argv */)
{
    boost::shared_ptr<test_data> s = boost::make_shared<test_data>();
    printf("cmake test %d\n", s->x);

    z_stream zstream;
    zstream.zalloc = Z_NULL;
    zstream.zfree = Z_NULL;
    zstream.opaque = Z_NULL;
    deflateInit(&zstream, Z_BEST_COMPRESSION);
    deflateEnd(&zstream);
    return 0;
}
