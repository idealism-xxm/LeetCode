// 链接：https://leetcode.com/problems/invert-binary-tree/
// 题意：给定一个二叉树，翻转所有节点的左右子节点？

// 输入：
//      4
//    /   \
//   2     7
//  / \   / \
// 1   3 6   9
// 输出：
//      4
//    /   \
//   7     2
//  / \   / \
// 9   6 3   1

// 思路： 递归
//
//      递归翻转即可，
//      当前节点是空节点，则直接返回，
//      当前节点不是空节点，则递归处理左右子节点，然后在再换左右子节点即可
//
//      时间复杂度： O(n)
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
    pub fn invert_tree(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        // 如果是空节点，则直接返回
        if root.is_none() {
            return root
        }

        // 递归处理左右子节点
        let mut root = root;
        let root_node = root.as_mut().unwrap();
        let left = Solution::invert_tree(root_node.borrow().left.clone());
        let right = Solution::invert_tree(root_node.borrow().right.clone());

        // 交换左右子节点
        root_node.borrow_mut().left = right;
        root_node.borrow_mut().right = left;

        root
    }
}
