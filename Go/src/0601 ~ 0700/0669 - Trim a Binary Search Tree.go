// 链接：https://leetcode.com/problems/trim-a-binary-search-tree/
// 题意：给定一棵二叉搜索树 root 和两个整数 low, high ，
//      需要裁剪这棵二叉树，仅保留结点值在 [low, high] 间的结点，
//      返回裁剪后形成的新的二叉搜索树的根结点。


// 数据限制：
//  这棵树的结点数在 [1, 10 ^ 4] 内
//  0 <= Node.val <= 10 ^ 4
//  这棵树所有结点的值均不相同
//  root 是一棵合法的二叉搜索树
//  0 <= low <= high <= 10 ^ 4


// 输入： root = [1,0,2], low = 1, high = 2
// 输出： [1,null,2]
// 解释： 1             1
//      / \      →      \
//     0   2             2

// 输入： root = [3,0,4,null,2,null,null,1], low = 1, high = 3
// 输出： [3,2,null,1]
// 解释： 3             3
//      / \           /
//     0   4         2
//      \       →   /
//       2         1
//      /
//     1


// 思路： 递归
//
//      如果当前子树的根结点 root 为空，则直接返回 null ，
//      否则根据 root.val 与 low 和 high 的关系递归处理：
//          1. root.val < low: 则该结点及其左子树都小于 low ，
//              都需要被裁减掉，那么裁剪结果必定是裁剪右子树的结果，
//              直接返回递归处理右子树的结果即可；
//          2. root.val > high: 则该结点及其右子树都大于 high ，
//              都需要被裁减掉，那么裁剪结果必定是裁剪左子树的结果，
//              直接返回递归处理左子树的结果即可；
//          3. low <= root.val <= high: 当前结点需要保留，
//              递归裁剪左右子树，然后返回当前根结点
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
func trimBST(root *TreeNode, low int, high int) *TreeNode {
    // 如果 root 为空，则直接返回 nil
    if root == nil {
        return nil
    }
    // 如果 root 非空，则继续裁剪
    if root.Val < low {
        // 如果当前结点的值小于 low ，
        // 则该结点及其左子树都小于 low ，都需要被裁减掉，
        // 那么裁剪结果必定是裁剪右子树的结果
        return trimBST(root.Right, low, high)
    }
    if root.Val > high {
        // 如果当前结点的值大于 high ，
        // 则该结点及其右子树都大于 high ，都需要被裁减掉，
        // 那么裁剪结果必定是裁剪左子树的结果
        return trimBST(root.Left, low, high)
    }

    // 此时当前结点需要保留，递归裁剪左右子树即可
    root.Left = trimBST(root.Left, low, high)
    root.Right = trimBST(root.Right, low, high)
    // 返回当前结点
    return root
}
