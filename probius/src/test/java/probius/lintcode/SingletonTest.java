package probius.lintcode;

import static org.junit.Assert.assertEquals;

import org.junit.Test;

public class SingletonTest {
  @Test
  public void test() {
    final Singleton instA = Singleton.getInstance();
    final Singleton instB = Singleton.getInstance();
    assertEquals(instA, instB);
  }
}
