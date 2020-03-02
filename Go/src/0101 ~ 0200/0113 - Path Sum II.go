// 链接：https://leetcode.com/problems/path-sum-ii/
// 题意：给定一个二叉树，找到所有根到叶子路径上所有值到和为指定到数的路径？

// 输入： sum = 22
//       5
//      / \
//     4   8
//    /   / \
//   11  13  4
//  /  \      \
// 7    2      1
// 输出：
// [
//    [5,4,11,2],
//    [5,8,4,5]
// ]

// 思路：递归
//
//		思路和 0113 一样，不过需要记录根结点到当前结点的所有值和总和，最后在叶子结点判断返回
//		1. 如果当前是叶子结点
//			(1) root.Val == sum, 则当前列表中的值为一个合法的答案
//			(2) 当前路径不满足题意，返回 nil
//		2. 如果左子结点存在，将左子结点产生的答案放入结果列表中
//		3. 若果右子结点存在，将右子结点产生的答案放入结果列表中
//		4. 返回本结点收集的结果列表
//

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	return dfs(root, sum, nil)
}

func dfs(root *TreeNode, sum int, list []int) [][]int {
	// 如果是叶子结点，则仅当值等于 sum 时，才满足题意，当前列表中的值为一个合法的答案
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			return [][]int{append(list[:0:0], append(list, root.Val)...)}
		}
		return nil
	}
	remainSum := sum - root.Val
	nextList := append(list, root.Val)
	var result [][]int
	// 当左子结点存在，将左子结点产生的答案放入结果列表中
	if root.Left != nil {
		result = append(result, dfs(root.Left, remainSum, nextList)...)
	}
	// 当右子结点存在，将右子结点产生的答案放入结果列表中
	if root.Right != nil {
		result = append(result, dfs(root.Right, remainSum, nextList)...)
	}
	return result
}
