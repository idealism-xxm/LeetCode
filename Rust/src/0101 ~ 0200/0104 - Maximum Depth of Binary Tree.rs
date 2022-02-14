// 链接：https://leetcode.com/problems/maximum-depth-of-binary-tree/
// 题意：给定一棵二叉树，返回其高度？


// 数据限制：
//  树中结点的数量范围是 [0, 10 ^ 4]
//  -100 <= Node.val <= 100


// 输入：root = [3,9,20,null,null,15,7]
// 输出：3
// 解释：
//    3
//   / \
//  9  20
//    /  \
//   15   7

// 输入：root = [1,null,2]
// 输出：2
// 解释：
//    1
//     \
//      2


// 思路：递归
//
//      从根结点开始递归处理，每一次进行如下判断：
//          1. 如果根结点不存在，则树的高度是 0 ，直接返回 0
//          2. 如果根结点存在，则树的高度是 1 + max(左子树高度, 右子树高度)
//
//      时间复杂度：O(n) 。每一个结点都需要便利，所以时间复杂度是 O(n) 。
//      空间复杂度：O(n) 。最差情况下，树的高度是 n ，需要开辟 n 层栈空间，所以空间复杂度是 O(n) 。

// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//   pub val: i32,
//   pub left: Option<Rc<RefCell<TreeNode>>>,
//   pub right: Option<Rc<RefCell<TreeNode>>>,
// }

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
    pub fn max_depth(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        if let Some(root) = root {
            // 如果根结点存在，则返回左右子树的最大深度加 1
            1 + i32::max(
                Self::max_depth(root.borrow().left.clone()), 
                Self::max_depth(root.borrow().right.clone()),
            )
        } else {
            // 如果根结点不存在，则返回 0
            0
        }
    }
}
