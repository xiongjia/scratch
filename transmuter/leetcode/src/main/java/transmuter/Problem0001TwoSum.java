package transmuter;

import java.util.HashMap;
import java.util.Map;

/* Problem 1: Two Sum (easy） https://leetcode.com/problems/two-sum/
 *
 * Given an array of integers, return indices of the two numbers
 * such that they add up to a specific target.
 * You may assume that each input would have exactly one solution,
 * and you may not use the same element twice.
 *
 * Example:
 *  Given nums = [2, 7, 11, 15], target = 9,
 *  Because nums[0] + nums[1] = 2 + 7 = 9,
 *  return [0, 1].
 */
public class Problem0001TwoSum {
  /** Solution 001 - HashMap to save the numbers. */
  public int[] solution001(int[] nums, int target) {
    final Map<Integer, Integer> map = new HashMap<>();

    for (int i = 0; i < nums.length; i++) {
      final int delta = target - nums[i];
      if (!map.containsKey(delta)) {
        map.put(nums[i], i);
        continue;
      }
      return new int[] {map.get(delta), i};
    }
    return null;
  }

  /** Solution 002 - 2 for loops. */
  public int[] solution002(int[] nums, int target) {
    for (int i = 0; i < nums.length; i++) {
      final int delta = target - nums[i];
      for (int j = i + 1; j < nums.length; j++) {
        if (delta != nums[j]) {
          continue;
        }
        return new int[] {i, j};
      }
    }
    return null;
  }
}
