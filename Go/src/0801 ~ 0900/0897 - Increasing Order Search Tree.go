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
//      为了方便处理，我们维护一个 rightMost 结点，
//      表示结果树中最右侧的结点，
//      初始化为 dummy 结点，这样就不需要特殊处理了。
//
//      使用 dfs(root) 闭包来处理当前结点，
//      同时该闭包能引用外部变量 rightMost ，
//      如果 root 为空，则直接返回。
//
//      否则，先递归处理左子树 dfs(root.left) ，
//      然后，将当前结点插入到 rightMost 的右侧，
//      即 rightMost.right = root 。
//
//      同时我们要更新 rightMost 为当前结点，
//      即 rightMost = root ，
//      最后再递归处理右子树 dfs(root.right) 。  
//
//
//      时间复杂度：O(n)
//          1. 需要遍历二叉搜索树中全部 O(n) 个结点
//      空间复杂度：O(n)
//          1. 栈递归深度就是树高，最差情况下，全部 O(n) 个结点在一条链上


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
    // 使用一个 dummy 结点，方便后续处理
    dummy := &TreeNode{Val: 0}
    // 定义 dfs 中序遍历函数
    rightMost := dummy
    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root != nil {
            // 先递归处理左子树
            dfs(root.Left)
            root.Left = nil
            // 然后把 root 挂在 rightMost 的右子结点上
            rightMost.Right = root
            // 移动 rightMost 到其右子结点上（也就是 root 结点）
            rightMost = rightMost.Right
            // 最后递归处理右子树
            dfs(root.Right)
        }
    }

    // 递归处理二叉树 root
    dfs(root)
    // dummy 的右子结点就是结果树的根节点
    return dummy.Right
}
