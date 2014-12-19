/**
 *
 */


#include "roach_misc.hxx"

namespace roach {

UVHandle::UVHandle(uv_handle_t *handle)
    : m_handle(handle)
{
    m_handle->data = this;
}

UVStreamWrap::UVStreamWrap(uv_stream_t *stream)
    : UVHandle(static_cast<uv_handle_t*>(static_cast<void*>(stream)))
{
    /* NOP */
}

UVTCPWrap::UVTCPWrap(uv_loop_t *uvLoop)
    : UVStreamWrap(static_cast<uv_stream_t*>(static_cast<void*>(&m_tcpHandle)))
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

} /* namespace roach */
