// 链接：https://leetcode.com/problems/house-robber-iii/
// 题意：给定一棵二叉树，不能选择相邻的两个结点的数，求选择某些数的和的最大值？


// 数据限制：
//  树的结点数范围为 [1, 10 ^ 4]
//  0 <= Node.val <= 10 ^ 4


// 输入： root = [3,2,3,null,3,null,1]
// 输出： 7
// 解释： 选择 3, 3, 1 ，最大和为 3 + 3 + 1 = 7 
//      (3)
//      / \
//     2   3
//      \   \
//      (3) (1)

// 输入： root = [3,4,5,1,3,null,1]
// 输出： 9
// 解释： 选择 4, 5 ，最大和为 4 + 5 = 9
//       3
//      / \
//    (4) (5)
//    / \   \
//   1   3   1


// 思路： 树形 DP
//
//      本题是 LeetCode 213 的加强版，将环变成了树。
//
//      原来在环中只需要考虑相邻的数，而在树中需要考虑相邻层的数。
//
//      设 dp[root] = (pick[root], skip[root]) ：
//          1. pick[root] 表示在子树 root 中，选择 root 时，选择的数的最大和
//          2. skip[root] 表示在子树 root 中，不选 root 时，选择的数的最大和
//
//      对树的空结点初始化为 (0, 0) ，方便状态转移。
//            
//      状态转移：
//          1. 选择 root 时，则必定不能选其子结点：
//              pick[root] = root.val + skip[root.left] + skip[root.right]
//              
//          2. 不选 root 时，则子结点可选可不选，即子结点取两者最大值：
//              skip[root] = max(pick[root.left], skip[root.left]) +
//                           max(pick[root.right], skip[root.right])           
//
//      则最终结果就是 max(pick[root], skip[root])
//
//
//      时间复杂度： O(n)
//          1. 需要遍历计算树中全部 O(n) 个状态
//      空间复杂度： O(n)
//          2. 栈递归深度就是树高，最差情况下，全部 O(n) 个结点在一条链上


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
    pub fn rob(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        // pick 表示在子树 root 中，选择 root 时，选择的数的最大和
        // skip 表示在子树 root 中，不选 root 时，选择的数的最大和
        let (pick, skip) = Self::dfs(root);
        // 两者取最大值返回即可
        pick.max(skip)
    }
    
    fn dfs(root: Option<Rc<RefCell<TreeNode>>>) -> (i32, i32) {
        if let Some(root) = root {
            // 递归计算左右子结点的状态
            let (left_pick, left_skip) = Self::dfs(root.borrow_mut().left.take());
            let (right_pick, right_skip) = Self::dfs(root.borrow_mut().right.take());
            
            // 选择 root 时，则必定不能选其子结点
            let pick = root.borrow().val + left_skip + right_skip;
            // 不选 root 时，则子结点可选可不选，即子结点取两者最大值
            let skip = left_pick.max(left_skip) + right_pick.max(right_skip);
            
            (pick, skip)
        } else {
            // 如果是空结点，则直接返回 (0, 0)
            (0, 0)
        }
    }
}
