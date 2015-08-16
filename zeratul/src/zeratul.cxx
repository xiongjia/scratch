/**
 * Zeratul - A simple socks proxy. (Boost ASIO Example)
 */

#include <vector>

#include "boost/make_shared.hpp"
#include "boost/enable_shared_from_this.hpp"
#include "boost/array.hpp"
#include "boost/asio/error.hpp"

#include "zeratul.hxx"

namespace asio     = boost::asio;
namespace asio_ip  = boost::asio::ip;

_ZERATUL_BEGIN_

class Session : public boost::enable_shared_from_this<Session>
{
private:
    enum { MAX_BUFF_LEN = 1024 };

    typedef enum 
    {
        VER_5 = 5
    } SocksProto;

    typedef enum 
    {
        NOAUTH = 0, GSSAPI = 1, USERPASS = 2
    } SocksAuthMethods;

    typedef enum 
    {
        CONNECT = 1, BIND = 2, UDP_ASSOCIATE = 3
    } SocksCmd;

    typedef enum
    {
        IP_V4 = 1,
        DOMAINNAME = 3,
        IP_V6 = 4
    } SocksAddrType;

private:
    asio::io_service::work m_wrkSvc;
    asio_ip::tcp::socket m_socket;
    std::vector<boost::shared_ptr<asio_ip::tcp::socket>> m_clients;

    unsigned char m_data[MAX_BUFF_LEN];
    unsigned char m_srcToDstBuf[MAX_BUFF_LEN];
    unsigned char m_dstToSrcBuf[MAX_BUFF_LEN];

public:
    Session(asio::io_service::work &wrkSvc,
            asio_ip::tcp::socket sock);

    void Start(void);
    void Close(void);

private:
    void OnHandshake(const unsigned char *data, std::size_t len);
    void OnRequest(const unsigned char *data, std::size_t len);

    void CmdConn(const asio_ip::address_v4 &dstAddr,
                 const unsigned short dstPort);


    void SrcToDest(boost::shared_ptr<asio_ip::tcp::socket> socket);
    void DestToSrc(const boost::array<unsigned char, 10> &rep,
                   boost::shared_ptr<asio_ip::tcp::socket> socket);
};

class ServerImpl : public Server
{
private:
    asio_ip::tcp::acceptor m_acceptor;
    asio_ip::tcp::socket   m_socket;
    asio::io_service::work m_wrkSvc;

public:
    ServerImpl(boost::asio::io_service &ioSvc, unsigned short port);

private:
    void DoAccept(void);
};

Session::Session(asio::io_service::work &wrkSvc, asio_ip::tcp::socket sock)
    : m_socket(std::move(sock))
    , m_wrkSvc(wrkSvc.get_io_service())
{
    /* NOP */
}
void Session::Close(void)
{
    BOOST_LOG_TRIVIAL(debug) << "Closing session";
    m_socket.close();
}

void Session::Start(void)
{
    BOOST_LOG_TRIVIAL(debug) << "Session start";

    auto self(shared_from_this());
    m_socket.async_read_some(boost::asio::buffer(m_data, MAX_BUFF_LEN),
        [this, self](boost::system::error_code errCode, std::size_t len)
    {
        if (errCode)
        {
            BOOST_LOG_TRIVIAL(error) << "Session Read error: " << errCode;
            self->Close();
            return;
        }
        BOOST_LOG_TRIVIAL(trace) << "Read Data (" << len << ")";
        OnHandshake(m_data, len);
    });
}

void Session::SrcToDest(boost::shared_ptr<asio_ip::tcp::socket> socket)
{
    /* Src => Dst */
    auto self(shared_from_this());
    m_socket.async_read_some(asio::buffer(m_srcToDstBuf, MAX_BUFF_LEN),
        [this, socket, self](boost::system::error_code errCode,
        std::size_t len)
    {
        if (errCode)
        {
            BOOST_LOG_TRIVIAL(error) << "Read error: " << errCode;
            return;
        }

        BOOST_LOG_TRIVIAL(trace) << "Src => Dst: " << len;
        socket->async_write_some(asio::buffer(m_srcToDstBuf, len),
            [this, socket, self](boost::system::error_code errCode, std::size_t len)
        {
            if (errCode)
            {
                BOOST_LOG_TRIVIAL(error) << "Write error: " << errCode;
                return;
            }
            SrcToDest(socket);
        });
    });
}

