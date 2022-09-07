// 链接：https://leetcode.com/problems/construct-string-from-binary-tree/
// 题意：给定一棵二叉树 root ，将其转化成由括号和数字组成的字符串。
//      空结点用一对空括号 "()" 表示，
//      转化后要省略所有不影响字符串与原始二叉树一对一映射的空括号对。


// 数据限制：
//  二叉树的结点范围为 [1, 10 ^ 4]
//  -1000 <= Node.val <= 1000


// 输入： root = [1,2,3,4]
// 输出： "1(2(4))(3)"
// 解释： 转化的原始字符串为 "1(2(4)())(3()())" ，
//       省略所有非必需的空括号对后，最终变为 "1(2(4))(3)"
//         1
//        / \
//       2   3
//      /
//     4

// 输入： root = [1,2,3,null,4]
// 输出： "1(2()(4))(3)"
// 解释： 转化的原始字符串为 "1(2()(4))(3()())" ，
//       省略所有非必需的空括号对后，最终变为 "1(2()(4))(3)"
//      （注意第一个空括号度不能省略，否则会影响一对一映射）
//         1
//        / \
//       2   3
//        \
//         4


// 思路： 递归/DFS
//
//      对于一棵树 root 来说，我们必定要优先转化其左右子树，
//      这就是一个递归处理的过程。
//
//      处理完左右子树后，我们会获得对应的字符串，记为 left 和 right ，
//      然后我们再生成当前结点值对应的字符串 val 。
//
//      此时我们可以根据 left 和 right 是否为空，来格式化树 root 对应的字符串：
//          1. left 和 right 均不为空：则三部分必定都存在，
//              结果格式为 {val}({left})({right})
//          2. left 为空，但 right 不为空：则三部分必定都存在，
//              结果格式为 {val}()({right})
//          3. left 不为空，但 right 为空：则只有前两部分存在，
//              结果格式为 {val}({left})
//          4. left 和 right 均为空：则只有第一部分存在，
//              结果格式为 {val}
//
//      其中 1 和 2 两种情况可以合并为一种，条件合并为 right 不为空，
//      结果格式合并为 {val}({left})({right})
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
    pub fn tree2str(mut root: Option<Rc<RefCell<TreeNode>>>) -> String {
        // 如果 root 为空，则直接返回空串
        if root.is_none() {
            return "".to_string()
        }

        // 提前计算好三部分的字符串（结点值字符串、左子树字符串、右子树字符串）
        let mut root = root.as_mut().unwrap().borrow_mut();
        let val = root.val.to_string();
        let left = Self::tree2str(root.left.take());
        let right = Self::tree2str(root.right.take());

        if !right.is_empty() {
            // 如果右子树字符串不为空，则三部分必定都要存在
            format!("{val}({left})({right})")
        } else if !left.is_empty() {
            // 如果右子树字符串为空，但左子树字符串不为空，则只有前两部分存在
            format!("{val}({left})")
        } else {
            // 如果左右子树字符串都为空，则只有第一部分存在
            val
        }
    }
}
