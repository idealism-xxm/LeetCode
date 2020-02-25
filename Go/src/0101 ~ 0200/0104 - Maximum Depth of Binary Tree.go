// 链接：https://leetcode.com/problems/maximum-depth-of-binary-tree/
// 题意：计算树的高度？

// 输入：[3,9,20,null,null,15,7]
// 输出：3
// 解释：
//    3
//   / \
//  9  20
//    /  \
//   15   7

// 思路：递归
//
//		从根结点开始
//		1. 如果当前结点为空，则返回 0
//		2. 如果当前结点不为空，则返回 1 + max(左子树高度, 右子树高度)
//
//		时间复杂度： O(n) 空间复杂度： O(n)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
