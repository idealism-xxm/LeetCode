// 链接：https://leetcode.com/problems/number-of-islands/
// 题意：给定一个 m * n 的矩阵 grid ，其中 "1" 表示陆地， "0" 表示水。
//      假设 grid 四周全被水环绕。
//      4 联通的陆地形成一座岛，求岛的数量？


// 数据限制：
//  m == grid.length
//  n == grid[i].length
//  1 <= m, n <= 300
//  grid[i][j] 是 '0' 或 '1'


// 输入： grid = [["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]
// 输出： 1
// 解释： 所有陆地都是 4 联通的，所以只有一座岛。 
//       11110
//       11010
//       11000
//       00000

// 输入： grid = [["1","1","0","0","0"],["1","1","0","0","0"],["0","0","1","0","0"],["0","0","0","1","1"]]
// 输出： 3
// 解释： 前两行的陆地是 4 联通的，第三行和第四行的陆地分别是 4 联通的，总共有 3 座岛。
//       11000
//       11000
//       00100
//       00011


// 思路： 递归/DFS
//
//      本题是 LeetCode 695 的变形，同样可以使用 DFS 处理，
//      稍作变形那题的思路和代码，就可以 AC 本题。
//
//      为了维护每个位置的访问状态，我们需要定义一个 m * n 的二维数组 visited ，
//      其中 visited[r][c] 表示位置 (r, c) 是否被访问过。
//
//      对于每个位置 (r, c) 直接使用 dfs(grid, r, c, visited) 来计算岛屿的面积：
//          1. grid: 题目给定的矩阵，透传即可
//          2. r: 当前访问的位置的行号
//          3. c: 当前访问的位置的列号
//          4. visited: 访问状态的二维数组，会在 dfs 中进行修改
//
//      在 dfs 中，如果位置 (r, c) 不合法 或者 是水 或者 已经被访问过，则直接返回 0 。
//
//      否则，将 (r, c) 标记为已访问，并且开始计算岛屿面积，递归计算其 4 联通的位置。
//
//      如果 dfs 出的岛面积不为 0 ，则说明当前的岛是一个全新的岛，需要计入结果中。
//
//      时间复杂度：O(m * n)
//          1. 需要遍历 grid 中的全部 O(m * n) 个元素
//      空间复杂度：O(m * n)
//          1. 需要维护 visited 中全部 O(m * n) 个元素


// 每个方向的位置改变量
//  0: 向上走 1 步
//  1: 向右走 1 步
//  2: 向下走 1 步
//  3: 向左走 1 步
static DIRS: [(i32, i32); 4] = [(-1, 0), (0, 1), (1, 0), (0, -1)];


impl Solution {
    pub fn num_islands(grid: Vec<Vec<char>>) -> i32 {
        let (m, n) = (grid.len(), grid[0].len());
        // visited[r][c] 表示 (r, c) 的位置是否被访问过
        let mut visited = vec![vec![false; n]; m];
        // 初始化岛数量为 0
        let mut ans = 0;
        for r in 0..m {
            for c in 0..n {
                // 遍历计算每个位置的岛面积
                let area = Self::dfs(&grid, r, c, &mut visited);
                // 如果存在岛屿，则 ans 加 1
                if area > 0 {
                    ans += 1;
                }
            }
        }

        ans
    }

    fn dfs(grid: &Vec<Vec<char>>, r: usize, c: usize, visited: &mut Vec<Vec<bool>>) -> i32 {
        // 如果位置不合法 或者 是水 或者 已访问过，则直接返回 0 。
        // 注意： r, c 的类型为 usize ，所以会下溢，无需判断是否小于 0 。
        if r >= grid.len() || c >= grid[r].len() || grid[r][c] == '0' || visited[r][c] {
            return 0;
        }

        // 标记该位置已被遍历
        visited[r][c] = true;
        // 当前陆地的面积为 1
        let mut area = 1;
        // 加上四个方向联通的陆地的面积
        for (dr, dc) in DIRS.iter() {
            let (rr, cc) = (r as i32 + dr, c as i32 + dc);
            area += Self::dfs(grid, rr as usize, cc as usize, visited);
        }

        area
    }
}