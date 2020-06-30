// 链接：https://leetcode.com/problems/kth-smallest-element-in-a-bst/
// 题意：给定一颗平衡二叉树，求第 k 小的节点的值？

// 输入： root = [3,1,4,null,2], k = 1
// 输出： true
// 解释：
//    3
//   / \
//  1   4
//   \
//    2

// 输入： root = [5,3,6,2,4,null,null,1], k = 3
// 输出： 3
// 解释：
//        5
//       / \
//      3   6
//     / \
//    2   4
//   /
//  1

// 思路： 递归
//
//      直接使用递归中序遍历即可，当遍历到第 k 个节点时，终止递归返回
//
//      【进阶】如果这个平衡二叉树会进行频繁的插入/删除操作，如何设计这个算法？
//          我们在进行刚才递归的操作就可以发现，我们每次递归前都要先计算左子树的节点数，
//          如果我们对每个节点维护一个值表示其左子树的节点数，
//          这样我们每次递归时都知道应该处理哪一个子树，不必全部遍历
//
//      时间复杂度： O(h + k)
//      空间复杂度： O(h)

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
    pub fn kth_smallest(root: Option<Rc<RefCell<TreeNode>>>, k: i32) -> i32 {
        // 递归处理
        let (_, result) = Solution::dfs(&root, k);
        result
    }

    // 最多遍历 k 个元素，返回还需要遍历的元素个数和最后一个元素的值
    fn dfs(root: &Option<Rc<RefCell<TreeNode>>>, k: i32) -> (i32, i32) {
        // 如果是空节点，则直接返回 k
        if root.is_none() {
            return (k, 0);
        }

        let root_node = root.as_ref().unwrap().borrow();
        // 递归遍历
        let (remain, last) = Solution::dfs(&root_node.left, k);
        // 若左子树中已经找到结果，则直接返回
        if remain == 0 {
            return (remain, last);
        }
        // 若还差一个元素，则必定是当前的根节点
        if remain == 1 {
            return (0, root_node.val);
        }
        // 若还差不止一个元素，则要么在右子树中，要么在其某个祖先节点的右子树中
        Solution::dfs(&root_node.right, remain - 1)
    }
}
