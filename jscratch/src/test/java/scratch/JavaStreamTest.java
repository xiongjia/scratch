package scratch;

import org.junit.Test;

import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

public class JavaStreamTest {
  final int testNumsSize = 1024 * 1024 * 5;
  final List<Integer> testNums = IntStream.range(0, testNumsSize)
      .boxed()
      .collect(Collectors.toList());

  @Test
  public void scanNumsForEach() {
    for (Integer num : testNums) {
      @SuppressWarnings("unused")
      int testVal = num;
    }
  }

  @Test
  public void scanNumsStreamForEach() {
    testNums.stream().forEach((num) -> {
      @SuppressWarnings("unused")
      int testVal = num;
    });
  }

  @Test
  public void scanNumsStreamParallelForEach() {
    testNums.stream().parallel().forEach((num) -> {
      @SuppressWarnings("unused")
      int testVal = num;
    });
  }
}
