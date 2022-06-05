// 链接：https://leetcode.com/problems/n-queens-ii/
// 题意：返回 n 皇后所有合法的放置方案数。
//
//      合法的放置方案就是：任意两个皇后不能在同一行、同一列、
//      同一左斜线和同一右斜线上。


// 数据限制：
//  1 <= n <= 9


// 输入： n = 4
// 输出： 2
// 解释： 只有两种合法的放置方案，分别是：
//          .Q..        ..Q.
//          ...Q        Q...
//          Q...        ...Q
//          ..Q.        .Q..

// 输入： n = 1
// 输出： 1
// 解释： 只有一种合法的放置方案，是： Q


// 思路： 递归/回溯/DFS
//
//      我们使用 used 数组来记录某一位置 (row, col) 三种情况的是否被占用：
//          1. used[0][col] 表示第 col 列是否被占用
//          2. used[1][row + col] 表示 (row, col) 所在的左斜线是否被占用
//          3. used[2][n + row - col] 表示 (row, col) 所在的右斜线是否被占用
//
//      使用 dfs(row, used) 回溯统计所有合法的放置方法数，其中：
//          1. row: 当前该放置棋盘的第 row 行
//          2. used: 当前所有位置三种情况的占用情况
//      返回值为当前状态下所有合法的放置方案数。
//
//      在 dfs 中，我们按照如下逻辑处理即可：
//          1. row == n: 已选放置完所有的皇后，则当前放置方法满足题意，返回 1
//          2. row < n: 则还需要继续在第 row 行放置皇后，
//              遍历第 row 行要放置皇后的列 col ，
//              如果 (row, col) 三种情况都未被占用，则在此放置皇后，
//              并调用 dfs 继续回溯，统计所有回溯的方案数之和并返回。
//
//
//      时间复杂度：O(n!)
//          1. 需要枚举全部 O(n!) 种可能的放置方案
//          2. 需要初始化 used 全部 O(n) 个值
//      空间复杂度：O(n)
//          1. 栈递归深度为 O(n)
//          2. 需要维护 used 全部 O(n) 个值


impl Solution {
    pub fn total_n_queens(n: i32) -> i32 {
        let n = n as usize;
        // used[0][col] 表示第 col 列是否被占用
        // used[1][row + col] 表示 (row, col) 所在的左斜线是否被占用
        // used[2][n + row - col] 表示 (row, col) 所在的右斜线是否被占用
        let mut used = vec![vec![false; n << 1]; 3];
        // 递归统计所有可能的放置结果
        Self::dfs(0, n, &mut used)
    }

    fn dfs(row: usize, n: usize, used: &mut Vec<Vec<bool>>) -> i32 {
        // 如果 row == n ，则说明当前 board 是一个合法的放置方法，直接返回 1
        if row == n {
            return 1;
        }

        let mut ans = 0;
        // 枚举第 row 行放置皇后的列下标
        for col in 0..n {
            // 如果 (row, col) 没有被占用，则可以在此放置一个皇后
            if !Self::is_used(used, n, row, col) {
                // 标记 (row, col) 为已被占用
                Self::set_used(used, n, row, col, true);
                // 递归处理第 row + 1 行
                ans += Self::dfs(row + 1, n, used);
                // 标记 (row, col) 为未被占用
                Self::set_used(used, n, row, col, false);
            }
        }

        ans
    }

    fn is_used(used: &mut Vec<Vec<bool>>, n: usize, row: usize, col: usize) -> bool {
        // 如果 (row, col) 对应的任意一种情况已有皇后，则该位置已被占用
        return used[0][col] || used[1][row + col] || used[2][n + row - col];
    }

    fn set_used(used: &mut Vec<Vec<bool>>, n: usize, row: usize, col: usize, value: bool) {
        // 依次标记 (row, col) 对应的三种情况的占用值为 value 。
        // 其中： value = true 代表已经使用，value = false 代表未使用。
        used[0][col] = value;
        used[1][row + col] = value;
        used[2][n + row - col] = value;
    }
}
