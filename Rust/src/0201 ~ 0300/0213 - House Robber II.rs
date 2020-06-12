// 链接：https://leetcode.com/problems/house-robber-ii/
// 题意：给定一个数组表示的环，不能选择相邻的两个数，求选择某些数的和的最大值？

// 输入： [2,3,2]
// 输出： 3
// 解释： 第一个数和第三个数在环中相邻，所以不能同时选，那么只能选第二个数

// 输入： [1,2,3,1]
// 输出： 4
// 解释： 选择第一个数和第三个数即可

// 思路： DP
//
//		0198 的加强版
//
//		变成环其实就多了第一个数和最后一个数不能同时选的限制，
//      我们可以复用 0198 的 DP 方法，
//      分别用 nums[..nums.len() - 1] 和 nums[1..] 调用两次取较大值即可
//
//      不过这里再用另一种 DP 形式
//      设 dp[i] 表示前 i 个数经过选择后和的最大值
//      初始化： dp[0] = nums[0], dp[1] = max(nums[0], nums[1])
//      状态转移：
//          1. 选择第 i 个数，则 dp[i] = dp[i - 2] + nums[i]
//          2. 不选择第 i 个数，则 dp[i] = dp[i - 1]
//      因此状态转移方程为： dp[i] = max(dp[i - 2] + nums[i], dp[i - 1])
//
//      时间复杂度： O(n)
//      空间复杂度： O(n) 【当然可以优化为 O(1) ，因为只有最近的两个才会用于转移，但没必要】

use std::cmp;

impl Solution {
    pub fn rob(nums: Vec<i32>) -> i32 {
        if nums.len() == 0 {
            return 0;
        }
        if nums.len() == 1 {
            return nums[0];
        }
        // 返回 nums[..] 和 nums[1..] 中最终结果的较大值
        cmp::max(Solution::do_rob(&nums[..nums.len() - 1]), Solution::do_rob(&nums[1..]))
    }

    pub fn do_rob(nums: &[i32]) -> i32 {
        if nums.len() == 0 {
            return 0;
        }
        if nums.len() == 1 {
            return nums[0];
        }
        // 初始化 dp 数组
        let mut dp = vec![0; nums.len()];
        dp[0] = nums[0];
        dp[1] = cmp::max(nums[0], nums[1]);
        // 状态转移
        for i in 2..nums.len() {
            dp[i] = cmp::max(dp[i - 2] + nums[i], dp[i - 1]);
        }
        // 返回所有数中选择的和的最大值
        dp[nums.len() - 1]
    }
}
