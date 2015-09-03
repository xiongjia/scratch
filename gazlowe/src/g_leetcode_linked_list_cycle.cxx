/**
 * gazlowe - My algorithm tests
 */

#include "boost/test/unit_test.hpp"
#include "g_misc.hxx"

namespace gazlowe
{
    /* Linked List Cycle
     * https://leetcode.com/problems/linked-list-cycle/
     *
     * Given a linked list, determine if it has a cycle in it.
     * Follow up:
     * Can you solve it without using extra space?
     */
    class LinkedListCycle
    {
    public:
        static bool hasCycle(ListNode *head)
        {
            if (nullptr == head)
            {
                return false;
            }
            ListNode *fast = head;
            ListNode *slow = head;
            for (;;)
            {
                slow = slow->next;

                if (fast->next == nullptr)
                {
                    return false;
                }
                fast = fast->next->next;
                if (nullptr == slow || nullptr == fast)
                {
                    return false;
                }
                if (slow == fast)
                {
                    return true;
                }
            }
        }
    };
}

BOOST_AUTO_TEST_SUITE(leetcode)
BOOST_AUTO_TEST_CASE(linked_list_cycle)
{
    BOOST_TEST_MESSAGE("Leetcode - Linked List Cycle");

    auto linkedList = gazlowe::LinkedList::Create();
    auto list = gazlowe::LinkedList::Load(linkedList,
        boost::array<int, 4>({ { 1, 2, 3, 4 } }));
    bool result = gazlowe::LinkedListCycle::hasCycle(list);
    BOOST_REQUIRE_EQUAL(result, false);

    result = gazlowe::LinkedListCycle::hasCycle(nullptr);
    BOOST_REQUIRE_EQUAL(result, false);

    /*   [1] -> [2] -> [3] -> [4] -> [5]
     *           ^                    |
     *           |--------------------|
     */
    auto invalidList = linkedList->AllocNode(1);
    auto listNode = linkedList->AllocNode(2);
    invalidList->next = listNode;
    listNode = linkedList->AllocNode(3);
    invalidList->next->next = listNode;
    listNode = linkedList->AllocNode(4);
    invalidList->next->next->next = listNode;
    listNode = linkedList->AllocNode(5);
    invalidList->next->next->next->next = listNode;
    listNode->next = invalidList->next;
    result = gazlowe::LinkedListCycle::hasCycle(invalidList);
    BOOST_REQUIRE_EQUAL(result, true);
}
BOOST_AUTO_TEST_SUITE_END()
