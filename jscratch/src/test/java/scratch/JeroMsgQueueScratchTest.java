package scratch;

import static org.junit.Assert.assertEquals;

import org.junit.After;
import org.junit.Before;
import org.junit.Test;

import scratch.JeroMsgQueueScratch.ZChannel;

import java.util.Map;
import java.util.concurrent.Callable;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.TimeoutException;

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

  @Test
  public void testPubSubPattern()
      throws InterruptedException, ExecutionException, TimeoutException {
    final String topic = "zscratch";
    final String testData = "test-data-123";
    final ExecutorService exec = Executors.newFixedThreadPool(1);

    final ZChannel zpub = jeromq.buildPublisher();
    final ZChannel zsub = jeromq.buildSubscriber(zpub.getBindEndPoint(), topic);

    class Recv implements Callable<Map<String, String>> {
      final long timeoutMilliseconds = 1000 * 30;
      final ZChannel zsub;

      public Recv(ZChannel zsub) {
        this.zsub = zsub;
      }

      @Override
      public Map<String, String> call() throws Exception {
        return this.zsub.recvStr(timeoutMilliseconds);
      }
    }

    Future<Map<String, String>> future = exec.submit(new Recv(zsub));

    /* zsub.recvStr() must started before zpub.send() */
    Thread.sleep(1000);
    zpub.send(topic, testData);

    final Map<String, String> recvData = future.get(30, TimeUnit.SECONDS);
    assertEquals(recvData.get("topic"), topic);
    assertEquals(recvData.get("data"), testData);

    zpub.close();
    zsub.close();
  }
}
