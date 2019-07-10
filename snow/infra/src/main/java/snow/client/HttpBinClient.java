package snow.client;

import com.google.common.util.concurrent.FutureCallback;
import com.google.common.util.concurrent.Futures;
import com.google.common.util.concurrent.ListenableFuture;
import com.google.common.util.concurrent.MoreExecutors;

import javax.ws.rs.core.Response;

import org.apache.http.client.HttpResponseException;

import org.glassfish.jersey.client.authentication.HttpAuthenticationFeature;
import org.glassfish.jersey.client.rx.guava.RxListenableFutureInvoker;
import snow.rpc.RestClient;

import java.util.concurrent.ExecutionException;
import java.util.concurrent.Executor;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.TimeoutException;

/** The REST client for httpbin ( https://httpbin.org/ ) */
public class HttpBinClient {
  private static final String DEFAULT_API_ROOT = "https://httpbin.org/";
  private final String apiRoot;
  private final RestClient restClient;

  public HttpBinClient(String apiRoot, RestClient restClient) {
    this.apiRoot = apiRoot;
    this.restClient = restClient;
  }

  public HttpBinClient() {
    this(DEFAULT_API_ROOT, new RestClient());
  }

  public HttpBinClient(String apiRoot) {
    this(apiRoot, new RestClient());
  }

  /** HttpBin Get request. */
  public ResponseHttpGet requestGet() throws HttpResponseException {
    final Response response = restClient.target(apiRoot)
        .path("get")
        .queryParam("test", "abc", "test2", 1, "test3", true)
        .register(HttpAuthenticationFeature.basic("username", "password"))
        .request()
        .get();

    final int responseStatus = response.getStatus();
    if (responseStatus != Response.Status.OK.getStatusCode()) {
      throw new HttpResponseException(responseStatus,
          String.format("HTTP Error: %d", responseStatus));
    }
    return response.readEntity(ResponseHttpGet.class);
  }

  /** HttpBin Get request (Async). */
  public void asyncRequestGet(FutureCallback<ResponseHttpGet> callback)
      throws HttpResponseException {
    final ListenableFuture<ResponseHttpGet> responseFuture = restClient.target(apiRoot)
        .path("get")
        .queryParam("test", "abc", "test2", 1, "test3", true)
        .register(HttpAuthenticationFeature.basic("username", "password"))
        .request()
        .rx(RxListenableFutureInvoker.class)
        .get(ResponseHttpGet.class);
    Futures.addCallback(responseFuture, callback, MoreExecutors.directExecutor());
  }
}
