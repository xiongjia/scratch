package probius.lintcode;

import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;

/**
 * LintCode exercise - 3 sum http://www.lintcode.com/en/problem/3sum/
 *  Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0?
 *  Find all unique triplets in the array which gives the sum of zero.
 * Notice
 *  Elements in a triplet (a,b,c) must be in non-descending order. (ie, a <= b <= c)
 *  The solution set must not contain duplicate triplets.
 * Example
 *  For example, given array S = {-1 0 1 2 -1 -4}, A solution set is:
 *  ```
 *  (-1, 0, 1)
 *  (-1, -1, 2)
 *  ```
 */
public class ThreeSum {
  /** solution. */
  public List<List<Integer>> threeSum(int[] numbers) {
    final List<List<Integer>> result = new LinkedList<>();
    if (numbers == null) {
      return result;
    }

    final int numsLen = numbers.length;
    if (numsLen <= 2) {
      return result;
    }

    Arrays.sort(numbers);
    final int min = numbers[0];
    final int max = numbers[numsLen - 1];
    if (min * 3 > 0 || max * 3 < 0) {
      return result;
    }

    for (int i = 0; i < numsLen - 2; ++i) {
      if (i != 0 && (numbers[i] == numbers[i - 1])) {
        continue;
      }

      int lo = i + 1;
      int hi = numsLen - 1;
      final int sum = 0 - numbers[i];
      while (lo < hi) {
        if (numbers[lo] + numbers[hi] == sum) {
          result.add(Arrays.asList(numbers[i], numbers[lo], numbers[hi]));
          while (lo < hi && numbers[lo] == numbers[lo + 1]) {
            lo++;
          }
          while (lo < hi && numbers[hi] == numbers[hi - 1]) {
            hi--;
          }
          lo++;
          hi--;
        } else if (numbers[lo] + numbers[hi] < sum) {
          lo++;
        } else {
          hi--;
        }
      }
    }
    return result;
  }
}
