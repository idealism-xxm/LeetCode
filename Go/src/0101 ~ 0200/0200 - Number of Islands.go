// 链接：https://leetcode.com/problems/number-of-islands/
// 题意：给定一个 m * n 的矩阵 grid ，其中 "1" 表示陆地， "0" 表示水。
//      假设 grid 四周全被水环绕。
//      4 联通的陆地形成一座岛，求岛的数量？


// 数据限制：
//  m == grid.length
//	n == grid[i].length
//	1 <= m, n <= 300
//	grid[i][j] 是 '0' 或 '1'


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
var DR = []int{-1, 0, 1, 0}
var DC = []int{0, 1, 0, -1}


func numIslands(grid [][]byte) int {
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
            // 遍历计算每个位置的岛面积
            area := dfs(grid, r, c, visited)
            // 如果存在岛屿，则 ans 加 1
            if area > 0 {
                ans += 1
            }
        }
    }

    return ans
}

func dfs(grid [][]byte, r int, c int, visited [][]bool) int {
    // 如果位置不合法 或者 是水 或者 已访问过，则直接返回 0
    if r < 0 || r >= len(grid) || 
        c < 0 || c >= len(grid[r]) || 
        grid[r][c] == '0' || visited[r][c] {
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
