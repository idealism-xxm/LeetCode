// 链接：https://leetcode.com/problems/insert-into-a-binary-search-tree/
// 题意：给定一棵二叉搜索树，将一个新值 val 插入其中，并使其仍然是二叉搜索树。

// 数据限制：
//  树的节点个数在 [0, 10 ^ 4] 内
//  -(10 ^ 8) <= Node.val <= 10 ^ 8
//  所有的 Node.val 都是唯一的
//  -(10 ^ 8) <= val <= 10 ^ 8
//  val 不在原始的二叉搜索树中

// 输入： root = [4,2,7,1,3], val = 5
// 输出： [4,2,7,1,3,5]
// 解释： 5 直接作为根节点也是合法的

// 输入： root = [40,20,60,10,30,50,70], val = 25
// 输出： [40,20,60,10,30,50,70,null,null,25]

// 输入： root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
// 输出： [4,2,7,1,3,5]


// 思路： 递归
//
//      二叉搜索树的插入和查询类似，因为 val 不在二叉搜索树中，
//      所以本质就是找到 val 应该在的那个节点（必定是 None ），
//      然后设置其值为 val 即可。
//
//      可以直接使用递归处理，刚好可以兼容空二叉搜索树的情况：
//
//      1. 当前子树 root 不存在，则 val 直接作为子树，建立节点返回即可
//      2. 当前子树 root 存在，
//          (1) val < root.val: 将 val 插入 root 的左子树中，
//              递归执行插入，并将最后的左子树节点赋值给 root.left
//          (2) val > root.val: 将 val 插入 root 的右子树中，
//              递归执行插入，并将最后的右子树节点赋值给 root.right
//
//      时间复杂度： O(n)
//      空间复杂度： O(n)

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
    pub fn insert_into_bst(root: Option<Rc<RefCell<TreeNode>>>, val: i32) -> Option<Rc<RefCell<TreeNode>>> {
        // 最后返回一个 Option 类型的根节点
        Some(match root {
            // 如果根节点不存在，则 val 直接成为根节点
            None => Rc::new(RefCell::new(TreeNode::new(val))),
            // 如果根节点存在，则根据 val 和根节点的值递归到左右子树处理
            Some(cur) => {
                // 如果 val 更小，则递归执行插入，并将最后的左子树节点赋值给 root.left
                if val < cur.borrow().val {
                    let left = Solution::insert_into_bst(cur.borrow().left.clone(), val);
                    cur.borrow_mut().left = left;
                } else {
                    // 如果 val 更大，则递归执行插入，并将最后的右子树节点赋值给 root.right
                    let right = Self::insert_into_bst(cur.borrow().right.clone(), val);
                    cur.borrow_mut().right = right;
                }

                // 返回根节点本身
                cur
            }
        })
    }
}
