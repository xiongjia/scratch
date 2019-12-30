package snow.tmp;

import java.util.Arrays;

import org.junit.Test;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class QuickSortTest {
  private static final Logger logger = LoggerFactory.getLogger(QuickSortTest.class);

  private void qsort(int[] data) {
    // TODO
  }

  @Test
  public void test() {
    final int[] data = new int[]{11, 2, 7, 8, 15};
    System.out.println("before sort: " + Arrays.toString(data));
    qsort(data);
    System.out.println("after sort: " + Arrays.toString(data));
  }
}
