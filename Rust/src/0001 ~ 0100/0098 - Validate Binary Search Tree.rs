// 链接：https://leetcode.com/problems/validate-binary-search-tree/
// 题意：给定一个二叉树，判断是不是二叉搜索树？


// 数据限制：
//  树的结点数在 [1, 10 ^ 4] 内
//  -(2 ^ 31) <= Node.val <= 2 ^ 31 - 1


// 输入：[2,1,3]
// 输出：true
// 解释： 2
//      / \
//     1   3

// 输入：[5,1,4,null,null,3,6]
// 输出：false
// 解释： 5
//      / \
//     1   4
//        / \
//       3   6


// 思路1：递归
//
//      二叉搜索树满足以下条件：左子树所有的值 < root.Val < 右子树所有的值。
//
//      那么我们可以用 dfs(root, low, high) 判断 root 是否是二叉搜索树。
//          1. root: 待判断的二叉搜索树的结点
//          2. low: root 结点祖先结点的最小值，空表示 -∞
//          3. high: root 结点祖先结点的最大值，空表示 +∞
//
//      然后在 dfs 内按照如下逻辑处理即可：
//          1. root 为空，则空子树必定满足题意，直接返回 true
//          2. low 不为空，若此时 root.val <= low，则返回 false
//          3. high 不为空，若此时 root.val >= high，则返回 false
//          4. 递归处理左右子树，左子树的所有结点值需要在 (low, root.val) 内，
//              右子树的所有结点值需要在 (root.val, high) 内
//
//
//      时间复杂度：O(n)
//          1. 需要遍历全部 O(n) 个结点
//      空间复杂度：O(h)
//          1. 栈递归深度就是树高 h ，最差情况下，全部 O(n) 个结点在一条链上     


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
    pub fn is_valid_bst(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        Self::dfs(&root, None, None)
    }

    fn dfs(root: &Option<Rc<RefCell<TreeNode>>>, low: Option<i32>, high: Option<i32>) -> bool {
        // 空子树必定满足题意
        if root.is_none() {
            return true
        }

        let root = root.as_ref().unwrap().borrow();
        // 如果存在最小值，则根结点的值不能 <= 最小值
        if low.is_some() && root.val <= low.unwrap() {
            return false
        }
        //如果存在最大值，则 根结点的值不能 >= 最大值
        if high.is_some() && root.val >= high.unwrap() {
            return false
        }

        // 此时需要递归处理左右子树：
        //  1. 左子树的最大值不能 >= min(low, root.val) = root.val ，
        //      则左子树的所有结点值需要在 (low, root.val) 内。
        //  2. 右子树的最小值不能 <= max(low, root.val) = root.val ，
        //      则右子树的所有结点值需要在 (root.val, high) 内。
        Self::dfs(&root.left, low, Some(root.val)) && Self::dfs(&root.right, Some(root.val), high)
    }
}
