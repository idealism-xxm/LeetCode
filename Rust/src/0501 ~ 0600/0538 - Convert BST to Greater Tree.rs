// 链接：https://leetcode.com/problems/convert-bst-to-greater-tree/
// 题意：给定一棵二叉搜索树 root ，
//      将这个树的每个结点的值加上所有大于它的结点的值之和，
//      最后再返回结果树的根节点。


// 数据限制：
//  这棵树的结点数在  [0, 10 ^ 4] 内
//  -(10 ^ 4) <= Node.val <= 10 ^ 4
//  这棵树所有结点的值均不相同
//  root 是一棵合法的二叉搜索树


// 输入： root = [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
// 输出： [30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
// 解释：       4                          30
//          /     \                    /      \
//         1       6                 36       21
//        / \     / \       →        / \      / \   
//       0   2   5   7             36  35    26  15
//          /         \                /          \
//         3           8              33           8

// 输入： root = [0,null,1]
// 输出： [1,null,1]
// 解释： 0            1
//        \     →      \
//         1            1


// 思路： 递归
//
//      二叉搜索树是左子树的值小于根节点，右子树的值大于根节点。
//
//      所以我们可以通过递归的方式来遍历这棵树，
//      并维护一个值 sum ，
//      表示大于等于当前结点值的所有结点值之和。
//
//      递归时，如果当前结点 root 为空，则直接返回。
//
//      如果当前结点 root 不为空，则先递归处理右子树，
//      因为右子树的结点值一定大于 root.val 。
//
//      然后再对 sum 加上当前结点值 root.val ，
//      这时 sum 变为所有大于等于 root.val 的结点值之和。
//
//      此时，按照题意 root.val 应该等于 sum ，
//      即 root.val = sum 。
//
//      最后再递归处理左子树即可。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历二叉搜索树中全部 O(n) 个结点
//      空间复杂度：O(n)
//          1. 栈递归深度就是树高，最差情况下，全部 O(n) 个结点在一条链上


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
    pub fn convert_bst(mut root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        // 为了避免使用 clone() ，采用 &mut 方式递归处理。
        // sum 维护大于等于当前结点值的所有结点值之和
        Self::convert(&mut root, &mut 0);
        root
    }

    pub fn convert(root: &mut Option<Rc<RefCell<TreeNode>>>, sum: &mut i32) {
        if let Some(root) = root.as_mut() {
            // 如果 root 不为空，则继续处理
            let mut root = root.borrow_mut();
            // 先递归处理右子树，因为右子树的结点值一定大于 root.val
            Self::convert(&mut root.right, sum);
            // 再加上当前结点的值
            *sum += root.val;
            // 此时 sum 是大于等于 root.val 的所有结点值之和，
            // 按照题意 root.val 应该等于 sum
            root.val = *sum;
            // 再递归处理左子树
            Self::convert(&mut root.left, sum);
        }
    }
}
