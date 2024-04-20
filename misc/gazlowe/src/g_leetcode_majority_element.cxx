/**
 * gazlowe - My algorithm tests
 */

#include <vector>
#include "boost/test/unit_test.hpp"

namespace gazlowe
{
    /* Majority Element
     * https://leetcode.com/problems/majority-element/
     *
     * Given an array of size n, find the majority element.
     * The majority element is the element that appears more than [n/2] times.
     *
     * You may assume that the array is non-empty and 
     * the majority element always exist in the array.
     */
    class MajorityElement
    {
    public:
        static int majorityElement(std::vector<int> &nums)
        {
            std::size_t cnt = 0;
            int candidate = 0;
            for (auto itr = nums.begin(); itr != nums.end(); ++itr)
            {
                if (0 == cnt)
                {
                    candidate = *itr;
                    cnt = 1;
                    continue;
                }

                cnt = (*itr == candidate) ? cnt + 1 : cnt - 1;
            }
            return candidate;
        }
    };

    /* Majority Element II
     * https://leetcode.com/problems/majority-element-ii/
     *
     * Given an integer array of size n, find all elements that appear
     * more than [n/3] times. The algorithm should run in linear time
     * and in O(1) space.
     *
     * Hint: How many majority elements could it possibly have?
     *
     */
    class MajorityElement2
    {
    public:
        static std::vector<int> majorityElement(std::vector<int> &nums)
        {
            int candidate1 = 0;
            int candidate2 = 0;
            std::size_t candidate1Cnt = 0;
            std::size_t candidate2Cnt = 0;
            for (auto itr = nums.begin(); itr != nums.end(); ++itr)
            {
                if (0 != candidate1Cnt && candidate1 == *itr)
                {
                    candidate1Cnt++;
                    continue;
                }
                if (0 != candidate2Cnt && candidate2 == *itr)
                {
                    candidate2Cnt++;
                    continue;
                }

                if (0 == candidate1Cnt)
                {
                    candidate1 = *itr;
                    candidate1Cnt = 1;
                    continue;
                }
                if (0 == candidate2Cnt)
                {
                    candidate2 = *itr;
                    candidate2Cnt = 1;
                    continue;
                }

                candidate1Cnt--;
                candidate2Cnt--;
            }

            std::size_t candidate1Nums = 0;
            std::size_t candidate2Nums = 0;
            for (auto itr = nums.begin(); itr != nums.end(); ++itr)
            {
                if (*itr == candidate1)
                {
                    candidate1Nums++;
                }
                if (*itr == candidate2)
                {
                    candidate2Nums++;
                }
            }

            std::vector<int> ret;
            if (candidate1Nums > nums.size() / 3)
            {
                ret.push_back(candidate1);
            }
            if (candidate2 != candidate1 && candidate2Nums > nums.size() / 3)
            {
                ret.push_back(candidate2);
            }
            return ret;
        }
    };
}

BOOST_AUTO_TEST_SUITE(leetcode)
BOOST_AUTO_TEST_CASE(majority_element)
{
    BOOST_TEST_MESSAGE("Leetcode - Majority Element");

    auto data = std::vector<int>({ 1 });
    int ret = gazlowe::MajorityElement::majorityElement(data);
    BOOST_REQUIRE_EQUAL(ret, 1);

    data = std::vector<int>({ 1, 2, 2 });
    ret = gazlowe::MajorityElement::majorityElement(data);
    BOOST_REQUIRE_EQUAL(ret, 2);

    data = std::vector<int>({ 1, 1, 1, 2, 2, 2, 2});
    ret = gazlowe::MajorityElement::majorityElement(data);
    BOOST_REQUIRE_EQUAL(ret, 2);
}

BOOST_AUTO_TEST_CASE(majority_element2)
{
    BOOST_TEST_MESSAGE("Leetcode - Majority Element II");

    auto data = std::vector<int>({ 1, 1, 1, 2, 2, 2 });
    auto ret = gazlowe::MajorityElement2::majorityElement(data);
    auto expect = std::vector<int>({ 1, 2 });
    BOOST_CHECK_EQUAL_COLLECTIONS(ret.begin(), ret.end(),
                                  expect.begin(), expect.end());

    data = std::vector<int>({ 3, 2, 3});
    ret = gazlowe::MajorityElement2::majorityElement(data);
    expect = std::vector<int>({ 3 });
    BOOST_CHECK_EQUAL_COLLECTIONS(ret.begin(), ret.end(),
                                  expect.begin(), expect.end());

    data = std::vector<int>({ 0, 0, 0 });
    ret = gazlowe::MajorityElement2::majorityElement(data);
    expect = std::vector<int>({ 0 });
    BOOST_CHECK_EQUAL_COLLECTIONS(ret.begin(), ret.end(),
                                  expect.begin(), expect.end());

    data = std::vector<int>({ 2, 2, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9 });
    ret = gazlowe::MajorityElement2::majorityElement(data);
    expect = std::vector<int>({ 9, 3 });
    BOOST_CHECK_EQUAL_COLLECTIONS(ret.begin(), ret.end(),
                                  expect.begin(), expect.end());
}
BOOST_AUTO_TEST_SUITE_END()
