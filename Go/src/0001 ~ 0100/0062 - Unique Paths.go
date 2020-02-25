// 链接：https://leetcode.com/problems/unique-paths/
// 题意：给定一个 m * n 的格子，一个机器人在左上角
//		它每一次可以向右或者向下移动一格，求移动到右下角有多少种路径？

// 输入：m = 3, n = 2
// 输出：3

// 输入：m = 7, n = 3
// 输出：28

// 思路：DP
//		dp[i][j] 表示运动到 (i, j) 是的路径数
//		初始状态： dp[1][1] = 1
//		状态转移方程： dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
//		时间复杂度： O(m * n) ，空间复杂度： O(m * n) 【可以用滚动数组优化为 O(n) 】

func uniquePaths(m int, n int) int {
	dp := make([][]int, m + 1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n + 1)
	}
	dp[1][1] = 1

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i != 1 || j != 1 {
				dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
			}
		}
	}
	return dp[m][n]
}
