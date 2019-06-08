package snow.client;


import com.google.common.util.concurrent.FutureCallback;
import org.apache.http.client.HttpResponseException;
import org.checkerframework.checker.nullness.qual.Nullable;
import org.junit.Test;

public class HttpBinClientTest {
  @Test
  public void testRequestGet() {
    final HttpBinClient httpBinClient = new HttpBinClient("http://192.168.1.37:8111/");

//    try {
//      final ResponseHttpGet responseHttpGet1 = httpBinClient.requestGet();
//      System.out.println(responseHttpGet1);
//      final ResponseHttpGet responseHttpGet2 = httpBinClient.requestGet();
//      System.out.println(responseHttpGet2);
//    } catch (HttpResponseException e) {
//      e.printStackTrace();
//    }

    try {
      httpBinClient.asyncRequestGet(new FutureCallback<ResponseHttpGet>() {
        @Override
        public void onSuccess(@Nullable ResponseHttpGet result) {
          if (result == null) {
            return;
          }
          System.out.println(result.toString());
        }

        @Override
        public void onFailure(Throwable t) {
          System.out.println("err: " + t.toString());
        }
      });
    } catch (HttpResponseException e) {
      e.printStackTrace();
    }

    try {
      Thread.sleep(1000 * 5);
    } catch (InterruptedException e) {
      e.printStackTrace();
    }
  }
}
