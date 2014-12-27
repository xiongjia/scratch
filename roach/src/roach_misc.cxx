/**
 * Roach - A simple HTTP Server. (libuv & http-parser & boost & C++11 & CMake)
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

class UVAddrImpl : public UVAddr
{
private:
    struct sockaddr_in m_addr;

public:
    UVAddrImpl(struct sockaddr_in *addr)
        : UVAddr()
    {
        memcpy(&m_addr, addr, sizeof(sockaddr_in));
    }

    virtual const struct sockaddr* GetSockAddr(void) const
    {
        return static_cast<const struct sockaddr*>(static_cast<const void*>(&m_addr));
    }

    virtual const struct sockaddr_in* GetSockAddrIn(void) const
    {
        return const_cast<const sockaddr_in*>(&m_addr);
    }
};

UVAddr::UVAddr(void)
{
    /* NOP */
}

std::pair<int, boost::shared_ptr<UVAddr>> UVAddr::CreateIP4(const char* ip,
                                                            int port)
{
    struct sockaddr_in addr;
    int uvErr = uv_ip4_addr(ip, port, &addr);
    if (0 != uvErr)
    {
        return std::make_pair(uvErr, boost::shared_ptr<UVAddr>(nullptr));
    }
    return std::make_pair(uvErr, boost::make_shared<UVAddrImpl>(&addr));
}

} /* namespace roach */
