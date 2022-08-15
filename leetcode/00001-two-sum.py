#!/usr/bin/env python


'''
Given an array of integers nums and an integer target, 
return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, 
and you may not use the same element twice.
You can return the answer in any order.

Example 1:
  Input: nums = [2,7,11,15], target = 9
  Output: [0,1]
  Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:
  Input: nums = [3,2,4], target = 6
  Output: [1,2]

Example 3:
  Input: nums = [3,3], target = 6
  Output: [0,1]


Constraints:
    2 <= nums.length <= 104
    -109 <= nums[i] <= 109
    -109 <= target <= 109
    Only one valid answer exists.

Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?

'''

from typing import List

class Solution:
  def twoSum(self, nums: List[int], target: int) -> List[int]:
    cache = {}
    for i in range(len(nums)):
      difference = target - nums[i]
      if difference in cache:
        return [cache[difference], i]
      cache[nums[i]] = i
    return [0, 0]

if __name__ == '__main__':
  s = Solution()
  val = s.twoSum([3, 2, 4], 6)
  print('val =  ' + ' '.join(str(item) for item in val))

