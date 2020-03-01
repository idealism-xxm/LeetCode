// 链接：https://leetcode.com/problems/balanced-binary-tree/
// 题意：给定一个二叉树，判断其是否高度平衡（每个结点的左右子树高度差不大于 1 ）？

// 输入：[3,9,20,null,null,15,7]
// 输出：true
// 输出：
//    3
//   / \
//  9  20
//    /  \
//   15   7

// 输入：[1,2,2,3,3,null,null,4,4]
// 输出：false
// 输出：
//       1
//      / \
//     2   2
//    / \
//   3   3
//  / \
// 4   4

// 思路：递归
//
//		递归统计左右子树高度，并判断左右子树是否为高度平衡
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
func isBalanced(root *TreeNode) bool {
	return judge(root) != -1
}

// 判断 root 子数是否高度平衡
//	若高度平衡：返回树的高度
//	若高度不平衡：返回 -1
func judge(root *TreeNode) (height int) {
	// 空结点，则必定高度平衡
	if root == nil {
		return 0
	}

	// 若左子数不高度平衡，则直接返回 -1
	leftHeight := judge(root.Left)
	if leftHeight == -1 {
		return -1
	}
	// 若右子数不高度平衡，则直接返回 -1
	rightHeight := judge(root.Right)
	if rightHeight == -1 {
		return -1
	}

	// 保证左边高度更大，方便后续处理
	if leftHeight < rightHeight {
		leftHeight, rightHeight = rightHeight, leftHeight
	}
	// 若当前子树高度平衡，则返回当前子树的高度
	if leftHeight - rightHeight <= 1 {
		return leftHeight + 1
	}
	// 否则返回 -1
	return -1
}
