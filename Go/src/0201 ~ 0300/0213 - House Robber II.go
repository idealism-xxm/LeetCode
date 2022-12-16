// 链接：https://leetcode.com/problems/house-robber-ii/
// 题意：给定一个数组表示的环，不能选择相邻的两个数，求选择某些数的和的最大值？


// 数据限制：
//  1 <= nums.length <= 100
//  0 <= nums[i] <= 1000


// 输入： [2,3,2]
// 输出： 3
// 解释： 第一个数和第三个数在环中相邻，所以不能同时选，那么只能选第二个数

// 输入： [1,2,3,1]
// 输出： 4
// 解释： 选择第一个数和第三个数即可


// 思路： DP
//
//      本题是 LeetCode 198 的加强版，将数组变成了环。
//
//      将数组变成环其实就只多了一个限制：第一个数和最后一个数不能同时选。
//
//      我们可以将其转化为数组的形式：
//          1. 必定不选 nums[0] ，那么相当于对数组 nums[1..] 求可选数的最大和
//          2. 必定不选 nums[n - 1] ，那么相当于对数组 nums[..n - 1] 求可选数的最大值
//
//      分别调用 LeetCode 198 的 DP 方法求最大和，然后再取两者最大值即可。
//
//
//      设 dp[i] 表示在 nums[..=i] 中选择的数的最大和。
//
//      初始化： 
//          1. dp[0] = nums[0]: nums[..=0] 中只能选择 nums[0]
//          2. dp[1] = max(nums[0], nums[1]): 
//              nums[..=1] 中不能两个都选，只能二选一，贪心选择最大的
//            
//      状态转移： dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])
//          1. 不选 nums[i] ，则 dp[i] 由 dp[i - 1] 转移而来
//          2. 选择 nums[i] ，则 dp[i] 由 dp[i - 2] + nums[i] 转移而来
//
//
//      DP 常见的三种优化方式见 LeetCode 583 这题的思路，
//      本题只能采用滚动数组的方式进行优化，将空间复杂度从 O(n) 优化为 O(1) 。
//      本实现为了便于理解，不做优化处理。
//
//
//      时间复杂度： O(n)
//          1. 需要遍历 dp 中全部 O(n) 个状态
//      空间复杂度： O(n)
//          2. 需要维护 dp 中全部 O(n) 个状态


func rob(nums []int) int {
    n := len(nums)
    // 如果只有一个数，那么必定是选择这个数
    if n == 1 {
        return nums[0]
    }

    // 返回 nums[..n - 1] 和 nums[1..] 中最终结果的较大值
    return max(doRob(nums[:n - 1]), doRob(nums[1:]))
}

func doRob(nums []int) int {
    n := len(nums)
    // 如果只有一个数，那么必定是选择这个数
    if n == 1 {
        return nums[0]
    }

    // dp[i] 表示从 nums[..=i] 中选择的数的最大和
    dp := make([]int, n)
    // nums[..=0] 中只能选择 nums[0]
    dp[0] = nums[0]
    // nums[..=1] 中不能两个都选，只能二选一，贪心选择最大的
    dp[1] = max(nums[0], nums[1])
    for i := 2; i < n; i++ {
        // 1. 不选 nums[i] ，则 dp[i] 由 dp[i - 1] 转移而来
        // 2. 选择 nums[i] ，则 dp[i] 由 dp[i - 2] + nums[i] 转移而来
        dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])
    }

    return dp[n - 1]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
