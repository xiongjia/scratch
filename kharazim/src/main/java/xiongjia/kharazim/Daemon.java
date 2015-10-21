package xiongjia.kharazim;

import java.io.IOException;
import java.io.InterruptedIOException;
import java.net.ServerSocket;
import java.net.Socket;
import java.util.logging.Logger;

public class Daemon extends Thread {
    private final static Logger log = Logger.getLogger(Daemon.class.getName());
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
        log.info("Proxy server started");

        running = true;
        try {
            while (running) {
                try {
                    Socket clientSocket = mainSocket.accept();
                    if (running) {
                        log.info("new connection");
                    }

                }
                catch (InterruptedIOException e) {
                    continue;
                }
            }
        } catch (Exception e) {
            log.warning("HTTP(S) Test Script Recorder stopped");
        }
    }

    public void stopServer() {
        running = false;
    }
}

