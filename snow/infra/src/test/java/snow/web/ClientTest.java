package snow.web;

import org.junit.Test;

import java.io.IOException;

public class ClientTest {
  @Test
  public void test() {
    final Client client = new Client();
    for (int i = 0; i < 100; ++i) {
      client.request("http://192.168.1.37:8200");
      client.request("https://www.douban.com/");
    }
  }
}
