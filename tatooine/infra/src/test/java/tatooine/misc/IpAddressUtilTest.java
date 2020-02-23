package tatooine.misc;

import org.junit.Test;

public class IpAddressUtilTest {
  @Test
  public void test() {
    IpAddressUtil.convertNonSequentialBlock("192.168.1.0/24");
  }
}
