/**
 * gazlowe - My algorithm tests
 */

#ifndef _G_MISC_HXX_
#define _G_MISC_HXX_  1


#include <istream>
#include <ostream>
#include "boost/utility.hpp"
#include "boost/shared_ptr.hpp"

namespace gazlowe
{
    /* TreeNode struct for leetcode problems */
    typedef struct _TreeNode  TreeNode;

    struct _TreeNode
    {
        int val;
        TreeNode *left;
        TreeNode *right;
    };

    /* creating a new node on the mem pool */
    class TreeNodes : boost::noncopyable
    {
    public:
        TreeNodes(void);

        static boost::shared_ptr<TreeNodes> Create(void);

    public:
        virtual TreeNode *AllocNode(int val) = 0;

        virtual void Dump(TreeNode *root, std::ostream &data) = 0;
        virtual TreeNode* Load(std::istream &data) = 0;
    };
}

#endif /* !defined(_G_MISC_HXX_) */

