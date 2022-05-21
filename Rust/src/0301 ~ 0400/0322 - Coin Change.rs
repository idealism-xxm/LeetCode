// 链接：https://leetcode.com/problems/coin-change/
// 题意：给定一个正整数数组 coins ，表示不同的硬币面值，
//      每种面值的硬币可以无限使用，
//      求能凑出 amount 所需的最少硬币数量？
//      如果无法凑出 amount ，则返回 -1 。


// 数据限制：
//  1 <= coins.length <= 12
//  1 <= coins[i] <= 2 ^ 31 - 1
//  0 <= amount <= 10 ^ 4


// 输入： coins = [1,2,5], amount = 11
// 输出： 3
// 解释： 11 = 5 + 5 + 1

// 输入： coins = [2], amount = 3
// 输出： -1

// 输入： coins = [2], amount = 0
// 输出： 0


// 思路： DP
//
//      设 dp[i] 表示凑出 i 所需的最少硬币数量。
//
//      初始状态：
//          1. dp[i] = amount + 1: amount + 1 表示当前还凑不出
//          2. dp[0] = 0: 最开始只能确定不需要任何硬币就可以凑出 0
//
//      状态转移方程：当 i >= coin[j] 时， dp[i] = min(dp[i], dp[i - coin[j]] + 1) ，
//          即凑出 i 所需的最少硬币数量 = 凑出 i - coin[j] 所需的最少硬币数量 + 1
//      
//      最后，根据 dp[amount] 返回答案：
//          1. dp[amount] == amount + 1: 说明无法凑出 amount ，返回 -1
//          2. dp[amount] != amount + 1: 则 dp[amount] 就是所需的最少硬币数量
//
//
//      时间复杂度：O(n * amount)
//          1. 需要遍历 dp 全部 O(amount) 个状态，
//              遍历每个状态时还需要遍历 coins 中全部 O(n) 个硬币
//      空间复杂度：O(amount)
//          1. 需要维护一个大小为 O(amount) 的数组 dp


impl Solution {
    pub fn coin_change(coins: Vec<i32>, amount: i32) -> i32 {
        // dp[i] 表示凑出 i 所需的最少硬币数量，
        // 初始化为 amount + 1 ，表示当前还凑不出
        let mut dp = vec![amount + 1; amount as usize + 1];
        // 最开始只能确认不需要任何硬币就可以凑出 0
        dp[0] = 0;
        // 遍历要凑的数字状态
        for i in 1..=amount {
            // 遍历当前选择的硬币
            for &coin in coins.iter() {
                // 如果 i >= coin 时，
                // 则可以选择 coin ，进行状态转移
                if i >= coin {
                    dp[i as usize] = dp[i as usize].min(dp[(i - coin) as usize] + 1);
                }
            }
        }

        if dp[amount as usize] == amount + 1 {
            // 无法凑出 amount ，返回 -1
            -1
        } else {
            // dp[amount] 就是所需的最少硬币数量
            dp[amount as usize]
        }
    }
}
