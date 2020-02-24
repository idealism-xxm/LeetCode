// 链接：https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
// 题意：对一个二叉树进行锯齿宽度遍历（一层从左向右，一层从右向左交替），每一层一个数组？

// 输入：[3,9,20,null,null,15,7]
// 输出：
// [
//   [3],
//   [20,9],
//   [15,7]
// ]
// 解释：
//    3
//   / \
//  9  20
//    /  \
//   15   7

// 思路：模拟
//
//		BFS 即可，直接复用 0102 的代码
//		提出一个方法用于遍历每一层的结点，然后返回本层结点对应的值和下一层的结点
//		每次将 values 根据当前记录的 reverse 的值处理
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var result [][]int
	reverse := false
	for nodes := []*TreeNode{root}; len(nodes) != 0; {
		values, nextNodes := iter(nodes)
		// 如果需要翻转，则先翻转
		if reverse {
			for i, j := 0, len(values) - 1; i < j; i, j = i + 1, j - 1 {
				values[i], values[j] = values[j], values[i]
			}
		}
		result = append(result, values)

		nodes = nextNodes
		reverse = !reverse
	}
	return result
}

func iter(nodes []*TreeNode) (values []int, nextNodes []*TreeNode) {
	for _, node := range nodes {
		values = append(values, node.Val)
		if node.Left != nil {
			nextNodes = append(nextNodes, node.Left)
		}
		if node.Right != nil {
			nextNodes = append(nextNodes, node.Right)
		}
	}
	return values, nextNodes
}
