/**
 * gazlowe - My algorithm tests
 */

#include "boost/test/unit_test.hpp"

namespace gazlowe
{
    /* Add Digits
     * https://leetcode.com/problems/add-digits/
     *
     * Given a non-negative integer num, repeatedly add all its
     * digits until the result has only one digit.
     * For example:
     * Given num = 38, the process is like: 3 + 8 = 11, 1 + 1 = 2.
     * Since 2 has only one digit, return it.
     *
     * Follow up:
     * Could you do it without any loop/recursion in O(1) runtime?
     */
    class AddDigits
    {
    public:
        static int addDigitsEnhance(int num)
        {
            return (num - 1) % 9 + 1;
        }

        static int addDigits(int num)
        {
            for (; num > 9;)
            {
                int tmp = 0;
                for (; num != 0;)
                {
                    tmp += num % 10;
                    num /= 10;
                }
                num = tmp;
            }
            return num;
        }
    };
}

BOOST_AUTO_TEST_SUITE(leetcode)
BOOST_AUTO_TEST_CASE(add_digits)
{
    BOOST_TEST_MESSAGE("Leetcode - Add Digits");
    int result = gazlowe::AddDigits::addDigits(38);
    int resultEnhance = gazlowe::AddDigits::addDigitsEnhance(38);
    BOOST_REQUIRE_EQUAL(result, 2);
    BOOST_REQUIRE_EQUAL(result, resultEnhance);

    result = gazlowe::AddDigits::addDigits(0);
    resultEnhance = gazlowe::AddDigits::addDigitsEnhance(0);
    BOOST_REQUIRE_EQUAL(result, 0);
    BOOST_REQUIRE_EQUAL(result, resultEnhance);

    result = gazlowe::AddDigits::addDigits(138);
    resultEnhance = gazlowe::AddDigits::addDigitsEnhance(138);
    BOOST_REQUIRE_EQUAL(result, 3);
    BOOST_REQUIRE_EQUAL(result, resultEnhance);
}
BOOST_AUTO_TEST_SUITE_END()
