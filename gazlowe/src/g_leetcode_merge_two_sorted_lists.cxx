/**
 * gazlowe - My algorithm tests
 */

#include <string>
#include "boost/test/unit_test.hpp"
#include "g_misc.hxx"

namespace gazlowe
{
    /* Merge Two Sorted Lists
     * https://leetcode.com/problems/merge-two-sorted-lists/
     *
     * Merge two sorted linked lists and return it as a new list.
     * The new list should be made by splicing together the nodes
     * of the first two lists.
     */
    class Merge2SortedLists
    {
    public:
        static ListNode* mergeTwoLists(ListNode *l1, ListNode *l2)
        {
            ListNode *hdr = nullptr;
            ListNode *cur = nullptr;
            ListNode *p1 = l1;
            ListNode *p2 = l2;
            for (; nullptr != p1 && nullptr != p2;)
            {
                ListNode *tmp;
                if (p1->val <= p2->val)
                {
                    tmp = p1;
                    p1 = p1->next;
                }
                else
                {
                    tmp = p2;
                    p2 = p2->next;
                }

                tmp->next = nullptr;
                if (nullptr == cur)
                {
                    hdr = tmp;
                    cur = hdr;
                }
                else
                {
                    cur->next = tmp;
                    cur = tmp;
                }
            }
            if (nullptr != p1 || nullptr != p2)
            {
                ListNode *tmp = (nullptr != p1) ? p1 : p2;
                if (nullptr == cur)
                {
                    hdr = tmp;
                }
                else
                {
                    cur->next = tmp;
                }
            }
            return hdr;
        }
    };
}

BOOST_AUTO_TEST_SUITE(leetcode)
BOOST_AUTO_TEST_CASE(merge_two_sorted_lists)
{
    BOOST_TEST_MESSAGE("Leetcode - Merge Two Sorted Lists");
    auto linkedList = gazlowe::LinkedList::Create();
    auto l1 = gazlowe::LinkedList::Load(linkedList, 
        boost::array<int, 4>({ { 1, 3, 5, 7 } }));
    auto l2 = gazlowe::LinkedList::Load(linkedList,
        boost::array<int, 4>({ { 2, 4, 6, 7 } }));
    auto ret = gazlowe::Merge2SortedLists::mergeTwoLists(l1, l2);
    std::string dump;
    linkedList->Dump(ret, dump);
    BOOST_REQUIRE_EQUAL(dump, "1 2 3 4 5 6 7 7");
}
BOOST_AUTO_TEST_SUITE_END()
