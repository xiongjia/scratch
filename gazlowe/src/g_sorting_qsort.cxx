/**
 * gazlowe - My algorithm tests
 */

#include <functional>
#include <algorithm>
#include <string>
#include <stack>
#include <utility>
#include "boost/test/unit_test.hpp"
#include "boost/array.hpp"

namespace gazlowe 
{
    /* Quicksort: https://en.wikipedia.org/wiki/Quicksort */

    class QSort
    {
    private:
        template<class T, std::size_t N, class C>
        static void SortPart(boost::array<T, N> &data,
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
            SortPart(data, begin, begin + last, cmp);
            SortPart(data, begin + last + 1, end, cmp);
        }

    public:
        template<class T, std::size_t N, class C>
        static void Sort(boost::array<T, N> &data, C cmp)
        {
            SortPart(data, 0, data.size(), cmp);
        }
    };

    class QSortLoop
    {
    public:
        template<class T, std::size_t N, class C>
        static void Sort(boost::array<T, N> &data, C cmp)
        {
            std::stack<std::pair<std::size_t, std::size_t>> ranges;
            ranges.push(std::make_pair(0, data.size()));
            while (!ranges.empty())
            {
                auto top = ranges.top();
                ranges.pop();

                std::size_t begin = top.first;
                std::size_t end = top.second;
                std::size_t len = end - begin;
                if (len <= 1)
                {
                    continue;
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

                ranges.push(std::make_pair(begin, begin + last));
                ranges.push(std::make_pair(begin + last + 1, end));
            }
        }
    };
}

BOOST_AUTO_TEST_SUITE(sorting)
BOOST_AUTO_TEST_CASE(sorting_qsort)
{
    BOOST_TEST_MESSAGE("Sorting - qsort");

    /* Testing datum */
    const std::size_t dataSz = 7;
    boost::array<int, dataSz> dataNum = { {2, 5, 3, 4, 8, 10, 2} };
    boost::array<std::string, dataSz> dataStr = {
        {"2", "1", "5", "3", "4", "5", "3"}
    };

    /* sorting the numbers */
    boost::array<int, dataSz> dataNums1(dataNum);
    gazlowe::QSort::Sort(dataNums1, std::greater<int>());
    boost::array<int, dataSz> dataNums2(dataNum);
    gazlowe::QSortLoop::Sort(dataNums2, std::greater<int>());
    boost::array<int, dataSz> dataNums3(dataNum);
    std::sort(dataNums3.begin(), dataNums3.end(), std::greater<int>());
    BOOST_CHECK_EQUAL_COLLECTIONS(dataNums1.begin(), dataNums1.end(),
                                  dataNums3.begin(), dataNums3.end());
    BOOST_CHECK_EQUAL_COLLECTIONS(dataNums2.begin(), dataNums2.end(),
                                  dataNums3.begin(), dataNums3.end());

    /* sorting the strings */
    boost::array<std::string, dataSz> dataStr1(dataStr);
    gazlowe::QSort::Sort(dataStr1, std::less<std::string>());
    boost::array<std::string, dataSz> dataStr2(dataStr);
    gazlowe::QSortLoop::Sort(dataStr2, std::less<std::string>());
    boost::array<std::string, dataSz> dataStr3(dataStr);
    std::sort(dataStr3.begin(), dataStr3.end(), std::less<std::string>());
    BOOST_CHECK_EQUAL_COLLECTIONS(dataStr1.begin(), dataStr1.end(),
                                  dataStr3.begin(), dataStr3.end());
    BOOST_CHECK_EQUAL_COLLECTIONS(dataStr2.begin(), dataStr2.end(),
                                  dataStr3.begin(), dataStr3.end());
}
BOOST_AUTO_TEST_SUITE_END()
