package xiongjia.kharazim;

import java.net.Socket;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

public class Proxy extends Thread {
    private final static Log log = LogFactory.getLog(Daemon.class);
    private Socket clientSocket = null;

    public Proxy(Socket sock) {
        log.debug("createing proxy");
        this.clientSocket = sock;
    }

    @Override
    public void run() {
        log.debug("proxy thread started");

    }
}
