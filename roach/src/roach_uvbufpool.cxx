/**
 * Roach - A simple HTTP Server. (libuv & http-parser & boost & C++11 & CMake)
 */

#include <list>
#include "boost/make_shared.hpp"
#include "boost/thread/thread.hpp"
#include "boost/thread/mutex.hpp"

#include "roach_uvbufpool.hxx"

namespace roach {

class UVBufPoolImpl : public UVBufPool
{
private:
    boost::mutex m_mutex;
    boost::shared_ptr<Context> m_context;
    boost::shared_ptr<Logger> m_logger;

    std::list<uv_buf_t> m_cache;
    const size_t m_maxCache;
    size_t m_cachedSz;

    void DoAlloc(size_t suggestedSize, uv_buf_t *buf);
    void DoFree(const uv_buf_t *buf);

public:
    UVBufPoolImpl(boost::shared_ptr<Context> ctx, const unsigned int maxCache)
        : m_maxCache(maxCache)
        , m_context(ctx)
        , m_logger(m_context->GetLogger())
        , m_cachedSz(0)
    {
        /* NOP */
    }

    virtual void Alloc(size_t suggestedSize, uv_buf_t *buf)
    {
        boost::mutex::scoped_lock scopedLock(m_mutex);
        DoAlloc(suggestedSize, buf);
    }

    virtual void Free(const uv_buf_t *buf)
    {
        boost::mutex::scoped_lock scopedLock(m_mutex);
        DoFree(buf);
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

void UVBufPoolImpl::DoAlloc(size_t suggestedSize, uv_buf_t *buf)
{
    auto cacheItem = std::find_if(m_cache.begin(), m_cache.end(),
        [=](uv_buf_t item) { return item.len >= suggestedSize; });
    if (cacheItem == m_cache.end())
    {
        /* create a new uv_buf_t */
        char *base = static_cast<char*>(malloc(suggestedSize));
        if (NULL == base)
        {
            m_logger->LogNoFmt(Logger::Err, "Cannot allocat uv_buf (No memory)");
            buf->base = NULL;
            buf->len = 0;
        }
        else
        {
            m_logger->Log(Logger::Dbg, "Create new uv_buf(%d)", suggestedSize);
            buf->base = base;
            buf->len = suggestedSize;
        }
    }
    else
    {
        /* reuse the cached uv_buf_t */
        buf->base = (*cacheItem).base;
        buf->len = (*cacheItem).len;
        m_cachedSz = (m_cachedSz >= buf->len) ? (m_cachedSz - buf->len) : 0;
        m_cache.erase(cacheItem);
        m_logger->Log(Logger::Dbg,
            "Reuse buffer: Suggested(%d), BufLen(%d), cachedSz(%d)",
            suggestedSize, buf->len, m_cachedSz);
    }
}

void UVBufPoolImpl::DoFree(const uv_buf_t *buf)
{
    if (NULL == buf || NULL == buf->base)
    {
        m_logger->LogNoFmt(Logger::War, "Freeing an invalid buffer");
        return;
    }

    m_logger->Log(Logger::Dbg, "Freeing uv_buf(%d)", buf->len);
    if ((buf->len + m_cachedSz) >= m_maxCache)
    {
        m_logger->Log(Logger::Dbg, "m_cachedSz(%d) + bufLen(%d) > m_maxCache",
            m_cachedSz, buf->len);
        free(buf->base);
        return;
    }

    m_logger->Log(Logger::Dbg, "Adding uv_buf(%d) to cache", buf->len);
    uv_buf_t cacheItem = uv_buf_init(buf->base, buf->len);
    m_cache.push_back(cacheItem);
    m_cachedSz += buf->len;
}

} /* namespace roach */
