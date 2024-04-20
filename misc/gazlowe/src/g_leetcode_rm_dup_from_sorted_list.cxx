/**
 * gazlowe - My algorithm tests
 */

#include <string>
#include "boost/test/unit_test.hpp"
#include "g_misc.hxx"

namespace gazlowe
{
    /* Remove Duplicates from Sorted List
     * https://leetcode.com/problems/remove-duplicates-from-sorted-list/
     *
     * Given a sorted linked list, delete all duplicates such
     * that each element appear only once. 
     *
     * For example,
     *   Given 1->1->2, return 1->2.
     *   Given 1->1->2->3->3, return 1->2->3. 
     */
    class RMDupFromSortedList
    {
    public:
        static ListNode* deleteDuplicates(ListNode *head)
        {
            ListNode *tmp = nullptr;
            for (ListNode *itr = head; nullptr != itr;)
            {
                if (nullptr == itr->next ||
                    itr->val != itr->next->val)
                {
                    tmp = itr;
                    itr = itr->next;
                    continue;
                }
                
                if (nullptr != tmp)
                {
                    tmp->next = itr->next;
                    itr = itr->next;
                }
                else
                {
                    head = itr->next;
                    itr = head;
                }
            }
            return head;
        }
    };
}

BOOST_AUTO_TEST_SUITE(leetcode)
BOOST_AUTO_TEST_CASE(rm_dup_from_sorted_list)
{
    BOOST_TEST_MESSAGE("Leetcode - Remove Duplicates from Sorted List");

    auto linkedList = gazlowe::LinkedList::Create();
    auto list = gazlowe::LinkedList::Load(linkedList,
        boost::array<int, 8>({ { 1, 1, 2, 3, 3, 4, 4, 5 } }));
    auto ret = gazlowe::RMDupFromSortedList::deleteDuplicates(list);
    std::string dump;
    linkedList->Dump(ret, dump);
    BOOST_REQUIRE_EQUAL(dump, "1 2 3 4 5");
}
BOOST_AUTO_TEST_SUITE_END()
