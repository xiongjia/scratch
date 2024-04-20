/**
 * gazlowe - My algorithm tests
 */

#include <stack>
#include <utility>
#include "boost/test/unit_test.hpp"
#include "g_misc.hxx"

namespace gazlowe
{
    /* Same Tree
     * https://leetcode.com/problems/same-tree/
     *
     * Given two binary trees, write a function to check
     * if they are equal or not. Two binary trees are considered equal
     * if they are structurally identical and the nodes have the same value.
     */
    class SameTree
    {
    public:
        static bool isSameTree(TreeNode *p, TreeNode *q)
        {
            if (nullptr == p && nullptr == q)
            {
                return true;
            }

            if (nullptr != p && nullptr != q)
            {
                return p->val == q->val &&
                    isSameTree(p->left, q->left) &&
                    isSameTree(p->right, q->right);
            }
            else
            {
                return false;
            }
        }

        static bool isSameTreeItr(TreeNode *p, TreeNode *q)
        {
            if (nullptr == p && nullptr == q)
            {
                return true;
            }

            std::stack<std::pair<TreeNode*, TreeNode*>> data;
            data.push(std::make_pair(p, q));
            while (!data.empty())
            {
                auto top = data.top();
                data.pop();
                if (nullptr == top.first && nullptr == top.second)
                {
                    continue;
                }

                if (!(nullptr != top.first && nullptr != top.second))
                {
                    return false;
                }

                if (top.first->val != top.second->val)
                {
                    return false;
                }

                data.push(std::make_pair(top.first->left, top.second->left));
                data.push(std::make_pair(top.first->right, top.second->right));
            }
            return true;
        }
    };
}

BOOST_AUTO_TEST_SUITE(leetcode)
BOOST_AUTO_TEST_CASE(same_tree)
{
    BOOST_TEST_MESSAGE("Leetcode - Same Tree");
    const char *tree1 = "4 2 1 # # 3 # # 7 6 # # 9 # #";
    const char *tree2 = "4 2 1 # # 3 # # 7 6 # # 7 # #";
    auto tree = gazlowe::TreeNodes::Create();

    bool result = gazlowe::SameTree::isSameTree(tree->Load(tree1),
        tree->Load(tree2));
    BOOST_REQUIRE_EQUAL(result, false);
    result = gazlowe::SameTree::isSameTreeItr(tree->Load(tree1),
        tree->Load(tree2));
    BOOST_REQUIRE_EQUAL(result, false);

    result = gazlowe::SameTree::isSameTree(tree->Load(tree1),
        tree->Load(tree1));
    BOOST_REQUIRE_EQUAL(result, true);
    result = gazlowe::SameTree::isSameTreeItr(tree->Load(tree1),
        tree->Load(tree1));
    BOOST_REQUIRE_EQUAL(result, true);

    result = gazlowe::SameTree::isSameTree(nullptr, nullptr);
    BOOST_REQUIRE_EQUAL(result, true);
    result = gazlowe::SameTree::isSameTreeItr(nullptr, nullptr);
    BOOST_REQUIRE_EQUAL(result, true);

    result = gazlowe::SameTree::isSameTree(tree->Load(tree1), nullptr);
    BOOST_REQUIRE_EQUAL(result, false);
    result = gazlowe::SameTree::isSameTreeItr(tree->Load(tree1), nullptr);
    BOOST_REQUIRE_EQUAL(result, false);
}
BOOST_AUTO_TEST_SUITE_END()
