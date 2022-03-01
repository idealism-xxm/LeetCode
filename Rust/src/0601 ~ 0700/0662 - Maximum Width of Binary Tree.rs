// 链接：https://leetcode.com/problems/maximum-width-of-binary-tree/
// 题意：给定一棵二叉树，求其最大宽度？
//
//      最大宽度：该二叉树所有层的宽度的最大值。
//      某一层的宽度：该层最左侧非空结点和最右侧非空结点之间的长度。
//                  （计算长度时，包含空结点）


// 数据限制：
//  树中的结点数量范围是 [1, 3000]
//  -100 <= Node.val <= 100


// 输入： root = [1,3,2,5,3,null,9]
// 输出： 4
// 解释： 第三层的宽度最大，结点为 (5,3,null,9) ，宽度为 4 
//      1
//     / \
//    3   2
//   / \   \
//  5   3   9

// 输入： root = [1,3,null,5,3]
// 输出： 2
// 解释： 第三层的宽度最大，结点为 (5,3) ，宽度为 2
//      1
//     /
//    3
//   / \
//  5   3

// 输入： root = [1,3,2,5]
// 输出： 2
// 解释： 第三层的宽度最大，结点为 (3,2) ，宽度为 2
//      1
//     / \
//    3   2
//   /
//  5


// 思路： DFS
//
//      我们可以直接使用 dfs 进行处理，并同时维护一个数组 level_left_index ，
//      其中 level_left_index[i] 表示第 i 层的最左侧非空结点的索引。
//
//      在递归时优先进行如下判断：
//          1. root 结点为空：认为以该结点为最右侧结点时，宽度为 0 ，直接返回即可
//          2. root 结点非空：
//              (1) 如果该层还没有遍历过，则 root 必定是该层的最左侧结点，
//                  将其索引放入 level_left_index 中
//              (2) 然后递归获取左右子树中的最大宽度，并取最大值为 max_width
//              (3) 最后以 root 为当前层最右侧结点，计算宽度，
//                  返回它与 max_width 的最大值即可
//		
//
//		时间复杂度： O(n)
//          1. 需要遍历全部 O(n) 结点一次
//		空间复杂度： O(n)
//          1. 栈递归深度就是树高，最差情况下，全部 O(n) 个结点在一条链上
//          2. 需要记录每一层最左侧结点的索引，也就是树高，
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
    pub fn width_of_binary_tree(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        // dfs 遍历二叉树
        //  初始传入根结点 root ，
        //  其层数为 0 ，索引为 0 ，
        //  level_left_index 为空列表
        Self::dfs(&root, 0, 0, &mut vec![])
    }

    //  递归计算每一层的宽度，并返回所有层宽度的最大值
    //
    //  root:               当前子树根结点
    //  level:              当前子树根结点的层数
    //  index:              当前子树根结点的索引
    //  level_left_index:   维护二叉树每一层的最左侧结点的索引
    fn dfs(root: &Option<Rc<RefCell<TreeNode>>>, level: usize, index: i32, level_left_index: &mut Vec<i32>) -> i32 {
        // 如果 root 非空，则可以继续递归处理
        if let Some(root) = root {
            // 如果当前层还没有最左侧结点的索引，
            // 那么当前结点就是该层的最左侧结点，放入到 level_left_index 中
            if level >= level_left_index.len() {
                level_left_index.push(index);
            }

            // 先递归处理，获取左右子树中的最大宽度，并取最大值
            let mut max_width = i32::max(
                Self::dfs(&root.borrow().left, level + 1, index * 2, level_left_index),
                Self::dfs(&root.borrow().right, level + 1, index * 2 + 1, level_left_index),
            );

            // 然后再计算以当前结点为最右侧结点时的宽度，取最大值并返回
            max_width.max(index - level_left_index[level] + 1)
        } else {
            // 如果 root 为空，则认为以该结点为最右侧结点时，宽度为 0
            0
        }
    }
}
