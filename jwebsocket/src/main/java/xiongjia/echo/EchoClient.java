package xiongjia.echo;

import java.net.URI;
import java.util.logging.Logger;

import org.java_websocket.client.WebSocketClient;
import org.java_websocket.drafts.Draft;
import org.java_websocket.handshake.ServerHandshake;

import xiongjia.echo.EchoEventListener;

public class EchoClient extends WebSocketClient {
    private final static Logger LOG = Logger.getLogger(EchoClient.class.getName());
    private EchoEventListener evtListener;
    private String echoMsg = "";

    private EchoClient(final URI srvURI, final Draft draft, final String msg, EchoEventListener listener) {
        super(srvURI, draft);
        LOG.info("Server URI: " + srvURI +
                 "; Draft: " + draft.toString() +
                 "; Message: " + msg);
        echoMsg = msg;
        evtListener = listener;
    }
    
    public static EchoClient create(final URI srvURI, final Draft draft, final String msg, EchoEventListener listener) {
        return new EchoClient(srvURI, draft, msg, listener);
    }

    public static EchoClient create(final URI srvURI, final Draft draft, final String msg) {
        return new EchoClient(srvURI, draft, msg, null);
    }
    
    @Override
    public void onOpen(ServerHandshake handshakedata) {
        LOG.info("onOpen: " + handshakedata.toString());
        this.send(echoMsg);
    }

    @Override
    public void onMessage(String message) {
        LOG.info("onMessage: " + message);
        if (evtListener != null) {
            evtListener.onMessage(message);
        }
        this.close();
    }

    @Override
    public void onClose(int code, String reason, boolean remote) {
        LOG.info("onClose: code = " + code + "; reason = " + reason +
                 "; remote = " + remote);
    }

    @Override
    public void onError(Exception ex) {
        LOG.warning("Client Error: " + ex.toString());
    }
}
