// 链接：https://leetcode.com/problems/increasing-order-search-tree/
// 题意：给定一棵二叉搜索树 root ，将其转换成最左侧结点为根结点，
//      且所有结点无左子结点的树，最后再返回结果树的根节点。


// 数据限制：
//  这棵树的结点数在 [0, 100] 内
//  0 <= Node.val <= 1000


// 输入： root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
// 输出： [1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
// 解释：       5                1
//          /     \              \
//         3       6              2
//        / \       \       →      \
//       2   4       8              3
//      /           / \              \
//     1           7   9              4
//                                     \
//                                      5
//                                       \
//                                        6
//                                         \
//                                          7
//                                           \
//                                            8
//                                             \
//                                              9

// 输入： root = [5,1,7]
// 输出： [1,null,5,null,7]
// 解释： 5            1
//      / \     →      \
//     1   7            5
//                       \
//                        7


// 思路： 递归
//
//      二叉搜索树是左子树的值小于根节点，右子树的值大于根节点。
//
//      所以可以通过递归的方式来中序遍历这棵树，
//      为了方便处理，我们维护一个 right_most 结点，
//      表示结果树中最右侧的结点，
//      初始化为 dummy 结点，这样就不需要特殊处理了。
//
//      使用 dfs(root, &mut right_most) 来处理当前结点，
//      如果 root 为空，则直接返回。
//
//      否则，先递归处理左子树 dfs(root.left, right_most) ，
//      然后，将当前结点插入到 right_most 的右侧，
//      即 right_most.right = root 。
//
//      同时我们要更新 right_most 为当前结点，
//      即 *right_most = root ，
//      最后再递归处理右子树 dfs(root.right, right_most) 。  
//
//
//      时间复杂度：O(n)
//          1. 需要遍历二叉搜索树中全部 O(n) 个结点
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
    pub fn increasing_bst(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        // 使用一个 dummy 结点，方便后续处理
        let mut dummy = Rc::new(RefCell::new(TreeNode::new(0)));
        // 递归处理二叉树 root
        Self::dfs(root, &mut Some(dummy.clone()));
        // dummy 的右子结点就是结果树的根节点
        let ans = dummy.borrow_mut().right.take();
        ans
    }

    pub fn dfs(mut root: Option<Rc<RefCell<TreeNode>>>, right_most: &mut Option<Rc<RefCell<TreeNode>>>) {
        if root.is_some() {
            // 先递归处理左子树
            Self::dfs(root.as_mut().unwrap().borrow_mut().left.take(), right_most);
            // 提前获取右子树，方便后续处理
            let right = root.as_mut().unwrap().borrow_mut().right.take();
            // 克隆 root ，方便后续处理
            let root_clone = root.clone();
            // 然后把 root 挂在 right_most 的右子结点上
            right_most.as_mut().unwrap().borrow_mut().right = root;
            // 移动 right_most 到其右子结点上（也就是 root 结点）
            *right_most = root_clone;
            // 最后递归处理右子树
            Self::dfs(right, right_most);
        }
    }
}
