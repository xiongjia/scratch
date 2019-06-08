package snow.rpc;

import org.junit.Test;

import javax.ws.rs.client.WebTarget;
import javax.ws.rs.core.Response;

public class RestClientTest {
  @Test
  public void test() {
    final RestClient client = new RestClient();
    final WebTarget target = client.makeTarget("http://192.168.1.37:8111", "get");
    final Response response = client.invokeGet(target);
    final String test = response.readEntity(String.class);
    System.out.println(test);
  }
}
