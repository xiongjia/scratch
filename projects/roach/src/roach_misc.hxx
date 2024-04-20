/**
 * Roach - A simple HTTP Server. (libuv & http-parser & boost & C++11 & CMake)
 */

#ifndef _ROACH_MISC_HXX_
#define _ROACH_MISC_HXX_ 1

#include <utility>
#include <string>
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
    uv_loop_t *m_loop;

public:
    UVTCPWrap(uv_loop_t *uvLoop, void *uvCtx = nullptr);
    ~UVTCPWrap(void);

    inline uv_tcp_t* GetTCP(void)
    {
        return &m_tcpHandle;
    }

    inline uv_loop_t* GetLoop(void)
    {
        return m_loop;
    }
};

class UVAddr;
typedef boost::shared_ptr<UVAddr> UVAddrPtr;

class UVAddr : boost::noncopyable
{
public:
    
    static std::pair<int, UVAddrPtr> CreateIP4(const char* ip, int port);

    virtual const struct sockaddr*    GetSockAddr(void) const = 0;
    virtual const struct sockaddr_in* GetSockAddrIn(void) const = 0;

protected:
    UVAddr(void);
};

void UVShutdownLoop(uv_loop_t *loop);

template<class _Clz>
void UVAsyncDelete(uv_loop_t *loop, _Clz *ptr)
{
    uv_async_t *async = static_cast<uv_async_t*>(malloc(sizeof(uv_async_t)));
    async->data = ptr;
    uv_async_init(loop, async, [](uv_async_t *handle) {
        _Clz *ptr = static_cast<_Clz*>(handle->data);
        delete ptr;
        free(handle);
    });
    uv_async_send(async);
}

} /* namespace roach */

#endif /* !defined(_ROACH_MISC_HXX_) */
