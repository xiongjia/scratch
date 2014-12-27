/**
 * Roach - A simple HTTP Server. (libuv & http-parser & boost & C++11 & CMake)
 */

#ifndef _ROACH_UVBUFPOOL_HXX_
#define _ROACH_UVBUFPOOL_HXX_ 1

#include "boost/utility.hpp"
#include "boost/shared_ptr.hpp"
#include "uv.h"

#include "roach_context.hxx"

namespace roach {

class UVBufPool : boost::noncopyable
{
private:
    static const size_t MAXCACHE = 1024 * 1024 * 5;

public:
    static boost::shared_ptr<UVBufPool> Create(boost::shared_ptr<Context> ctx,
                                               const size_t maxCache = MAXCACHE);

    virtual void Alloc(size_t suggestedSize, uv_buf_t *buf) = 0;
    virtual void Free(const uv_buf_t *buf) = 0;

protected:
    UVBufPool(void);
};

} /* namespace roach */

#endif /* !defined(_ROACH_UVBUFPOOL_HXX_) */
