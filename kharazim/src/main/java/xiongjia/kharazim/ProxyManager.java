package xiongjia.kharazim;

import java.io.IOException;
import java.util.logging.Logger;

public class ProxyManager {
    private final static Logger log = Logger.getLogger(ProxyManager.class.getName());
    private Daemon daemon;

    ProxyManager(int port) throws IOException {
        this.daemon = new Daemon(port);
    }

    void startProxy() {
        log.info("proxy start");
        daemon.start();
    }
}
