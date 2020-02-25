// 链接：https://leetcode.com/problems/n-queens-ii/
// 题意：求解 n 皇后所有可行解的个数？

// 输入：4
// 输出：2
// 解释：总共有两个可行解
// [
//   [".Q..",  // Solution 1
//     "...Q",
//     "Q...",
//     "..Q."],
//
//   ["..Q.",  // Solution 2
//     "Q...",
//     "...Q",
//     ".Q.."]
// ]


// 思路：递归模拟即可（上一题的简化版）
//		每一层枚举时判断当前行、列、左斜和右斜是否能够放置

func totalNQueens(n int) int {
	// 初始化棋盘
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		board[i] = row
	}
	// 初始化 used 数组
	used := make([][]bool, 4)
	for i := 0; i < 4; i++ {
		used[i] = make([]bool, n * 3)
	}
	return dfs(board, n, 0, used)
}

func dfs(board [][]byte, n, row int, used [][]bool) int {
	if row >= n {
		return 1
	}

	// 收集所有可能结果
	result := 0
	for i := 0; i < n; i++ {
		if !isUsed(used, n, row, i) {
			setUsed(used, n, row, i, true)
			board[row][i] = 'Q'
			result += dfs(board, n, row + 1, used)
			board[row][i] = '.'
			setUsed(used, n, row, i, false)
		}
	}
	return result
}

func isUsed(used[][]bool, n, row, col int) bool {
	return  used[0][n + row] ||   // 行
			used[1][n + col] ||   // 列
			used[2][n + row + col] ||   // 左斜
			used[3][n + row - col]  // 右斜
}

func setUsed(used [][]bool, n, row, col int, value bool) {
	used[0][n + row] = value  // 行
	used[1][n + col] = value  // 列
	used[2][n + row + col] = value  // 左斜
	used[3][n + row - col] = value  // 右斜
}
