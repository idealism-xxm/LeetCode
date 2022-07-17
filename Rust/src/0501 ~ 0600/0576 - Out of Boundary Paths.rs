// 链接：https://leetcode.com/problems/out-of-boundary-paths/
// 题意：有一个 m * n 的网格，有一个球初始在 (startRow, startColumn) 处，
//      每次可以将球移动到 4 个相邻的单元格中，最多可移动 maxMove 次。
//      求有多少种移动方案能将球移出网格？


// 数据限制：
//  1 <= m, n <= 50
//  0 <= maxMove <= 50
//  0 <= startRow < m
//  0 <= startColumn < n


// 输入： m = 2, n = 2, maxMove = 2, startRow = 0, startColumn = 0
// 输出： 6
// 解释： 移动 1 次： (0, 0) -> (0, -1)
//                  (0, 0) -> (-1, 0)
//       移动 2 次： (0, 0) -> (0, 1) -> (0, 2)
//                  (0, 0) -> (0, 1) -> (-1, 1)
//                  (0, 0) -> (1, 0) -> (1, -1)
//                  (0, 0) -> (1, 0) -> (2, 0)

// 输入： m = 1, n = 3, maxMove = 3, startRow = 0, startColumn = 1
// 输出： 12


// 思路： DP
//
//      由于移动的方向有 4 个，且无任何限制，所以无法按照某种顺序处理，
//      也就无法使用二维 DP ，需要增加一维使用三维 DP ，方便相邻移动次数的状态进行转移。
//
//      设 dp[i][r][c] 表示球移动 i 次到位置 (r, c) 的方案数。
//
//      如果移动不超过 maxMove - 1 次时，球能到网格边界，
//      那么下一次移动必定可以将球移出网格，所以我们只需要统计处理到达边界的方案数即可。
//
//      初始化： dp[0][r][c] = 0, dp[0][startRow][startColumn] = 1 。
//          即移动 0 次时，只有 (startRow, startColumn) 位置可能有球，方案数为 1 。
//
//      每次进行状态转移前，都需要统计合法的方案数：计入移动 i 次到达网格边界的方案数。
//      注意：球在网格四个角时，下一次移动有两种方案移出网格，所以需要统计两次。
//
//      状态转移：计算当前位置 (r, c) 可以移动到的 4 个位置 (rr, cc) ，
//              如果 (rr, cc) 合法，则有 dp[i + 1][rr][cc] += dp[i][r][c]
//
//      DP 常见的三种优化方式见 LeetCode 583 这题的思路，
//      本题只能采用滚动数组的方式进行优化，将空间复杂度从 O(N * mn) 优化为 O(mn) 。
//      因为第 dp[i][r][c] 以来 dp[i - 1] 中 4 个方向上的状态。
//
//
//      设 N 为 max_move 。
//
//      时间复杂度：O(N * mn)
//          1. 需要进行 O(n) 次移动，每次移动时，
//              都需要对全部 O(mn) 个单元格进行状态转移
//      空间复杂度：O(mn)
//          1. 需要维护两个大小为 O(mn) 的状态数组   


// 每个方向的位置改变量
//  0: 向上走 1 步
//  1: 向右走 1 步
//  2: 向下走 1 步
//  3: 向左走 1 步
static DIRS: [(i32, i32); 4] = [(-1, 0), (0, 1), (1, 0), (0, -1)];
const MOD: i32 = 1_000_000_007;


impl Solution {
    pub fn find_paths(m: i32, n: i32, max_move: i32, start_row: i32, start_column: i32) -> i32 {
        let (m, n) = (m as usize, n as usize);
        // dp[r][c] 表示球移动 i 次到位置 (r, c) 的方案数
        let mut dp = vec![vec![0; n as usize]; m as usize];
        // 初始时，移动 0 次，只有 (start_row, start_column) 位置可能有球，方案数为 1
        dp[start_row as usize][start_column as usize] = 1;
        let mut ans = 0;
        // 【注意】这里只统计转移到 max_move - 1 次移动，因为需要额外的一次移动移出网格
        for _ in 0..max_move {
            // 定义第 i + 1 次时网格的状态，初始化都为 0
            let mut next_dp = vec![vec![0; n as usize]; m as usize];
            for r in 0..m {
                for c in 0..n {
                    // 如果第 i 步已到达边界，则下一步能通过下面 4 个对应的边界移出网格。
                    // 【注意】必须每个边界都判断一次，防止遗漏。
                    //      因为此时如果在网格的四个角，那么下一步有两种方案移出网格。
                    if r == 0 {
                        ans = (ans + dp[r][c]) % MOD;
                    }
                    if r == m - 1 {
                        ans = (ans + dp[r][c]) % MOD;
                    }
                    if c == 0 {
                        ans = (ans + dp[r][c]) % MOD;
                    }
                    if c == n - 1 {
                        ans = (ans + dp[r][c]) % MOD;
                    }
                    // 当前 (r, c) 处的状态转移至第 i + 1 次移动的 (rr, cc) 的状态
                    for (dr, dc) in DIRS.iter() {
                        let rr = r as i32 + dr;
                        let cc = c as i32 + dc;
                        // 位置 (rr, cc) 合法时才能转移成功
                        if (0..m as i32).contains(&rr) && (0..n as i32).contains(&cc) {
                            next_dp[rr as usize][cc as usize] = (next_dp[rr as usize][cc as usize] + dp[r][c]) % MOD;
                        }
                    }
                }
            }
            // 滚动数组切换，下次处理第 i + 1 次移动的情况
            dp = next_dp;
        }

        ans
    }
}
