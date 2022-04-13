// 链接：https://leetcode.com/problems/spiral-matrix-ii/
// 题意：给定一个正整数 n ，生成一个 n * n 的矩阵，
//      按照顺时针螺旋的填入 1 ~ n ^ 2 。


// 数据限制：
//  1 <= n <= 20


// 输入： n = 3
// 输出： [[1,2,3],[8,9,4],[7,6,5]]
// 解释： 螺旋矩阵如下：
//      [
//       [ 1, 2, 3 ],
//       [ 8, 9, 4 ],
//       [ 7, 6, 5 ],
//      ]

// 输入： n = 1
// 输出： 1


// 思路： 模拟
//
//      我们模拟螺旋矩阵设置的过程，
//      从 1 ~ n ^ 2 中遍历需要设置的数 num ，
//      每次将 matrix[r][c] 设置为 num ，
//      然后移动到下一个位置。
//
//      注意当超出当前层时，需要右转换方向。
//
//
//      初始时 r = 0, c = 0 ，表示我们从左上角开始设置；
//      方向 direction = 0 ，表示下一步向右移动。
//
//      每次移动时，先计算下一个可能的位置，
//      即 rr = r + DR[direction], cc = c + DC[direction] ，
//      当超出当前层时（即满足以下两个条件之一）时，
//      需要右转换方向（即 direction = (direction + 1) % 4 ）：
//          1. rr 或 cc 越界，此时说明超出了最外层
//          2. matrix[rr][cc] 不为 0 ，
//              此时说明超出了非最外层
//
//      注意我们是右转换方向，
//      即按照 右 -> 下 -> 左 -> 上 的顺序不断循环
//
//
//      时间复杂度：O(n ^ 2)
//          1. 需要遍历矩阵中全部 O(n ^ 2) 个位置
//      空间复杂度：O(1)
//          1. 只需要维护常数个额外变量


// 4 个方向的行下标改变量
static DR: [i32; 4] = [0, 1, 0, -1];
// 4 个方向的列下标改变量
static DC: [i32; 4] = [1, 0, -1, 0];


impl Solution {
    pub fn generate_matrix(n: i32) -> Vec<Vec<i32>> {
        let n = n as usize;
        // 定义 n * n 的结果矩阵，初始化都是 0
        let mut matrix = vec![vec![0; n]; n];
        // 定义最开始的行下标和列下标
        let (mut r, mut c) = (0, 0);
        // 定义最开始的方向
        let mut direction = 0;
        // 遍历要生成的所有数
        for num in 1..=n*n {
            // 将当前数字放入到 matrix[r][c] 处
            matrix[r][c] = num as i32;
            // 计算下一个数字所在位置的可能的行下标和列下标
            let rr = r as i32 + DR[direction];
            let cc = c as i32 + DC[direction];
            // 以下两种情况需要右转换方向：
            //  1. rr 或 cc 越界，此时说明超出了最外层
            //  2. matrix[rr][cc] 不为 0 ，
            //      此时说明超出了非最外层
            if rr < 0 || rr >= n as i32 || cc < 0 || cc >= n as i32 || matrix[rr as usize][cc as usize] != 0 {
                // 右转换方向，
                // 按照 右 -> 下 -> 左 -> 上 的顺序不断循环
                direction = (direction + 1) % 4;
            }
            // 计算当前下一个位置的真正行下标和列下标
            r += DR[direction] as usize;
            c += DC[direction] as usize;
        }

        matrix
    }
}
