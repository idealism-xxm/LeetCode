// 链接：https://leetcode.com/problems/triangle/
// 题意：给定一个数字三角形，求顶到底部到路径上数字和最小到路径到和 （使用 O(numRows) 的额外空间）？

// 输入：
// [
//      [2],
//     [3,4],
//    [6,5,7],
//   [4,1,8,3]
// ]
// 输出： 11

// 思路：DP
//
//		应该是刚入门 DP 时做得题目
//		从下到上更新最后不用再循环一遍，方便一点
//		初始化：多放入一行空的，方便迁移至其他语言
//			dp = make([]int, numRows + 1)
//		状态转移：第 i 层第 j 列可从 第 i + 1 层第 j 列 和第 i + 1 层第 j + 1 列 中较小值得到
//			dp[j] = min(dp[j], dp[j + 1]) + triangle[i][j]
//
//		时间复杂度： O(numRows ^ 2)
//		空间复杂度： O(numRows)

func minimumTotal(triangle [][]int) int {
	numRows := len(triangle)
	dp := make([]int, numRows + 1)
	for i := numRows - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			// 上一行算出来的较小值算入本行的路径中
			dp[j] = min(dp[j], dp[j + 1]) + triangle[i][j]
		}
	}
	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
