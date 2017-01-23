package scratch;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.fail;

import org.junit.After;
import org.junit.Before;
import org.junit.Ignore;
import org.junit.Test;

import scratch.JeroMsgQueueScratch.ZChannel;

import java.util.Map;

public class JeroMsgQueueScratchTest {
  private JeroMsgQueueScratch jeromq;

  @Before
  public final void setUp() {
    jeromq = new JeroMsgQueueScratch();
  }

  @After
  public final void tearDown() {
    jeromq.destoryContext();
  }

  @Ignore("Test is ignored")
  @Test
  public void testPubSubPattern() {
    final String topic = "zscratch";
    final String testData = "test-data-123";
    final long timeoutMilliseconds = 1000 * 30;

    ZChannel zpub = jeromq.buildPublisher();
    ZChannel zsub = jeromq.buildSubscriber(zpub.getBindEndPoint(), topic);

    zpub.send(topic, testData);
    final Map<String, String> rcv = zsub.recvStr(timeoutMilliseconds);
    if (rcv == null) {
      fail("recvStr is <null>");
      return;
    }
    assertEquals(rcv.get("topic"), topic);
    assertEquals(rcv.get("data"), testData);

    zpub.close();
    zsub.close();
  }
}
