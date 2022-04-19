// 链接：https://leetcode.com/problems/recover-binary-search-tree/
// 题意：给定一棵二叉搜索树 root ，其中恰好有两个结点被互换了，
//      将这颗二叉搜索树复原。


// 数据限制：
//  树的结点数为 n
//  1 <= k <= n <= 10 ^ 4
//  0 <= Node.val <= 10 ^ 4


// 输入： root = [1,3,null,null,2]
// 输出： [3,1,null,null,2]
// 解释： 3 不能作为 1 的左子结点，因为 3 > 1 
//       1            3
//      /            /
//     3      →     1
//      \            \
//       2            2
//

// 输入： root = [3,1,4,null,null,2]
// 输出： [2,1,4,null,null,3]
// 解释： 2 不能在 3 的右子树中，因为 2 < 3
//       3           2
//      / \         / \
//     1   4   →   1   4
//        /           /
//       2           3


// 思路： 递归
//
//      一个二叉搜索树的中序遍历值必定严格递增，
//      所以我们只要进行中序遍历，然后比较相邻两个结点的值。
//
//      我们可以使用 dfs 递归中序遍历处理，并维护三个结点：
//          1. previous: 表示中序遍历的前一个结点
//          2. first:    表示互换结点的前者，该结点必定比后一个结点大
//          3. second:   表示互换结点的后者，前一个结点必定比该结点大
//
//      可以发现两个互换的结点，必定出现在中序遍历时大小不对的位置处，
//      所以在 dfs 中，如果前一个结点 previous 的值大于当前结点 root 的值，
//      则找到了一个互换的结点。
//
//      1. 如果这样的位置有 1 处，那么 first 必定是 previous ，
//          second 必定是 root 
//      2. 如果这样的位置有 2 处，那么 first 一定是第一处的 previous ，
//          second 一定是第二处的 root
//
//      综上： first 必定是第一处的 previous ，
//           second 必定是最后一处的 root
//
//
//      时间复杂度：O(n)
//          1. 需要遍历找到两个互换的结点，最差情况下，
//              最后一个结点被换了，需要遍历全部 O(n) 个结点
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
use std::mem::swap;
impl Solution {
    pub fn recover_tree(root: &mut Option<Rc<RefCell<TreeNode>>>) {
        // 定义 first 和 second ，用于维护交换的两个结点
        let (mut first, mut second) = (None, None);
        // 递归中序遍历找到互换的结点
        Self::dfs(root, &mut None, &mut first, &mut second);
        // 交换 first 和 second 两个结点的值
        swap(&mut first.as_mut().unwrap().borrow_mut().val, &mut second.as_mut().unwrap().borrow_mut().val);
    }

    fn dfs(root: &Option<Rc<RefCell<TreeNode>>>, previous: &mut Option<Rc<RefCell<TreeNode>>>, first: &mut Option<Rc<RefCell<TreeNode>>>, second: &mut Option<Rc<RefCell<TreeNode>>>) {
        // 对 root 子树进行递归中序遍历，找到两个互换的结点。
        // previous: 表示中序遍历的前一个结点。
        // first:    表示互换结点的前者，该结点必定比后一个结点大。
        // second:   表示互换结点的后者，前一个结点必定比该结点大。

        // 如果当前结点为空，则直接返回
        if root.is_none() {
            return;
        }

        // 先递归处理左子树
        Self::dfs(&root.as_ref().unwrap().borrow().left, previous, first, second);

        // 如果前一个结点的值大于当前结点，则找到了一个互换的结点
        if previous.is_some() && previous.as_ref().unwrap().borrow().val > root.as_ref().unwrap().borrow().val {
            if first.is_none() {
                // 如果第 1 个结点未找到，则设置第 1 个结点为前一个结点
                *first = previous.clone();
            }
            if first.is_some() {
                // 如果第 1 个结点已找到，则当前找到的是第 2 个结点，
                // 设置第 2 个结点为当前结点
                *second = root.clone();
            }
        }

        // 设置前一个结点为当前结点
        *previous = root.clone();
        // 继续递归处理右子树
        Self::dfs(&root.as_ref().unwrap().borrow().right, previous, first, second);
    }
}
