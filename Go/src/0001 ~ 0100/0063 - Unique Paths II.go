// 链接：https://leetcode.com/problems/unique-paths-ii/
// 题意：给定一个 m * n 的格子，一个机器人在左上角
//		它每一次可以向右或者向下移动一格，其中有些格子不能进入
//		求移动到右下角有多少种路径？

// 输入：
// [
//   [0,0,0],
//   [0,1,0],
//   [0,0,0]
// ]
// 输出：2

// 思路：DP
//		dp[i][j] 表示运动到 (i, j) 是的路径数
//		初始状态： dp[1][1] = 1
//		状态转移方程：如果 (i, j) 没有障碍物，则： dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
//		时间复杂度： O(m * n) ，空间复杂度： O(m * n) 【可以用滚动数组优化为 O(n) 】

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m + 1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n + 1)
	}
	dp[1][1] = 1 - obstacleGrid[0][0]

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if obstacleGrid[i - 1][j - 1] != 1 && (i != 1 || j != 1) {
				dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
			}
		}
	}
	return dp[m][n]
}
