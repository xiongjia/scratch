/**
 * gazlowe - My algorithm tests
 */

#include <algorithm>
#include <vector>
#include "boost/test/unit_test.hpp"

namespace gazlowe
{
    /* Single Number:
     * https://leetcode.com/problems/single-number/
     *
     * Given an array of integers, every element appears twice except for one.
     * Find that single one.
     */
    class SingleNumber 
    {
    public:
        static int singleNumber(std::vector<int> &nums)
        {
            int ret = 0;
            std::for_each(nums.begin(), nums.end(),
                [&](int &num) { ret ^= num; });
            return ret;
        }
    };

    /* Single Number II
     * https://leetcode.com/problems/single-number-ii/
     * 
     * Given an array of integers, every element appears 
     * three times except for one. Find that single one. 
     *
     * Note:
     *  Your algorithm should have a linear runtime complexity. 
     *  Could you implement it without using extra memory? 
     */
    class SingleNumber2
    {
    public:
        static int singleNumber(std::vector<int> &nums)
        {
            std::size_t bitsMap[32] = { 0 };
            int ret = 0;
            for (std::size_t i = 0; i < 32; ++i)
            {
                std::for_each(nums.begin(), nums.end(), [&](const int num)
                {
                    bitsMap[i] += (num >> i) & 0x01;
                });

                ret |= (bitsMap[i] % 3 << i);
            }
            return ret;
        }

        static int singleNumber2(std::vector<int> &nums)
        {
            int one = 0;
            int two = 0;
            int three = 0;
            std::for_each(nums.begin(), nums.end(), [&](const int num)
            {
                two |= one & num;
                one ^= num;
                three = one & two;
                one &= ~three;
                two &= ~three;
            });
            return one;
        }
    };

    /* Single Number III 
     * https://leetcode.com/problems/single-number-iii/
     *
     * Given an array of numbers nums, in which exactly two elements appear 
     * only once and all the other elements appear exactly twice. Find the 
     * two elements that appear only once.
     *
     * For example: 
     * Given nums = [1, 2, 1, 3, 2, 5], return [3, 5]. 
     * Note:
     *   1. The order of the result is not important. 
     *      So in the above example, [5, 3] is also correct.
     *   2. Your algorithm should run in linear runtime complexity. 
     *      Could you implement it using only constant space complexity?
     */
    class SingleNumber3
    {
    public:
        static std::vector<int> singleNumber(std::vector<int> &nums)
        {
            int res = 0;
            std::for_each(nums.begin(), nums.end(), 
                         [&](const int num) { res ^= num; });
            int cmp = (res & (res - 1)) ^ res;
            int a = 0;
            int b = 0;
            std::for_each(nums.begin(), nums.end(), [&](const int num)
            { 
                if (cmp & num)
                {
                    a ^= num;
                }
                else
                {
                    b ^= num;
                }
            });
            return std::vector<int>({ {a, b } });
        }
    };
}

BOOST_AUTO_TEST_SUITE(leetcode)
BOOST_AUTO_TEST_CASE(single_number)
{
    /* Testing gazlowe::single_number() */
    BOOST_TEST_MESSAGE("Leetcode - Single Number");
    std::vector<int> testData = { 1, 1, 2, 2, 3, 3, 5 };
    int ret = gazlowe::SingleNumber::singleNumber(testData);
    BOOST_REQUIRE_EQUAL(ret, 5);

    std::vector<int> testData2 = { 2, 3, 2, 2};
    ret = gazlowe::SingleNumber2::singleNumber(testData2);
    BOOST_REQUIRE_EQUAL(ret, 3);
    ret = gazlowe::SingleNumber2::singleNumber2(testData2);
    BOOST_REQUIRE_EQUAL(ret, 3);

    std::vector<int> testData3 = { 1, 2, 1, 3, 2, 5 };
    auto retData = gazlowe::SingleNumber3::singleNumber(testData3);
    BOOST_REQUIRE_EQUAL(retData.size(), 2);
    BOOST_REQUIRE((retData[0] == 3 && retData[1] == 5) ||
                  (retData[0] == 5 && retData[1] == 3));
}
BOOST_AUTO_TEST_SUITE_END()
