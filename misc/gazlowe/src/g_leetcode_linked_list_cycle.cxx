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
            if (nullptr == head ||
                nullptr == head->next ||
                nullptr == head->next->next)
            {
                return false;
            }

            ListNode *fast = head->next;
            ListNode *slow = head->next;
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

    /* Linked List Cycle II
     * https://leetcode.com/problems/linked-list-cycle-ii/
     *
     * Given a linked list, return the node where the cycle begins.
     * If there is no cycle, return null. 
     *
     * Note: Do not modify the linked list.
     * Follow up: Can you solve it without using extra space? 
     */
    class LinkedListCycle2
    {
    public:
        static ListNode* detectCycle(ListNode *head)
        {
            if (nullptr == head || nullptr == head->next)
            {
                return nullptr;
            }
            ListNode *fast = head;
            ListNode *slow = head;
            for (;;)
            {
                slow = slow->next;
                if (fast->next == nullptr)
                {
                    return nullptr;
                }
                fast = fast->next->next;
                if (nullptr == slow || nullptr == fast)
                {
                    return nullptr;
                }
                if (fast == slow)
                {
                    break;
                }
            }
            if (nullptr == slow || nullptr == fast)
            {
                return nullptr;
            }
            slow = head;
            while (slow != fast)
            {
                slow = slow->next;
                fast = fast->next;
            }
            return slow;
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
    auto detect = gazlowe::LinkedListCycle2::detectCycle(list);
    BOOST_REQUIRE(detect == nullptr);

    result = gazlowe::LinkedListCycle::hasCycle(nullptr);
    BOOST_REQUIRE_EQUAL(result, false);
    detect = gazlowe::LinkedListCycle2::detectCycle(nullptr);
    BOOST_REQUIRE(detect == nullptr);

    result = gazlowe::LinkedListCycle::hasCycle(linkedList->AllocNode(1));
    BOOST_REQUIRE_EQUAL(result, false);
    detect = gazlowe::LinkedListCycle2::detectCycle(linkedList->AllocNode(1));
    BOOST_REQUIRE(detect == nullptr);


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
    detect = gazlowe::LinkedListCycle2::detectCycle(invalidList);
    BOOST_REQUIRE(detect != nullptr);
    BOOST_REQUIRE_EQUAL(detect->val, 2);
}
BOOST_AUTO_TEST_SUITE_END()
