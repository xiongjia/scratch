/**
 * gazlowe - My algorithm tests
 */

#include <algorithm>
#include <utility>
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

        static int maxDepthLoop(TreeNode *root)
        {
            if (nullptr == root)
            {
                return 0;
            }

            int depth = 0;
            std::stack<std::pair<TreeNode*, int>> data;
            data.push(std::make_pair(root, 1));
            while (!data.empty())
            {
                auto top = data.top();
                data.pop();
                if (nullptr == top.first)
                {
                    continue;
                }

                depth = std::max(depth, top.second);
                data.push(std::make_pair(top.first->left, top.second + 1));
                data.push(std::make_pair(top.first->right, top.second + 1));
            }

            return depth;
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
    depth = gazlowe::MaxDepthBTree::maxDepthLoop(root);
    BOOST_REQUIRE_EQUAL(depth, 5);

    depth = gazlowe::MaxDepthBTree::maxDepth(nullptr);
    BOOST_REQUIRE_EQUAL(depth, 0);
    depth = gazlowe::MaxDepthBTree::maxDepthLoop(nullptr);
    BOOST_REQUIRE_EQUAL(depth, 0);
}
BOOST_AUTO_TEST_SUITE_END()
