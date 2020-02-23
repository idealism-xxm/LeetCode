// 链接：https://leetcode.com/problems/same-tree/
// 题意：判断一个二叉树是否对称？

// 输入：[1,2,2,3,4,4,3]
// 输出：true
// 解释：
//     1
//    / \
//   2   2
//  / \ / \
// 3  4 4  3

// 输入：[1,2,2,null,3,null,3]
// 输出：false
// 解释：
//     1
//    / \
//   2   2
//    \   \
//    3    3

// 思路1：递归
//
//		和 0100 一样递归处理即可，即判断直接判断根结点都左右子树是否对称
//		1. 若当前结点的值相同，则前者左子树与后者右子树、前者右子树与后者左子树都对称才返回 true
//		2. 若当前结点都值不同，则直接返回 false
//
//		时间复杂度： O(n) 空间复杂度： O(n)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return dfs(root, root)
}

func dfs(p, q *TreeNode) bool {
	// 都是空结点，则相同
	if p == nil && q == nil {
		return true
	}
	// 值不一样就不同
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	// 递归处理左右子树
	return dfs(p.Left, q.Right) && dfs(p.Right, q.Left)
}

// 思路2：循环
//
//		思路同上，使用栈用循环模拟递归即可
//
//		时间复杂度： O(n)  空间复杂度： O(n)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	pStack, qStack := []*TreeNode{root}, []*TreeNode{root}
	// 若都还有结点未遍历，则继续处理
	for ; len(pStack) != 0 && len(qStack) != 0; {
		// 分别出栈
		lengthP, lengthQ := len(pStack), len(qStack)
		p, q := pStack[lengthP - 1], qStack[lengthQ - 1]
		pStack, qStack = pStack[:lengthP - 1], qStack[:lengthQ - 1]
		// 都是空结点，则相同，继续处理下一个结点
		if p == nil && q == nil {
			continue
		} else if p == nil || q == nil || p.Val != q.Val {
			// 值不一样就不同
			return false
		}

		// 将左右子树结点放入
		pStack = append(pStack, p.Left)
		qStack = append(qStack, q.Right)
		pStack = append(pStack, p.Right)
		qStack = append(qStack, q.Left)
	}
	// 若还存在结点未遍历，则必定不对称
	return len(pStack) == len(qStack)
}
