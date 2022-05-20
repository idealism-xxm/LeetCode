// 链接：https://leetcode.com/problems/unique-paths-ii/
// 题意：给定一个 m * n 的 grid ，一个机器人在左上角 (0, 0) 处，
//      它每一次可以向右或者向下移动一格，
//      其中有些格子不能进入，在 grid 中标记为 1 ，
//      求移动到右下角 (m - 1, n - 1) 有多少种路径？


// 数据限制：
//  m == obstacleGrid.length
//  n == obstacleGrid[i].length
//  1 <= m, n <= 100
//  obstacleGrid[i][j] 是 0 或 1


// 输入： obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
// 输出： 2
// 解释： 有两种路径：
//          1. 向右 -> 向右 -> 向下 -> 向下
//          2. 向下 -> 向下 -> 向右 -> 向右

// 输入： obstacleGrid = [[0,1],[0,0]]
// 输出： 1
// 解释： 有一种路径： 向下 -> 向右


// 思路： DP
//
//      设 dp[i][j] 表示从 (0, 0) 到 (i - 1, j - 1) 的不同路径数，
//      这里 i 和 j 都比对应格子的下标大 1 ，
//      是为了方便处理第一行和第一列的情况，避免越界判断。
//
//      初始状态： dp[i][j] = 0; dp[0][1] = 1
//          注意，这里也可以令 dp[1][0] = 1 ，但两者不能同时为 1 ，
//          这样初始化也是为了方便后续处理，并能转移出 dp[1][1] = 1 。
//
//      状态转移方程：
//          1. obstacle_grid[i - 1][j - 1] == 0:
//              dp[i][j] = dp[i - 1][j] + dp[i][j - 1] ，
//              即可以从上边或左边的格子走过来
//          2. obstacle_grid[i - 1][j - 1] == 1:
//              dp[i][j] = 0 ，实际上不用处理，因为初始化时就是 0
//
//      最后， dp[m][n] 就是从 (0, 0) 到 (m - 1, n - 1) 的不同路径数。
//
//
//      时间复杂度：O(mn)
//          1. 需要遍历 dp 全部 O(mn) 个状态
//      空间复杂度：O(mn)
//          1. 需要维护一个大小为 O(mn) 的数组 dp


impl Solution {
    pub fn unique_paths_with_obstacles(obstacle_grid: Vec<Vec<i32>>) -> i32 {
        let (m, n) = (obstacle_grid.len(), obstacle_grid[0].len());
        // dp[i][j] 表示从 (0, 0) 到 (i - 1, j - 1) 的不同路径数，
        // 初始化均为 0
        let mut dp = vec![vec![0; n + 1]; m + 1];
        // 为了方便后续获得 dp[1][1] 为 1 ，这里将 dp[0][1] 设置为 1
        dp[0][1] = 1;
        // 遍历当前要走到的格子下标 (i, j)
        for i in 1..=m {
            for j in 1..=n {
                // 如果当前格子为空，则可以从上边和左边走过来
                if obstacle_grid[i - 1][j - 1] == 0 {
                    dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
                }
            }
        }

        // dp[m][n] 就是从 (0, 0) 到 (m - 1, n - 1) 的不同路径数
        dp[m][n]
    }
}
