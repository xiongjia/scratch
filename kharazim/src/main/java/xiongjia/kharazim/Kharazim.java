package xiongjia.kharazim;

import org.apache.commons.logging.Log;

public class Kharazim {
    private final static Log log = LoggingManager.getLoggerForClass();

    public void start(String[] args) {
        log.debug("Starting...");

    	/* TODO read the port and other settings from args */
        try {
            ProxyManager proxyMgr = new ProxyManager.Builder()
                .port(8080)
                .build();
            proxyMgr.startProxy();
        }
        catch (Exception ex) {
        	log.fatal("proxy server error: ", ex);
        }
    }

    public static void main(String[] args) {
        (new Kharazim()).start(args);
    }
}
