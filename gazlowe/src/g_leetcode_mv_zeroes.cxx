/**
 * gazlowe - My algorithm tests
 */

#include <vector>
#include "boost/test/unit_test.hpp"
#include "boost/foreach.hpp"

namespace gazlowe
{
    /* Move Zeroes
     * https://leetcode.com/problems/move-zeroes/
     *
     * Given an array nums, write a function to move all 0's to
     * the end of it while maintaining the relative order of
     * the non-zero elements. 
     *
     * For example, given nums = [0, 1, 0, 3, 12], after calling
     * your function, nums should be [1, 3, 12, 0, 0]. 
     * Note:
     *   - You must do this in-place without making a copy of the array.
     *   - Minimize the total number of operations.
     */
    class MVZeroes
    {
    public:
        static void moveZeroes(std::vector<int> &nums)
        {
            std::size_t cnt = 0;
            std::size_t sz = nums.size();
            for (std::size_t itr = 0; itr < sz;)
            {
                if (nums[itr++] != 0)
                {
                    std::swap(nums[cnt++], nums[itr - 1]);
                }
            }
        }
    };
}

BOOST_AUTO_TEST_SUITE(leetcode)
BOOST_AUTO_TEST_CASE(mv_zeroes)
{
    BOOST_TEST_MESSAGE("Leetcode - Move Zeroes");

    auto data = std::vector<int>({ 0, 1, 0, 3, 12 });
    gazlowe::MVZeroes::moveZeroes(data);
    auto result = std::vector<int>({ 1, 3, 12, 0, 0 });
    BOOST_CHECK_EQUAL_COLLECTIONS(data.begin(), data.end(),
                                  result.begin(), result.end());

}
BOOST_AUTO_TEST_SUITE_END()
