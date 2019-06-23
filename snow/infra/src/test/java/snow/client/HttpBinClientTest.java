package snow.client;

import com.google.common.util.concurrent.FutureCallback;

import org.apache.http.client.HttpResponseException;

import org.checkerframework.checker.nullness.qual.Nullable;

import org.junit.Ignore;
import org.junit.Test;

public class HttpBinClientTest {
  @Ignore
  @Test
  public void testRequestGet() {
    final HttpBinClient httpBinClient = new HttpBinClient("http://192.168.1.37:8111/");
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
      System.out.println("call ended");
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
