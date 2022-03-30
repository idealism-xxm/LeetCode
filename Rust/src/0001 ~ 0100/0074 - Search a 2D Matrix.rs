// 链接：https://leetcode.com/problems/search-a-2d-matrix/
// 题意：给定一个 m * n 的整数矩阵 matrix ，每一行从左到右升序，
//		每一行的第一个数字比上一行的最后一个数字大，
//		判断 target 是否在这个矩阵中？


// 数据限制：
//  m == matrix.length
//  n == matrix[i].length
//  1 <= m, n <= 100
//  -(10 ^ 4) <= matrix[i][j], target <= 10 ^ 4


// 输入： matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// 输出： true
// 解释： target 在 matrix 中
// [
//   [1, (3),  5,  7],
//   [10, 11, 16, 20],
//   [23, 30, 34, 50]
// ]

// 输入： matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
// 输出： false
// 解释： target 不在 matrix 中
// [
//   [1,   3,  5,  7],
//   [10, 11, 16, 20],
//   [23, 30, 34, 50]
// ]


// 思路： 一次二分
//
//      这个矩阵每行从左到右升序，每行的第一个数字比上一行的最后一个数字大，
//      所以我们把这个矩阵看成一个数组，那么这个数组就是升序排序的。
//
//      这样我们就可以使用一次二分就能判断 target 是否在这个矩阵中。
//
//      二分区间初始化为 [0, m * n - 1] ，二分结束条件为 l < r 。
//      每次二分时计算 mid 位置：
//          1. matrix[mid / n][mid % n] < target: 
//              target 存在的话必定在 [mid + 1, r] 之间，
//              令 l = mid + 1
//          2. matrix[mid / n][mid % n] >= target: 
//              target 存在的话必定在 [l, mid] 之间，
//              令 r = mid
//
//      二分结束时， matrix[l / n][l % n] 是大于等于 target 的第一个数，
//      只有 matrix[l / n][l % n] == target 时，
//      target 才存在于矩阵中。
//
//      时间复杂度：O(log(mn))
//          1. 需要二分区间 [0, m * n - 1] ，时间复杂度为 O(log(mn))
//      空间复杂度：O(1)
//          1. 只需要维护常数个额外变量


impl Solution {
    pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
        // 矩阵行数
        let mut m = matrix.len();
        // 矩阵列数
        let mut n = matrix[0].len();

        // 二分区间左边界初始化为 0
        let mut l = 0;
        // 二分区间右边界初始化为 m * n - 1
        let mut r = m * n - 1;
        // 当二分区间至少还存在 2 个数时，继续二分
        while l < r {
            // 计算区间中点
            let mid = (l + r) >> 1;
            if matrix[mid / n][mid % n] < target {
                // 如果中点对应的数小于 target ，
                // 那么 target 存在的话，必定在 [mid + 1, r] 内
                l = mid + 1;
            } else {
                // 否则认为 target 在 [l, mid] 内
                // （注意这里不会取到 mid - 1 ，
                // 因为二分终止条件是 l < r ）
                r = mid;
            }
        }

        // 现在 matrix[l / n][l % n] 是大于等于 target 的第一个数，
        // 只有 matrix[l / n][l % n] == target 时，
        // target 才存在于矩阵中。
        matrix[l / n][l % n] == target
    }
}
