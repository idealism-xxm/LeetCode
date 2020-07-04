// 链接：https://leetcode.com/problems/number-of-digit-one/
// 题意：求 [1, n] 内每个数字含有的 1 的个数之和？

// 输入： 13
// 输出： 6
// 解释： 1, 10, 11, 12, 13 总共有 6 个 1

// 思路： 数位 DP
//
//      设 dp[i][j] 表示长度为 i 的数字（含前导零）
//      且前面含有 j 个数位 1 （不含这 i 个数位）时的所有数字数位 1 之和
//
//      使用 dfs 进行状态转移，其中：
//          index 表示当前剩余的数字的长度，
//          one_count 表示前面含有的数位 1 的个数（不含这 i 个数位）
//          limit 表示是否存在限制（当前面考虑的数字是 n 的前缀时就存在限制）
//      那么每一层，如果：
//          1. index == 0: 没数字了，直接返回当前数字中数位 1 的个数
//          2. !limit && dp[index][one_count] != -1: 已经记忆化过了，直接返回 dp[index][one_count]
//              注意：只有没有限制才行，因为限制的情况对于多个数字并不相同（在多测试用例复用的情况下）
//          3. 确定当前数位能到达的最大值 max
//              当有限制时就是 digit[index] ，没有限制时就是 9
//          4. 计算本层产生的结果，遍历当前数位的值 0..=max
//              递归处理即可，注意数位位 1 时要对 one_count + 1
//          5. 若不存在限制，则进行记忆化
//          6. 返回本层结果
//
//
//      时间复杂度： O((lgn) ^ 2)
//      空间复杂度： O((lgn) ^ 2)

impl Solution {
    pub fn count_digit_one(n: i32) -> i32 {
        let mut n = n;
        // dp[i][j] 表示长度为 i 的数字（含前导零）且前面含有 j 个数位 1 的数字所有数位 1 之和
        let dp: [[i32; 33]; 33] = [[-1; 33]; 33];
        // 从 index = 1 开始按低位到高位存储
        let mut digit: [usize; 33] = [0; 33];
        let mut len: usize = 1;
        while n > 0 {
            digit[len] = (n % 10) as usize;
            n /= 10;
            len += 1;
        }

        // 进行 数位 DP
        let (_, result) = Solution::dfs(dp, &digit, len, 0, true);
        result
    }

    fn dfs(dp: [[i32; 33]; 33], digit: &[usize; 33], index: usize, one_count: usize, limit: bool) -> ([[i32; 33]; 33], i32) {
        // 所有位已考虑完，直接返回
        if index == 0 {
            return (dp, one_count as i32);
        }
        // 如果这个状态已经计算过，则直接返回
        if !limit && dp[index][one_count] != -1 {
            return (dp, dp[index][one_count]);
        }

        // 计算当前状态下，当前位能到达的最大值，
        // 如果有限制，则只能到达 digit[index] ，
        // 否则可以到达 9
        let max = if limit { digit[index] } else { 9 as usize };
        let mut dp = dp;
        // 遍历所有可能的值
        let mut result = 0;
        for i in 0..=max {
            // 如果遍历的数字位 1 ，则下一层状态又多了一个数位 1
            let next_one_count = one_count + if i == 1 { 1 } else { 0 };
            // 递归处理
            let (dfs_dp, dfs_result) = Solution::dfs(dp, digit, index - 1, next_one_count, limit && i == max);
            dp = dfs_dp;
            result += dfs_result;
        }

        // 如果没有限制，则当前结果可以记忆化
        if !limit {
            dp[index][one_count] = result;
        }
        (dp, result)
    }
}
