// 链接：https://leetcode.com/problems/jump-game-ii/
// 题意：给定一个非负整数数组 nums，初始位于数组第一位，
//      若当前在 nums[i] ，则最远可以跳到 i + nums[i] ，
//      求最少几步能跳到数组最后一位？


// 输入：[2,3,1,1,4]
// 输出：2

// 思路1：DP
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

// 思路2：贪心
//
//      看了题解才发现还有很多没有利用的信息
//
//      我们对 思路 1 中得到的 dp 数组进行分析，
//      可以发现： dp 数组是非递减的，并且每次从 i 变为 i + 1 时，
//      当前位置都是 i - 1 中所有位置能跳一次到的最远的位置 + 1 ，
//      而这个位置我们不需要嵌套循环即可直接计算出来
//
//      我们可以维护三个值：
//          count: 跳到当前位置所需的最小次数
//          last: 跳跃 count 次能到达的最远位置
//          maxIndex: 跳跃 count 次，下一次跳跃后能到达的最远位置
//      当 i > last 时，表明跳跃 count 次能到达的一段已经处理完毕，
//      当前位置开始需要跳跃 count + 1 次，所以需要更新：
//          last = maxIndex
//          count++
//      而每次都需要更新 maxIndex ，以便获取跳跃 count + 1 次能到达的最远位置：
//          maxIndex = max(maxIndex, i + nums[i])
//
//      时间复杂度： O(n)
//      空间复杂度： O(1)

func jump(nums []int) int {
    // 开始在 index = 0 处，且不需要跳跃
    // 所以跳跃次数为 0 的一段，在 last = 0 处结束
    last, count := 0, 0
    // maxIndex 表示跳跃次数为 count 时，下一次跳跃能到达的最远位置
    maxIndex := nums[0]
    for i := 1; i < len(nums); i++ {
        // 如果跳跃为 count 时的一段已经结束，
        // 则接下来处理跳跃次数为 count 的一段，
        // 其对应的结束位置为 maxIndex
        if i > last {
            last = maxIndex
            count++
        }
        // 跳跃次数为 count 时，下一次能跳到的最远位置
        maxIndex = max(maxIndex, i + nums[i])
    }
    return count
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
