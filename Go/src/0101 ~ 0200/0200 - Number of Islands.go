// 链接：https://leetcode.com/problems/number-of-islands/
// 题意：给定一个由 '0' （代表水） 和 '1' （代表陆地） 组成的地图，
//		求地图中岛屿的数量？

// 输入：
//	11110
//	11010
//	11000
//	00000
// 输出： 1

// 输入：
//	11000
//	11000
//	00100
//	00011
// 输出： 3

// 思路： DFS
//
//		遍历所有为访问过的格子，若为 1 ，则对结果加 1 ，
//		然后用 dfs 把对应岛屿的 1 都标记为已访问过即可
//
//		时间复杂度： O(m * n)
//		空间复杂度： O(m * n)

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	// 1. 初始化标记数组
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for r := 0; r < m; r++ {
		visited[r] = make([]bool, n)
	}

	result := 0
	// 2. 遍历格子
	for r, row := range grid {
		for c, cell := range row {
			// 如果未访问过，且当前是陆地
			if !visited[r][c] && cell == '1' {
				// 结果岛屿 + 1
				result++
				// 标记该岛屿已访问过
				dfs(grid, visited, r, c)
			}
		}
	}
	return result
}

func isValid(m, n int, r, c int) bool {
	return 0 <= r && r < m && 0 <= c && c < n
}

var dr = []int{-1, 0, 1, 0}
var dc = []int{0, 1, 0, -1}

func dfs(grid [][]byte, visited [][]bool, r, c int) {
	// 标记为已访问过
	visited[r][c] = true

	m, n := len(grid), len(grid[0])
	for i := 0; i < 4; i++ {
		rr, cc := r + dr[i], c + dc[i]
		// 如果相邻的下一个在范围内，没访问过，且是陆地，则继续递归访问
		if isValid(m, n, rr, cc) && !visited[rr][cc] && grid[rr][cc] == '1' {
			dfs(grid, visited, rr, cc)
		}
	}
}