void Session::DestToSrc(const boost::array<unsigned char, 10> &rep,
                        boost::shared_ptr<asio_ip::tcp::socket> socket)
{
    /* Dst => Src */
    auto self(shared_from_this());
    socket->async_read_some(asio::buffer(m_dstToSrcBuf, MAX_BUFF_LEN),
        [this, rep, socket, self](boost::system::error_code errCode,
                                  std::size_t len)
    {
        if (errCode)
        {
            BOOST_LOG_TRIVIAL(error) << "Read error: " << errCode;
            return;
        }
        m_socket.async_write_some(asio::buffer(m_dstToSrcBuf, len),
            [this, socket, rep, self](boost::system::error_code errCode,
                                        std::size_t len)
        {
            if (errCode)
            {
                BOOST_LOG_TRIVIAL(error) << "Write error: " << errCode;
                return;
            }
            DestToSrc(rep, socket);
        });
    });
}

void Session::CmdConn(const asio_ip::address_v4 &dstAddr, 
                      const unsigned short dstPort)
{
    BOOST_LOG_TRIVIAL(trace) << "Connecting to: "
        << dstAddr << " (" << dstPort << ")";
    auto dst = asio_ip::tcp::endpoint(dstAddr, dstPort);
    boost::shared_ptr<asio_ip::tcp::socket> socket =
        boost::make_shared<asio_ip::tcp::socket>(m_wrkSvc.get_io_service());
    m_clients.push_back(socket);
    auto self(shared_from_this());
    socket->async_connect(dst,
        [this, socket, dst, self](boost::system::error_code errCode)
    {
        auto addrBytes = dst.address().to_v4().to_bytes();
        unsigned short nport = htons(dst.port());
        boost::array<unsigned char, 2> portNum = {
            static_cast<unsigned char>(nport & 0x00FF),
            static_cast<unsigned char>((nport & 0xFF00) >> 8),
        };
        if (errCode)
        {            
            BOOST_LOG_TRIVIAL(warning) << "Cannot connect to " 
                << dst.address().to_string() << " (" << dst.port() << ") : "
                << errCode;

            /* Connection refused */
            boost::array<unsigned char, 10> rep =
            {
                SocksProto::VER_5, 0x05, 0,
                SocksAddrType::IP_V4, 
                addrBytes[0], addrBytes[1], addrBytes[2], addrBytes[3],
                portNum[0], portNum[1]
            };
            m_socket.async_write_some(asio::buffer(rep),
                [this, self](boost::system::error_code errCode, std::size_t len)
            {
                if (errCode)
                {
                    BOOST_LOG_TRIVIAL(error) << 
                        "Cannot write rep: " << errCode;
                    return;
                }
                self->Close();
            });
            return;
        }

        /* connected */
        boost::array<unsigned char, 10> rep =
        {
            SocksProto::VER_5, 0x00, 0,
            SocksAddrType::IP_V4,
            addrBytes[0], addrBytes[1], addrBytes[2], addrBytes[3],
            portNum[0], portNum[1]
        };
        m_socket.async_write_some(asio::buffer(rep),
            [this, socket, rep, self](boost::system::error_code errCode,
                                      std::size_t len)
        {
            if (errCode)
            {
                BOOST_LOG_TRIVIAL(error) <<
                    "Cannot write rep: " << errCode;
                return;
            }
            /* Src => Dst */
            SrcToDest(socket);
            /* Dst => Src */
            DestToSrc(rep, socket);
        });
    });
}

