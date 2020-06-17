// 链接：https://leetcode-cn.com/problems/count-complete-tree-nodes/
// 题意：给定一个完全二叉树，求节点个数？

// 输入：
//     1
//    / \
//   2   3
//  / \  /
// 4  5 6
// 输出： 6

// 思路： 二分
//
//		直接递归很简单，但不是这中等题想要的解法，
//      就直接看了题解，看到标题说明要用二分就立刻明白了
//
//      先求出二叉树的高度 height ，那么我们每次递归二分即可，
//      定义一个函数 dfs(root, height) 表示求以 root 为根的
//          完全二叉树（高度为 height ）的节点数量，
//      我们求其右子树的高度 right_height ，
//      1. 若为 height - 1 ，则左子树为高度 height - 1 的满二叉树，
//          可以直接计算节点数为 2 ^ (height - 1) - 1 ，
//          右子树可以递归处理 dfs(root.right, right_height)
//      2. 若不为 height - 1 ，则右子树为高度 right_height 的满二叉树，
//          可以直接计算节点数为 2 ^ right_height - 1 ，
//          左子树可以递归处理 dfs(root.left, height - 1)
//      （注意上面的结果还需要加上根节点 root 的数量）
//
//      时间复杂度： O(height ^ 2)
//      空间复杂度： O(height) 【栈的开销】

// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//   pub val: i32,
//   pub left: Option<Rc<RefCell<TreeNode>>>,
//   pub right: Option<Rc<RefCell<TreeNode>>>,
// }
//
// impl TreeNode {
//   #[inline]
//   pub fn new(val: i32) -> Self {
//     TreeNode {
//       val,
//       left: None,
//       right: None
//     }
//   }
// }
use std::rc::Rc;
use std::cell::RefCell;
impl Solution {
    pub fn count_nodes(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        // 先计算 root 树的高度
        let mut height = Solution::count_height(root.clone());

        // 二分计算 root 树的节点数
        Solution::dfs(root, height)
    }

    fn count_height(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let mut height = 0;

        // 计算树高
        let mut cur = root;
        while let Some(cur_node) = cur {
            height += 1;
            cur = cur_node.as_ref().borrow().left.clone();
        }

        height
    }

    fn dfs(root: Option<Rc<RefCell<TreeNode>>>, height: i32) -> i32 {
        // 空子树节点数量为 0
        if root.is_none() {
            return 0;
        }
        let root_node = root.as_ref().unwrap().borrow();
        // 获取右子树的高度
        let right_height = Solution::count_height(root_node.right.clone());
        // 左子树高度必定为 height - 1
        return if right_height == height - 1 {
            // 如果右子树高度等于 root 树高度 - 1 ，那么左子树必定是满二叉树，
            // 则左子树节点数量为 2 ^ (height - 1) - 1 ，右子树节点数量递归处理
            (1 << (height - 1)) + Solution::dfs(root_node.right.clone(), right_height)
        } else {
            // 如果右子树高度不等于 root 树高度 - 1 ，那么右子树必定是满二叉树，
            // 则右子树节点数量为 2 ^ right_height - 1 ，左子树节点数量递归处理
            Solution::dfs(root_node.left.clone(), height - 1) + (1 << right_height)
        }
    }
}
