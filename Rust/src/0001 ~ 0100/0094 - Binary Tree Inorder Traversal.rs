// 链接：https://leetcode.com/problems/binary-tree-inorder-traversal/
// 题意：给定一个二叉树，返回中序遍历的结果？
//
//      进阶：使用循环的方式进行中序遍历。


// 数据限制：
//  二叉树的结点范围为 [0, 100]
//  -100 <= Node.val <= 100


// 输入： root = [1,null,2,3]
// 输出： [1,3,2]
// 解释： 1
//        \
//         2
//        /
//       3

// 输入： root = []
// 输出： []

// 输入： root = [1]
// 输出： [1]


// 思路1：递归/DFS
//
//      递归版本的中序遍历非常简单，核心就是要保证收集顺序如下：
//          1. 先收集左子树的值
//          2. 再收集当前结点的值
//          3. 最后收集右子树的值
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
    pub fn inorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        // ans 用于收集中序遍历的结果，传入 dfs 中递归处理即可
        let mut ans = vec![];
        Self::dfs(&root, &mut ans);

        ans
    }

    fn dfs(root: &Option<Rc<RefCell<TreeNode>>>, ans: &mut Vec<i32>) {
        // 如果当前子树为空，则直接返回
        if root.is_none() {
            return;
        }

        let mut root = root.as_ref().unwrap().borrow();
        // 中序遍历就是需要先收集左子树的值，再收集当前结点的值，最后收集右子树的值
        Self::dfs(&root.left, ans);
        ans.push(root.val);
        Self::dfs(&root.right, ans);
    }
}


//      递归很容易就能写出来，转换成循环就需要用栈记录信息
//      若当前结点不为 nil ，则入栈，然后处理左子结点
//      若当前结点为 nil ，则出栈，记录出栈结点的值，并将当前结点指向其右子结点