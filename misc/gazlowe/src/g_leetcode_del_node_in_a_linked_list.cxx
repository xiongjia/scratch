/**
 * gazlowe - My algorithm tests
 */

#include "boost/test/unit_test.hpp"
#include "g_misc.hxx"

namespace gazlowe
{
    /* Delete Node in a Linked List
     * https://leetcode.com/problems/delete-node-in-a-linked-list/
     *
     *  Write a function to delete a node (except the tail) in a singly
     *  linked list, given only access to that node.
     *
     *  Supposed the linked list is 1 -> 2 -> 3 -> 4 and you are given
     *  the third node with value 3, the linked list should become 1 -> 2 -> 4
     *  after calling your function.
     */
    class DelNodeLinkList
    {
    public:
        static void deleteNode(ListNode *node)
        {
            if (nullptr == node || nullptr == node->next)
            {
                return;
            }
            node->val = node->next->val;
            node->next = node->next->next;
        }
    };
}

BOOST_AUTO_TEST_SUITE(leetcode)
BOOST_AUTO_TEST_CASE(del_node_in_a_linked_list)
{
    BOOST_TEST_MESSAGE("Leetcode - Delete Node in a linked list");
    auto linkedList = gazlowe::LinkedList::Create();
    auto list = gazlowe::LinkedList::Load(linkedList,
        boost::array<int, 4>({ { 1, 2, 3, 4 } }));
    gazlowe::DelNodeLinkList::deleteNode(linkedList->Search(list, 3));
    std::string result;
    linkedList->Dump(list, result);
    BOOST_REQUIRE_EQUAL(result, "1 2 4");
}
BOOST_AUTO_TEST_SUITE_END()
