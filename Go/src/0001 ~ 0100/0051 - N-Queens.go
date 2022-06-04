// 链接：https://leetcode.com/problems/n-queens/
// 题意：返回 n 皇后所有合法的放置方案。
//
//      合法的放置方案就是：任意两个皇后不能在同一行、同一列、
//      同一左斜线和同一右斜线上。


// 数据限制：
//  1 <= n <= 9


// 输入： n = 4
// 输出： [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
// 解释： 只有两种合法的放置方案，分别是：
//          .Q..        ..Q.
//          ...Q        Q...
//          Q...        ...Q
//          ..Q.        .Q..

// 输入： n = 1
// 输出： [["Q"]]
// 解释： 只有一种合法的放置方案，是： Q


// 思路： 递归/回溯/DFS
//
//      我们使用 used 数组来记录某一位置 (row, col) 三种情况的是否被占用：
//          1. used[0][col] 表示第 col 列是否被占用
//          2. used[1][row + col] 表示 (row, col) 所在的左斜线是否被占用
//          3. used[2][n + row - col] 表示 (row, col) 所在的右斜线是否被占用
//
//      同时初始化一个空棋盘 board （全部为 '.' ），
//      然后使用 dfs(row, board, used, ans) 回溯收集所有合法的放置方法，其中：
//          1. row: 当前该放置棋盘的第 row 行
//          2. board: 当前棋盘状态
//          3. used: 当前所有位置三种情况的占用情况
//          4. ans: 当前收集到的所有合法的放置方法
//
//      在 dfs 中，我们按照如下逻辑处理即可：
//          1. row == n: 已选放置完所有的皇后，
//              则当前放置方法 board 满足题意，收集到 ans 中，然后返回
//          2. row < n: 则还需要继续在第 row 行放置皇后，
//              遍历第 row 行要放置皇后的列 col ，
//              如果 (row, col) 三种情况都未被占用，则在此放置皇后，
//              并调用 dfs 继续回溯。
//
//
//      设 n 皇后共有 C 种合法的放置方案数。
//
//      时间复杂度：O(n! + C * n ^ 2)
//          1. 需要枚举全部 O(n!) 种可能的放置方案
//          2. 需要收集全部 O(C) 种合法的放置方案，
//             每次收集时都需要遍历棋盘全部 O(n ^ 2) 字符
//			3. 需要初始化 board 全部 O(n ^ 2) 字符
//			4. 需要初始化 used 全部 O(n) 个值
//      空间复杂度：O(C * n ^ 2)
//          1. 栈递归深度为 O(n)
//          2. 需要收集全部 O(C) 种合法的放置方案，
//             每次方案都需要存储棋盘全部 O(n ^ 2) 字符
//			3. 需要维护 board 全部 O(n ^ 2) 字符
//			4. 需要维护 used 全部 O(n) 个值


func solveNQueens(n int) [][]string {
	// 初始化棋盘，最开始没有放置任何 Queen ，都为 '.'
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	// used[0][col] 表示第 col 列是否被占用
	// used[1][row + col] 表示 (row, col) 所在的左斜线是否被占用
	// used[2][n + row - col] 表示 (row, col) 所在的右斜线是否被占用
	used := make([][]bool, 3)
	for i := 0; i < 3; i++ {
		used[i] = make([]bool, n << 1)
	}
	ans := make([][]string, 0)
	// 递归收集所有可能的放置结果
	dfs(0, board, used, &ans)

	return ans
}

func dfs(row int, board [][]byte, used [][]bool, ans *[][]string) {
	// 如果 row == n ，则说明当前 board 是一个合法的放置方法，收集入 ans 中
	n := len(board)
	if row == n {
		// 收集当前棋盘 board 对应的字符串结果 res ，然后放入 ans 中
		res := make([]string, n)
        for i := 0; i < n; i++ {
			res[i] = string(board[i])
		}
		*ans = append(*ans, res)
		return
	}

	// 枚举第 row 行放置皇后的列下标
	for col := 0; col < n; col++ {
		// 如果 (row, col) 没有被占用，则可以在此放置一个皇后
		if !isUsed(used, n, row, col) {
			// 标记 (row, col) 为已被占用
			setUsed(used, n, row, col, true)
			board[row][col] = 'Q'
			// 递归处理第 row + 1 行
			dfs(row + 1, board, used, ans)
			// 标记 (row, col) 为未被占用
			board[row][col] = '.'
			setUsed(used, n, row, col, false)
		}
	}
}

func isUsed(used [][]bool, n int, row int, col int) bool {
	// 如果 (row, col) 对应的任意一种情况已有皇后，则该位置已被占用
	return used[0][col] || used[1][row + col] || used[2][n + row - col]
}

func setUsed(used [][]bool, n int, row int, col int, value bool) {
	// 依次标记 (row, col) 对应的三种情况的占用值为 value 。
	// 其中： value = true 代表已经使用，value = false 代表未使用。
	used[0][col] = value
	used[1][row + col] = value
	used[2][n + row - col] = value
}
