package xiongjia.kharazim;

import java.io.IOException;
import java.io.InterruptedIOException;
import java.net.ServerSocket;
import java.net.Socket;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

public class Daemon extends Thread {
    private final static Log log = LogFactory.getLog(Daemon.class);
    private final int daemonPort;
    private final ServerSocket mainSocket;
    private volatile boolean running = false;

    public Daemon(int port) throws IOException {
        this.daemonPort = port;
        log.info("Proxy server port is " + daemonPort);
        mainSocket = new ServerSocket(daemonPort);
        mainSocket.setSoTimeout(1000);
    }

    @Override
    public void run() {
        log.debug("starting proxy");
        running = true;
        try {
            while (running) {
                try {
                    Socket clientSocket = mainSocket.accept();
                    if (running) {
                        log.info("new connection");
                        Proxy proxy = new Proxy(clientSocket);
                        proxy.start();
                    }
                }
                catch (InterruptedIOException err) {
                    continue;
                }
            }
        }
        catch (Exception ex) {
            log.warn("proxy server error", ex);
        }
    }

    public void stopServer() {
        running = false;
    }
}

