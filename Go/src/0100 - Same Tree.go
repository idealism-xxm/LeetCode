// 链接：https://leetcode.com/problems/same-tree/
// 题意：判断两个二叉树是否一样？

// 输入：[1,2,3],   [1,2,3]
// 输出：true

// 输入：[1,2],     [1,null,2]
// 输出：false

// 输入：[1,2,1],   [1,1,2]
// 输出：false

// 思路：递归
//
//		递归处理即可
//		1. 若当前结点的值相同，则左子树和右子树都相同才返回 true
//		2. 若当前结点都值不同，则直接返回 false
//
//		时间复杂度： O(n)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 都是空结点，则相同
	if p == nil && q == nil {
		return true
	}
	// 值不一样就不同
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	// 递归处理左右子树
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
