// 链接：https://leetcode.com/problems/path-with-minimum-effort/
// 题意：给定一个二维数组 heights ，其中 heights[r][c] 表示 (r, c) 处的高度，
//      现在需要从左上角 (0, 0) 处走到右下角 (m - 1, n - 1) 处，
//      每次可以向周围四个方向走。
//
//      设一条路径上相邻两点的高度差绝对值的最大值为 diff ，
//      求所有路径中 diff 的最小值？


// 数据限制：
//  rows == heights.length
//  columns == heights[i].length
//  1 <= rows, columns <= 100
//  1 <= heights[i][j] <= 10 ^ 6


// 输入： heights = [[1,2,2],[3,8,2],[5,3,5]]
// 输出： 2
// 解释： (1)  2   2
//       (3)  8   2
//       (5) (3) (5)
//       
//       路径 [1,3,5,3,5] 上相邻两点的高度差绝对值的最大值为 2 ，
//       这是所有路径中的最小值。

// 输入： heights = [[1,2,3],[3,8,4],[5,3,5]]
// 输出： 1
// 解释： (1) (2) (3)
//        3   8  (4)
//        5   3  (5)
//
//       路径 [1,2,3,4,5] 上相邻两点的高度差绝对值的最大值为 1 ，
//       这是所有路径中的最小值。

// 输入： heights = [[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]
// 输出： 0
// 解释： (1)  2  (1) (1) (1)
//       (1)  2  (1)  2  (1)
//       (1)  2  (1)  2  (1)
//       (1)  2  (1)  2  (1)
//       (1) (1) (1)  2  (1)
//
//       这条路径上相邻两点的高度差绝对值的最大值为 0 ，
//       这是所有路径中的最小值。


// 思路： 二分 + DFS
//
//      一般这种求最大值的最小值题目，都可以用二分来解决。
//
//      我们可以二分路径上最大高度差的绝对值，初始区间为 [0, H] ，
//      其中 H 为所有点的最大高度。
//
//      因为我们采用的是闭后闭的区间，所以当二分区间为空时才结束二分，
//      此时 l 就是最大高度差绝对值的最小值。
//
//      每次二分时使用 dfs(mid) 来搜索，返回值为 bool 类型，
//      判断最大高度差的绝对值不超过 mid 的路径是否存在。
//
//          1. 如果存在，则说明最高度的最小值在左边区间，
//              二分区间变为 [l, mid - 1] 。
//          2. 如果不存在，则说明最高度的最小值在右边区间，
//              二分区间变为 [mid + 1, r] 。
//
//
//      设 heights 的最大值为 H 。
//
//      时间复杂度：O(m * n * log(H))
//          1. 二分区间为 [0, H] ，二分的时间复杂度为 O(log(H))
//          2. 每次二分判断时， dfs 的时间复杂度为 O(m * n)
//      空间复杂度：O(m * n)
//          1. 需要维护一个大小为 O(m * n) 的二维数组 visited ，
//              用来记录每个点是否被访问过
//          2. 需要递归搜索可能的路径，路径长度就是栈深度，
//              最差情况下，全部 O(m * n) 个点都在路径上


// 每个方向的位置改变量
//  0: 向上走 1 步
//  1: 向右走 1 步
//  2: 向下走 1 步
//  3: 向左走 1 步
var DR = []int{-1, 0, 1, 0}
var DC = []int{0, 1, 0, -1}


func minimumEffortPath(heights [][]int) int {
    m, n := len(heights), len(heights[0])
    // 获取 heights 中最大的高度
    maxHeight := 0
    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            if maxHeight < heights[r][c] {
                maxHeight = heights[r][c]
            }
        }
    }
    // 二分相邻位置高度差，最小为 0 ，最大为 max_height
    l, r := 0, maxHeight
    // visited[r][c] 表示 (r, c) 是否访问过
    visited := make([][]bool, m)
    for r := 0; r < m; r++ {
        visited[r] = make([]bool, n)
    }
    // 当二分区间不为空时，还需要继续处理
    for l <= r {
        // 计算区间中点下标 mid
        mid := (l + r) >> 1
        // 清空 visited
        for _, row := range visited {
            for j := range row {
                row[j] = false
            }
        }
        if dfs(heights, visited, mid, 0, 0) {
            // 如果能在最大高度差不超过 mid 的情况下走到右下角，
            // 则说明最高度的最小值在左边区间
            r = mid - 1
        } else {
            // 如果不能走到右下角，
            // 说明最高度的最小值在右边区间
            l = mid + 1
        }
    }
    // l 就是最大高度的最小值
    return l
}

func dfs(heights [][]int, visited [][]bool, maxDiff int, r int, c int) bool {
    m, n := len(heights), len(heights[0])
    // 如果已经到了右下角，则最大差为 max_diff 时，
    // 能成功从左上角走到右下角，直接返回 true
    if r == m - 1 && c == n - 1 {
        return true
    }

    // 标记 (r, c) 已访问，
    // 由于我们已经必定能抵达 (r, c) ，
    // 后续只需要判断能否从 (r, c) 到达右下角即可，
    // 所以无需取消标记
    visited[r][c] = true
    // 获取 (r, c) 处的高度
    curHeight := heights[r][c]
    // 遍历 4 个方向
    for i := 0; i < 4; i++ {
        // 计算该方向的下一个位置
        rr := r + DR[i]
        cc := c + DC[i]
        // 如果下一个位置合法 
        // 且 未访问过 
        // 且 当前高度和下一个位置的高度差小于 max_diff 
        // 且 从 (rr, cc) 处能走到右下角，
        // 则返回 true
        if 0 <= rr && rr < m &&
            0 <= cc && cc < n &&
            !visited[rr][cc] &&
            abs(heights[rr][cc] - curHeight) <= maxDiff &&
            dfs(heights, visited, maxDiff, rr, cc) {
                return true
        }
    }
    // 此时不能成功走到右下角，返回 false
    return false
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
