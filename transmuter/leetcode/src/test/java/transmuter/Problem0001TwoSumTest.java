package transmuter;

import org.junit.Assert;
import org.junit.Test;

public class Problem0001TwoSumTest {
  final Problem0001TwoSum problem = new Problem0001TwoSum();

  @Test
  public void twoSumTest1() {
    final int[] testSrc = {2, 7, 11, 15};
    final int[] result = problem.twoSum(testSrc, 9);
    Assert.assertEquals(testSrc[result[0]] + testSrc[result[1]], 9);
  }

  @Test
  public void twoSumTest2() {
    final int[] testSrc = {};
    final int[] result = problem.twoSum(testSrc, 0);
    Assert.assertNull(result);
  }

  @Test
  public void twoSumTest3() {
    final int[] testSrc = {2, 7, 8, 8, 0};
    final int[] result = problem.twoSum(testSrc, 10);
    Assert.assertEquals(testSrc[result[0]] + testSrc[result[1]], 10);
  }

  @Test
  public void twoSumTest4() {
    final int[] testSrc = {1, 3, 7, 9, 0};
    final int[] result = problem.twoSum(testSrc, 11);
    Assert.assertNull(result);
  }
}
