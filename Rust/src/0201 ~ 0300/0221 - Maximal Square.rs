// 链接：https://leetcode.com/problems/maximal-square/
// 题意：给定一个 01 矩阵，求全为 1 的正方形的最大面积？

// 输入：
// 1 0 1 0 0
// 1 0 1 1 1
// 1 1 1 1 1
// 1 0 0 1 0
// 输出： 4

// 思路： DP
//
//		设 dp[i][j] 表示以 matrix[i][j] 为右下角的最大正方形的边长，
//      状态转移：
//          1. 若 matrix[i][j] = 0 ，则 dp[i][j] = 0
//          2. 若 matrix[i][j] = 1 ，则可以进行状态转移：
//              以 matrix[i][j] 为右下角的最大正方形的边长 取决于
//              以 matrix[i - 1][j] 为右下角的最大正方形的边长 和
//              以 matrix[i][j - 1] 为右下角的最大正方形的边长 的较小值 length
//              设： length = min(dp[i - 1][j], dp[i][j - 1])
//              则： dp[i][j] = length + (matrix[i - length][j - length] - '0')
//
//              这个正方形是最大的，不可能再往外扩大了，若要变长再多 1 ，
//              则左边正方形和右边正方形的变成都需要再多 1 ，然而从前面已经推出来，
//              再变成再多 1 就会有一个正方形的左边或者上面存在 0
//
//      时间复杂度： O(m * n)
//      空间复杂度： O(m * n) 【注意：当然可以直接使用一维数组优化为 O(1) ，但没必要】

use std::cmp;

impl Solution {
    pub fn maximal_square(matrix: Vec<Vec<char>>) -> i32 {
        if matrix.len() == 0 || matrix[0].len() == 0 {
            return 0;
        }

        let mut dp = vec![vec![0 as usize; matrix[0].len()]; matrix.len()];
        let mut result = 0;
        for i in 0..matrix.len() {
            for j in 0..matrix[0].len() {
                // 若为 0 ，则不可能组成正方形，直接处理下一个
                if matrix[i][j] == '0' {
                    dp[i][j] = 0;
                    continue;
                }
                // 获取左边和上边正方形边长的较小值
                let mut length: usize = 0;
                // 若两者有一个不存在，则就是 0
                if i > 0 && j > 0 {
                    length = cmp::min(dp[i - 1][j], dp[i][j - 1]);
                }
                // 状态转移
                dp[i][j] = length + (matrix[i - length][j - length] as usize - '0' as usize);

                result = cmp::max(result, dp[i][j])
            }
        }
        (result * result) as i32
    }
}
