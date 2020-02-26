// 链接：https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
// 题意：通过中序遍历列表和后序遍历列表重建二叉树？

// 输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
// 输出：
//    3
//   / \
//  9  20
//    /  \
//   15   7

// 思路：递归
//		和 0105 一样的思路，只是每次从 postorder 拿最后一个结点在 inorder 中寻找
//
//		1. 若 postorder 为空，则返回 空结点
//		2. 若 postorder 不为空，则拿出 preorder 的最后一个结点 x 作为为当前子树的根结点
//			在 inorder 中找到 x 的下标为 i ，则
//			(1) 左子树由 inorder[:i] 和 postorder[:i] 进行构建
//			(2) 右子树由 inorder[i + 1:] 和 postorder[i:len(postorder) - 1] 进行构建
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	// 如果为空，则为空子树
	length := len(postorder)
	if length == 0 {
		return nil
	}

	// 在 inorder 中找到 postorder 最后一个结点的值
	for i, cur := range inorder {
		if cur == postorder[length - 1] {
			return &TreeNode {
				Val: cur,
				Left: buildTree(inorder[:i], postorder[:i]),
				Right: buildTree(inorder[i + 1:], postorder[i:length - 1]),
			}
		}
	}
	// 不会走到这
	return nil
}
