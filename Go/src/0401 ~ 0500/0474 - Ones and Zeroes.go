// 链接：https://leetcode.com/problems/ones-and-zeroes/
// 题意：给定一个 01 串数组 strs ，以及两个整数 m 和 n 。
//      返回 strs 满足以下要求的最大子集的大小：
//          1. 子集中所有字符串中的 0 的个数之和不超过 m
//          2. 子集中所有字符串中的 1 的个数之和不超过 n


// 数据限制：
//  1 <= strs[i].length <= 100
//  strs[i] 只含有数位 '0' and '1'
//  1 <= m, n <= 100


// 输入： strs = ["10","0001","111001","1","0"], m = 5, n = 3
// 输出： 4
// 解释： 满足题意的最大子集是： {"10", "0001", "1", "0"} ，所以答案是 4
//       其他合法但是小一点的子集有： {"0001", "1"} 和 {"10", "1", "0"} 。
//       注意 {"111001"} 不满足题意，因为 1 的个数是 4 ，超过了 3 个。

// 输入： strs = ["10","0","1"], m = 1, n = 1
// 输出： 2
// 解释： 满足题意的最大子集是： {"0", "1"} ，所以答案是 2


// 思路： DP
//
//      本题是一个 01 背包问题，只不过限制条件变成了 2 个，所以增加一维状态即可。
//
//      设 dp[k][i][j] 表示在前 k 个字符串中，
//      满足集合中 0 的个数不超过 i 个、 1 的个数不超过 j 个的最大子集的大小。
//
//      初始化都为 0 ，表示空集必定满足题意。
//
//      设第 k 个字符串中 0 和 1 的个数分别为 zeroCount 和 oneCount ，
//      则有以下状态转移：
//          1. 不选第 k 个字符串： dp[k][i][j] = dp[k - 1][i][j]
//          2. 选第 k 个字符串： dp[k][i][j] = dp[k - 1][i - zeroCount][j - oneCount] + 1
//
//      即： dp[k][i][j] = max(dp[k - 1][i][j], dp[k - 1][i - zeroCount][j - oneCount] + 1)
//
//      可以使用优化消除 dp 的第一维，因为本题是 01 背包，所以每个字符串只能选择一次，
//      那么消除第一维后，状态 i 和 j 均需要从大到小遍历，防止一个字符串被多次选择。
//
//
//      设 strs 的长度为 l 。
//
//      时间复杂度：O(l * m * n)
//          1. 需要遍历 strs 全部 O(l) 个字符串，
//              遍历每个字符串时还需要遍历 dp 中全部 O(mn) 个状态
//      空间复杂度：O(mn)
//          1. 需要维护一个大小为 O(mn) 的数组 dp


func findMaxForm(strs []string, m int, n int) int {
    // dp[i][j] 表示在当前已遍历的字符中，
    // 满足集合中 0 的个数不超过 i 个、 1 的个数不超过 j 个的最大子集的大小。
    // 初始化都为 0 ，表示空集必定满足题意。
    dp := make([][]int, m + 1)
    for i := 0; i <= m; i++ {
        dp[i] = make([]int, n + 1)
    }
    // 遍历每个字符串
    for _, cur := range strs {
        // 统计当前字符串中 0 和 1 的个数
        zeroCount := strings.Count(cur, "0")
        oneCount := len(cur) - zeroCount

        // 由于是 01 背包，每个字符串只能使用一次，所以要倒着遍历
        for i := m; i >= zeroCount; i-- {
            for j := n; j >= oneCount; j-- {
                // 1. 不选当前字符串，则 dp[i][j] = dp[i][j]
                // 2. 选当前字符串，则 dp[i][j] = dp[i - zeroCount][j - oneCount] + 1
                dp[i][j] = max(dp[i][j], dp[i - zeroCount][j - oneCount] + 1)
            }
        }
    }

    return dp[m][n]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}