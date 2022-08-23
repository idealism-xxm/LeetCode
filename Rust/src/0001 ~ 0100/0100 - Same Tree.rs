// 链接：https://leetcode.com/problems/same-tree/
// 题意：判断两个二叉树是否一样？


// 数据限制：
//  树的结点数范围为 [0, 100]
//  -(10 ^ 4) <= Node.val <= 10 ^ 4


// 输入： p = [1,2,3], q = [1,2,3]
// 输出： true
// 解释： 1             1
//      / \           / \
//     2   3         2   3

// 输入： p = [1,2], q = [1,null,2]
// 输出： false
// 解释： 1             1
//      /               \
//     2                 2

// 输入： p = [1,2,1], q = [1,1,2]
// 输出： false
// 解释： 1             1
//      / \           / \
//     2   1         1   2

// 思路：递归/DFS
//
//      递归处理即可
//          1. 若当前结点的值相同，则左子树和右子树都相同才返回 true
//          2. 若当前结点都值不同，则直接返回 false
//      
//      时间复杂度： O(n)
//          1. 需要遍历全部 O(n) 个结点
//      空间复杂度： O(h)
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
    pub fn is_same_tree(p: Option<Rc<RefCell<TreeNode>>>, q: Option<Rc<RefCell<TreeNode>>>) -> bool {
        // 采用借用的形式处理，避免递归中使用 clone
        Self::dfs(&p, &q)
    }

    pub fn dfs(p: &Option<Rc<RefCell<TreeNode>>>, q: &Option<Rc<RefCell<TreeNode>>>) -> bool {
        match (p, q) {
            // 都是空结点，则当前子树相同，直接返回 true
            (None, None) => true,
            // 都是非空结点，则需要判断当前结点的值是否相同，以及左右子树是否相同
            (Some(p), Some(q)) => {
                let p = p.borrow();
                let q = q.borrow();
                p.val == q.val && Self::dfs(&p.left, &q.left) && Self::dfs(&p.right, &q.right),
            },
            // 一个非空，另一个为空，则当前子树不同，直接返回 false
            _ => false,
        }
    }
}