void Session::OnRequest(const unsigned char *data, std::size_t len)
{
    if (data[0] != SocksProto::VER_5)
    {
        BOOST_LOG_TRIVIAL(warning) << "Don't support proto vers: "
            << static_cast<unsigned int>(data[0]);
        Close();
        return;
    }

    int cmd = data[1];
    int addrType = data[3];
    if (addrType != SocksAddrType::IP_V4)
    {
        BOOST_LOG_TRIVIAL(warning) << "Don't support addr type: " << addrType;
        Close();
        return;
    }

    asio_ip::address_v4::bytes_type addr = {
        data[4], data[5], data[6], data[7] };
    auto dstAddr = asio_ip::address_v4(addr);
    unsigned short dstPort = ((unsigned short)data[9] << 8) |
                             (unsigned short)data[8];
    dstPort = ntohs(dstPort);

    BOOST_LOG_TRIVIAL(trace) << "Cmd = " << cmd << "; Dest = "
        << dstAddr << " (" << dstPort << ")";
    switch (cmd)
    {
    case SocksCmd::CONNECT:
        CmdConn(dstAddr, dstPort);
        break;
    default:
        BOOST_LOG_TRIVIAL(warning) << "Don't support cmd: " << cmd;
        Close();
        break;
    }
}

void Session::OnHandshake(const unsigned char *data, std::size_t len)
{
    if (2 > len)
    {
        BOOST_LOG_TRIVIAL(warning) << "Invalid Handshake. len = " << len;
        Close();
        return;
    }

    if (data[0] != SocksProto::VER_5)
    {
        BOOST_LOG_TRIVIAL(warning) << "Don't support proto vers: " 
            << static_cast<unsigned int>(data[0]);
        Close();
        return;
    }
    unsigned int methodsNum = data[1];
    BOOST_LOG_TRIVIAL(trace) << "Autho Methods Num = " << methodsNum;

    if ((2 + methodsNum) > len)
    {
        BOOST_LOG_TRIVIAL(warning) << "Invalid Handshake. len = " << len 
                                   << "; Methods Num = " << methodsNum;
        Close();
        return;
    }

    std::vector<int> authMethods;
    for (unsigned int i = 0; i < methodsNum; ++i)
    {
        unsigned int method = data[2 + i];
        BOOST_LOG_TRIVIAL(trace) << "Autho Method[" << i << "] = " << method;
        authMethods.push_back(method);
    }

    /* In this version, it only supports SocksAuthMethods::NOAUTH */    
    boost::array<unsigned char, 2> rep = {
        SocksProto::VER_5,
        SocksAuthMethods::NOAUTH
    };
    auto self(shared_from_this());
    m_socket.async_write_some(asio::buffer(rep), 
        [this, self](boost::system::error_code errCode, std::size_t len)
    {
        if (errCode)
        {
            BOOST_LOG_TRIVIAL(error) << "Handshake write error: " << errCode;
            return;
        }

        BOOST_LOG_TRIVIAL(trace) << "Handshake write data (" << len << ")";
        /* only support IPv4 and it needs 10 bytes at least */
        asio::async_read(m_socket, asio::buffer(m_data, MAX_BUFF_LEN),
            asio::transfer_at_least(10),
            [this, self](boost::system::error_code errCode, std::size_t len)
        {
            if (errCode)
            {
                BOOST_LOG_TRIVIAL(error) << "Session Read error: " << errCode;
                self->Close();
                return;
            }
            BOOST_LOG_TRIVIAL(trace) << "Read Data (" << len << ")";
            OnRequest(m_data, len);
        });
    });
}

ServerImpl::ServerImpl(boost::asio::io_service &ioSvc, unsigned short port)
    : Server()
    , m_acceptor(ioSvc, asio_ip::tcp::endpoint(asio_ip::tcp::v4(), port))
    , m_socket(ioSvc)
    , m_wrkSvc(ioSvc)
{
    DoAccept();
}

void ServerImpl::DoAccept(void)
{
    BOOST_LOG_TRIVIAL(debug) << "Server DoAccept()";
    m_acceptor.async_accept(m_socket,
        [this](boost::system::error_code errCode)
    {
        if (!errCode)
        {
            BOOST_LOG_TRIVIAL(debug) << "New connection";
            auto sesson = boost::make_shared<Session>(m_wrkSvc,
                std::move(m_socket));
            sesson->Start();
        }
        else
        {
            BOOST_LOG_TRIVIAL(error) << "Accept Error: " << errCode;
        }
        DoAccept();
    });
}

Server::Server(void)
{
    /* NOP */
}

boost::shared_ptr<Server> Server::Create(boost::asio::io_service &ioSvc,
                                         unsigned short port)
{
    return boost::make_shared<ServerImpl>(ioSvc, port);
}
 
_ZERATUL_END_
