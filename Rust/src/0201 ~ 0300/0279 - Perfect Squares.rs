// 链接：https://leetcode.com/problems/perfect-squares/
// 题意：给定一个正整数 n ，求 n 最少能表示成多少个完全平方数的和？


// 数据限制：
//  1 <= n <= 10 ^ 4


// 输入： n = 12
// 输出： 3
// 解释： 12 = 4 + 4 + 4

// 输入： n = 13
// 输出： 2
// 解释： 12 = 4 + 9


// 思路： DP
//
//      设 dp[i] 表示 i 最少能表示成 dp[i] 个完全平方数之和。
//
//      初始化： dp[i] = n + 1: 表示暂时还不确定，同时方便后续处理。
//              dp[0] = 0: 0 最少能表示成 0 个完全平方数之和。
//      状态转移：dp[i] = min(dp[i], dp[i - square])
//              遍历每个状态 i ，然后遍历小于等于 i 的完全平方数 square ，
//              则 i 可由 i - square 转移而来。
//
//
//      时间复杂度： O(n * sqrt(n))
//          1. 需要遍历 dp 中全部 O(n) 个状态，
//              每次都需要遍历全部 O(sqrt(n)) 个完全平方数
//      空间复杂度： O(n)
//          1. 需要维护 dp 中全部 O(n) 个状态


impl Solution {
    pub fn num_squares(n: i32) -> i32 {
        let n = n as usize;
        // dp[i] 表示 i 最少能表示成 dp[i] 个完全平方数之和。
        // 初始化为 n + 1 ，表示暂时还不确定，同时方便后续处理
        let mut dp = vec![n as i32 + 1; n + 1];
        // 0 最少能表示成 0 个完全平方数之和
        dp[0] = 0;
        // 遍历所有状态 i
        for i in 1..=n {
            // 遍历所有小于等于 i 的完全平方数 square ，则 i 可由 i - square 转移而来。
            for square in (1..=i).map(|j| j * j).take_while(|&square| square <= i) {
                dp[i] = dp[i].min(dp[i - square] + 1);
            }
        }

        dp[n]
    }
}
