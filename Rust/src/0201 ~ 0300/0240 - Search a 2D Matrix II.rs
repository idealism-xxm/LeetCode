// 链接：https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
// 题意：给定一个排序后的 m * n 的矩阵（每行递增，每列递增），
//      判断指定的数字 target 是否存在？


// 数据限制：
//  m == matrix.length
//  n == matrix[i].length
//  1 <= n, m <= 300
//  -(10 ^ 9) <= matrix[i][j] <= 10 ^ 9
//  每一行的数字都是递增的
//  每一列的数字都是递增的
//  -(10 ^ 9) <= target <= 10 ^ 9


// 输入： matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
// 输出： true
// 解释： [1,   4,  7, 11, 15]
//       [2,   5,  8, 12, 19]
//       [3,   6,  9, 16, 22]
//       [10, 13, 14, 17, 24]
//       [18, 21, 23, 26, 30]

// 输入： matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
// 输出： false


// 思路2： 减治
//
//      我们从左下角 (r, c) 开始不断进行判断：
//            1. target < matrix[r][c]: 则 r 这一行都不满足题意，
//                可以转换为求解子矩阵 [:r-1][c:]
//            2. target > matrix[r][c]: 则 c 这一列都不满足题意，
//                可以转换为求解子矩阵 [:r][c+1:]
//            3. target == matrix[r][c]: 直接返回 true
//
//        最后矩阵为空仍未找到答案，则返回 false
//
//        时间复杂度： O(m + n)
//          1. 最差情况下，需要遍历全部 O(m) 行的各一个数字
//          2. 最差情况下，需要遍历全部 O(n) 列的各一个数字
//        空间复杂度： O(1)
//          1. 只需要使用常数个额外变量


impl Solution {
    pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
        let (m, n) = (matrix.len(), matrix[0].len());
		// 从左下角 (m - 1, 0) 处开始减治
        let (mut r, mut c) = (m - 1, 0);
        // 若矩阵中还有元素，则继续减治处理
		// 【注意】这里 r 的类型是 usize ，所以小于 0 时会下溢到最大值，
		//		所以用 r < m 替代 r >= 0 进行判断
        while r < m && c < n {
            if target < matrix[r][c] {
                // r 这一行都不满足题意
                r -= 1;
            } else if target > matrix[r][c] {
                // c 这一列都不满足题意
                c += 1;
            } else {
                // target == matrix[r][c] ，直接返回 true
                return true;
            }
        }
        // 矩阵中不存在 target
        false
    }
}
