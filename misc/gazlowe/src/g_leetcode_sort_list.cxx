/**
 * gazlowe - My algorithm tests
 */

#include "boost/test/unit_test.hpp"
#include "g_misc.hxx"

namespace gazlowe
{
    /* Sort List
     * https://leetcode.com/problems/sort-list/
     *
     * Sort a linked list in O(n log n) time using constant space complexity.
     */
    class SortList
    {
    private:
        static ListNode* mergeLists(ListNode *list1, ListNode *list2)
        {
            if (nullptr == list1)
            {
                return list2;
            }
            if (nullptr == list2)
            {
                return list1;
            }

            ListNode *ret = nullptr;
            if (list1->val < list2->val)
            {
                ret = list1;
                list1 = list1->next;
            }
            else
            {
                ret = list2;
                list2 = list2->next;
            }
            ListNode *tmp = ret;
            while (nullptr != list1 && nullptr != list2)
            {
                if (list1->val < list2->val)
                {
                    tmp->next = list1;
                    list1 = list1->next;
                }
                else
                {
                    tmp->next = list2;
                    list2 = list2->next;
                }
                tmp = tmp->next;
            }
            if (nullptr != list1)
            {
                tmp->next = list1;
            }
            else if (nullptr != list2)
            {
                tmp->next = list2;
            }
            return ret;
        }

    public:
        static ListNode *sortList(ListNode *head)
        {
            if (nullptr == head || nullptr == head->next)
            {
                return head;
            }

            ListNode *slow = head;
            ListNode *fast = head;
            while (fast->next && fast->next->next)
            {
                slow = slow->next;
                fast = fast->next->next;
            }
            ListNode *list1 = head;
            ListNode *list2 = slow->next;
            slow->next = nullptr;
            return mergeLists(sortList(list1), sortList(list2));
        }
    };
}

BOOST_AUTO_TEST_SUITE(leetcode)
BOOST_AUTO_TEST_CASE(sort_list)
{
    BOOST_TEST_MESSAGE("Leetcode - Sort List");

    auto linkedList = gazlowe::LinkedList::Create();
    auto list = gazlowe::LinkedList::Load(linkedList,
        boost::array<int, 6>({ { 5, 4, 2, 3, 10, 1 } }));
    auto ret = gazlowe::SortList::sortList(list);
    std::string dump;
    linkedList->Dump(ret, dump);
    BOOST_REQUIRE_EQUAL(dump, "1 2 3 4 5 10");
}
BOOST_AUTO_TEST_SUITE_END()
