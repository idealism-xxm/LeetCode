// 链接：https://leetcode.com/problems/maximum-length-of-repeated-subarray/
// 题意：给定两个数组 nums1 和 nums2 ，求其最长公共子数组的长度？


// 数据限制：
//  1 <= nums1.length, nums2.length <= 1000
//  0 <= nums1[i], nums2[i] <= 100


// 输入： nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
// 输出： 3
// 解释： 最长公共子数组为 [3,2,1] ，长度为 3

// 输入： nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
// 输出： 5
// 解释： 最长公共子数组为 [0,0,0,0,0] ，长度为 5


// 思路： DP
//
//		本题其实就是求最长公共子串 LCS ，可以使用 DP 进行处理。
//
//      设 dp[i][j] 表示 nums1[:i] 和 nums2[:j] 的最长后缀的长度。
//
//      初始化： dp[0][j] = 0; dp[i][0] = 0; 
//          表示空数组的最长后缀必定是 0 。
//      状态转移：
//          1. nums[i - 1] == nums[j - 1]: dp[i][j] = dp[i - 1][j - 1] + 1;
//              最后一个数字相同，则状态可由 dp[i - 1][j - 1] 转移而来
//          2. nums[i - 1] != nums[j - 1]: dp[i][j] = 0;
//              最后一个数字不同，则其最长后缀的长度必定是 0
//
//      注意到初始化和状态转移中的分支 2 都是将对应的值赋值为 0 ，
//      那么我们在最开始就可以初始化全部 dp[i][j] = 0 ，
//      状态转移时就无需处理分支 2 的情况。
//
//      同时我们要维护所有这些最长后缀长度的最大值 ans ，
//      那么 ans 就是题目所求最长公共子数组的长度。
//
//
//      DP 常见的三种优化方式见 LeetCode 583 这题的思路，
//      本题可以采用一维数组 + 倒序转移的方式进行优化，
//      能将空间复杂度从 O(mn) 优化为 O(n) ，
//      本实现为了便于理解，不做优化处理。
//
//
//      时间复杂度：O(mn)
//          1. 需要对 dp 中全部 O(mn) 个状态进行转移
//      空间复杂度：O(mn)
//          1. 需要维护 dp 中全部 O(mn) 个状态


func findLength(nums1 []int, nums2 []int) int {
    m, n := len(nums1), len(nums2)
    // dp[i][j] 表示 nums1[:i] 和 nums2[:j] 的最长后缀的长度
    dp := make([][]int, m + 1)
    dp[0] = make([]int, n + 1)
    // ans 维护所有 dp[i][j] 的最大值，即最长公共子数组的长度
    ans := 0
    for i := 1; i <= m; i++ {
        dp[i] = make([]int, n + 1)
        for j := 1; j <= n; j++ {
            // 如果最后一个数相同，则状态可由 dp[i - 1][j - 1] 转移而来
            if nums1[i - 1] == nums2[j - 1] {
                dp[i][j] = dp[i - 1][j - 1] + 1
                ans = max(ans, dp[i][j])
            }
        }
    }

    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}