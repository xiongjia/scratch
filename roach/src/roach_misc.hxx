/**
 *
 */

#ifndef _ROACH_MISC_HXX_
#define _ROACH_MISC_HXX_ 1

#include <utility>
#include "boost/utility.hpp"
#include "boost/shared_ptr.hpp"
#include "uv.h"

namespace roach {

class UVHandle : boost::noncopyable
{
protected:
    uv_handle_t *m_handle;
    void *m_uvCtx;

public:
    UVHandle(uv_handle_t *handle, void *uvCtx = nullptr);

    inline uv_handle_t *GetHandle(void)
    {
        return m_handle;
    }

    inline void* GetUVCtx(void)
    {
        return m_uvCtx;
    }

    inline void SetUVCtx(void *uvCtx)
    {
        m_uvCtx = uvCtx;
    }
};

class UVStreamWrap : public UVHandle
{
public:
    UVStreamWrap(uv_stream_t *stream, void *uvCtx = nullptr);

    inline uv_stream_t *GetStream(void)
    {
        return static_cast<uv_stream_t*>(static_cast<void*>(GetHandle()));
    }
};

class UVTCPWrap : public UVStreamWrap
{
protected:
    uv_tcp_t  m_tcpHandle;

public:
    UVTCPWrap(uv_loop_t *uvLoop, void *uvCtx = nullptr);
    ~UVTCPWrap(void);

    inline uv_tcp_t* GetTCP(void)
    {
        return &m_tcpHandle;
    }
};

class UVBufPool : boost::noncopyable
{
private:
    static const unsigned int DEFAULT_MAXCACHE = 1024 * 5;

public:
    static boost::shared_ptr<UVBufPool> Create(const unsigned int maxCache = DEFAULT_MAXCACHE);

    virtual void Alloc(size_t suggestedSize, uv_buf_t *buf) = 0;
    virtual void Free(uv_buf_t *buf) = 0;

protected:
    UVBufPool(void);
};


} /* namespace roach */

#endif /* !defined(_ROACH_MISC_HXX_) */
