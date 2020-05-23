// 链接：https://leetcode.com/problems/binary-tree-upside-down/
// 题意：给定一颗二叉树，保证所有右节点要么为空，要么不为空且有兄弟节点，
//		将它上下翻转成一个另一颗二叉树，原来的右节点转换成左节点，返回新的根？

// 输入： [1,2,3,4,5]
//        1
//       / \
//      2   3
//     / \
//    4   5
// 输出： [4,5,2,#,#,3,1]
//      4
//     / \
//    5   2
//       / \
//      3   1


// 思路： 模拟
//
//		按照题意模拟即可
//		1. 当前节点有没有左子节点，则不需要翻转，直接返回即可
//			（且它必定是新树的根）
//		2. 当前节点有左子节点，则需要翻转，
//			先断绝父子关系，递归处理左子树，右子树保持原样，
//			然后当前节点变为左子节点的右子节点，
//			右子节点变为左子节点的左子节点
//
//		时间复杂度： O(n)
//		空间复杂度： O(1)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	return dfs(root)
}

func dfs(root *TreeNode) *TreeNode {
	// 左子节点没有，则 root 为叶节点，直接返回即可
	if root.Left == nil {
		return root
	}

	// 保存左右子节点，并断绝父子关系
	left, right := root.Left, root.Right
	root.Left, root.Right = nil, nil

	// 将左子树翻转，并获得新左子树的根
	leftRoot := upsideDownBinaryTree(left)

	// 当前节点变为左子节点的右子节点，右子节点变为左子节点的左子节点
	left.Left, left.Right = right, root

	return leftRoot
}
