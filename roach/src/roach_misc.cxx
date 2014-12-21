/**
 *
 */

#include "boost/make_shared.hpp"
#include "roach_misc.hxx"

namespace roach {

UVHandle::UVHandle(uv_handle_t *handle, void *uvCtx)
    : m_handle(handle)
    , m_uvCtx(uvCtx)
{
    m_handle->data = this;
}

UVStreamWrap::UVStreamWrap(uv_stream_t *stream, void *uvCtx)
    : UVHandle(static_cast<uv_handle_t*>(static_cast<void*>(stream)), uvCtx)
{
    /* NOP */
}

UVTCPWrap::UVTCPWrap(uv_loop_t *uvLoop, void *uvCtx)
    : UVStreamWrap(static_cast<uv_stream_t*>(static_cast<void*>(&m_tcpHandle)), uvCtx)
{
    /* Don't need handle the error for uv_tcp_init(). It'll never fail. 
     * (libuv v0.10 Linux&Windows) 
     */
    uv_tcp_init(uvLoop, &m_tcpHandle);
}

UVTCPWrap::~UVTCPWrap(void)
{
    uv_close(static_cast<uv_handle_t*>(static_cast<void*>(&m_tcpHandle)), NULL);
}


class UVBufPoolImpl : public UVBufPool
{
private:
    const unsigned int m_maxCache;

public:
    UVBufPoolImpl(const unsigned int maxCache)
        : m_maxCache(maxCache)
    {
        /* NOP */
    }

    virtual void Alloc(size_t suggestedSize, uv_buf_t *buf)
    {
        buf->base = static_cast<char*>(malloc(suggestedSize));
        buf->len = suggestedSize;
    }

    virtual void Free(uv_buf_t *buf)
    {
        if (NULL == buf || NULL == buf->base)
        {
            return;
        }
        /* TODO cache this buf */
        free(buf->base);
    }
};

UVBufPool::UVBufPool(void)    
{
    /* NOP */
}

boost::shared_ptr<UVBufPool> UVBufPool::Create(const unsigned int maxCache)
{
    return boost::make_shared<UVBufPoolImpl>(maxCache);
}

} /* namespace roach */
