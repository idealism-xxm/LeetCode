// 链接：https://leetcode.com/problems/count-good-nodes-in-binary-tree/
// 题意：给定一棵二叉树 root ，求有多少个结点 x 满足以下条件？
//      从根结点 root 到结点 x 的路径上所有结点的值都不大于 x.val 。


// 数据限制：
//  二叉树的结点数范围为 [1, 10 ^ 5]
//  每个结点值的范围为 [-(10 ^ 4), 10 ^ 4]


// 输入： root = [3,1,4,3,null,1,5]
// 输出： 4
// 解释： 以下 4 个带括号的结点满足题意。
//         (3)
//         / \
//        1  (4)
//       /   / \
//     (3)  1  (5)

// 输入： root = [3,3,null,4,2]
// 输出： 3
// 解释： 以下 3 个带括号的结点满足题意。
//         (3)
//         /
//       (3)
//       / \
//     (4)  2


// 思路： 递归/DFS
//
//      我们在递归遍历的时候，维护根结点到结点 x 路径上的最大值 max_val 。
//
//      只要 x.val >= max_val ，那么结点 x 就满足题意，计入 ans 中。
//
//      然后用 x.val 更新 max_val 的最大值，再递归处理即可。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历二叉树中全部 O(n) 个结点
//      空间复杂度：O(h)
//          1. 栈递归深度就是树高 O(h) ，最差情况下，全部 O(n) 个结点在一条链上


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
    pub fn good_nodes(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        Self::dfs(&root, i32::MIN)
    }

    // dfs(root, max_val) 递归返回子树 root 中满足以下条件的结点 x 的数量：
    //  1. root 到 x 的路径上所有结点的值都不大于 max_val
    //  2. root 到 x 的路径上所有结点的值都不大于 x.val
    fn dfs(root: &Option<Rc<RefCell<TreeNode>>>, mut max_val: i32) -> i32 {
        // 如果当前结点为空，则直接返回 0
        if root.is_none() {
            return 0;
        }

        let root = root.as_ref().unwrap().borrow();
        // ans 表示 root 是否为满足题意的结点
        let mut ans = 0;
        // 如果实际上 root 满足题意，设置 ans 为 1 ，
        // 同时更新 max_val 的最大值为 root.val 。
        if root.val >= max_val {
            ans = 1;
            max_val = root.val;
        }

        // 递归处理左右子树，返回其结果之和
        ans + Self::dfs(&root.left, max_val) + Self::dfs(&root.right, max_val)
    }
}
