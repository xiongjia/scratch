/**
 * gazlowe - My algorithm tests
 */

#include <set>
#include <vector>
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
            for (std::vector<int>::iterator itr = nums.begin();
                itr != nums.end(); ++itr)
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
BOOST_AUTO_TEST_SUITE_END()
