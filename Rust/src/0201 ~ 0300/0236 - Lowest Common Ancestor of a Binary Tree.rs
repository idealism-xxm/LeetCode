// 链接：https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
// 题意：给定一棵二叉树，找到两个结点的最近公共祖先？


// 数据限制：
//  树的结点数在 [2, 10 ^ 5] 之间
//	-(10 ^ 9) <= Node.val <= 10 ^ 9
//	所有的 Node.val 都各不相同
//	p != q
//	p 和 q 必定在树中


// 输入： root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// 输出： 3
// 解释： 结点 5 和 1 的 LCA 是 3
//       3
//     /   \
//   5       1
//  / \     / \
// 6   2   0   8
//    / \
//   7   4

// 输入： root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
// 输出： 5
// 解释： 结点 5 和 4 的 LCA 是 5
//       3
//     /   \
//   5       1
//  / \     / \
// 6   2   0   8
//    / \
//   7   4


// 思路： 递归
//
//      我们使用递归处理即可，返回值有两个：
//           cnt 表示当前子树中含有的 p 和 q 的结点数
//          node 表示 cnt 为 2 时当前子树中 p 和 q 的 LCA ，其他时候无意义
//
//      然后按照以下步骤继续处理：
//          1. root 为空：直接返回 None, 0
//          2. 递归计算左子树结果 left_node, left_cnt 。
//              若 left_cnt == 2 ，则 left_node 就是 LCA ，直接返回即可
//          3. 递归计算右子树结果 right_node, right_cnt 。
//              若 right_cnt == 2 ，则 right_node 就是 LCA ，直接返回即可
//          4. 计算当前子树含有的 p 和 q 的结点数 cur_cnt = left_cnt + right_cnt 。
//              若 root == p || root == q ，则还需要对 cur_cnt + 1
//          5. 返回 root, cur_cnt （若 cur_cnt == 2 ，则 root 就是 LCA ）
//
//      如果本题是一棵树有多次查询，那么可以使用 tarjan 或者倍增进行处理。
//
//
//      时间复杂度： O(n)
//            1. 需要遍历全部 O(n) 个结点一次
//      空间复杂度： O(h)
//            1. 栈递归深度就是树高，最差情况下，全部 O(n) 个结点在一条链上


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
    pub fn lowest_common_ancestor(root: Option<Rc<RefCell<TreeNode>>>, p: Option<Rc<RefCell<TreeNode>>>, q: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        Self::dfs(root, &p, &q).0
    }

    fn dfs(root: Option<Rc<RefCell<TreeNode>>>, p: &Option<Rc<RefCell<TreeNode>>>, q: &Option<Rc<RefCell<TreeNode>>>) -> (Option<Rc<RefCell<TreeNode>>>, i32) {
        // root 为空，则直接返回 None, 0
        if root.is_none() {
            return (None, 0);
        }

        // 递归计算左子树结果 left_node, left_cnt
        let (left_node, left_cnt) = Self::dfs(root.as_ref().unwrap().borrow().left.clone(), p, q);
        // 若 left_cnt == 2 ，则 left_node 就是 LCA ，直接返回即可
        if left_cnt == 2 {
            return (left_node, 2);
        }

        // 递归计算右子树结果 right_node, right_cnt
        let (right_node, right_cnt) = Self::dfs(root.as_ref().unwrap().borrow().right.clone(), p, q);
        // 若 right_cnt == 2 ，则 right_node 就是 LCA
        if right_cnt == 2 {
            return (right_node, 2);
        }

        // 计算当前子树含有的 p 和 q 的结点数
        let mut cur_cnt = left_cnt + right_cnt;
        // 同时要统计 root 可能为 p 或 q 的情况
        if &root == p || &root == q {
            cur_cnt += 1;
        }

		// 若 cur_cnt == 2 ，则 root 就是 LCA
        (root, cur_cnt)
    }
}
