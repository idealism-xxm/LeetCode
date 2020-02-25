// 链接：https://leetcode.com/problems/jump-game-ii/
// 题意：给定一个非负整数数组 nums，初始位于数组第一位，
//      若当前在 nums[i] ，则最远可以跳到 i + nums[i] ，
//      求最少几步能跳到数组最后一位？


// 输入：[2,3,1,1,4]
// 输出：2

// 思路：DP
//      很容易就能看出来是 DP
//      dp[i] 表示 跳到 i 处最少所需的步数
//      初始化:
//      dp[0] = 0
//
//      若已知 dp[i] ，则对于 ∀j ∈ (i, i + nums[i]] 有 dp[i + j] = min(dp[i + j], dp[i] + 1)
//      时间复杂度： O(n^2) ，空间复杂度： O(n)

func jump(nums []int) int {
    length := len(nums)
    dp := make([]int, length)
    for i := 0; i < length; i++ {
        dp[i] = length  // 因为一定存在答案，所以最多 length - 1 必定能够跳到最后一位，取第一个不可达的值
    }
    dp[0] = 0
    for i := 0; i < length; i++ {
        num := nums[i]
        for j := 1; j <= num && i + j < length; j++ {
            dp[i + j] = min(dp[i + j], dp[i] + 1)
        }
    }
    return dp[length - 1]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
