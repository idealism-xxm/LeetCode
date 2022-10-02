// 链接：https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/
// 题意：有 n 个骰子，每个骰子都有 k 面，上面的数字分别是 1 ~ k ，
//      求使得所有骰子朝上一面的数字之和恰好为 target 的方案数？
//      结果模 10 ^ 9 + 7 。


// 数据限制：
//  1 <= n, k <= 30
//  1 <= target <= 1000


// 输入： n = 1, k = 6, target = 3
// 输出： 1
// 解释： 有 1 个 6 面骰子，那么该骰子朝上的面只能是 3 。

// 输入： n = 2, k = 6, target = 7
// 输出： 6
// 解释： 有 2 个 6 面骰子，总共有以下 6 种方案使得朝上的面的数字之和为 7 ：
//       1 + 6, 2 + 5, 3 + 4, 4 + 3, 5 + 2, 6 + 1

// 输入： n = 30, k = 30, target = 500
// 输出： 222616187
// 解释： 结果需要模 10 ^ 9 + 7


// 思路： DP
//
//      一般这种求方案数的都能使用 DP 进行处理，本题的状态也非常好定义。
//
//      设 dp[i][j] 表示前 i 个骰子的数字之和为 j 时的方案数。
//
//      初始化：初始只能确定前 0 个骰子的数字之和为 0 这个空状态是一个合法的方案。
//          即 dp[i][j] = 0; dp[0][0] = 1
//      状态转移：第 i 个骰子的数字为 l 时，
//          状态 dp[i][j] 可由状态 dp[i - 1][j - l] 转移而来。
//          即 dp[i][j] = dp[i][j] + dp[i - 1][j - l]  
//
//
//      DP 常见的三种优化方式见 LeetCode 583 这题的思路，
//      本题只能采用滚动数组的方式进行优化（因为 dp[i] 依赖 dp[i - 1] 的全部状态），
//      能将空间复杂度从 O(n * target) 优化为 O(target) 。
//      本实现为了便于理解，不做优化处理。
//
//
//      时间复杂度：O(nk * target)
//          1. 需要遍历全部 O(n) 个骰子，
//              遍历每个骰子时都需要遍历该骰子的全部 O(k) 个数字，
//              遍历每个数字时都需要遍历全部 O(target) 个数字之和。
//      空间复杂度：O(n * target)
//          1. 需要维护 dp 中全部 O(n * target) 个状态


const MOD: i32 = 1_000_000_007;


impl Solution {
    pub fn num_rolls_to_target(n: i32, k: i32, target: i32) -> i32 {
        let (n, k, target) = (n as usize, k as usize, target as usize);
        // dp[i][j] 表示前 i 个骰子的数字之和为 j 时的方案数
        let mut dp = vec![vec![0; target + 1]; n + 1];
        // 初始只能确定前 0 个骰子的数字之和为 0 这个空状态是一个合法的方案
        dp[0][0] = 1;
        for i in 1..=n {
            for l in 1..=k {
                for j in l..=target {
                    // 第 i 个骰子的数字为 l 时，
                    // 状态 dp[i][j] 可由状态 dp[i - 1][j - l] 转移而来
                    dp[i][j] = (dp[i][j] + dp[i - 1][j - l]) % MOD;
                }
            }
        }

        dp[n][target]
    }
}
