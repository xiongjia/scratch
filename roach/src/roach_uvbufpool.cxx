/**
 * Roach - A simple HTTP Server. (libuv & http-parser & boost & C++11 & CMake)
 */

#include "boost/make_shared.hpp"
#include "roach_uvbufpool.hxx"

namespace roach {

class UVBufPoolImpl : public UVBufPool
{
private:
    boost::shared_ptr<Context> m_context;
    boost::shared_ptr<Logger> m_logger;
    const unsigned int m_maxCache;

public:
    UVBufPoolImpl(boost::shared_ptr<Context> ctx, const unsigned int maxCache)
        : m_maxCache(maxCache)
        , m_context(ctx)
        , m_logger(m_context->GetLogger())
    {
        /* NOP */
    }

    virtual void Alloc(size_t suggestedSize, uv_buf_t *buf)
    {
        char *base = static_cast<char*>(malloc(suggestedSize));
        if (NULL == base)
        {
            m_logger->LogNoFmt(Logger::Err, "Cannot allocat uv_buf (No memory)");
            buf->base = NULL;
            buf->len = 0;
            return;
        }
        m_logger->Log(Logger::Dbg, "Create new uv_buf(%d)", suggestedSize);
        buf->base = base;
        buf->len = suggestedSize;
    }

    virtual void Free(const uv_buf_t *buf)
    {
        if (NULL == buf || NULL == buf->base)
        {
            return;
        }

        m_logger->Log(Logger::Dbg, "Free uv_buf(%d)", buf->len);

        /* TODO cache this buf */
        free(buf->base);
    }
};

UVBufPool::UVBufPool(void)
{
    /* NOP */
}

boost::shared_ptr<UVBufPool> UVBufPool::Create(boost::shared_ptr<Context> ctx,
                                               const size_t maxCache)
{
    return boost::make_shared<UVBufPoolImpl>(ctx, maxCache);
}

} /* namespace roach */
