/**
 * gazlowe - My algorithm tests
 */

#include "boost/test/unit_test.hpp"
#include "g_misc.hxx"

namespace gazlowe
{
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
}

BOOST_AUTO_TEST_SUITE(leetcode)
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
BOOST_AUTO_TEST_SUITE_END()
