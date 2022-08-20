// 链接：https://leetcode.com/problems/evaluate-boolean-binary-tree/
// 题意：给定一棵含有以下性质的满二叉树，求根节点 root 的运算结果？
//          1. 叶子结点的值为 0 或 1 ，其中 0 表示 False ，1 表示 True
//          2. 非叶子结点的值为 2 或 3 ，其中 2 表示或， 3 表示与
//
//      一个结点 node 的运算结果如下：
//          1. 如果 node 是叶子结点，则运算结果就是 node 的值，即 True 或 False
//          2. 如果 node 非叶子结点，则运算结果就是 node 的左右子树的运算结果的或/与结果
//
//      一颗满二叉树的子结点含有 0 个或 2 个子结点。


// 数据限制：
//  二叉树的结点数范围为 [1, 1000]
//  0 <= Node.val <= 3
//  每个子结点有 0 个或 2 个子结点
//  叶子结点的值为 0 或 1
//  非叶子结点的值为 2 或 3


// 输入： root = [2,1,3,null,null,0,1]
// 输出： true
// 解释： OR                 OR                True
//      /  \               /  \
//   True  AND      ->   True  False    ->
//        /   \
//      False True

// 输入： root = [0]
// 输出： false


// 思路： 递归/DFS
//
//      由于题目给定的数据合法，所以我们无需判断每个结点是否为叶子结点，
//      直接针对每个结点的值 node.val 进行判断，执行对应的操作即可：
//          1. node.val == 0 ， node 是叶子结点，直接返回 False
//          2. node.val == 1 ， node 是叶子结点，直接返回 True
//          3. node.val == 2 ， node 是非叶子结点，返回左右子树的运算结果的或结果
//          4. node.val == 3 ， node 是非叶子结点，返回左右子树的运算结果的与结果
//
//
//      时间复杂度： O(n)
//          1. 需要遍历全部 O(n) 个结点
//      空间复杂度： O(h)
//          1. 栈递归深度就是树高 h ，最差情况下， O(n) 个结点在一条链上


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
    pub fn evaluate_tree(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        let mut node = root.as_ref().unwrap().borrow();
        // 根据结点的值进行处理
        match node.val {
            // 叶子结点，直接返回 false
            0 => false,
            // 叶子结点，直接返回 true
            1 => true,
            // 非叶子结点，返回左右子树的运算结果的或结果
            2 => Solution::evaluate_tree(node.left.clone()) || Solution::evaluate_tree(node.right.clone()),
            // 非叶子结点，返回左右子树的运算结果的与结果
            3 => Solution::evaluate_tree(node.left.clone()) && Solution::evaluate_tree(node.right.clone()),
            // 其他情况不存在
            _ => unreachable!(),
        }
    }
}
