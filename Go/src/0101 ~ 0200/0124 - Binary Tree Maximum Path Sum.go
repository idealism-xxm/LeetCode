// 链接：https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/
// 题意：给定一颗二叉树，求树上一个结点到另一个结点到所有路径中，路径和最大的是多少？

// 输入： [1,2,3]
// 输出： 6
// 解释：
//       1
//      / \
//     2   3

// 输入： [-10,9,20,null,null,15,7]
// 输出： 42
// 解释：
//   -10
//   / \
//  9  20
//    /  \
//   15   7

// 思路： 树形 DP
//		看到这题第一反应还是 DP ，
//		首先还是要考虑为了得到题目的答案所需的数据是什么，
//		这题很容易就能想到：如果我们 dfs 的时候计算每颗子树从根结点开始的最大路径和
//		那么当 root 结点的左右子结点更新完毕后，以 root 开始的最大路径和也可直接更新，
//		同时：root 子树中包含当前结点的最大路径和也可以直接算出来
//		为了维护子树中的最大路径和（不一定包含 root 结点），
//		那么 dfs 的时候就需要返回两个值：
//			1. 每颗子树从根结点开始的最大路径和
//			2. 每颗子树的最大路径和（不一定包含根结点）
//
//		设：
//		dp[cur][0] 表示 cur 子树从 cur 开始的最大路径和
//		dp[cur][1] 表示 cur 子树中最大路径和（不一定包含根结点）
//
//		初始化： dp[leaf][0] = leaf.Val, dp[leaf][1] = leaf.Val （叶子结点只能包含本身）
//		状态转移：
//			对于非叶子结点 cur 有：
//			1. 左右子树取较大值，如果都为负则不取子结点
//			dp[cur][0] = cur.Val + max(0, dp[cur.Left][0], dp[cur.Right][0])
//			2. 左右子树取较大值，且再从包含根结点的最大路径和中取较大值
//			dp[cur][1] = max(dp[cur.Left][1], dp[cur.Right][1], cur.Val + max(0, dp[cur.Left][0]) + max(0, dp[cur.Right][0]))
//
//		由于 dfs 的时候可以直接返回两个值，所以不需要额外变量存储
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
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	_, result := dfs(root)
	return result
}

func dfs(root *TreeNode) (maxSumIncludeRoot, maxSum int) {
	// 空结点则直接返回 math.MinInt32 ，方便后续不必判断
	if root == nil {
		return math.MinInt32, math.MinInt32
	}
	// 叶子结点只能包含自身，则直接返回自身的值
	if root.Left == nil && root.Right == nil {
		return root.Val, root.Val
	}

	// 递归计算左右子结点对应的值
	lMaxSumIncludeRoot, lMaxSum := dfs(root.Left)
	rMaxSumIncludeRoot, rMaxSum := dfs(root.Right)

	// 左右子树取较大值，如果都为负则不取子结点
	maxSumIncludeRoot = root.Val + max(0, lMaxSumIncludeRoot, rMaxSumIncludeRoot)
	// 左右子树取较大值，且再从包含根结点的最大路径和中取较大值
	maxSum = max(lMaxSum, rMaxSum, root.Val + max(0, lMaxSumIncludeRoot) + max(0, rMaxSumIncludeRoot))
	return maxSumIncludeRoot, maxSum
}

func max(firstNum int, remainNums ...int) int {
	for _, num := range remainNums {
		if firstNum < num {
			firstNum = num
		}
	}
	return firstNum
}
