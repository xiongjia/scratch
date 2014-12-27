/**
 * Roach - A simple HTTP Server. (libuv & http-parser & boost & C++11 & CMake)
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

class UVAddr : boost::noncopyable
{
public:
    static std::pair<int, boost::shared_ptr<UVAddr>> CreateIP4(const char* ip, int port);

    virtual const struct sockaddr*    GetSockAddr(void) const = 0;
    virtual const struct sockaddr_in* GetSockAddrIn(void) const = 0;

protected:
    UVAddr(void);
};

} /* namespace roach */

#endif /* !defined(_ROACH_MISC_HXX_) */
