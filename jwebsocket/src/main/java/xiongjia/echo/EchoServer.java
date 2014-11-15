package xiongjia.echo;

import java.net.UnknownHostException;
import java.net.InetSocketAddress;
import java.util.logging.Logger;

import org.java_websocket.WebSocket;
import org.java_websocket.handshake.ClientHandshake;
import org.java_websocket.server.WebSocketServer;

public class EchoServer extends WebSocketServer {
    private final static Logger LOG = Logger.getLogger(EchoServer.class.getName()); 

    public EchoServer(int port) throws UnknownHostException {
        super(new InetSocketAddress(port));
        LOG.info("EchoServer() : " + port);
    }

    @Override
    public void onOpen(WebSocket conn, ClientHandshake handshake) {
        LOG.info("New connection: " + handshake.getResourceDescriptor() +
                 "; Remote Address: " + conn.getRemoteSocketAddress().getAddress().getHostAddress());
    }

    @Override
    public void onClose(WebSocket conn, int code, String reason, boolean remote) {
        LOG.info("Close connection: " + conn + "; Reason: " + reason +
                 "; Code: " + code + "; Remote: " + remote);
    }

    @Override
    public void onMessage(WebSocket conn, String message) {
        LOG.info("New Message, connection: " + conn +
                 "Message: " + message);
        /* send this message back */
        conn.send(message);
    }

    @Override
    public void onError(WebSocket conn, Exception ex) {
        LOG.warning("EchoServer error : " + ex.toString());
    }
}
