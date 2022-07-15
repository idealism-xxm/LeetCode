// 链接：https://leetcode.com/problems/matchsticks-to-square/
// 题意：给定一个数组 matchsticks ，其中 matchsticks[i] 表示第 i 根火柴棍的长度。
//      求恰好用每根火柴棍一次，能否拼出一个正方形？


// 数据限制：
//  1 <= matchsticks.length <= 15
//  1 <= matchsticks[i] <= 10 ^ 8


// 输入：matchsticks = [1,1,2,2,2]
// 输出：true
// 解释：可以拼出一个正方形，每边长度为 2 。

// 输入：matchsticks = [3,3,3,3,4]
// 输出：false
// 解释：无法使用全部火柴棍拼出一个正方形。


// 思路1：递归/回溯/DFS
//
//      如果这些火柴棍能拼出一个正方形，必定先满足下面的两个条件：
//          1. 火柴棍的个数至少为 4 根
//          2. 所有火柴棍的长度之和 total 能被 4 整除
//
//      然后我们使用 dfs(matchsticks, i, remain) 回溯遍历所有可能的组合：
//          1. matchsticks: 所有火柴棍，直接透传
//          2. i: 当前需要放置的火柴棍的下标，初始化为 len(matchsticks) - 1
//          3. remain: 长度为 4 的整型数组， remain[j] 表示第 j 条边还需的长度，
//                  初始化均为 total / 4
//
//      在 dfs 中，我们按照如下逻辑处理：
//          1. i < 0: 已放置完全部火柴棍，则必定有 remain[0~4] = 0 ，
//                  即此时可以拼出一个正方形，直接返回 true
//          2. i >= 0: 枚举第 i 根火柴棍放置的边 j ，如果 remain[j] >= 0 ，
//                  则将其放置在第 j 条边，然后递归处理下根火柴棍。
//
//                  如果递归返回结果为 true ，则说明此时能拼出一个正方形，直接返回 true ；
//                  否则恢复状态，继续处理下一条边。
//
//                  如果放置在所有边时，都不能拼出正方形，则返回 false 。
//
//      我们在调用 dfs 前可以进一步优化：先对 matchsticks 按照升序排序，
//      然后在回溯时先枚举更长的火柴棍，这样能减小搜索空间。
//
//
//		时间复杂度： O(4 ^ n)
//          1. 需要对 matchsticks 全部 O(n) 个元素排序，时间复杂度为 O(nlogn)
//          2. 需要对 matchsticks 全部 O(n) 个元素回溯，
//              每次都有 4 种选择，时间复杂度为 O(4 ^ n)
//		空间复杂度： O(n)
//          1. 栈递归深度最大为 O(n)


func makesquare(matchsticks []int) bool {
    // 如果不足 4 个火柴棍，则不能拼出正方形
    if len(matchsticks) < 4 {
        return false
    }
    // 计算所有火柴棍的长度之和，如果不能被 4 整除，则不能拼出正方形
    total := 0
    for _, matchstick := range matchsticks {
        total += matchstick
    }
    if total % 4 != 0 {
        return false
    }
    // 初始化 4 条边都还需要 total / 4 的长度
    remain := make([]int, 4)
    for i := 0; i < 4; i++ {
        remain[i] = total / 4
    }
    // 将火柴棍长度升序排序，回溯时先枚举更长的火柴棍，能减小搜索空间
    sort.SliceStable(matchsticks, func(i, j int) bool { return matchsticks[i] < matchsticks[j] })
    
    // 回溯搜索
    lastIndex := len(matchsticks) - 1
    return dfs(matchsticks, lastIndex, remain)
}

func dfs(matchsticks []int, i int, remain []int) bool {
    // 如果放置完全部火柴棍，则必定有 remain[0~4] = 0 ，
    // 即此时可以拼出一个正方形，直接返回 true
    if i < 0 {
        return true
    }

    matchstick := matchsticks[i]
    // 枚举第 i 根火柴棍放置的边
    for j := range remain {
        // 如果第 j 条边的所需的长度小于当前火柴棍的长度，则直接处理下一条边
        if remain[j] < matchstick {
            continue
        }

        // 当前边所需的长度减去当前火柴棍的长度
        remain[j] -= matchstick
        // 递归处理下根火柴棍，如果能成功拼出正方形，则直接返回
        if dfs(matchsticks, i - 1, remain) {
            return true
        }
        // 此时不能拼出正方形，需要回溯，将当前边所需的长度加回去
        remain[j] += matchstick
    }

    // 当前情况下，第 i 根火柴棍放在哪条边都不能拼出正方形，返回 false
    return false
}


