// 链接：https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// 题意：通过先序遍历列表和中序遍历列表重建二叉树？

// 输入：preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// 输出：
//    3
//   / \
//  9  20
//    /  \
//   15   7

// 思路：递归
//
//		1. 若 preorder 为空，则返回 空结点
//		2. 若 preorder 不为空，则拿出 preorder 的第一个结点 x 作为为当前子树的根结点
//			在 inorder 中找到 x 的下标为 i ，则
//			(1) 左子树由 preorder[1:i + 1] 和 inorder[:i] 进行构建
//			(2) 右子树由 preorder[i + 1:] 和 inorder[i + 1:] 进行构建
//
//		时间复杂度： O(n^2) 【可以使用 map 优化为 O(n)】
//		空间复杂度： O(n)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 如果为空，则为空子树
	if len(preorder) == 0 {
		return nil
	}

	// 在 inorder 中找到 preorder 第一个结点的值
	for i, cur := range inorder {
		if cur == preorder[0] {
			return &TreeNode {
				Val: cur,
				Left: buildTree(preorder[1:i + 1], inorder[:i]),
				Right: buildTree(preorder[i + 1:], inorder[i + 1:]),
			}
		}
	}
	// 不会走到这
	return nil
}
