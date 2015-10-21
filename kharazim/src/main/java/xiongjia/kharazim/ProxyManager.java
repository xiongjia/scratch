package xiongjia.kharazim;

import java.io.IOException;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

public class ProxyManager {
    private final static Log log = LogFactory.getLog(ProxyManager.class);
    private Daemon daemon;

    ProxyManager(int port) throws IOException {
        this.daemon = new Daemon(port);
    }

    void startProxy() {
        log.info("proxy start");
        daemon.start();
    }
}
