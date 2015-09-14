/**
 * gazlowe - My algorithm tests
 */

#include <queue>

#include "boost/shared_ptr.hpp"
#include "boost/make_shared.hpp"
#include "boost/utility.hpp"
#include "boost/test/unit_test.hpp"

namespace gazlowe
{
    /* Implement Stack using Queues
     * https://leetcode.com/problems/implement-stack-using-queues/
     *
     * Implement the following operations of a stack using queues.
     *   - push(x) -- Push element x onto stack.
     *   - pop() -- Removes the element on top of the stack.
     *   - top() -- Get the top element.
     *   - empty() -- Return whether the stack is empty.
     *
     * Notes:
     *   - You must use only standard operations of a queue -- which
     *     means only push to back, peek/pop from front, size, and is
     *     empty operations are valid.
     *   -  Depending on your language, queue may not be supported natively.
     *      You may simulate a queue by using a list or deque (double-ended
     *      queue), as long as you use only standard operations of a queue.
     *   - You may assume that all operations are valid (for example, no pop
     *     or top operations will be called on an empty stack).
     */

    /* 2 std::queue<int> */
    class StackUsingQueues : boost::noncopyable
    {
    public:
        static boost::shared_ptr<StackUsingQueues> Create(void)
        {
            return boost::make_shared<StackUsingQueues>();
        }

    private:
        std::queue<int> m_pushQ;
        std::queue<int> m_popQ;

    public:
        /* Push element x onto stack. */
        void push(int x)
        {
            m_pushQ.push(x);
            while (!empty())
            {
                m_pushQ.push(top());
                pop();
            }
            std::swap(m_popQ, m_pushQ);
        }

        /* Removes the element on top of the stack. */
        void pop(void)
        {
            m_popQ.pop();
        }

        /* Get the top element. */
        int top(void) 
        {
            return m_popQ.front();
        }

        /* Return whether the stack is empty. */
        bool empty(void)
        {
            return m_popQ.empty();
        }
    };

    /* 1 std::deque<int> */
    class StackUsingDQueue : boost::noncopyable
    {
    public:
        static boost::shared_ptr<StackUsingDQueue> Create(void)
        {
            return boost::make_shared<StackUsingDQueue>();
        }

    private:
        std::deque<int> m_que;

    public:
        /* Push element x onto stack. */
        void push(int x)
        {
            m_que.push_back(x);
        }

        /* Removes the element on top of the stack. */
        void pop(void)
        {
            m_que.pop_back();
        }

        /* Get the top element. */
        int top(void)
        {
            return m_que.back();
        }

        /* Return whether the stack is empty. */
        bool empty(void)
        {
            return m_que.empty();
        }
    };
}

BOOST_AUTO_TEST_SUITE(leetcode)
BOOST_AUTO_TEST_CASE(stack_using_queues)
{
    BOOST_TEST_MESSAGE("Leetcode - Implement Stack using Queues");
    auto stackq = gazlowe::StackUsingQueues::Create();
    BOOST_REQUIRE_EQUAL(stackq->empty(), true);
    stackq->push(100);
    BOOST_REQUIRE_EQUAL(stackq->top(), 100);
    BOOST_REQUIRE_EQUAL(stackq->empty(), false);
    stackq->push(101);
    BOOST_REQUIRE_EQUAL(stackq->top(), 101);
    stackq->pop();
    BOOST_REQUIRE_EQUAL(stackq->top(), 100);
    stackq->pop();
    BOOST_REQUIRE_EQUAL(stackq->empty(), true);

    auto stackdq = gazlowe::StackUsingDQueue::Create();
    BOOST_REQUIRE_EQUAL(stackdq->empty(), true);
    stackdq->push(100);
    BOOST_REQUIRE_EQUAL(stackdq->top(), 100);
    BOOST_REQUIRE_EQUAL(stackdq->empty(), false);
    stackdq->push(101);
    BOOST_REQUIRE_EQUAL(stackdq->top(), 101);
    stackdq->pop();
    BOOST_REQUIRE_EQUAL(stackdq->top(), 100);
    stackdq->pop();
    BOOST_REQUIRE_EQUAL(stackdq->empty(), true);
}
BOOST_AUTO_TEST_SUITE_END()
