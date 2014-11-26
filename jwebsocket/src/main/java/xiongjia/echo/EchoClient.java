package xiongjia.echo;

import java.net.URI;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.java_websocket.client.WebSocketClient;
import org.java_websocket.drafts.Draft;
import org.java_websocket.drafts.Draft_10;
import org.java_websocket.handshake.ServerHandshake;

public class EchoClient extends WebSocketClient {
    private final static Log LOG = LogFactory.getLog(EchoClient.class);

    public interface Listener {
        void onOpen(EchoClient client);
        void onClose(EchoClient client);
        void onError(EchoClient client, Exception ex);

        void onMessage(EchoClient client, String message);
    }

    private Listener listener;

    private EchoClient(final URI srvURI, final Draft draft, Listener listener) {
        super(srvURI, draft);
        if (LOG.isDebugEnabled()) {
            LOG.debug("Server URI: " + srvURI + "; Draft: " + draft.toString());
        }
        this.listener = listener;
    }

    public static EchoClient create(final URI srvURI, Listener listener) {
        return new EchoClient(srvURI, new Draft_10(), listener);
    }

    @Override
    public void onOpen(ServerHandshake handshakedata) {
        if (LOG.isDebugEnabled()) {
            LOG.debug("onOpen: " + handshakedata.toString());
        }

        if (listener != null) {
            listener.onOpen(this);
        }
    }

    @Override
    public void onMessage(String message) {
        if (LOG.isDebugEnabled()) {
            LOG.debug("onMessage: " + message);
        }

        if (listener != null) {
            listener.onMessage(this, message);
        }
    }

    @Override
    public void onClose(int code, String reason, boolean remote) {
        if (LOG.isDebugEnabled()) {
            LOG.debug("onClose: code = " + code +
                      "; reason = " + reason + 
                      "; remote = " + remote);
        }

        if (listener != null) {
            listener.onClose(this);
        }
    }

    @Override
    public void onError(Exception ex) {
        if (LOG.isDebugEnabled()) {
            LOG.debug("Client Error: " + ex.toString());
        }

        if (listener != null) {
            listener.onError(this, ex);
        }
    }
}
