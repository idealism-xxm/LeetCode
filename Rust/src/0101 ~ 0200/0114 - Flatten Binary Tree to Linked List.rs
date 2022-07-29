// 链接：https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
// 题意：给定一个二叉树，将其转换成链表形式？

// 输入： root = [1,2,5,3,4,null,6]
// 输出： [1,null,2,null,3,null,4,null,5,null,6]
// 解释：
//     1             1
//    / \             \
//   2   5             2
//  / \   \      →      \
// 3   4   6             3
//                        \
//                         4
//                          \
//                           5
//                            \
//                             6

// 输入： root = []
// 输出： []

// 输入： root = [0]
// 输出： [0]


// 思路1：递归
//
//		对于子树 root 来说，可以递归调用处理成三部分
//			1. 当前根结点
//			2. 左子树形成的链表（可能为空链表）
//			3. 右子树形成的链表（可能为空链表）
//
//		然后将三部分按顺序连接起来即可。
//
//      最后返回链表的头结点和尾结点，方便上层处理。
//
//
//		时间复杂度： O(n)
//          1. 需要遍历全部 O(n) 个结点一次
//		空间复杂度： O(n)
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
    pub fn flatten(root: &mut Option<Rc<RefCell<TreeNode>>>) {
        Self::dfs(root.clone());
    }

    fn dfs(mut root: Option<Rc<RefCell<TreeNode>>>) -> (Option<Rc<RefCell<TreeNode>>>, Option<Rc<RefCell<TreeNode>>>) {
        if root.is_none() {
            return (None, None);
        }
        // 递归处理左右结点，并获取对应链表的头结点和尾结点
        let (left_head, left_tail) = Self::dfs(root.as_mut().unwrap().borrow_mut().left.take());
        let (right_head, right_tail) = Self::dfs(root.as_mut().unwrap().borrow_mut().right.take());

        // 当前结点目前既是头结点，也是尾结点
        let mut head = root.clone();
        let mut tail = root;
        // 将左半部分挂在链表尾部
        if left_head.is_some() {
            tail.as_mut().unwrap().borrow_mut().right = left_head;
            tail = left_tail;
        }
        // 将右半部分挂在链表尾部
        if right_head.is_some() {
            tail.as_mut().unwrap().borrow_mut().right = right_head;
            tail = right_tail;
        }

        // 返回当前子树转换成的链表头结点和尾结点
        (head, tail)
    }
}
