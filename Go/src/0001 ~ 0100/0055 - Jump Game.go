// 链接：https://leetcode.com/problems/jump-game/
// 题意：给定一个非负整数数组 nums，初始位于数组第一位，
//      若当前在 nums[i] ，则最远可以跳到 i + nums[i] ，
//      求是否能跳到数组最后一位？

// 输入：[2,3,1,1,4]
// 输出：true

// 输入：[3,2,1,0,4]
// 输出：false

// 思路1：DP
//		Jump Game II（https://leetcode.com/problems/jump-game-ii/）的简化版
//      dp[i] 表示 能否达到 i 处
//      初始化:
//      dp[0] = true
//
//      若已知 dp[i] = true ，则对于 ∀j ∈ (i, i + nums[i]] 有 dp[i + j] = true
//      时间复杂度： O(n^2) ，空间复杂度： O(n)

func canJump(nums []int) bool {
	length := len(nums)
	dp := make([]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = false
	}
	dp[0] = true
	for i := 0; i < length; i++ {
		num := nums[i]
		if dp[i] {
			for j := 1; j <= num && i + j < length; j++ {
				dp[i + j] = true
			}
		}

	}
	return dp[length - 1]
}

// 思路2：贪心
//      维护一个最远能到达的下标，最后看这个下标是否大于最后一个下标
//      时间复杂度： O(n) ，空间复杂度： O(n)

func canJump(nums []int) bool {
	length := len(nums)
	maxIndex := 0  // 初始只能到第一个元素
	for i := 0; i < length; i++ {
		if i <= maxIndex {
			maxIndex = max(maxIndex, i + nums[i])
		}
	}
	return maxIndex >= length - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
