// 链接：https://leetcode.com/problems/sudoku-solver/
// 题意：给定一个有唯一解的合法数独的初始状态，返回该数独的解。
//      合法数独的状态如下：
//      1. 每一行出现且仅出现 1-9 的每一个数
//      2. 每一列出现且仅出现 1-9 的每一个数
//      3. 每一个九宫格出现且仅出现 1-9 的每一个数

// 输入：
// [
//  ["5","3",".",".","7",".",".",".","."],
//  ["6",".",".","1","9","5",".",".","."],
//  [".","9","8",".",".",".",".","6","."],
//  ["8",".",".",".","6",".",".",".","3"],
//  ["4",".",".","8",".","3",".",".","1"],
//  ["7",".",".",".","2",".",".",".","6"],
//  [".","6",".",".",".",".","2","8","."],
//  [".",".",".","4","1","9",".",".","5"],
//  [".",".",".",".","8",".",".","7","9"]
//]
// 输出：
// [
//  ["5","3","4","6","7","8","9","1","2"],
//  ["6","7","2","1","9","5","3","4","8"],
//  ["1","9","8","3","4","2","5","6","7"],
//  ["8","5","9","7","6","1","4","2","3"],
//  ["4","2","6","8","5","3","7","9","1"],
//  ["7","1","3","9","2","4","8","5","6"],
//  ["9","6","1","5","3","7","2","8","."],
//  ["2","8","7","4","1","9","6","3","5"],
//  ["3","4","5","2","8","6","1","7","9"]
//]

// 思路：递归模拟即可，使用上一题的判断函数可以简化为只用写递归
//      可用优化方法：每次选下一个格子时，要选可用数字最少的，以减少无用递归
//      当然还可以使用 Dancing Links 算法求解

func solveSudoku(board [][]byte) {
    var remainGrids [][]int
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] == '.' {
                remainGrids = append(remainGrids, []int{i, j})  // 收集所有待填充格子
            }
        }
    }
    dfs(board, remainGrids)
}

func dfs(board [][]byte, remainGrids [][]int) bool {
    if len(remainGrids) == 0 { // 若没有剩余的格子需要填充，则已解出数独
        return true
    }

    r, c := remainGrids[0][0], remainGrids[0][1]
    for i := '1'; i <= '9'; i++ {
        board[r][c] = byte(i) // 假设当前格子放 i
        squareR, squareC := r / 3 * 3, c / 3 * 3
        if isValidRow(board, r) &&
            isValidColumn(board, c) &&
            isValidSquare(board, squareR, squareC) { // 如果当前格子所在行、列和九宫格均合法，则递归调用
            if dfs(board, remainGrids[1:]) { // 如果已解出数独，则返回 true
                return true
            }
        }
        board[r][c] = '.' // 当前位置重置
    }
    return false // 还未解出，返回 false
}

func isValidRow(board [][]byte, i int) bool {
    exists := [127]bool{}
    for j := 0; j < 9; j++ {
        if board[i][j] == '.' { // 如果是空，则继续
            continue
        }
        if exists[board[i][j]] { // 如果当前数字已出现，则当前行不合法
            return false
        }
        exists[board[i][j]] = true // 标记当前数字已出现
    }
    return true // 所有数字否不重复，则当前行合法
}

func isValidColumn(board [][]byte, j int) bool {
    exists := [127]bool{}
    for i := 0; i < 9; i++ {
        if board[i][j] == '.' { // 如果是空，则继续
            continue
        }
        if exists[board[i][j]] { // 如果当前数字已出现，则当前列不合法
            return false
        }
        exists[board[i][j]] = true // 标记当前数字已出现
    }
    return true // 所有数字否不重复，则当前列合法
}

func isValidSquare(board [][]byte, r, c int) bool {
    exists := [127]bool{}
    for i := r + 2; i >= r; i-- {
        for j := c + 2; j >= c; j-- {
            if board[i][j] == '.' { // 如果是空，则继续
                continue
            }
            if exists[board[i][j]] { // 如果当前数字已出现，则当前九宫格不合法
                return false
            }
            exists[board[i][j]] = true // 标记当前数字已出现
        }
    }
    return true // 所有数字否不重复，则当前九宫格合法
}
