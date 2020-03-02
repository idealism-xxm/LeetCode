// 链接：https://leetcode.com/problems/path-sum/
// 题意：给定一个二叉树，判断是否存在一条根到叶子路径上所有值到和为指定到数？

// 输入： sum = 22
//       5
//      / \
//     4   8
//    /   / \
//   11  13  4
//  /  \      \
// 7    2      1
// 输出：true

// 思路：递归
//
//		1. 如果当前是叶子结点，则返回 root.Val == sum
//		2. 如果左子结点存在，递归处理左子结点，并将 sum 减去当前结点的值
//			(1) 递归处理结果为 true ，直接返回 true
//			(2) 递归处理结果为 false ，继续执行以下逻辑
//		3. 若果右子结点存在，递归处理右子结点，并将 sum 减去当前结点的值
//			(1) 递归处理结果为 true ，直接返回 true
//  		(2) 递归处理结果为 false ，继续执行以下逻辑
//		4. 左右子结点都不满足题意，返回 false
//
//		时间复杂度： O(n)
//		空间复杂度： O(n)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	// 如果是叶子结点，则仅当值等于 sum 时，才满足题意
	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}
	remainSum := sum - root.Val
	// 当左子结点存在，且存在一条左子结点到叶子路径上所有值到和为 remainSum ，则满足题意
	if root.Left != nil && hasPathSum(root.Left, remainSum) {
		return true
	}
	// 当右子结点存在，且存在一条右子结点到叶子路径上所有值到和为 remainSum ，则满足题意
	if root.Right != nil && hasPathSum(root.Right, remainSum) {
		return true
	}
	return false
}
