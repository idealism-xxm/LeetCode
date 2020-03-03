// 链接：https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
// 题意：给定一个二叉树，将其转换成链表形式？

// 输入：
//     1
//    / \
//   2   5
//  / \   \
// 3   4   6
// 输出：
// 1
//  \
//   2
//    \
//     3
//      \
//       4
//        \
//         5
//          \
//           6

// 思路1：递归
//
//		对于子树 root 来说，可以递归调用处理成三部分
//		1. 当前根结点
//		2. 左子树形成的链表（可能为空链表）
//		3. 右子树形成的链表（可能为空链表）
//		然后将三部分按顺序连接起来即可
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
func flatten(root *TreeNode) {
	dfs(root)
}

func dfs(root *TreeNode) (head, tail *TreeNode) {
	if root == nil {
		return nil, nil
	}
	// 递归处理左右结点
	lHead, lTail := dfs(root.Left)
	rHead, rTail := dfs(root.Right)
	// 清空左子结点
	root.Left = nil

	// 当前结点是目前既是头结点，也是尾结点
	head, tail = root, root
	// 将左半部分挂在链表尾
	if lHead != nil {
		tail.Right = lHead
		tail = lTail
	}
	// 将右半部分挂在链表尾
	if rHead != nil {
		tail.Right = rHead
		tail = rTail
	}
	// 返回当前子树转换成的链表表头和表尾
	return head, tail
}

// 思路2：Morris
//
//		看见题解说可以使用 Morris 算法
//		就想到 Morris 算法不仅可以将二叉树转换成链表，还只需要 O(1) 的额外空间
//		只需要在 Morris 的算法上改变一下，并不将其还原即可
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
func flatten(root *TreeNode) {
	for ; root != nil; {
		// 如果左子树不为空，则将其挂在右子结点，原有的右子结点变为左子树树上最右端的结点的右子结点
		if root.Left != nil {
			// 找到左子树上最右端的结点
			tail := root.Left
			for ; tail.Right != nil; {
				tail = tail.Right
			}
			// 左子树挂在右子结点
			tail.Right = root.Right
			root.Right = root.Left
			root.Left = nil
		}
		// 进入右子结点继续处理
		root = root.Right
	}
}
