package snow.web;

import org.junit.Test;

import java.io.IOException;
import java.util.stream.IntStream;

public class ClientTest {
  @Test
  public void test() {
    final Client client = new Client();

    for (int i = 0; i < 10; ++i) {
      IntStream.rangeClosed(1, 10).parallel().forEach(item -> {
        client.request("http://192.168.1.37:8200");
        client.request("https://www.douban.com/");
      });

    }
  }
}
