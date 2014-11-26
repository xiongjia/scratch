package xiongjia.echo;

import java.net.URI;
import java.net.URISyntaxException;
import java.net.UnknownHostException;

import org.junit.Test;
import org.junit.Before;
import org.junit.After;

import static org.junit.Assert.*;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

public class EchoTests {
    private final static Log LOG = LogFactory.getLog(EchoTests.class);
    private static int ECHO_TEST_SRV_PORT = 9501;
    private static String ECHO_TEST_SRV_URI = "ws://localhost:9501/";
    private EchoServer testSrv;
            
    @Before
    public void initialize() {
        try {
            testSrv = EchoServer.create(ECHO_TEST_SRV_PORT);
            assertNotEquals(testSrv, null);
        }
        catch (UnknownHostException ex) {
            fail();
        }
    }

    @Test
    public void sendMessage() {
        final String testMsg = "EchoTest";
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
                }

                @Override
                public void onError(EchoClient client, Exception ex) {
                    LOG.error("Echo client error: " + ex.toString());
                    client.close();
                }

                @Override
                public void onClose(EchoClient client) {
                    LOG.debug("Closing echo client");
                }
            });
            client.connect();
        }
        catch (URISyntaxException uriErr) {
            fail();
        }
    }

    @After
    public void cleanup() {
        if (testSrv != null) {
            try {
                testSrv.stop(1000 * 3);
            }
            catch (Exception ignore) {
                LOG.error("Test Server stop error: " + ignore.toString());
            }
        }
    }
}

