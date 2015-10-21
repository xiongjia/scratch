package xiongjia.kharazim;

import java.util.logging.Level;
import java.util.logging.Logger;

public class Kharazim {
    private final static Logger log = Logger.getLogger(Kharazim.class.getName());

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
        log.setLevel(Level.ALL);
        log.info("Kharazim started");
        Kharazim kharazim = new Kharazim();
        kharazim.run();
    }
}
