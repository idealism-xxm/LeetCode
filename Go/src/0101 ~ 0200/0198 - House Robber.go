// 链接：https://leetcode.com/problems/house-robber/
// 题意：给定一个数组，不能选择相邻的两个数，求选择某些数的和的最大值？

// 输入： [1,2,3,1]
// 输出： 4
// 解释： 选择第 1 个和第 3 个， 1 + 3 = 4

// 输入： [2,7,9,3,1]
// 输出： 12
// 解释： 选择第 1 个、第 3 个和第 5 个， 2 + 9 + 1 = 12

// 思路： DP
//
//		很容易就能想到 DP
//		设 dp[i] 表示在前 i 个数中选择，
//		并且最后一个选择的数是第 i 个时能选择的数的和的最大值
//		初始化： dp[0] = nums[0], dp[1] = nums[1]
//		状态转移： dp[i] = max(dp[i - 2], dp[i - 3]) + nums[i]
//			只需在 dp[i - 2] 和 dp[i - 3] 中选择较大值即可，
//			如果还需要考虑 dp[i - 4] ，那么必定还可以选择第 i - 2 个数，
//			因此，结果不会更优，再前面的同理
//
//		时间复杂度： O(n)
//		空间复杂度： O(n) 【当然可以优化为 O(1) ，因为每次只用到最近的 3 个 dp 值】

func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}

	dp := make([]int, length)
	dp[0], dp[1] = nums[0], nums[1]
	for i := 2; i < length; i++ {
		dp[i] = dp[i - 2] + nums[i]
		if i > 2 {
			dp[i] = max(dp[i], dp[i - 3] + nums[i])
		}
	}
	return max(dp[length - 2], dp[length - 1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
