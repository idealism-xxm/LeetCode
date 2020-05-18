// 链接：https://leetcode.com/problems/binary-tree-right-side-view/
// 题意：给定一颗二叉树，返回其从右侧看时看到的节点值形成的数组？

// 输入： [1,2,3,null,5,null,4]
// 输出： [1, 3, 4]
// 解释：
//       1            <---
//     /   \
//    2     3         <---
//     \     \
//      5     4       <---

// 思路： DFS
//
//		直接 DFS 即可，访问顺序为 左右中，
//		每一层维护其该放入的数组下标值，在叶子节点判断数组长度是否够用，
//		不够用则新建一个数组，抛弃原有的值
//		（因为右边后遍历，所以如果更长的话，则先前数组中的值都会被盖住）
//
//		时间复杂度： O(n)
//		空间复杂度： O(h)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	return dfs(root, nil, 0)
}

func dfs(root *TreeNode, result []int, index int) []int {
	if root.Left == nil && root.Right == nil {
		// 叶子节点，如果其更深，则新建一个数组收集结果，抛弃原收集的内容
		// （因为右边后遍历，所以如果更长的话，则先前数组中的值都会被盖住）
		if len(result) <= index {
			result = make([]int, index + 1)
		}
		result[index] = root.Val
		return result
	}

	// 收集左子树的值
	if root.Left != nil {
		result = dfs(root.Left, result, index + 1)
	}
	// 收集右子树的值
	if root.Right != nil {
		result = dfs(root.Right, result, index + 1)
	}
	// 收集当前节点的值
	result[index] = root.Val
	return result
}
