/**
 * gazlowe - My algorithm tests
 */

#include <sstream>
#include <string>
#include <algorithm>

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

        virtual TreeNode* Load(std::istream &data)
        {
            std::for_each(std::istream_iterator<std::string>(data),
                          std::istream_iterator<std::string>(), 
                          [](std::string const &item) 
            {
                /* TODO Create the node again */
            });
            return nullptr;
        }

    private:
        virtual void DumpEntry(TreeNode *node, std::ostream &data)
        {
            if (nullptr == node)
            {
                data << " # ";
                return;
            }
            data << " " << node->val << " ";
            DumpEntry(node->left, data);
            DumpEntry(node->right, data);
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

    std::stringstream data;
    tree->Dump(root, data);

    std::stringstream content(data.str());
    root = tree->Load(content);
}
BOOST_AUTO_TEST_SUITE_END()
