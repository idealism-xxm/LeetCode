// 链接：https://leetcode.com/problems/kth-smallest-element-in-a-bst/
// 题意：给定一棵二叉搜索树 root 和一个正整数 k ，
//      返回二叉搜索树中第 k 小的结点值。


// 数据限制：
//  树的结点数为 n
//  1 <= k <= n <= 10 ^ 4
//  0 <= Node.val <= 10 ^ 4


// 输入： root = [3,1,4,null,2], k = 1
// 输出： 1
// 解释： 3
//      / \
//     1   4
//      \
//       2

// 输入： root = [5,3,6,2,4,null,null,1], k = 3
// 输出： 3
// 解释： 5
//      / \
//     3   6
//    / \
//   2   4
//  /
// 1


// 思路： 递归
//
//      二叉搜索树是左子树的值小于根节点，右子树的值大于根节点。
//
//      可以使用 dfs(root, &remain, &ans) 来递归中序遍历，
//      其中， remain 表示还需遍历的剩余结点数，
//      ans 表示第 k 小的结点值。
//
//      返回值 true 表示已找到第 k 小的结点；
//      返回值 false 表示还没找到第 k 小的结点。
//
//      在 dfs 中，如果 root 为空，则直接返回。
//
//      否则，先递归处理左子树 dfs(root.left, remain, ans) ，
//      如果 dfs 返回 true，则表示已找到第 k 小的结点，
//      直接返回 true 。
//
//      然后,把当前结点纳入考量，即 *remain -= 1 ，
//      如果此时 remain 变为 0 ，则当前结点就是第 k 小的结点，
//      更新 ans 为当前结点的值，并返回 true 。
//
//      最后，还未返回的话，第 k 小的结点必定在右子树中，
//      继续递归处理右子树 dfs(root.right, remain, ans) 。
//
//
//      时间复杂度：O(n + k)
//          1. 需要先遍历到最小的结点，最差情况下，全部 O(n) 个结点在一条链上
//          2. 需要遍历最小的 O(k) 个节点
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
    pub fn kth_smallest(root: Option<Rc<RefCell<TreeNode>>>, mut k: i32) -> i32 {
        // ans 用于记录第 k 小的结点值
        let mut ans = 0;
        // 递归寻找第 k 小的结点
        Self::dfs(&root, &mut k, &mut ans);
        ans
    }

    fn dfs(root: &Option<Rc<RefCell<TreeNode>>>, remain: &mut i32, ans: &mut i32) -> bool {
        // 如果当前结点为空，则直接返回
        if root.is_none() {
            return false;
        }
        // 优先递归处理左子树，寻找第 k 小的
        if Self::dfs(&root.as_ref().unwrap().borrow().left, remain, ans) {
            // 如果在左子树中找到了第 k 小的，则直接返回
            return true;
        }

        // 把当前结点纳入考量
        *remain -= 1;
        // 如果剩余需要考虑的结点数为 0 ，
        // 则说明找到了第 k 小的结点
        if *remain == 0 {
            // 记录当前结点值，并返回
            *ans = root.as_ref().unwrap().borrow().val;
            return true;
        }

        // 此时第 k 小的结点还未找到，继续递归处理右子树
        Self::dfs(&root.as_ref().unwrap().borrow().right, remain, ans)
    }
}
