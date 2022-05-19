// 链接：https://leetcode.com/problems/longest-increasing-path-in-a-matrix/
// 题意：给定一个矩阵 m * n 的 matrix ，返回其最长上升路径的长度。
//
//      对于每个单元格，每次可以向上下左右四个方向移动，不能对角移动或越界。


// 数据限制：
//  m == matrix.length
//  n == matrix[i].length
//  1 <= m, n <= 200
//  0 <= matrix[i][j] <= 2 ^ 31 - 1


// 输入： matrix = [[9,9,4],[6,6,8],[2,1,1]]
// 输出： 4
// 解释： 最长上升路径是 [1, 2, 6, 9]
//
//       9   9  4
//       ↑
//       6   6  8
//       ↑
//       2 ← 1  1

// 输入： matrix = [[3,4,5],[3,2,6],[2,2,1]]
// 输出： 4
// 解释： 最长上升路径是 [3, 4, 5, 6] ，注意不能对角移动。
//
//       3 → 4 → 5
//               ↓
//       3   2   6
//
//       2   2   1

// 输入： matrix = [[1]]
// 输出： 1


// 思路： DP + 排序
//
//      这道题其实不容易想到 DP ，
//      因为常见的 DP 都可以直接按照题目给定数据的顺序处理，
//      例如：位置上从左到右，从上到下，从左上到右下等。
//
//      本题支持 4 个方向转移状态，所以不能直接用 DP 来解决，
//      但本题也加了其他限制，所以我们可以自定顺序，
//      即可以通过值从小到大的顺序处理，来进行状态转移。
//
//      设 dp[r][c] 表示以 (r, c) 为终点的最长上升路径的长度，
//      由于上升路径的中的值是严格单调递增的，
//      所以我们只要保证在处理 (r, c) 之前，
//      所有值小于 matrix[r][c] 的单元格都已处理完成，
//      那么 dp[r][c] 就能通过相邻的单元格的 dp 值转移得到。
//
//      matrix 不会改变，所以我们最开始就将 matrix 转成单元格数组 cells ，
//      其中 cells[i] = (matrix[r][c], r, c) ，
//      然后按照 matrix[r][c] 升序排序。
//
//      初始化令所有 dp[r][c] = 1 ，即都能以自身为终点。
//
//      然后遍历 cells 进行状态转移，此时能保证在处理 (r, c) 之前，
//      所有值小于 matrix[r][c] 的单元格都已处理完成。
//
//
//      时间复杂度：O(mn * log(mn))
//          1. 需要收集 matrix 中全部 O(mn) 个元素到数组 cells 中
//          2. 需要对数组 cells 进行排序，时间复杂度为 O(mn * log(mn))
//          3. 需要遍历数组 cells 全部 O(mn) 个元素
//      空间复杂度：O(mn)
//          1. 需要收集 matrix 中全部 O(mn) 个元素到数组 cells 中
//          2. 需要维护一个大小为 O(mn) 的数组 dp


// 每个方向的位置改变量
//  0: 向上走 1 步
//  1: 向右走 1 步
//  2: 向下走 1 步
//  3: 向左走 1 步
var dr = []int{-1, 0, 1, 0}
var dc = []int{0, 1, 0, -1}


func longestIncreasingPath(matrix [][]int) int {
    m, n := len(matrix), len(matrix[0])
    // 将矩阵转换为单元格数组
    cells := CellSlice(make([]*Cell, 0, m * n))
    // dp[r][c] 表示以单元格 (r, c) 为终点的最长上升路径的长度
    dp := make([][]int, m)
    for r := 0; r < m; r++ {
        dp[r] = make([]int, n)
        for c := 0; c < n; c++ {
            // 初始化都为 1 ，即都能以自身为终点
            dp[r][c] = 1
            // 将 matrix[r][c] 转换成 Cell 放入到 cells 中
            cells = append(cells, &Cell{matrix[r][c], r, c})
        }
    }
    // 按照单元格的值从小到大排序
    sort.Sort(cells)

    // ans 维护当前最长上升路径的长度
    ans := 0
    // 按照单元格的值从小到大遍历，
    // 保证处理 (r, c) 时，值小于 v 的单元格已经处理过了
    for _, cell := range cells {
        v, r, c := cell.v, cell.r, cell.c
        for i := 0; i < 4; i++ {
            // 计算相邻单元格的坐标
            rr := r + dr[i]
            cc := c + dc[i]
            if 0 <= rr && rr < m &&
                0 <= cc && cc < n &&
                matrix[rr][cc] < v {
                // 如果 (rr, cc) 的值比 (r, c) 的值小，
                // 则 (r, c) 的前一个单元格可以是 (rr, cc) ，
                // 即 dp[r][c] 可以由 dp[rr][cc] + 1 转移而来
                dp[r][c] = max(dp[r][c], dp[rr][cc] + 1)
            }
        }
        // 更新 ans
        ans = max(ans, dp[r][c])
    }

    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

type Cell struct {
    v, r, c int
}

type CellSlice []*Cell

func (s CellSlice) Len() int {
    return len(s)
}

func (s CellSlice) Less(i, j int) bool {
    return s[i].v < s[j].v
}

func (s CellSlice) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
