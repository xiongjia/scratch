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
  /** Solution for problem "two sum". */
  public int[] twoSum(int[] nums, int target) {
    final Map<Integer, Integer> map = new HashMap<>();

    for (int i = 0; i < nums.length; i++) {
      if (!map.containsKey(target - nums[i])) {
        map.put(nums[i], i);
        continue;
      }

      final int[] result = new int[2];
      result[0] = map.get(target - nums[i]);
      result[1] = i;
      return result;
    }
    return null;
  }
}
