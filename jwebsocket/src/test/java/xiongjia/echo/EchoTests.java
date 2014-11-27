package xiongjia.echo;

import java.net.URI;
import java.net.UnknownHostException;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;

import org.java_websocket.WebSocket;
import org.junit.Test;
import org.junit.Before;
import org.junit.After;

import static org.junit.Assert.*;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

public class EchoTests {
    private static final Log LOG = LogFactory.getLog(EchoTests.class);
    private static final int ECHO_TEST_SRV_PORT = 9501;
    private static final String ECHO_TEST_SRV_URI = "ws://localhost:9501/";
    private static final String testMsg = "EchoTest";

    private EchoServer testSrv;
    private String srvRcvMsg = "";

    @Before
    public void initialize() {
        try {
            testSrv = EchoServer.create(ECHO_TEST_SRV_PORT, new EchoServer.Listener() {
                @Override
                public void onOpen(EchoServer serv, WebSocket conn) {
                    LOG.debug("new connection");
                }

                @Override
                public void onMessage(EchoServer serv, WebSocket conn, String message) {
                    LOG.debug("Rcv: [" + message + "]");
                    srvRcvMsg = message;
                }

                @Override
                public void onError(EchoServer serv, WebSocket conn, Exception ex) {
                    /* We need ignore this exception.
                     * In some cases, the server will throw a "java.nio.channels.ClosedByInterruptException"
                     * when we close this server by server.stop(timeout).
                     */
                    LOG.error("Test Server stop error: " + ex.toString());                    
                }

                @Override
                public void onClose(EchoServer serv, WebSocket conn) {
                    LOG.debug("close connection");
                }
            });
            assertNotEquals(testSrv, null);
            testSrv.start();
        }
        catch (UnknownHostException ex) {
            fail();
        }
    }

    @Test
    public void sendMessage() {
        final CountDownLatch lock = new CountDownLatch(1);
        try {
            final EchoClient client = EchoClient.create(new URI(ECHO_TEST_SRV_URI), new EchoClient.Listener() {
                @Override
                public void onOpen(EchoClient client) {
                    /* sending message */
                    LOG.debug("Sending [" + testMsg + "] to " + ECHO_TEST_SRV_URI);
                    client.send(testMsg);
                }

                @Override
                public void onMessage(EchoClient client, String message) {
                    LOG.info("Echo client receive: [" + message + "]");
                    assertEquals(message, testMsg);
                    client.close();
                    lock.countDown();
                }

                @Override
                public void onError(EchoClient client, Exception ex) {
                    LOG.error("Echo client error: " + ex.toString());
                    fail();
                    client.close();
                }

                @Override
                public void onClose(EchoClient client) {
                    LOG.debug("Closing echo client");
                }
            });
            client.connect();
            lock.await(3000, TimeUnit.MILLISECONDS);
        }
        catch (Exception ex) {
            fail();
        }
    }

    @After
    public void cleanup() {
        assertEquals(srvRcvMsg, testMsg);
        if (testSrv != null) {
            try {
                testSrv.stop(1000 * 3);
            }
            catch (Exception ignore) {
                LOG.error("Test Server stop error: " + ignore.toString());
                fail();
            }
        }
    }
}
