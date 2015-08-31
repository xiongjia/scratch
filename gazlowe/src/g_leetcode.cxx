/**
 * gazlowe - My algorithm tests
 */

#include <algorithm>
#include <vector>

#include "g_misc.hxx"
#include "boost/test/unit_test.hpp"
#include "boost/pool/pool.hpp"

namespace gazlowe
{
    /* Single Number:
     * https://leetcode.com/problems/single-number/
     *
     * Given an array of integers, every element appears twice except for one.
     * Find that single one.
     */
    int single_number(std::vector<int> &nums)
    {
        int ret = 0;
        std::for_each(nums.begin(), nums.end(),
            [&](int &num) { ret ^= num; });
        return ret;
    }

    /* Maximum Depth of Binary Tree:
     * https://leetcode.com/problems/maximum-depth-of-binary-tree/
     *
     * Given a binary tree, find its maximum depth.
     * The maximum depth is the number of nodes along the longest path from
     * the root node down to the farthest leaf node
     *
     * Definition for a binary tree node.
     * struct TreeNode {
     *     int val;
     *     TreeNode *left;
     *     TreeNode *right;
     *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
     * };
     */
    class MaxDepthBTree
    {
    public:
        static int maxDepth(TreeNode *root)
        {
            if (nullptr == root)
            {
                return 0;
            }
            return 1 + std::max(maxDepth(root->left), maxDepth(root->right));
        }
    };

    /* Invert a binary tree:
     * https://leetcode.com/problems/invert-binary-tree/
     *
     *      4
     *    /   \
     *   2     7
     *  / \   / \
     * 1   3 6   9
     *
     * to
     *
     *       4
     *     /   \
     *    7     2
     *   / \   / \
     *  9   6 3   1
     */
    class InvertBTree
    {
    public:
        static TreeNode *invertTree(TreeNode *root)
        {
            if (nullptr == root)
            {
                return root;
            }
            TreeNode *tmp = root->left;
            root->left = invertTree(root->right);
            root->right = invertTree(tmp);
            return root;
        }
    };

    /* Delete Node in a Linked List
     * https://leetcode.com/problems/delete-node-in-a-linked-list/ 
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

    /* Add Digits  
     * Given a non-negative integer num, repeatedly add all its 
     * digits until the result has only one digit.
     * For example:
     * Given num = 38, the process is like: 3 + 8 = 11, 1 + 1 = 2.
     * Since 2 has only one digit, return it.
     *
     * Follow up:
     * Could you do it without any loop/recursion in O(1) runtime? 
     */
    class AddDigits  
    {
    public:
        static int addDigitsEnhance(int num)
        {
            return (num - 1) % 9 + 1;
        }

        static int addDigits(int num)
        {
            for (; num > 9;)
            {
                int tmp = 0;
                for (; num != 0;)
                {
                    tmp += num % 10;
                    num /= 10;
                }
                num = tmp;
            }
            return num;
        }
    };
}

BOOST_AUTO_TEST_SUITE(leetcode)
BOOST_AUTO_TEST_CASE(single_number)
{
    /* Testing gazlowe::single_number() */
    BOOST_TEST_MESSAGE("Leetcode - Single Number");
    std::vector<int> testData = { 1, 1, 2, 2, 3, 3, 5 };
    int ret = gazlowe::single_number(testData);
    BOOST_REQUIRE_EQUAL(ret, 5);
}

BOOST_AUTO_TEST_CASE(max_depth_btree)
{
    /* Testing data:
     *             0
     *            / \
     *           1   2
     *              / \
     *             3   4
     *            / \
     *           5   6
     *          /
     *         7
     * [0 1 # # 2 3 5 7 # # # 6 # # 4 # #]
     */
    BOOST_TEST_MESSAGE("Leetcode - Maximum Depth of Binary Tree");
    auto tree = gazlowe::TreeNodes::Create();
    auto root = tree->Load("0 1 # # 2 3 5 7 # # # 6 # # 4 # #");
    int depth = gazlowe::MaxDepthBTree::maxDepth(root);
    BOOST_REQUIRE_EQUAL(depth, 5);
}

BOOST_AUTO_TEST_CASE(invert_btree)
{
    /* Testing data:
     *      4
     *    /   \
     *   2     7
     *  / \   / \
     * 1   3 6   9
     * [4 2 1 # # 3 # # 7 6 # # 9 # #]
     *
     * to
     *
     *       4
     *     /   \
     *    7     2
     *   / \   / \
     *  9   6 3   1
     * [4 7 9 # # 6 # # 2 3 # # 1 # #]
     */
    BOOST_TEST_MESSAGE("Leetcode - Maximum Depth of Binary Tree");
    auto tree = gazlowe::TreeNodes::Create();
    auto rootSrc = tree->Load("4 2 1 # # 3 # # 7 6 # # 9 # #");
    auto rootRet = gazlowe::InvertBTree::invertTree(rootSrc);
    std::string results;
    tree->Dump(rootRet, results);
    BOOST_REQUIRE_EQUAL(results, "4 7 9 # # 6 # # 2 3 # # 1 # #");
}

BOOST_AUTO_TEST_CASE(del_node_in_a_linked_list)
{
    auto linkedList = gazlowe::LinkedList::Create();
    auto list = gazlowe::LinkedList::Load(linkedList,
        boost::array<int, 4>({{1, 2, 3, 4 }}));
    gazlowe::DelNodeLinkList::deleteNode(linkedList->Search(list, 3));
    std::string result;
    linkedList->Dump(list, result);
    BOOST_REQUIRE_EQUAL(result, "1 2 4");
}

BOOST_AUTO_TEST_CASE(add_digits)
{
    int result = gazlowe::AddDigits::addDigits(38);
    int resultEnhance = gazlowe::AddDigits::addDigitsEnhance(38);
    BOOST_REQUIRE_EQUAL(result, 2);
    BOOST_REQUIRE_EQUAL(result, resultEnhance);

    result = gazlowe::AddDigits::addDigits(0);
    resultEnhance = gazlowe::AddDigits::addDigitsEnhance(0);
    BOOST_REQUIRE_EQUAL(result, 0);
    BOOST_REQUIRE_EQUAL(result, resultEnhance);

    result = gazlowe::AddDigits::addDigits(138);
    resultEnhance = gazlowe::AddDigits::addDigitsEnhance(138);
    BOOST_REQUIRE_EQUAL(result, 3);
    BOOST_REQUIRE_EQUAL(result, resultEnhance);
}
BOOST_AUTO_TEST_SUITE_END()
