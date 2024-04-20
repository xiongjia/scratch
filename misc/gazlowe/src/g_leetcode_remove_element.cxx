/**
 * gazlowe - My algorithm tests
 */

#include <vector>
#include "boost/test/unit_test.hpp"

namespace gazlowe
{
    /* Remove Element
     * https://leetcode.com/problems/remove-element/
     *
     * Given an array and a value, remove all instances of 
     * that value in place and return the new length. 
     *
     * The order of elements can be changed. It doesn't 
     * matter what you leave beyond the new length.
     */
    class RemoveElement
    {
    public:
        static int removeElement(std::vector<int> &nums, int val)
        {
            std::size_t pivot = nums.size();
            for (std::size_t i = 0; i < pivot; ++i)
            {
                if (nums[i] == val)
                {
                    nums[i] = nums[pivot - 1];
                    i--;
                    pivot--;
                }
            }
            return pivot;
        }
    };
}

BOOST_AUTO_TEST_SUITE(leetcode)
BOOST_AUTO_TEST_CASE(remove_element)
{
    BOOST_TEST_MESSAGE("Leetcode - Remove Element");

    auto data = std::vector<int>({ 1, 2, 3, 4, 5 });
    int len = gazlowe::RemoveElement::removeElement(data, 1);
    BOOST_REQUIRE_EQUAL(len, 4);
    BOOST_REQUIRE(std::find(data.begin(), data.begin() + len - 1, 1) == 
                  (data.begin() + len - 1));

    data = std::vector<int>();
    len = gazlowe::RemoveElement::removeElement(data, 1);
    BOOST_REQUIRE_EQUAL(len, 0);
}
BOOST_AUTO_TEST_SUITE_END()
