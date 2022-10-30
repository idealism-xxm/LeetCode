// 链接：https://leetcode.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/
// 题意：给定一个 m * n 的矩阵 grid ，每个单元格的值是 0 （空） 或 1 （障碍） 。
//      最多可以将 k 个障碍移除，使其变为空单元格。
//      每一步可以向四个方向移动到相邻的空单元格。
//      求从左上角 (0, 0) 到右下角 (m - 1, n - 1) 最少要走多少步？
//      如果无法从左上角走到右下角，则返回 -1 。


// 数据限制：
//  m == grid.length
//  n == grid[i].length
//  1 <= m, n <= 40
//  1 <= k <= m * n
//  grid[i][j] 是 0 或 1
//  grid[0][0] == grid[m - 1][n - 1] == 0


// 输入： grid = [[0,0,0],[1,1,0],[0,0,0],[0,1,1],[0,0,0]], k = 1
// 输出： 6
// 解释： 移除 (3, 2) 处的障碍，只需要 6 步就能从左上角走到右下角。
//       000           000
//       110           110
//       000    ->     000 
//       011           010
//       000           000

// 输入： grid = [[0,1,1],[1,1,1],[1,0,0]], k = 1
// 输出： -1
// 解释： 至少要移除两处障碍才能从左上角走到右下角


// 思路： BFS
//
//      本题在普通 BFS 的基础上，增加了可移除障碍的特性，并且限制了移除障碍的最大数。
//
//      所以我们也要在 visited 和 q 中维护剩余可移除障碍数这个信息，
//      以保证移除障碍数不超过 k 。
//
//      其中 visited[r][c][remain] 表示移动到 (r, c) 处时，
//      还能移除 remain 个障碍这个状态是否已经访问过。
//
//      q 中的每个元素维护坐标 (r, c) ，到该处的步数 steps ，以及剩余可移除障碍数 remain 。
//
//      在从 (r, c) 移动到 (rr, cc) 时，如果 (rr, cc) 有障碍，
//      则 remain 至少为 1 才能移除该障碍，否则无法移动到 (rr, cc) 。
//
//
//      时间复杂度：O(mnk)
//          1. 需要从左上角走到右下角，最差情况下需要遍历全部 O(mnk) 个状态
//      空间复杂度：O(mnk)
//          1. 需要维护一个大小为 O(mnk) 的三维数组 visited ，
//              用来记录每个状态是否被访问过
//          2. 需要维护队列 q ，最差情况下有 O(mnk) 个元素


use std::collections::VecDeque;


// 每个方向的位置改变量
//  0: 向上走 1 步
//  1: 向右走 1 步
//  2: 向下走 1 步
//  3: 向左走 1 步
static DIRS: [(i32, i32); 4] = [(-1, 0), (0, 1), (1, 0), (0, -1)];


impl Solution {
    pub fn shortest_path(grid: Vec<Vec<i32>>, k: i32) -> i32 {
        let (m, n) = (grid.len(), grid[0].len());
        // visited[r][c][remain] 表示移动到 (r, c) 处时，
        //  还能移除 remain 个障碍这个状态是否已经访问过
        let mut visited = vec![vec![vec![false; k as usize + 1]; n]; m];
        // q 维护 BFS 的队列
        let mut q = VecDeque::new();
        // 初始在 (0, 0) 处，走了 0 步，还能移除 k 个障碍
        q.push_back((0, 0, 0, k));
        visited[0][0][k as usize] = true;
        // 如果队列还有元素，则取出继续处理
        while let Some((r, c, steps, remain)) = q.pop_front() {
            // 如果走到了右下角，则此时的步数 steps 是最小的
            if r == m - 1 && c == n - 1 {
                return steps
            }

            // 遍历 4 个方向
            for &(dr, dc) in DIRS.iter() {
                // 计算该方向的下一个位置
                let rr = (r as i32 + dr) as usize;
                let cc = (c as i32 + dc) as usize;
                // 如果下一个位置合法，则准备移动到 (rr, cc)
                // （注意： r, c 的类型为 usize ，所以会下溢，无需判断是否小于 0 ）
                if rr < m && cc < n {
                    // 计算移动到 (rr, cc) 后还能移除多少个障碍
                    let next_remain = remain - grid[rr][cc];
                    // 如果剩余的移除障碍数合法 且 当前状态还未访问过，则可以移动到 (rr, cc)
                    if next_remain >= 0 && !visited[rr][cc][next_remain as usize] {
                        visited[rr][cc][next_remain as usize] = true;
                        q.push_back((rr, cc, steps + 1, next_remain));
                    }
                }
            }
        }

        // 此时无法从左上角走到右下角
        -1
    }
}
