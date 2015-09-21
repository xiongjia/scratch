/**
 * gazlowe - My algorithm tests
 */

#include <vector>
#include "boost/test/unit_test.hpp"

namespace gazlowe
{
    /* Find Minimum in Rotated Sorted Array
     * https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
     *
     * Suppose a sorted array is rotated at some pivot 
     * unknown to you beforehand.
     * (i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).
     *
     * Find the minimum element. 
     * You may assume no duplicate exists in the array.
     */
    class FindMinInRotatedSortedArray
    {
    public:
        static int findMin(std::vector<int> &nums)
        {
            std::size_t len = nums.size();
            if (nums[0] <= nums[len - 1])
            {
                return nums[0];
            }
            std::size_t itr = 0;
            std::size_t tmp = len - 1;
            for (; itr + 1 < tmp;)
            {
                std::size_t k = (itr + tmp) / 2;
                if (nums[itr] > nums[k])
                {
                    tmp = k;
                }
                else 
                {
                    itr = k;
                }
            }
            return nums[tmp];
        }
    };
}

BOOST_AUTO_TEST_SUITE(leetcode)
BOOST_AUTO_TEST_CASE(find_min_in_rotated_sorted_array)
{
    BOOST_TEST_MESSAGE("Leetcode - Find Minimum in Rotated Sorted Array");
    auto data = std::vector<int>({{ 4, 5, 6, 7, 0, 1, 2 }});
    int ret = gazlowe::FindMinInRotatedSortedArray::findMin(data);
    BOOST_REQUIRE_EQUAL(ret, 0);
}
BOOST_AUTO_TEST_SUITE_END()
