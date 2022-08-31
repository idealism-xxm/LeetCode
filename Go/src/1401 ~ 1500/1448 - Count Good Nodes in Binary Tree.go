// 链接：https://leetcode.com/problems/count-good-nodes-in-binary-tree/
// 题意：给定一棵二叉树 root ，求有多少个结点 x 满足以下条件？
//      从根结点 root 到结点 x 的路径上所有结点的值都不大于 x.val 。


// 数据限制：
//  二叉树的结点数范围为 [1, 10 ^ 5]
//  每个结点值的范围为 [-(10 ^ 4), 10 ^ 4]


// 输入： root = [3,1,4,3,null,1,5]
// 输出： 4
// 解释： 以下 4 个带括号的结点满足题意。
//         (3)
//         / \
//        1  (4)
//       /   / \
//     (3)  1  (5)

// 输入： root = [3,3,null,4,2]
// 输出： 3
// 解释： 以下 3 个带括号的结点满足题意。
//         (3)
//         /
//       (3)
//       / \
//     (4)  2


// 思路： 递归/DFS
//
//      我们在递归遍历的时候，维护根结点到结点 x 路径上的最大值 maxVal 。
//
//      只要 x.val >= maxVal ，那么结点 x 就满足题意，计入 ans 中。
//
//      然后用 x.val 更新 maxVal 的最大值，再递归处理即可。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历二叉树中全部 O(n) 个结点
//      空间复杂度：O(h)
//          1. 栈递归深度就是树高 O(h) ，最差情况下，全部 O(n) 个结点在一条链上


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    return dfs(root, math.MinInt32)
}

// dfs(root, maxVal) 递归返回子树 root 中满足以下条件的结点 x 的数量：
//  1. root 到 x 的路径上所有结点的值都不大于 maxVal
//  2. root 到 x 的路径上所有结点的值都不大于 x.val
func dfs(root *TreeNode, maxVal int) int {
    // 如果当前结点为空，则直接返回 0
    if root == nil {
        return 0
    }

    // ans 表示 root 是否为满足题意的结点
    ans := 0
    // 如果实际上 root 满足题意，设置 ans 为 1 ，
    // 同时更新 maxVal 的最大值为 root.val 。
    if root.Val >= maxVal {
        ans = 1
        maxVal = root.Val
    }

    // 递归处理左右子树，返回其结果之和
    return ans + dfs(root.Left, maxVal) + dfs(root.Right, maxVal)
}
