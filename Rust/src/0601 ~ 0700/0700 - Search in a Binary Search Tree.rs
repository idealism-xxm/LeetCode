// 链接：https://leetcode.com/problems/search-in-a-binary-search-tree/
// 题意：给定一棵二叉搜索树 root 和一个整数 val ，
//      如果这棵二叉搜索树 root 中存在结点 node 的值是 val ，
//      则返回以该结点为根结点的子树，否则返回 null 。


// 数据限制：
//  这棵树的结点数在 [1, 5000] 内
//  1 <= Node.val <= 10 ^ 7
//  root 是一棵二叉搜索树
//  1 <= val <= 10 ^ 7


// 输入： root = [4,2,7,1,3], val = 2
// 输出： [2,1,3]
// 解释： 4
//      / \
//     2   7
//    / \
//   1   3

// 输入： root = [4,2,7,1,3], val = 5
// 输出： 5
// 解释： 4
//      / \
//     2   7
//    / \
//   1   3


// 思路： 迭代
//
//      使用根结点 root 不断迭代处理，直至 root 为空：
//          1. root.val == val: 已找到目标子树，直接返回 root
//          2. root.val >  val: 若 val 在这棵二叉搜索树中的话，
//              必定在当前结点的左子树中，因此 root = root.left
//          3. root.val <  val: 若 val 在这棵二叉搜索树中的话，
//              必定在当前结点的右子树中，因此 root = root.right
//
//      最后 root 为空，则直接返回 null
//
//
//      时间复杂度：O(n)
//          1. 最差情况下，需要遍历二叉搜索树中全部 O(n) 个结点
//      空间复杂度：O(1)
//          1. 只需要维护常数个额外变量


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
    pub fn search_bst(mut root: Option<Rc<RefCell<TreeNode>>>, val: i32) -> Option<Rc<RefCell<TreeNode>>> {
        // 如果当前根结点不为空，则继续处理
        while let Some(node) = root {
            if node.borrow().val == val {
                // 如果根结点的值就是 val ，则直接返回
                return Some(node);
            } else if node.borrow().val > val {
                // 如果根结点的值大于 val ，
                // 若 val 在这棵二叉搜索树中的话，
                // 必定在当前结点的左子树中
                root = node.borrow_mut().left.take();
            } else {
                // 如果根结点的值小于 val ，
                // 若 val 在这棵二叉搜索树中的话，
                // 必定在当前结点的右子树中
                root = node.borrow_mut().right.take();
            }
        }

        // 没有找到值为 val 的结点，则返回空
        None
    }
}
