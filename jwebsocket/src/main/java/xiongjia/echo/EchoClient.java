package xiongjia.echo;

import java.net.URI;
import java.util.logging.Logger;

import org.java_websocket.client.WebSocketClient;
import org.java_websocket.drafts.Draft;
import org.java_websocket.handshake.ServerHandshake;

public class EchoClient extends WebSocketClient {
    private final static Logger LOG = Logger.getLogger(EchoClient.class.getName());
    private String echoMsg = "";
    
    public EchoClient(final URI srvURI, final Draft draft, final String msg) {
        super(srvURI, draft);
        LOG.info("Server URI: " + srvURI +
                 "; Draft: " + draft.toString() +
                 "; Message: " + msg);
        echoMsg = msg;
    }

    @Override
    public void onOpen(ServerHandshake handshakedata) {
        LOG.info("onOpen: " + handshakedata.toString());
        this.send(echoMsg);
    }

    @Override
    public void onMessage(String message) {
        LOG.info("onMessage: " + message);
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
