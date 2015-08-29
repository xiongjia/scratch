/**
 * gazlowe - My algorithm tests
 */

#include <sstream>
#include <string>
#include <algorithm>

#include "boost/algorithm/string.hpp"
#include "boost/make_shared.hpp"
#include "boost/pool/pool.hpp"
#include "boost/test/unit_test.hpp"
#include "g_misc.hxx"

namespace gazlowe
{
    class TreeNodesImpl : public TreeNodes
    {
    private:
        boost::pool<> m_mpool;

    public:
        TreeNodesImpl(void)
            : TreeNodes()
            , m_mpool(sizeof(TreeNode))
        {
            /* NOP */
        }

        virtual TreeNode *AllocNode(int val)
        {
            TreeNode *node = reinterpret_cast<TreeNode*>(m_mpool.malloc());
            node->val = val;
            node->left = nullptr;
            node->right = nullptr;
            return node;
        }

        virtual void Dump(TreeNode *root, std::ostream &data)
        {
            DumpEntry(root, data);
        }

        virtual void Dump(TreeNode *root, std::string &data)
        {
            std::stringstream content;
            Dump(root, content);
            data = boost::algorithm::trim_copy(content.str());
        }

        virtual TreeNode* Load(std::istream &data)
        {
            std::istream_iterator<std::string> itr(data);
            std::istream_iterator<std::string> eos;
            TreeNode *root = NULL;
            LoadEntry(&root, itr, eos);
            return root;
        }

        virtual TreeNode* Load(const char *data)
        {
            std::stringstream content(data == NULL ? "" : data);
            return Load(content);
        }

    private:
        void DumpEntry(TreeNode *node, std::ostream &data)
        {
            if (nullptr == node)
            {
                data << "# ";
                return;
            }
            data << node->val << " ";
            DumpEntry(node->left, data);
            DumpEntry(node->right, data);
        }
        
        void LoadEntry(TreeNode **node,
                       std::istream_iterator<std::string> &itr,
                       std::istream_iterator<std::string> &eos)
        {
            if (itr == eos)
            {
                return;
            }

            if (itr->compare("#") == 0)
            {
                itr++;
                return;
            }

            *node = AllocNode(std::atoi(itr->c_str()));
            itr++;
            LoadEntry(&(*node)->left, itr, eos);
            LoadEntry(&(*node)->right, itr, eos);
        }
    };

    TreeNodes::TreeNodes(void)
    {
        /* NOP */
    }

    boost::shared_ptr<TreeNodes> TreeNodes::Create(void)
    {
        return boost::make_shared<TreeNodesImpl>();
    }
}

BOOST_AUTO_TEST_SUITE(misc)
BOOST_AUTO_TEST_CASE(tree_nodes)
{
    auto tree = gazlowe::TreeNodes::Create();
    /* Testing tree
     *          0
     *       /      \
     *      1        2
     *     / \      / \
     *    3   4    5   #
     *   / \ / \  / \
     *  #  # # # #   #
     *
     * The dump result:
     * [0 1 3 # # 4 # # 2 5 # # #]
     */
    auto root = tree->AllocNode(0);
    auto node = tree->AllocNode(1);
    root->left = node;
    node = tree->AllocNode(2);
    root->right = node;
    node = tree->AllocNode(3);
    root->left->left = node;
    node = tree->AllocNode(4);
    root->left->right = node;
    node = tree->AllocNode(5);
    root->right->left = node;

    std::string data;
    tree->Dump(root, data);
    BOOST_REQUIRE_EQUAL(data,
                        "0 1 3 # # 4 # # 2 5 # # #");

    root = tree->Load(data.c_str());
    data = "";
    tree->Dump(root, data);
    BOOST_REQUIRE_EQUAL(data,
                        "0 1 3 # # 4 # # 2 5 # # #");
}
BOOST_AUTO_TEST_SUITE_END()
