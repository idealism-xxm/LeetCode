// 链接：https://leetcode.com/problems/climbing-stairs/
// 题意：有一个 n 级楼梯，一个人每次可以上一级或者两级，
//		求有多少种方法可以到顶部？

// 输入：2
// 输出：2

// 输入：3
// 输出：3

// 思路：DP
//		初始化： dp[0] = dp[1] = 1
//		状态转移方程：dp[i] = dp[i - 1] + dp[i - 2]
//		时间复杂度： O(n) ，空间复杂度： O(n) 【使用模运算可以优化为 O(1)】

func climbStairs(n int) int {
	dp := make([]int, n + 1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i - 1] + dp[i - 2]
	}
	return dp[n]
}
