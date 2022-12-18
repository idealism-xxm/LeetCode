// 链接：https://leetcode.com/problems/longest-common-subsequence/
// 题意：求两个字符串的最长公共子序列的长度？


// 数据限制：
//  1 <= text1.length, text2.length <= 1000
//  text1 和 text2 仅由英文小写字母组成


// 输入： text1 = "abcde", text2 = "ace"
// 输出： 3
// 解释： 最长公共子序列为 "ace" ，长度为 3

// 输入： text1 = "abc", text2 = "abc"
// 输出： 3
// 解释： 最长公共子序列为 "abc" ，长度为 3

// 输入： text1 = "abc", text2 = "def"
// 输出： 0
// 解释： 没有公共子序列，长度为 0


// 思路： DP
//
//      设 dp[i][j] 表示 text1[..i] 和 text2[..j] 的最长公共子序列的长度。
//
//      初始化： dp[i][0] = dp[0][j] = 0 ，即空串的最长公共子序列的长度为 0 。
//      状态转移方程：
//          1. 如果 text1[i] == text2[j] ，则必定选择这两个字符作为最长公共子序列的结尾。
//             则 dp[i + 1][j + 1] 由 dp[i][j] + 1 转移而来，
//             即 dp[i + 1][j + 1] = dp[i][j] + 1
//          2. 如果 text1[i] != text2[j] ，
//             则 dp[i + 1][j + 1] 只能从 dp[i + 1][j] 和 dp[i][j + 1] 直接转移，
//             取两者最大值转移即可，
//             即 dp[i + 1][j + 1] = max(dp[i + 1][j], dp[i][j + 1])       
//
//      最后， dp[m][n] 就是 text1 和 text2 的最长公共子序列的长度。
//
//
//      DP 常见的三种优化方式见 LeetCode 583 这题的思路，
//      本题可以采用这三种方式进行优化，将空间复杂度从 O(mn) 优化为 O(n) 。
//      本实现为了便于理解，不做优化处理。
//
//
//      时间复杂度：O(mn)
//          1. 需要遍历计算 dp 中全部 O(mn) 个状态
//      空间复杂度：O(mn)
//          1. 需要维护 dp 中全部 O(mn) 个状态


func longestCommonSubsequence(text1 string, text2 string) int {
    m, n := len(text1), len(text2)
    // dp[i][j] 表示 text1[..i] 和 text2[..j] 的最长公共子序列的长度
    dp := make([][]int, m + 1)
    dp[0] = make([]int, n + 1)
    for i := range text1 {
        dp[i + 1] = make([]int, n + 1)
        for j := range text2 {
            if text1[i] == text2[j] {
                // 如果 text1[i] == text2[j] ，则必定选择这两个字符作为最长公共子序列的结尾，
                // 那么状态 dp[i + 1][j + 1] 可由 dp[i][j] + 1 转移而来
                dp[i + 1][j + 1] = dp[i][j] + 1
            } else {
                // 如果 text1[i] != text2[j] ，
                // 则 dp[i + 1][j + 1] 只能从 dp[i + 1][j] 和 dp[i][j + 1] 直接转移
                dp[i + 1][j + 1] = max(dp[i + 1][j], dp[i][j + 1])
            }
        }
    }

    // dp[m][n] 就是 text1 和 text2 的最长公共子序列的长度
    return dp[m][n]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
