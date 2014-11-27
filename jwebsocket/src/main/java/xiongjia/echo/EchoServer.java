package xiongjia.echo;

import java.net.UnknownHostException;
import java.net.InetSocketAddress;

import org.java_websocket.WebSocket;
import org.java_websocket.handshake.ClientHandshake;
import org.java_websocket.server.WebSocketServer;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;


public class EchoServer extends WebSocketServer {
    private static final Log LOG = LogFactory.getLog(EchoServer.class);

    public interface Listener {
        void onOpen(EchoServer serv, WebSocket conn);
        void onClose(EchoServer serv, WebSocket conn);
        void onError(EchoServer serv, WebSocket conn, Exception ex);

        void onMessage(EchoServer serv, WebSocket conn, String message);
    }

    private Listener listener;

    private EchoServer(int port, Listener listener) throws UnknownHostException {
        super(new InetSocketAddress(port));
        this.listener = listener;
        if (LOG.isDebugEnabled()) {
            LOG.debug("EchoServer() : " + port);
        }
    }

    public static EchoServer create(int port, Listener listener)
            throws UnknownHostException {
        return new EchoServer(port, listener);
    }

    public static EchoServer create(int port)
            throws UnknownHostException {
        return new EchoServer(port, null);
    }

    @Override
    public void onOpen(WebSocket conn, ClientHandshake handshake) {
        if (LOG.isDebugEnabled()) {
            LOG.debug("New connection: " + handshake.getResourceDescriptor() +
                      "; Remote Address: " +
                      conn.getRemoteSocketAddress().getAddress().getHostAddress());
        }

        if (listener != null) {
            listener.onOpen(this, conn);
        }
    }

    @Override
    public void onClose(WebSocket conn, int code, String reason, boolean remote) {
        if (LOG.isDebugEnabled()) {
            LOG.debug("Close connection: " + conn +
                      "; Reason: " + reason +
                      "; Code: " + code + "; Remote: " + remote);
        }
        if (listener != null) {
            listener.onClose(this, conn);
        }
    }

    @Override
    public void onMessage(WebSocket conn, String message) {
        if (LOG.isDebugEnabled()) {
            LOG.debug("New Message, connection: " + conn +
                      "Message: " + message);
        }

        if (listener != null) {
            listener.onMessage(this, conn, message);
        }
        conn.send(message);
    }

    @Override
    public void onError(WebSocket conn, Exception ex) {
        if (LOG.isDebugEnabled()) {
            LOG.debug("EchoServer error : " + ex.toString());
        }

        if (listener != null) {
            listener.onError(this, conn, ex);
        }
    }
}
