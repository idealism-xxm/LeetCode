// 链接：https://leetcode.com/problems/surrounded-regions/
// 题意：给定一个矩阵，包含字母 `X` 和 `O` ，将完全被 `X` 包围起来的 `O` 变成 `X` ？

// 输入：
// X X X X
// X O O X
// X X O X
// X O X X
// 输出：
// X X X X
// X X X X
// X X X X
// X O X X
// 解释：
// 		任何在边界上或者与边界上 `O` 相连的 `O` 不算被 `X` 包围

// 思路： DFS + 遍历
//
//		先用 DFS 将边界上及其相连的 `O` 改为 `*` ，
//		然后遍历修改：将 `O` 改为 `X` ，将 `*` 给为 `O`
//
//		时间复杂度： O(n ^ 2)
//		空间复杂度： O(n ^ 2)

func solve(board [][]byte)  {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	// 将边界上及相连的 `O` 改为 `*`
	n, m := len(board), len(board[0])
	for r := 0; r < n; r++ {
		dfs(board, r, 0)
		dfs(board, r, m - 1)
	}
	for c := 0; c < m; c++ {
		dfs(board, 0, c)
		dfs(board, n - 1, c)
	}

	// 将 `O` 改为 `X` ，将 `*` 给为 `O`
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			} else if board[r][c] == '*' {
				board[r][c] = 'O'
			}
		}
	}
}

var dr = []int{0, 1, 0, -1}
var dc = []int{1, 0, -1, 0}

func isValid(board [][]byte, r, c int) bool {
	return 0 <= r && r < len(board) && 0 <= c && c < len(board[0])
}

func dfs(board [][]byte, r, c int) {
	if board[r][c] != 'O' {
		return
	}

	board[r][c] = '*'
	for i := 0; i < 4; i++ {
		rr, cc := r + dr[i], c + dc[i]
		if isValid(board, rr, cc) {
			dfs(board, rr, cc)
		}
	}
}