// 思路2：状压 DP
//
//      如果这些火柴棍能拼出一个正方形，必定先满足下面的两个条件：
//          1. 火柴棍的个数至少为 4 根
//          2. 所有火柴棍的长度之和 total 能被 4 整除
//
//      设正方形边长为 target = total / 4 。
//
//      由于火柴棍最多只有 15 根，所以可以考虑用状压 DP 进行处理。
//
//      考虑状压 DP 时很容易就能想到如下状态：
//          设 dp[i][k] 表示使用状态 i 中的火柴棍时，
//          已经拼出 k 条完整的边，最新一条边拼出的长度， -1 表示暂时无法拼出。
//
//      初始化：
//          1. dp[i][k] = -1: 表示还无法拼出满足题意的正方形的边
//          2. dp[0][0] = 0: 表示已拼出 0 条完整的边，最新一条边的长度为 0
//
//      状态转移：从 0 到 1 << n 枚举状态 i ，然后枚举下一根加入的火柴棍 j 。
//
//      如果 i & (1 << j) == 0 ，则说明 j 可以尝试加入其中：
//          1. dp[i][k] + matchsticks[j] < target: 没有拼完一条边，
//              该边还可以继续加入火柴棍。
//              则有 dp[i & (1 << j)][k] = dp[i][k] + matchsticks[j]
//          2. dp[i][k] + matchsticks[j] == target: 已经拼出完整的一条边，
//              后续加入的火柴棍需要拼新的边。
//              则有 dp[i & (1 << j)][k + 1] = 0
//          3. dp[i][k] + matchsticks[j] > target: 不满足题意，直接处理下一根火柴棍。
//
//      最后，如果 dp[(1 << n) - 1][4] == 0 ，则说明可以拼出正方形的全部 4 条边。
//
//      继续分析，我们就能发现第二维状态其实可以优化掉。
//
//      因为状态 i 中的火柴棍已确定，那么这些火柴棍的长度之和 sum 也已确定，
//      如果能拼出满足题意的正方形的边，则必定已拼出 sum / target 条完整的边，
//      且最新一条边的长度为 sum % target 。
//
//      则有 dp[i][k] = -1, dp[i][sum / target] = sum % target 。
// 
//      所以我们可以将状态简化为： dp[i] 表示使用状态 i 中的火柴棍，
//      拼出满足题意的正方形的边时，最新的一条边的长度。
//
//      初始化： dp[i] = -1, dp[0] = 0
//      状态转移：从 0 到 1 << n 枚举状态 i ，然后枚举下一根加入的火柴棍 j 。
//
//      如果 i & (1 << j) == 0 ，则说明 j 可以尝试加入其中：
//          1. dp[i] + matchsticks[j] < target: 没有拼完一条边，
//              该边还可以继续加入火柴棍。
//              则有 dp[i & (1 << j)] = dp[i] + matchsticks[j]
//          2. dp[i] + matchsticks[j] == target: 已经拼出完整的一条边，
//              后续加入的火柴棍需要拼新的边。
//              则有 dp[i & (1 << j)] = 0
//          3. dp[i] + matchsticks[j] > target: 不满足题意，直接处理下一根火柴棍
//
//          综上：当 dp[i] + matchsticks[j] <= target 时，
//               有 dp[i & (1 << j)] = (dp[i] + matchsticks[j]) % target 。
//
//      最后，如果 dp[(1 << n) - 1] == 0 ，则说明可以拼出正方形的全部 4 条边。
//
//
//		时间复杂度： O(n * 2 ^ n)
//          1. 需要枚举全部 O(2 ^ n) 个状态
//          2. 枚举每个状态时，都需要枚举全部 O(n) 根火柴棍
//		空间复杂度： O(2 ^ n)
//          1. 需要维护 dp 全部 O(2 ^ n) 个状态


func makesquare(matchsticks []int) bool {
    n := len(matchsticks)
    // 如果不足 4 个火柴棍，则不能拼出正方形
    if n < 4 {
        return false
    }
    // 计算所有火柴棍的长度之和，如果不能被 4 整除，则不能拼出正方形
    total := 0
    for _, matchstick := range matchsticks {
        total += matchstick
    }
    if total % 4 != 0 {
        return false
    }
    // 正方形的边长为 total / 4
    target := total / 4
    // 总共有 mx = 2 ^ n 个状态
    mx := 1 << n
    // dp[i] 表示使用状态 i 中的火柴棍，拼出满足题意的正方形的边时，
    // 最新的一条边的长度。
    //
    // 因为状态 i 中的火柴棍已确定，那么这些火柴棍的长度之和 sum 也已确定，
    // 所以 dp[i] 的值只有两种情况，没有其他值的可能性：
    //  1. dp[i] == -1: 则使用状态 i 中的火柴棍，无法拼出满足题意的正方形的边
    //  2. dp[i] == sum % target: 则使用状态 i 中的火柴棍，
    //      已拼出 sum / target 条完整的边，最新一条边的长度为 sum % target
    dp := make([]int, mx)
    // 初始化 dp[0] 为 0 ，表示已拼出 0 条完整的边，最新一条边的长度为 0 。
    // 其他 dp 值都为 -1 ，表示还无法拼出满足题意的正方形的边，等后续状态转移时更新。
    for i := 1; i < mx; i++ {
        dp[i] = -1
    }

    for i := range dp {
        // 如果 dp[i] 为 -1 ，则使用状态 i 中的火柴棍，
        // 无法拼出满足题意的正方形的边，直接处理下一个状态
        if dp[i] == -1 {
            continue
        }
        // 枚举接下来要使用的火柴棍
        for j := range matchsticks {
            // 如果第 j 根火柴棍不在状态 i 中，且加入后拼出的边长不超过 target ，
            // 则状态 i | (1 << j) 可以拼出满足题意的正方形的边，更新状态
            if (i & (1 << j)) == 0 && dp[i] + matchsticks[j] <= target {
                // 1. dp[i] + matchsticks[j] < target: 没有拼完一条边，
                //      该边还可以继续加入火柴棍
                // 2. dp[i] + matchsticks[j] == target: 已经拼出完整的一条边，
                //      后续加入的火柴棍需要拼新的边，即需要令 dp[i | (1 << j)] = 0
                //
                // 综上：直接对 dp[i] + matchsticks[j] 取模即可
                dp[i | (1 << j)] = (dp[i] + matchsticks[j]) % target
            }
        }
    }

    // 如果使用完全部火柴棍，拼出了 4 条边，则必定有 dp[mx - 1] == 0
    return dp[mx - 1] == 0
}
