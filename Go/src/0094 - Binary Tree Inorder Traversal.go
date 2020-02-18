// 链接：https://leetcode.com/problems/binary-tree-inorder-traversal/
// 题意：给定一个二叉树，返回中序遍历的结果？

// 输入：[1,null,2,3]
// 输出：[1,3,2]

// 思路：模拟
//
//		递归很容易就能写出来，转换成循环就需要用栈记录信息
//		若当前结点不为 nil ，则入栈，然后处理左子结点
//		若当前结点为 nil ，则出栈，记录出栈结点的值，并将当前结点指向其右子结点

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	// stack 存储左子结点都已处理的结点
	var stack []*TreeNode
	cur := root
	for ; cur != nil || len(stack) != 0; {
		// 若当前结点存在，则把当前结点放入，处理左子结点
		for ; cur != nil; {
			stack = append(stack, cur)
			cur = cur.Left
		}

		// 栈顶元素出栈
		length := len(stack)
		cur = stack[length - 1]
		stack = stack[:length - 1]

		// 左子树已处理完，放入当前结点值
		result = append(result, cur.Val)
		// 接下来处理右子树
		cur = cur.Right
	}
	return result
}
