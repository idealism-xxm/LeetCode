// 链接：https://leetcode.com/problems/toeplitz-matrix/
// 题意：给定一个矩阵 matrix ，判断是否为常对角矩阵？
//
//      常对角矩阵的每一条从左上到右下的对角线上的数都相同。


// 数据限制：
//  m == matrix.length
//  n == matrix[i].length
//  1 <= m, n <= 20
//  0 <= matrix[i][j] <= 99


// 输入： matrix = [[1,2,3,4],[5,1,2,3],[9,5,1,2]]
// 输出： true
// 解释： 1234
//       5123
//       9512
//
//       所有的对角线为： [9], [5,5], [1,1,1], [2,2,2], [3,3], [4]
//       每一条对角线上的数都相同。

// 输入： matrix = [[1,2],[2,2]]
// 输出： false
// 解释： 对角线 [1,2] 上的数不相同。


// 思路： 模拟
//
//      等式具有传递性，即 a == b, b == c ，那么必定有 a == c 。
//
//      所以我们可以只用比较同一对角线上相邻的两个数的是否相同。
//
//      为了方便处理边界情况，我们遍历除最后一行和最后一列的所有数，
//      判断 matrix[r][c] 与 matrix[r + 1][c + 1] 是否相同。
//
//      
//      时间复杂度：O(mn)
//          1. 需要遍历 matrix 中全部 O(mn) 个数
//      空间复杂度：O(1)
//          1. 只需要使用常数个额外变量即可


impl Solution {
    pub fn is_toeplitz_matrix(matrix: Vec<Vec<i32>>) -> bool {
        let (m, n) = (matrix.len(), matrix[0].len());
        // 遍历除最后一行和最后一列的所有数，方便处理边界情况
        for r in 0..m-1 {
            for c in 0..n-1 {
                // 如果该数与其右下角的数不同，则该对角上存在不同的数字
                if matrix[r][c] != matrix[r + 1][c + 1] {
                    return false;
                }
            }
        }

        // 此时所有对角线上的数字都相同
        true
    }
}
