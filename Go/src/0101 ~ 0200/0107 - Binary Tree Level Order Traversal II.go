// 链接：https://leetcode.com/problems/binary-tree-level-order-traversal-ii/
// 题意：对一个二叉树进行宽度遍历，每一层一个数组，从最低一层到最高一层放置？

// 输入：[3,9,20,null,null,15,7]
// 输出：
// [
//   [15,7]
//   [9,20],
//   [3],
// ]
// 解释：
//    3
//   / \
//  9  20
//    /  \
//   15   7

// 思路：模拟
//
//		BFS 即可，同 0102 一样，直接提出一个方法用于遍历每一层的结点
//		然后返回本层结点对应的值和下一层的结点，最后返回前，翻转整个数组
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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var result [][]int
	for nodes := []*TreeNode{root}; len(nodes) != 0; {
		values, nextNodes := iter(nodes)
		result = append(result, values)

		nodes = nextNodes
	}
	for i, j := 0, len(result) - 1; i < j; i, j = i + 1, j - 1 {
		result[i], result[j] = result[j], result[i]
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
