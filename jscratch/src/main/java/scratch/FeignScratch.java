package scratch;

import feign.Client;
import feign.Feign;
import feign.Headers;
import feign.Request;
import feign.Request.Options;
import feign.RequestLine;
import feign.Response;
import feign.gson.GsonDecoder;
import feign.gson.GsonEncoder;

import lombok.Builder;
import lombok.Data;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;

class FeignScratch {
  private static final Logger log = LoggerFactory.getLogger(FeignScratch.class);

  public static class MockClient implements Client {
    @Override
    public Response execute(Request input, Options options) throws IOException {
      log.debug("Request: {} [{}]",  input.url(), input.method());
      if (input.body() != null) {
        log.debug("Sending: {}",
            input.charset() != null ? new String(input.body(), input.charset()) : "BIN");
      }
      return Response.create(200, "reason", input.headers(), "mock".getBytes());
    }
  }

  @Builder
  @Data
  public static class PostData {
    private String data;
    private int num;
  }

  public static interface TestApi {
    @RequestLine("GET /api/v1.0/ping")
    String ping();

    @RequestLine("POST /api/v1.0/pong")
    @Headers("Content-Type: application/json")
    void pong(PostData data);
  }

  static void feignTest() {
    final int connectTimeoutMillis = 1000 * 20;
    final int readTimeoutMillis = 1000 * 30;

    final TestApi api = Feign.builder()
        .options(new Options(connectTimeoutMillis, readTimeoutMillis))
        .decoder(new GsonDecoder())
        .encoder(new GsonEncoder())
        .client(new MockClient())
        .target(TestApi.class, "http://localhost:8090");

    final String ping = api.ping();
    log.debug("mock api result: {}", ping);

    api.pong(new PostData.PostDataBuilder().data("string").num(100).build());
  }
}
