// 链接：https://leetcode.com/problems/number-of-submatrices-that-sum-to-target/
// 题意：给定一个矩阵 matrix 和一个整数 target ，求有多少个非空子矩阵的和为 target ？


// 数据限制：
//  1 <= matrix.length <= 100
//  1 <= matrix[0].length <= 100
//  -1000 <= matrix[i] <= 1000
//  -(10 ^ 8) <= target <= 10 ^ 8


// 输入： matrix = [[0,1,0],[1,1,1],[0,1,0]], target = 0
// 输出： 4
// 解释： 0 1 0
//       1 1 1
//       0 1 0
//
//       只有 4 个角上的 1 x 1 子矩阵的和为 0 。


// 输入： matrix = [[1,-1],[-1,1]], target = 0
// 输出： 5
// 解释： 1 -1
//      -1  1
//
//       2 个 1 x 2 子矩阵的和为 0 ，
//       2 个 2 x 1 子矩阵的和为 0 ，
//       1 个 2 x 2 子矩阵的和为 0 。

// 输入： matrix = [[904]], target = 0
// 输出： 0


// 思路： 前缀和 + Map 
//
//      最朴素的想法就是先在 O(mn) 内计算矩阵的前缀和，
//      枚举所有 O(m ^ 2 * n ^ 2) 个子矩阵，
//      每次运用前缀和在 O(1) 内求出子矩阵的和，然后判断是否要进行统计。
//
//      这样的时间复杂度为 O(m ^ 2 * n ^ 2) ，在题目给定的数据范围内必定无法通过，
//      所以需要想办法优化掉一维的枚举。
//
//
//      我们可以注意到题目只让我们求和为 target 的子矩阵数量，不需要具体的子矩阵，
//      那么就能通过 map 进行优化处理。
//
//      这一点与 LeetCode 560 相同，所以可以采用前缀和 + Map 的方法统计数量。
//
//      其实我们只要枚举矩阵的上边界 i 和下边界 j ，
//      再维护子矩阵 (i, 0, j, n - 1) 的列和 col_sum 。
//
//      这时子矩阵的上下边界都已确定，且每一列的和都已知，
//      相当于给定一个长度为 n 的数组，求和为 target 的子数组的个数。
//
//      这个子问题就是 LeetCode 560 这题，可以直接服用相关代码计算即可。
//
//
//      时间复杂度：O(m ^ 2 * n)
//          1. 需要枚举子矩阵上边界 i ，下边界 j 和右边界 k ，
//              时间复杂度为 O(m ^ 2 * n)
//      空间复杂度：O(n)
//          1. 需用维护子矩阵 (i, 0, j, n - 1) 全部 O(n) 列的和
//          2. 需要维护子矩阵 (i, 0, j, k) 的不同和的个数，最差情况下有 O(n) 个不同和


use std::collections::HashMap;
use std::ops::AddAssign;


impl Solution {
    pub fn num_submatrix_sum_target(matrix: Vec<Vec<i32>>, target: i32) -> i32 {
        // ans 维护满足题意的子矩阵数量
        let mut ans = 0;
        let (m, n) = (matrix.len(), matrix[0].len());
        // 枚举子矩阵上边界下标 i
        for i in 0..m {
            // col_sum[k] 表示子矩阵 (i, 0, j, n - 1) 中第 k 列的和，
            // 即 col_sum[k] = matrix[i][k] + ... + matrix[j][k]
            let mut col_sum = vec![0; n];
            // 枚举子矩阵下边界下标 j
            for j in i..m {
                // 计入下边界上的数，更新子矩阵每一列的和
                for k in 0..n {
                    col_sum[k] += matrix[j][k];
                }

                // 计入 子问题——和为 target 的子数组 的数量
                ans += Self::subarray_sum(&col_sum, target);
            }
        }

        ans
    }

    fn subarray_sum(nums: &Vec<i32>, k: i32) -> i32 {
        // pre_sum_to_cnt[pre_sum] 表示前缀和 pre_sum 出现的次数
        let mut pre_sum_to_cnt = HashMap::new();
        // 最开始初始化前缀和 pre_sum[-1] = 0 出现过一次，
        // 即假设存在一个和为 0 的空子数组，为了方便统计和为 k 的子数组 nums[..=j] 的情况
        pre_sum_to_cnt.insert(0, 1);
        // 维护当前的前缀和
        let mut pre_sum = 0;
        // 维护满足题意的子数组个数
        let mut ans = 0;
        // 遍历数组
        for num in nums {
            // 前缀和加上当前的数字
            pre_sum += num;
            //  前缀和 pre_sum - k 出现的次数就是以当前数为结尾的和为 k 的子数组个数。
            //
            //  子数组 nums[i..=j] 的和为 pre_sum[j] - pre_sum[i - 1] ，
            //  那么根据题意，我们需要让这个和为 k ，
            //  即 pre_sum[j] - pre_sum[i - 1] = k
            //      => pre_sum[j] - k = pre_sum[i - 1]
            //  也就是要找到在 j 之前，前缀和为 pre_sum[j] - k 数量
            ans += pre_sum_to_cnt.get(&(pre_sum - k)).unwrap_or(&0);
            // 当前前缀和 pre_sum 的出现次数 +1
            pre_sum_to_cnt.entry(pre_sum).or_insert(0).add_assign(1);
        }

        ans
    }
}
