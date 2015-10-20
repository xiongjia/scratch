package xiongjia.kharazim;

import java.util.logging.Logger;

public class Daemon extends Thread {
    private final static Logger log = Logger.getLogger(Daemon.class.getName());

    @Override
    public void run() {
        log.info("Proxy server started");
    }
}

