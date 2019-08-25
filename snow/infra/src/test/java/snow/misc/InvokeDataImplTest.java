package snow.misc;

import org.junit.Test;

import java.lang.reflect.Proxy;

public class InvokeDataImplTest {
  @Test
  public void test() {
    final InvokeData invokeData = (InvokeData)Proxy.newProxyInstance(InvokeDataImplTest.class.getClassLoader(),
        new Class<?>[] { InvokeData.class },
        new InvokeDataImpl());
    System.out.println(invokeData.echo("test"));
  }
}
