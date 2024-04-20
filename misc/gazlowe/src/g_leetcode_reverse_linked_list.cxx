/**
 * gazlowe - My algorithm tests
 */

#include <string>
#include "boost/test/unit_test.hpp"
#include "g_misc.hxx"

namespace gazlowe
{
    /* Reverse Linked List
     * https://leetcode.com/problems/reverse-linked-list/
     *
     * Reverse a singly linked list.
     * Hint: 
     *      A linked list can be reversed either iteratively 
     *      or recursively. Could you implement both?
     */
    class ReverseLinkedList
    {
    public:
        static ListNode* reverseList(ListNode *head)
        {
            ListNode *root = nullptr;
            for (ListNode *itr = head; itr != nullptr;)
            {
                if (nullptr == root)
                {
                    root = itr;
                    itr = itr->next;
                    root->next = nullptr;
                }
                else
                {
                    ListNode *tmp = itr->next;
                    itr->next = root;
                    root = itr;
                    itr = tmp;
                }
            }
            return root;
        }

        static ListNode* reverseListRecursively(ListNode *head)
        {
            if (nullptr == head || nullptr == head->next)
            {
                return head;
            }

            ListNode *tmp = head->next;
            head->next = nullptr;
            ListNode *ret = reverseListRecursively(tmp);
            tmp->next = head;
            return ret;
        }
    };
}

BOOST_AUTO_TEST_SUITE(leetcode)
BOOST_AUTO_TEST_CASE(reverse_linked_list)
{
    BOOST_TEST_MESSAGE("Leetcode - Reverse Linked List");

    auto linkedList = gazlowe::LinkedList::Create();
    auto list = gazlowe::LinkedList::Load(linkedList,
        boost::array<int, 4>({ { 1, 2, 3, 4 } }));
    auto ret = gazlowe::ReverseLinkedList::reverseList(list);
    std::string dump;
    linkedList->Dump(ret, dump);
    BOOST_REQUIRE_EQUAL(dump, "4 3 2 1");

    ret = gazlowe::ReverseLinkedList::reverseList(nullptr);
    BOOST_REQUIRE(ret == nullptr);

    list = gazlowe::LinkedList::Load(linkedList,
        boost::array<int, 4>({ { 1, 2, 3, 4 } }));
    ret = gazlowe::ReverseLinkedList::reverseListRecursively(list);
    linkedList->Dump(ret, dump);
    BOOST_REQUIRE_EQUAL(dump, "4 3 2 1");
}
BOOST_AUTO_TEST_SUITE_END()
