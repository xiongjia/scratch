/**
 * gazlowe - My algorithm tests
 */

#ifndef _G_MISC_HXX_
#define _G_MISC_HXX_  1


#include <istream>
#include <ostream>

#include "boost/utility.hpp"
#include "boost/shared_ptr.hpp"
#include "boost/array.hpp"
#include "boost/foreach.hpp"

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

    /* Tree nodes manager */
    class TreeNodes : boost::noncopyable
    {
    public:
        TreeNodes(void);

        static boost::shared_ptr<TreeNodes> Create(void);
    public:
        virtual TreeNode *AllocNode(int val) = 0;

        virtual void Dump(TreeNode *root, std::ostream &data) = 0;
        virtual void Dump(TreeNode *root, std::string &data) = 0;

        virtual TreeNode* Load(std::istream &data) = 0;
        virtual TreeNode* Load(const char *data) = 0;
    };

    /* ListNode struct for leetcode problems */
    typedef struct _ListNode ListNode;
    struct _ListNode 
    {
        int val;
        ListNode *next;
    };

    /* List nodes manager */
    class ListNodes : boost::noncopyable
    {
    public:
        ListNodes(void);

        static boost::shared_ptr<ListNodes> Create(void);
    public:
        virtual ListNode *AllocNode(int val) = 0;

        virtual void Dump(ListNode *list, std::string &result) = 0;
        virtual ListNode* Search(ListNode *list, int val) = 0;

        template<std::size_t N>
        static ListNode* Load(boost::shared_ptr<ListNodes> nodes,
                              const boost::array<int, N> &data)
        {
            ListNode *root = nullptr;
            ListNode *cur = nullptr;
            BOOST_FOREACH(int val, data)
            {
                ListNode *node = nodes->AllocNode(val);
                if (nullptr == root)
                {
                    root = node;
                    cur = node;
                }
                else
                {
                    cur->next = node;
                    cur = node;
                }
            }
            if (nullptr != cur)
            {
                cur->next = root;
            }
            return root;
        }
    };
}

#endif /* !defined(_G_MISC_HXX_) */

