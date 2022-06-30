// 链接：https://leetcode.com/problems/pascals-triangle-ii/
// 题意：给定一个整数 rowIndex ，返回杨辉三角的第 rowIndex 行。
//
//      在杨辉三角中，每一个数是它左上方和右上方的数之和。
//
//      进阶：使用空间复杂度为 O(n) 的方法。


// 数据限制：
//  1 <= rowIndex <= 33


// 输入： rowIndex = 3
// 输出： [1,3,3,1]

// 输入： rowIndex = 0
// 输出： [1]

// 输入： rowIndex = 1
// 输出： [1,1]


// 思路： 模拟
//
//      本题是 LeetCode 118 这题的加强版，需要将空间复杂度优化为 O(n) 。
//
//      空间复杂度为 O(n ^ 2) 的方法很简单，按照题意从第一层开始计算即可，
//      对于每一个位置 (i, j) 有 dp[i][j] = dp[i - 1][j - 1] + dp[i][j] 。
//
//      注意边界情况，当 j == 0 || j == i 时， dp[i][j] = 1 。
//
//      DP 常见的三种优化方式见 LeetCode 583 这题的思路，
//      本题直接使用一维数组 + 倒序转移这种方法进行优化。 
//
//      因为状态 dp[i][j] 仅由 dp[i - 1][..=j] 中的状态转移而来时，
//      可以使用倒序转移，从后往前计算，以免使用当前行的状态进行转移。
//
//
//      时间复杂度：O(n ^ 2)
//          1. 需要遍历计算全部 O(n ^ 2) 个位置的数
//      空间复杂度：O(n)
//          1. 需要用 dp 维护全部 O(n) 个位置的数


impl Solution {
    pub fn get_row(row_index: i32) -> Vec<i32> {
        // 默认全都是 1 ，直接所有行的边界情况设置好
        let mut dp = vec![1; row_index as usize + 1];
        for i in 0..=row_index as usize {
            // 倒序转移，仅处理第 i 层中间非边界情况的数，即它等于左上和右上的数之和
            for j in (1..i).rev() {
                dp[j] += dp[j - 1];
            }
        }

        dp
    }
}
