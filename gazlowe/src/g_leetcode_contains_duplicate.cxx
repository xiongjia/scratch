/**
 * gazlowe - My algorithm tests
 */

#include <vector>
#include <set>
#include <map>
#include "boost/test/unit_test.hpp"
#include "g_misc.hxx"

namespace gazlowe
{
    /* Contains Duplicate
     * https://leetcode.com/problems/contains-duplicate/
     *
     * Given an array of integers, find if the array contains any duplicates.
     * Your function should return true if any value appears at least twice
     * in the array, and it should return false if every element is distinct.
     */
    class ContainsDuplicate
    {
    public:
        static bool containsDuplicate(std::vector<int>& nums)
        {
            std::set<int> contains;
            for (auto itr = nums.begin(); itr != nums.end(); ++itr)
            {
                if (contains.find(*itr) != contains.end())
                {
                    return true;
                }
                contains.insert(*itr);
            }
            return false;
        }
    };

    /* Contains Duplicate II
     * https://leetcode.com/problems/contains-duplicate-ii/ 
     * 
     * Given an array of integers and an integer k, find out whether
     * there are two distinct indices i and j in the array such that
     * nums[i] = nums[j] and the difference between i and j is at most k. 
     */
    class ContainsDuplicate2 
    {
    public:
        static bool containsNearbyDuplicate(std::vector<int> &nums, int k)
        {
            if (k <= 0 || nums.empty())
            {
                return false;
            }

            std::map<int, int> contains;
            int idx = 0;
            for (auto itr = nums.begin(); itr != nums.end(); ++itr, ++idx)
            {
                auto find = contains.find(*itr);
                if (find != contains.end() &&
                    (idx - find->second) <= k)
                {
                    return true;
                }
                contains[*itr] = idx;
            }
            return false;
        }
    };

    /* Contains Duplicate III
     * https://leetcode.com/problems/contains-duplicate-iii/
     *
     * Given an array of integers, find out whether there are two distinct
     * indices i and j in the array such that the difference between nums[i]
     * and nums[j] is at most t and the difference between i and j is at most k. 
     */
    class ContainsDuplicate3
    {
    public:
        static bool containsNearbyAlmostDuplicate(std::vector<int> &nums,
                                                  int k, int t)
        {
            if (k <= 0)
            {
                return false;
            }

            std::multiset<int64_t> contains;
            for (std::size_t i = 0; i < nums.size(); ++i)
            {
                if (contains.size() > static_cast<unsigned>(k))
                {
                    contains.erase(contains.find(nums[i - k - 1]));
                }

                auto lower = contains.lower_bound(nums[i] - t);
                if (lower != contains.end() && *lower - nums[i] <= t)
                {
                    return true;
                }

                contains.insert(nums[i]);
            }
            return false;
        }
    };
}

BOOST_AUTO_TEST_SUITE(leetcode)
BOOST_AUTO_TEST_CASE(contains_duplicate)
{
    BOOST_TEST_MESSAGE("Leetcode - Contains Duplicate");

    bool result = gazlowe::ContainsDuplicate::containsDuplicate(
        std::vector<int>({ { 1, 2, 3, 4, 5 } }));
    BOOST_REQUIRE_EQUAL(result, false);

    result = gazlowe::ContainsDuplicate::containsDuplicate(
        std::vector<int>({ { 1, 2, 3, 4, 1 } }));
    BOOST_REQUIRE_EQUAL(result, true);

    result = gazlowe::ContainsDuplicate::containsDuplicate(
        std::vector<int>({}));
    BOOST_REQUIRE_EQUAL(result, false);
}

BOOST_AUTO_TEST_CASE(contains_duplicate2)
{
    BOOST_TEST_MESSAGE("Leetcode - Contains Duplicate II");

    bool result = gazlowe::ContainsDuplicate2::containsNearbyDuplicate(
        std::vector<int>({ { 1, 2, 3, 4, 5 } }), 2);
    BOOST_REQUIRE_EQUAL(result, false);

    result = gazlowe::ContainsDuplicate2::containsNearbyDuplicate(
        std::vector<int>({ { 1, 2, 1, 4, 5 } }), 2);
    BOOST_REQUIRE_EQUAL(result, true);

    result = gazlowe::ContainsDuplicate2::containsNearbyDuplicate(
        std::vector<int>({ { 1, 2, 1, 4, 5 } }), 1);
    BOOST_REQUIRE_EQUAL(result, false);

    result = gazlowe::ContainsDuplicate2::containsNearbyDuplicate(
        std::vector<int>({ { 1, 2, 1, 4, 5 } }), 0);
    BOOST_REQUIRE_EQUAL(result, false);

    result = gazlowe::ContainsDuplicate2::containsNearbyDuplicate(
        std::vector<int>({}), 1);
    BOOST_REQUIRE_EQUAL(result, false);
}

BOOST_AUTO_TEST_CASE(contains_duplicate3)
{
    BOOST_TEST_MESSAGE("Leetcode - Contains Duplicate III");

    bool result = gazlowe::ContainsDuplicate3::containsNearbyAlmostDuplicate(
        std::vector<int>({ { 1, 2, 3, 4, 5 } }), 2, 3);
    BOOST_REQUIRE_EQUAL(result, true);

    result = gazlowe::ContainsDuplicate3::containsNearbyAlmostDuplicate(
        std::vector<int>({ { 1, 5, 9, 13, 17 } }), 2, 3);
    BOOST_REQUIRE_EQUAL(result, false);
}
BOOST_AUTO_TEST_SUITE_END()

