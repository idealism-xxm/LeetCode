// 链接：https://leetcode.com/problems/binary-tree-pruning/
// 题意：给定一棵二叉树 root ，将所有不含值为 1 的结点的子树移除，
//      返回结果二叉树的根。


// 数据限制：
//  二叉树的结点范围为 [0, 200]
//  Node.val 的值是 0 或 1


// 输入： root = [1,null,0,0,1]
// 输出： [1,null,0,null,1]
// 解释： 带括号的子树不含值为 1 的子结点，需要修剪掉。 
//         1              1
//          \              \
//           0      →       0
//          / \              \
//        (0)  1              1

// 输入： root = [1,0,1,0,0,0,1]
// 输出： [1,null,1,null,1]
// 解释： 带括号的子树不含值为 1 的子结点，需要修剪掉。
//         1              1
//       /   \             \
//      0     1     →       1
//     / \   / \           / \
//    0   0 1   1         1   1

// 输入： root = [1,1,0,1,1,0,1,0]
// 输出： [1,1,0,1,1,null,1]
// 解释： 带括号的子树不含值为 1 的子结点，需要修剪掉。
//          1              1
//        /   \           / \
//       1     0     →   1   0
//      / \   / \           / \
//     1   1 0   1         1   1
//    /
//   0


// 思路： 递归/DFS
//
//      对于一棵树 root 来说，我们必定要优先修剪其左右子树，
//      这就是一个递归处理的过程。
//
//      修剪左右子树完成后，我们可以确定 root.left/root.right 的状态，
//      有两种情况：
//          1. 为空：则表明左/右子树中的不含值为 1 的子结点，则整棵子树已被修剪掉
//          2. 非空：则表明左/右子树中的含有值为 1 的子结点，但子树内部已被修剪过
//
//      此时就可以确定是否需要修剪掉整棵树 root ，
//      如果 root.val == 1 或者 其左/右子树不为空，
//      则整棵树 root 无需修剪掉（但其内部已被修剪过），
//      直接返回 root 即可。
//
//      否则，整棵树 root 需要被修剪掉，返回空。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历全部 O(n) 个结点
//      空间复杂度：O(n)
//          1. 栈递归深度就是树高 O(h) ，
//              最差情况下，全部 O(n) 个结点在一条链上


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
    pub fn prune_tree(mut root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        // 如果 root 子树为空，则直接返回空
        if root.is_none() {
            return None
        }
        
        let mut node = root.as_mut().unwrap().borrow_mut();
        // 修剪左右子树，并将修剪后的子树设置到对应位置
        node.left = Self::prune_tree(node.left.take());
        node.right = Self::prune_tree(node.right.take());
        // 如果当前结点的值为 1 或者左/右子树存在，
        // 则当前子树必定含有值为 1 的子结点，直接返回 root 即可
        if node.val == 1 || node.left.is_some() || node.right.is_some() {
            // 手动丢弃 node 的所有权，避免借用 root 而无法正常返回
            drop(node);
            return root;
        }

        // 此时当前子树中的所有子结点都是 0 ，需要修剪掉
        None
    }
}
