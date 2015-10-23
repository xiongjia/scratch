package xiongjia.kharazim;

import java.io.IOException;

import org.apache.commons.logging.Log;

public class ProxyManager {
	private static final Log log = LoggingManager.getLoggerForClass();
    private Daemon daemon;

    public static class Builder {
        private int servPort = 8080;

        public Builder port(int val) {
            servPort = val;
            return this;
        }

        public ProxyManager build() throws IOException {
            return new ProxyManager(servPort);
        }
    }

    private ProxyManager(int port) throws IOException {
        this.daemon = new Daemon(port);
    }

    void startProxy() {
        log.info("proxy start");
        daemon.start();
    }
}
