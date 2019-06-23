package snow.rpc;

import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;

import org.junit.Ignore;
import org.junit.Test;

public class RestClientTest {
  @Ignore
  @Test
  public void test() {
    final RestClient client = new RestClient();
    final Response response = client.target("http://192.168.1.37:8111/")
        .path("get")
        .request(MediaType.APPLICATION_JSON)
        .get();
    System.out.println("status: " + response.getStatus());
    final String result = response.readEntity(String.class);
    System.out.println("content: " + result);
  }
}
