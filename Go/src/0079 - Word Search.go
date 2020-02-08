// 链接：https://leetcode.com/problems/word-search/
// 题意：给定一个字符矩阵，求一个单词 word 是否出现在其中？
//		（每个位置的字符最多可使用一次，当前字符必须在下一个字符的上下左右）

// board =
// [
//   ['A','B','C','E'],
//   ['S','F','C','S'],
//   ['A','D','E','E']
// ]
//
// word = "ABCCED", return true.
// word = "SEE", return true.
// word = "ABCB", return false.

// 思路：递归
//		尝试以每一个位置开始，递归判断是否满足题意即可

func exist(board [][]byte, word string) bool {
	// 尝试以每一个位置起始，有一个符合题意则直接返回 true
	for i := len(board) - 1; i >= 0; i-- {
		for j := len(board[i]) - 1; j >= 0; j-- {
			if dfs(board, word, i, j) {
				return true
			}
		}
	}
	return false
}

var dr = []int{0, 1, 0, -1}
var dc = []int{1, 0, -1, 0}

func dfs(board [][]byte, word string, r, c int) bool {
	if board[r][c] != word[0] {  // 当前字符和单词第一个字符不对应，直接返回 false
		return false
	}
	if len(word) == 1 {
		return true
	}
	tmp := board[r][c]
	board[r][c] = 0  // 标记当前字符已使用

	m, n := len(board), len(board[0])
	for i := 0; i < 4; i++ {  // 尝试四个方向是否满足后续的字符串
		rr, cc := r + dr[i], c + dc[i]
		// 如果下标在范围内，且后续的字符串满足题意，则直接返回
		if isValid(m, n, rr, cc) && dfs(board, word[1:], rr, cc) {
			return true
		}
	}

	board[r][c] = tmp  // 当前字符不再使用
	return false
}

func isValid(m, n int, r, c int) bool {
	return 0 <= r && r < m && 0 <= c && c < n
}
