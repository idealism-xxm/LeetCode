// 链接：https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
// 题意：给定一个排好序的数组，将其转换成一颗二叉搜索树，并保证每个结点的左右子树高度差不大于 1 ？

// 输入：[-10,-3,0,5,9],
// 输出：
//      0
//     / \
//   -3   9
//   /   /
// -10  5

// 思路：递归
//
//		如果当前数组长度为 length ，则让中间的数作为当前子树的根结点，
//		该数左边的数递归形成左子树，该数右边的数递归形成右子树
//
//		时间复杂度： O(n)
//		空间复杂度： O(n)  实际只需要 O(logn) 的额外空间

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}

	mid := length >> 1
	return &TreeNode {
		Val: nums[mid],
		Left: sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid + 1:]),
	}
}
