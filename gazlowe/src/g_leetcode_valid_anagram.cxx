/**
 * gazlowe - My algorithm tests
 */

#include <algorithm>
#include <string>
#include "boost/test/unit_test.hpp"

namespace gazlowe
{
    /* Valid Anagram
     * https://leetcode.com/problems/valid-anagram/
     *
     * Given two strings s and t, write a function to 
     * determine if t is an anagram of s.
     *
     * For example,
     *      s = "anagram", t = "nagaram", return true.
     *      s = "rat", t = "car", return false.
     *
     * Note:
     *      You may assume the string contains only 
     *      lowercase alphabets.
     */
    class ValidAnagram
    {
    public:
        static bool isAnagram(std::string s, std::string t)
        {
            if (s.size() != t.size())
            {
                return false;
            }
            int map[26] = { 0 };
            for (std::size_t i = 0; i < s.size(); ++i)
            {
                ++map[s[i] - 'a'];
            }
            for (std::size_t i = 0; i < t.size(); ++i)
            {
                if (--map[t[i] - 'a'] < 0)
                {
                    return false;
                }
            }
            return true;
        }
    };
}

BOOST_AUTO_TEST_SUITE(leetcode)
BOOST_AUTO_TEST_CASE(valid_anagram)
{
    BOOST_TEST_MESSAGE("Leetcode - Valid Anagram");
    BOOST_REQUIRE(gazlowe::ValidAnagram::
        isAnagram(std::string("abc"), std::string("abc")) == true);
    BOOST_REQUIRE(gazlowe::ValidAnagram::
        isAnagram(std::string("abc"), std::string("cba")) == true);
    BOOST_REQUIRE(gazlowe::ValidAnagram::
        isAnagram(std::string("abc"), std::string("bac")) == true);
    BOOST_REQUIRE(gazlowe::ValidAnagram::
        isAnagram(std::string("abc"), std::string("abd")) == false);
}
BOOST_AUTO_TEST_SUITE_END()
