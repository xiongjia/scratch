/**
 * gazlowe - My algorithm tests
 */

#include <algorithm>
#include <vector>

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
        typedef struct _TreeNode
        {
            int val;
            struct _TreeNode *left;
            struct _TreeNode *right;
        } TreeNode;

        static int maxDepth(TreeNode *root)
        {
            if (nullptr == root)
            {
                return 0;
            }
            return 1 + std::max(maxDepth(root->left), maxDepth(root->right));
        }

        static TreeNode* alloc_node(boost::pool<> &mpool, int val)
        {
            TreeNode *node = reinterpret_cast<TreeNode*>(mpool.malloc());
            node->val = val;
            node->left = nullptr;
            node->right = nullptr;
            return node;
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
    BOOST_CHECK(ret == 5);
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
     */
    boost::pool<> mpool(sizeof(gazlowe::MaxDepthBTree::TreeNode));    
    auto root = gazlowe::MaxDepthBTree::alloc_node(mpool, 0);
    auto node = gazlowe::MaxDepthBTree::alloc_node(mpool, 1);
    root->left = node;
    node = gazlowe::MaxDepthBTree::alloc_node(mpool, 2);
    root->right = node;
    node = gazlowe::MaxDepthBTree::alloc_node(mpool, 3);
    root->right->left = node;
    node = gazlowe::MaxDepthBTree::alloc_node(mpool, 4);
    root->right->right = node;
    node = gazlowe::MaxDepthBTree::alloc_node(mpool, 5);
    root->right->left->left = node;
    node = gazlowe::MaxDepthBTree::alloc_node(mpool, 6);
    root->right->left->right = node;
    node = gazlowe::MaxDepthBTree::alloc_node(mpool, 7);
    root->right->left->left->left = node;
    int depth = gazlowe::MaxDepthBTree::maxDepth(root);
    BOOST_CHECK(depth == 5);
}
BOOST_AUTO_TEST_SUITE_END()

