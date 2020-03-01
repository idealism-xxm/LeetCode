// 链接：https://leetcode.com/problems/minimum-depth-of-binary-tree/
// 题意：给定一个二叉树，计算最小深度？

// 输入：[3,9,20,null,null,15,7]
// 输出：2
// 解释：
//    3
//   / \
//  9  20
//    /  \
//   15   7

// 思路1：递归
//
//		递归统计左右子树的最小深度为 lMinDepth 和 rMinDepth
//		1. 左右子树均非空，则当前子树的最小深度为 min(lMinDepth, rMinDepth) + 1
//		2. 左右子树存在空，则当前子树的最小深度为 lMinDepth + rMinDepth + 1
//
//		当然还可以用 BFS 进行搜索，找到第一个叶子结点即可结束，不用遍历全部结点
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
func minDepth(root *TreeNode) int {
	// 空结点深度为 0
	if root == nil {
		return 0
	}
	// 计算左右子树最小深度
	lMinDepth, rMinDepth := minDepth(root.Left), minDepth(root.Right)
	// 若左右子树均非空，则当前子树的最小深度等于左右子树最小深度的较小值 + 1
	if lMinDepth != 0 && rMinDepth != 0 {
		return min(lMinDepth, rMinDepth) + 1
	}

	// 由于空结点深度为 0 ，所以可以合并剩余所有情况
	return lMinDepth + rMinDepth + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 思路2：BFS
//
//		用 BFS 进行搜索，找到第一个叶子结点即可结束，不用遍历全部结点
//
//		时间复杂度： O(n)
//		空间复杂度： O(n)

type Node struct {
	Current *TreeNode
	Depth int
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	// 空结点深度为 0
	if root == nil {
		return 0
	}

	// bfs 找到第一个叶子结点
	queue := []*Node {
		{
			Current: root,
			Depth: 1,
		},
	}
	head := 0
	for ; head < len(queue); {
		// 队首元素出队
		cur := queue[head]
		head++
		if cur.Current.Left == nil && cur.Current.Right == nil {
			return cur.Depth
		}
		// 左右子结点依次入队
		if cur.Current.Left != nil {
			queue = append(queue, &Node {
				Current: cur.Current.Left,
				Depth: cur.Depth + 1,
			})
		}
		if cur.Current.Right != nil {
			queue = append(queue, &Node {
				Current: cur.Current.Right,
				Depth: cur.Depth + 1,
			})
		}
	}
	// 永远不会走这
	return 0
}
