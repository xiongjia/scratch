/**
 * gazlowe - My algorithm tests
 */

#include <cstdlib>
#include <algorithm>
#include <functional>
#include <string>

#include "boost/test/unit_test.hpp"
#include "boost/array.hpp"

/* Quicksort: https://en.wikipedia.org/wiki/Quicksort */
namespace gazlowe 
{
    template<class T, std::size_t N, class C>
    void quick_sort_part(boost::array<T, N> &data, 
                         std::size_t begin, std::size_t end, C cmp)
    {
        std::size_t len = end - begin;
        if (len <= 1)
        {
            return;
        }

        std::swap(data[begin], data[begin + (std::rand() % len)]);
        std::size_t last = 0;
        for (std::size_t i = begin + 1; i < (begin + len); ++i)
        {
            if (cmp(data[i], data[begin]))
            {
                std::swap(data[i], data[(++last) + begin]);
            }
        }
        std::swap(data[begin], data[last + begin]);
        quick_sort_part(data, begin, begin + last, cmp);
        quick_sort_part(data, begin + last + 1, end, cmp);
    }

    template<class T, std::size_t N, class C>
    void quick_sort(boost::array<T, N> &data, C cmp)
    {
        quick_sort_part(data, 0, data.size(), cmp);
    }
}

BOOST_AUTO_TEST_SUITE(sorting)
BOOST_AUTO_TEST_CASE(sorting_qsort)
{
    /* Testing datum */
    const std::size_t dataSz = 7;
    boost::array<int, dataSz> dataNum = { {2, 5, 3, 4, 8, 10, 2} };
    boost::array<std::string, dataSz> dataStr = {
        {"2", "1", "5", "3", "4", "5", "3"}
    };

    /* sorting the numbers */
    BOOST_TEST_MESSAGE("Sorting numbers array");
    boost::array<int, dataSz> dataNums1(dataNum);
    gazlowe::quick_sort(dataNums1, std::greater<int>());
    boost::array<int, dataSz> dataNums2(dataNum);
    std::sort(dataNums2.begin(), dataNums2.end(), std::greater<int>());
    BOOST_CHECK_EQUAL_COLLECTIONS(dataNums1.begin(), dataNums1.end(),
                                  dataNums2.begin(), dataNums2.end());

    /* sorting the strings */
    BOOST_TEST_MESSAGE("Sorting string array");
    boost::array<std::string, dataSz> dataStr1(dataStr);
    gazlowe::quick_sort(dataStr1, std::less<std::string>());
    boost::array<std::string, dataSz> dataStr2(dataStr);
    std::sort(dataStr2.begin(), dataStr2.end(), std::less<std::string>());
    BOOST_CHECK_EQUAL_COLLECTIONS(dataStr1.begin(), dataStr1.end(),
                                  dataStr2.begin(), dataStr2.end());
}
BOOST_AUTO_TEST_SUITE_END()
