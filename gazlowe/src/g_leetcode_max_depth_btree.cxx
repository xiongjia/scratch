/**
 * gazlowe - My algorithm tests
 */

#include <algorithm>
#include <stack>
#include "boost/test/unit_test.hpp"
#include "g_misc.hxx"

namespace gazlowe
{
    /* Maximum Depth of Binary Tree:
     * https://leetcode.com/problems/maximum-depth-of-binary-tree/
     *
     * Given a binary tree, find its maximum depth.
     * The maximum depth is the number of nodes along the longest path from
     * the root node down to the farthest leaf node
     */
    class MaxDepthBTree
    {
    public:
        static int maxDepth(TreeNode *root)
        {
            return (nullptr == root) ?
                0 : 1 + std::max(maxDepth(root->left), maxDepth(root->right));
        }
    };
}

BOOST_AUTO_TEST_SUITE(leetcode)
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

    depth = gazlowe::MaxDepthBTree::maxDepth(nullptr);
    BOOST_REQUIRE_EQUAL(depth, 0);
}
BOOST_AUTO_TEST_SUITE_END()
