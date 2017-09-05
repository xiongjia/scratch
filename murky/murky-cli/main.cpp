/**
 *
 */

#include "zlib.h"

int main(const int argc, const char **argv)
{
    z_stream zstream;
    zstream.zalloc = Z_NULL;
    zstream.zfree = Z_NULL;
    zstream.opaque = Z_NULL;
    deflateInit(&zstream, Z_BEST_COMPRESSION);
    deflateEnd(&zstream);
    return 0;
}
