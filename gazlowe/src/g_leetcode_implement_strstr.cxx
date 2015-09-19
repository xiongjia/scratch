/**
 * gazlowe - My algorithm tests
 */

#include <string>
#include "boost/test/unit_test.hpp"

namespace gazlowe
{
    class ImplementStrstr
    {
    public:
        static int strStr(std::string haystack, std::string needle)
        {
            std::size_t needleLen = needle.size();
            std::size_t haystackLen = haystack.size();
            if (0 == needleLen )
            {
                return 0;
            }
            if (needleLen > haystackLen)
            {
                return -1;
            }

            for (std::size_t itr = 0; itr <= (haystackLen - needleLen); ++itr)
            {
                std::size_t idx = 0;
                for (; idx < needleLen; ++idx)
                {
                    if (haystack[itr + idx] != needle[idx])
                    {
                        break;
                    }
                }
                if (needleLen == idx)
                {
                    return itr;
                }
            }
            return -1;
        }
    };
}

BOOST_AUTO_TEST_SUITE(leetcode)
BOOST_AUTO_TEST_CASE(implement_strstr)
{
    BOOST_TEST_MESSAGE("Leetcode - Implement strStr()");
    int ret = gazlowe::ImplementStrstr::strStr(std::string("abcdef"),
                                               std::string("123"));
    BOOST_REQUIRE_EQUAL(ret, -1);
    ret = gazlowe::ImplementStrstr::strStr(std::string("abcdef"),
                                           std::string("bcd"));
    BOOST_REQUIRE_EQUAL(ret, 1);
    ret = gazlowe::ImplementStrstr::strStr(std::string("aaa"),
                                           std::string("aaaa"));
    BOOST_REQUIRE_EQUAL(ret, -1);
    ret = gazlowe::ImplementStrstr::strStr(std::string("a"),
                                           std::string("a"));
    BOOST_REQUIRE_EQUAL(ret, -1);
}
BOOST_AUTO_TEST_SUITE_END()
