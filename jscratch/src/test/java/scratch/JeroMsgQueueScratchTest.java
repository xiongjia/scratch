package scratch;

import static org.junit.Assert.assertEquals;

import org.junit.After;
import org.junit.Before;
import org.junit.Ignore;
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

  @Ignore("Test is ignored")
  @Test
  public void testPubSubPattern()
      throws InterruptedException, ExecutionException, TimeoutException {
    final String topic = "zscratch";
    final String testData = "test-data-123";

    ZChannel zpub = jeromq.buildPublisher();
    ZChannel zsub = jeromq.buildSubscriber(zpub.getBindEndPoint(), topic);
    zpub.send(topic, testData);

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

    ExecutorService exec = Executors.newFixedThreadPool(1);
    Callable<Map<String, String>> recv = new Recv(zsub);
    Future<Map<String, String>> future = exec.submit(recv);
    final Map<String, String> recvData = future.get(30, TimeUnit.SECONDS);
    assertEquals(recvData.get("topic"), topic);
    assertEquals(recvData.get("data"), testData);

    zpub.close();
    zsub.close();
  }
}
