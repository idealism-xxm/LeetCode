// 链接：https://leetcode.com/problems/binary-tree-right-side-view/
// 题意：给定一棵二叉树，想象你站在它右侧，
//      按照从顶部到底部的顺序，返回你能看见的结点值的数组。


// 数据限制：
//  二叉树的结点数的范围为 [0, 100]
//  -100 <= Node.val <= 100


// 输入： root = [1,2,3,null,5,null,4]
// 输出： [1,3,4]
// 解释： 1
//      / \
//     2   3
//      \   \
//       5   4

// 输入： root = [1,null,3]
// 输出： [1,3]
// 解释： 1
//        \
//         3

// 输入： root = []
// 输出： []


// 思路： 递归/DFS
//
//      用 ans 维护每一层最右侧结点的值，
//      然后使用 dfs(root, depth, ans) 先序遍历收集每一层最右侧结点的值即可。
//
//      dfs(root, depth, ans) 的参数含义如下：
//          1. root: 当前处理的 root 子树，初始为二叉树的根
//          2. depth: root 结点的深度，初始二叉树的根的深度为 1
//          3. ans: 已收集的每一层最右侧结点的值，初始为空
//
//      在 dfs 中，如果 root 不为空，则 root 有可能是当前层最右侧的结点，
//      可以通过 depth 和 len(ans) 的关系来判断：
//          1. depth <= len(ans): 已收集了前 len(ans) 最右侧结点的值，
//              则此时 root 不是当前层最右侧的结点，不做处理
//          2. depth > len(ans): 则必有 depth = len(ans) + 1 ，
//              此时 root 就是当前层最右侧的结点，将其值收集到 ans 中
//
//      然后递归遍历子树，注意先递归遍历右子树，再递归遍历左子树，
//      这样后续遇到每一层的第一个结点，必定是该层最右侧结点。
//
//
//		时间复杂度： O(n)
//          1. 需要遍历全部 O(n) 个结点一次
//		空间复杂度： O(n)
//          1. 栈递归深度就是树高，最差情况下，全部 O(n) 个结点在一条链上
//          2. 需要记录每一层最右侧结点的值，也就是树高，
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
    pub fn right_side_view(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        // ans 用于收集每一层最右侧结点的值
        let mut ans = vec![];
        // dfs 先序遍历收集每一层最右侧结点的值
        Self::dfs(&root, 1, &mut ans);

        ans
    }

    fn dfs(root: &Option<Rc<RefCell<TreeNode>>>, depth: usize, ans: &mut Vec<i32>) {
        // 如果 root 不为空，则继续先序遍历 root 子树
        if let Some(root) = root {
            // 如果深度大于 ans 的长度，则 root.val 是当前层最右侧结点的值，需要收集
            if depth > ans.len() {
                ans.push(root.borrow().val);
            }

            // 先递归遍历右子树，在递归遍历左子树，
            // 这样后续遇到每一层的第一个结点，必定是该层最右侧结点
            Self::dfs(&root.borrow().right, depth + 1, ans);
            Self::dfs(&root.borrow().left, depth + 1, ans);
        }
    }
}
