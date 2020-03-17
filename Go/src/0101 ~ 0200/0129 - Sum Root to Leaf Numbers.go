// 链接：https://leetcode.com/problems/sum-root-to-leaf-numbers/
// 题意：给定一颗二叉树，每个结点有 0 ～ 9 的数字，现将所有根到叶子结点上的路径转换成数字，
//		求所有路径代表的数字的和？

// 输入： [1,2,3]
//    1
//   / \
//  2   3
// 输出： 25
// 解释：
// 		路径 1 -> 2 代表数字 12
//		路径 1 -> 3 代表数字 13
//		所有路径代表的数字的和 = 12 + 13 = 25

// 输入： [4,9,0,5,1]
//     4
//    / \
//   9   0
//  / \
// 5   1
// 输出： 1026
// 解释：
// 		路径 4 -> 9 -> 5 代表数字 495
//		路径 4 -> 9 -> 1 代表数字 491
//		路径 4 -> 0 代表数字 40
//		所有路径代表的数字的和 = 495 + 491 + 40 = 1026

// 思路： 递归
//
//		从根结点开始，记录根结点到当前结点的所代表的数字，
//		1. 如果当前结点是叶子结点，则将答案返回
//		2. 如果当前结点非叶子结点，则递归收集所有子结点的和，然后返回
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
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root, root.Val)
}

func dfs(root *TreeNode, num int) int {
	// 如果是叶子结点，则直接返回
	if root.Left == nil && root.Right == nil {
		return num
	}

	// 记录本结点子结点产生的结果的和
	result := 0
	// 先乘以 10
	num *= 10
	if root.Left != nil {
		result += dfs(root.Left, num + root.Left.Val)
	}
	if root.Right != nil {
		result += dfs(root.Right, num + root.Right.Val)
	}
	return result
}
