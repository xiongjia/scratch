package transmuter;

import org.junit.Test;

import static org.junit.Assert.*;

public class Problem0001TwoSumTest {
  @Test
  public void twoSum() {
    final Problem0001TwoSum problem = new Problem0001TwoSum();

    final int[] testSorce1 = {2, 7, 11, 15};
    final int[] result = problem.twoSum(testSorce1, 9);
    assertTrue(testSorce1[result[0]] + testSorce1[result[1]] == 9);
  }
}
