package xiongjia.kharazim;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;


public class Kharazim {
    private final static Log log = LogFactory.getLog(Kharazim.class);


    public final static int DEFAULT_PORT = 8080;

    private int servPort = DEFAULT_PORT;


    public int getServPort() {
        return servPort;
    }

    public void setServPort(int servPort) {
        this.servPort = servPort;
    }

    public void run() {
        try {
            ProxyManager proxyMgr = new ProxyManager(servPort);
            proxyMgr.startProxy();
        }
        catch (Exception e) {
            log.info("io err");
        }
    }

    public static void main(String[] args) {
        log.info("Kharazim started");
        Kharazim kharazim = new Kharazim();
        kharazim.run();
    }
}
