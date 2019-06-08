package snow.client;


import org.apache.http.client.HttpResponseException;
import org.junit.Test;

public class HttpBinClientTest {
  @Test
  public void testRequestGet() {
    final HttpBinClient httpBinClient = new HttpBinClient("http://192.168.1.37:8111/");

    try {
      final ResponseHttpGet responseHttpGet1 = httpBinClient.requestGet();
      System.out.println(responseHttpGet1);

      final ResponseHttpGet responseHttpGet2 = httpBinClient.requestGet();
      System.out.println(responseHttpGet2);
    } catch (HttpResponseException e) {
      e.printStackTrace();
    }
  }
}
