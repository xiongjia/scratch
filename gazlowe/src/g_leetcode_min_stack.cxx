/**
 * gazlowe - My algorithm tests
 */

#include <stack>
#include <utility>
#include "boost/shared_ptr.hpp"
#include "boost/make_shared.hpp"
#include "boost/test/unit_test.hpp"

namespace gazlowe
{
    /* Min Stack
     * https://leetcode.com/problems/min-stack/
     * 
     * Design a stack that supports push, pop, top, 
     * and retrieving the minimum element in constant time. 
     *   - push(x) -- Push element x onto stack. 
     *   - pop() -- Removes the element on top of the stack. 
     *   - top() -- Get the top element. 
     *   - getMin() -- Retrieve the minimum element in the stack.
     */
    class MinStack
    {
    private:
        std::stack<std::pair<int, int>> m_data;

    public:
        static boost::shared_ptr<MinStack> Craete(void)
        {
            return boost::make_shared<MinStack>();
        }

    public:
        void push(int x)
        {
            int min = m_data.empty() ? x : std::min(m_data.top().second, x);
            m_data.push(std::make_pair(x, min));
        }

        void pop(void)
        {
            m_data.pop();
        }

        int top(void)
        {
            return m_data.top().first;
        }

        int getMin(void)
        {
            return m_data.top().second;
        }
    };
}

BOOST_AUTO_TEST_SUITE(leetcode)
BOOST_AUTO_TEST_CASE(min_stack)
{
    BOOST_TEST_MESSAGE("Leetcode - Min Stack");
    auto minStack = gazlowe::MinStack::Craete();
    minStack->push(3);
    BOOST_REQUIRE_EQUAL(minStack->top(), 3);
    BOOST_REQUIRE_EQUAL(minStack->getMin(), 3);
    minStack->push(1);
    BOOST_REQUIRE_EQUAL(minStack->top(), 1);
    BOOST_REQUIRE_EQUAL(minStack->getMin(), 1);
    minStack->push(2);
    BOOST_REQUIRE_EQUAL(minStack->top(), 2);
    BOOST_REQUIRE_EQUAL(minStack->getMin(), 1);

    minStack->pop();
    BOOST_REQUIRE_EQUAL(minStack->top(), 1);
    BOOST_REQUIRE_EQUAL(minStack->getMin(), 1);
    minStack->pop();
    BOOST_REQUIRE_EQUAL(minStack->top(), 3);
    BOOST_REQUIRE_EQUAL(minStack->getMin(), 3);
}
BOOST_AUTO_TEST_SUITE_END()
