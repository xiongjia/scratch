/**
*
*/

#ifndef _ROACH_MISC_HXX_
#define _ROACH_MISC_HXX_ 1

#include "boost/utility.hpp"
#include "uv.h"

namespace roach {

class UVHandle : boost::noncopyable
{
protected:
    uv_handle_t *m_handle;

public:
    UVHandle(uv_handle_t *handle);

    inline uv_handle_t *GetHandle(void)
    {
        return m_handle;
    }
};

class UVStreamWrap : public UVHandle
{
public:
    UVStreamWrap(uv_stream_t *stream);

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
    UVTCPWrap(uv_loop_t *uvLoop);
    ~UVTCPWrap(void);

    inline uv_tcp_t* GetTCP(void)
    {
        return &m_tcpHandle;
    }
};

} /* namespace roach */

#endif /* !defined(_ROACH_MISC_HXX_) */
