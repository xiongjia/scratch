package xiongjia.echo;

import java.net.UnknownHostException;
import java.net.InetSocketAddress;
import java.util.logging.Logger;

import org.java_websocket.WebSocket;
import org.java_websocket.handshake.ClientHandshake;
import org.java_websocket.server.WebSocketServer;

import xiongjia.echo.EchoEventListener;

public class EchoServer extends WebSocketServer {
    private final static Logger LOG = Logger.getLogger(EchoServer.class.getName()); 
    private EchoEventListener evtListener;

    private EchoServer(int port, EchoEventListener listener) 
            throws UnknownHostException {
        super(new InetSocketAddress(port));
        evtListener = listener;
        LOG.info("EchoServer() : " + port);
    }

    public static EchoServer create(int port, EchoEventListener listener)
            throws UnknownHostException {
        return new EchoServer(port, listener);
    }

    public static EchoServer create(int port)
            throws UnknownHostException {
        return new EchoServer(port, null);
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
        if (evtListener != null) {
            evtListener.onMessage(message);
        }
        /* send this message back */
        conn.send(message);
    }

    @Override
    public void onError(WebSocket conn, Exception ex) {
        LOG.warning("EchoServer error : " + ex.toString());
    }
}
