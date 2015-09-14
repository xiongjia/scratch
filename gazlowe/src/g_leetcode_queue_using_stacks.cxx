/**
 * gazlowe - My algorithm tests
 */

#include <stack>

#include "boost/shared_ptr.hpp"
#include "boost/make_shared.hpp"
#include "boost/test/unit_test.hpp"

namespace gazlowe
{
    /* Implement Queue using Stacks
     * https://leetcode.com/problems/implement-queue-using-stacks/
     *
     * Implement the following operations of a queue using stacks.
     *   - push(x) -- Push element x to the back of queue.
     *   - pop() -- Removes the element from in front of queue.
     *   - peek() -- Get the front element.
     *   - empty() -- Return whether the queue is empty.
     *
     * Note:
     *   - You must use only standard operations of a stack -- which
     *     means only push to top, peek/pop from top, size, and is empty
     *     operations are valid.
     *   - Depending on your language, stack may not be supported natively.
     *     You may simulate a stack by using a list or deque (double-ended
     *     queue), as long as you use only standard operations of a stack.
     *   - You may assume that all operations are valid (for example, no
     *     pop or peek operations will be called on an empty queue).
     */
    class QueueUsingStacks
    {
    public:
        static boost::shared_ptr<QueueUsingStacks> Create(void)
        {
            return boost::make_shared<QueueUsingStacks>();
        }

    private:
        std::stack<int> m_push;
        std::stack<int> m_pop;

        void PopStackPrepare(void)
        {
            if (!m_pop.empty())
            {
                return;
            }
            while (!m_push.empty())
            {
                m_pop.push(m_push.top());
                m_push.pop();
            }
        }

    public:
        /* Push element x to the back of queue. */
        void push(int x) 
        {
            m_push.push(x);
        }

        /* Removes the element from in front of queue. */
        void pop(void)
        {
            PopStackPrepare();
            m_pop.pop();
        }

        /* Get the front element. */
        int peek(void)
        {
            PopStackPrepare();
            return m_pop.top();
        }

        /* Return whether the queue is empty. */
        bool empty(void)
        {
            return m_push.empty() && m_pop.empty();
        }
    };
}

BOOST_AUTO_TEST_SUITE(leetcode)
BOOST_AUTO_TEST_CASE(queue_using_stacks)
{
    BOOST_TEST_MESSAGE("Leetcode - Implement Queue using Stacks");
    auto queue = gazlowe::QueueUsingStacks::Create();
    BOOST_REQUIRE_EQUAL(queue->empty(), true);

    queue->push(100);
    BOOST_REQUIRE_EQUAL(queue->peek(), 100);
    BOOST_REQUIRE_EQUAL(queue->empty(), false);

    queue->pop();
    BOOST_REQUIRE_EQUAL(queue->empty(), true);  
}
BOOST_AUTO_TEST_SUITE_END()
