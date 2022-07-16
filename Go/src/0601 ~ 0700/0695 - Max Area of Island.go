// 链接：https://leetcode.com/problems/baseball-game/
// 题意：给定一个 m * n 的矩阵 grid ，其中 1 表示陆地，0 表示水。
//      假设 grid 四周全被水环绕。
//
//      4 联通的陆地形成一座岛，求所有岛中最大岛的面积？


// 数据限制：
//  m == grid.length
//  n == grid[i].length
//  1 <= m, n <= 50
//  grid[i][j] 是 0 或 1


// 输入： grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
// 输出： 6
// 解释： 注意，答案不是 11 ，因为岛必须是 4 联通的。
//       0010000100000
//       0000000111000
//       0110100000000
//       0100110010100
//       0100110011100
//       0000000000100
//       0000000111000
//       0000000110000

// 输入： grid = [[0,0,0,0,0,0,0,0]]
// 输出： 0
// 解释： 没有岛，所以返回 0 。


// 思路： 递归/DFS
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
var DR = []int{-1, 0, 1, 0}
var DC = []int{0, 1, 0, -1}


func maxAreaOfIsland(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    // visited[r][c] 表示 (r, c) 的位置是否被访问过
    visited := make([][]bool, m)
    for i := range visited {
        visited[i] = make([]bool, n)
    }
    // 初始化最大岛面积为 0
    ans := 0
    for r := range grid {
        for c := range grid[r] {
            // 遍历计算每个位置岛岛面积，并更新 ans 的最大值
            area := dfs(grid, r, c, visited)
            ans = max(ans, area)
        }
    }

    return ans
}

func dfs(grid [][]int, r int, c int, visited [][]bool) int {
    // 如果位置不合法 或者 是水 或者 已访问过，则直接返回 0 。
    if r < 0 || r >= len(grid) || 
        c < 0 || c >= len(grid[r]) || 
        grid[r][c] == 0 || visited[r][c] {
            return 0
    }

    // 标记该位置已被遍历
    visited[r][c] = true
    // 当前陆地的面积为 1
    area := 1
    // 加上四个方向联通的陆地的面积
    for i := range DR {
        rr, cc := r + DR[i], c + DC[i]
        area += dfs(grid, rr, cc, visited)
    }

    return area
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
