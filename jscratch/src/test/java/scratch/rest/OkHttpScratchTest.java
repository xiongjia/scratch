package scratch.rest;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertFalse;
import static org.junit.Assert.assertNotEquals;

import okhttp3.HttpUrl;

import okhttp3.mockwebserver.MockResponse;
import okhttp3.mockwebserver.MockWebServer;

import org.junit.After;
import org.junit.Before;
import org.junit.Test;

import scratch.rest.OkHttpScratch.ResponseData;

import java.io.IOException;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.TimeoutException;

public class OkHttpScratchTest {
  private final MockWebServer mockServ = new MockWebServer();

  /** creating mock server. */
  @Before
  public final void setUp() throws IOException {
    /* disable mock server logger and start the mock server */
    java.util.logging.Logger.getLogger(MockWebServer.class.getName())
      .setLevel(java.util.logging.Level.OFF);
    mockServ.start();
  }

  @After
  public final void tearDown() throws IOException {
    /* stop mock server */
    mockServ.shutdown();
  }

  @Test
  public void baseTest() throws IOException {
    final String mockBody = this.getClass().getName() + "baseTest";
    final HttpUrl mockUrl = mockServ.url("/");
    mockServ.enqueue(new MockResponse().setResponseCode(200).setBody(mockBody));

    OkHttpScratch okHttp = new OkHttpScratch.Builder().build();
    OkHttpScratch.ResponseData rep = okHttp.get(mockUrl.toString());

    assertEquals(rep.getCode(), 200);
    assertEquals(rep.getBody(), mockBody);
    assertFalse("Response Headers", rep.getHeaders().isEmpty());
    assertEquals(rep.getHeaders().get("Content-Length"), String.format("%d", mockBody.length()));
  }

  @Test
  public void asyncTest() throws IOException,
          InterruptedException, ExecutionException, TimeoutException {
    final String mockBody = this.getClass().getName() + "asyncTest";
    final HttpUrl mockUrl = mockServ.url("/");
    mockServ.enqueue(new MockResponse().setResponseCode(200).setBody(mockBody));

    OkHttpScratch okHttp = new OkHttpScratch.Builder().build();
    CompletableFuture<ResponseData> rep = new CompletableFuture<>();
    okHttp.getAsync(mockUrl.toString(), new OkHttpScratch.AsyncCallback() {
      @Override
      public void onResponse(ResponseData response) {
        rep.complete(response);
      }

      @Override
      public void onFailure(IOException err) {
        rep.complete(null);
      }
    });
    ResponseData ret = rep.get(30, TimeUnit.SECONDS);
    assertNotEquals(ret, null);

    assertEquals(ret.getCode(), 200);
    assertEquals(ret.getBody(), mockBody);
    assertFalse("Response Headers", ret.getHeaders().isEmpty());
    assertEquals(ret.getHeaders().get("Content-Length"),
            String.format("%d", mockBody.length()));

    okHttp.shutdownClient();
  }

  @Test(expected = IOException.class)
  public void timeoutTest() throws IOException {
    final HttpUrl mockUrl = mockServ.url("/");
    OkHttpScratch okHttp = new OkHttpScratch.Builder()
        .setTimeoutConn(3).setTimeoutRd(3).setTimeoutWr(3)
        .build();
    okHttp.get(mockUrl.toString());
  }
}
