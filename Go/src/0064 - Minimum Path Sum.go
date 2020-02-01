// 链接：https://leetcode.com/problems/minimum-path-sum/
// 题意：给定一个 m * n 的格子，一个机器人在左上角
//		它每一次可以向右或者向下移动一格，其中每个格子又一个非负整数
//		求移动到右下角的所有路径中，数字和最小路径的和？

// 输入：
// [
//   [0,0,0],
//   [0,1,0],
//   [0,0,0]
// ]
// 输出：2

// 思路：DP
//		dp[i][j] 表示运动到 (i, j) 的路径中，数字和最小路径的和
//		状态转移方程：dp[i][j] = min(dp[i - 1][j], dp[i][j - 1]) + grid[i][j]
//		时间复杂度： O(m * n) ，空间复杂度： O(1)

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i - 1][0]
	}
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j - 1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = min(grid[i - 1][j], grid[i][j - 1]) + grid[i][j]
		}
	}
	return grid[m - 1][n - 1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
