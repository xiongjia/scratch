package snow.client;

import javax.ws.rs.core.Response;

import org.apache.http.client.HttpResponseException;

import org.glassfish.jersey.client.authentication.HttpAuthenticationFeature;
import snow.rpc.RestClient;

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
}
