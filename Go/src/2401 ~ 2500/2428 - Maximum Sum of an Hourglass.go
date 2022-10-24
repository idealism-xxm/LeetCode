// 链接：https://leetcode.com/problems/maximum-sum-of-an-hourglass/
// 题意：给定一个 m * n 的矩阵 grid ，求所有“工”字形的数字之和的最大值？
//      一个“工”字形是如下 3 * 3 的子矩阵的一部分：
//      A B C
//      _ D _
//      E F G
//
//      注意：“工”字形不能旋转，比如按照以上形式才有效。


// 数据限制：
//  m == grid.length
//  n == grid[i].length
//  3 <= m, n <= 150
//  0 <= grid[i][j] <= 10 ^ 6


// 输入： grid = [[6,2,1,3],[4,2,1,5],[9,2,8,7],[4,1,2,9]]
// 输出： 30
// 解释： 数字和最大的“工”字形为： 6 + 2 + 1 + 2 + 9 + 2 + 8 = 30

// 输入： grid = [[1,2,3],[4,5,6],[7,8,9]]
// 输出： 35
// 解释： 只存在一个“工”字形，和为： 1 + 2 + 3 + 5 + 7 + 8 + 9 = 35


// 思路： 模拟
//
//      遍历所有“工”字形的右下角坐标 (i, j) ，
//      然后计算该“工”字形的 7 个数字之和 total ，
//      那么所有 total 的最大值就是答案。
//
//
//      时间复杂度：O(mn)
//          1. 需要遍历计算全部 O(mn) 个“工”字形的 7 个数字之和
//      空间复杂度：O(1)
//          1. 只需要使用常数个额外变量即可


func maxSum(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    // ans 维护所有“工”字形数字和的最大值
    ans := 0
    // 遍历所有“工”字形的右下角坐标 (i, j) ，方便处理边界情况
    for i := 2; i < m; i++ {
        for j := 2; j < n; j++ {
            // 计算当前“工”字形的数字和，并更新 ans 最大值
            total := grid[i][j] + grid[i][j - 1] + grid[i][j - 2] + 
                grid[i - 1][j - 1] + 
                grid[i - 2][j] + grid[i - 2][j - 1] + grid[i - 2][j - 2];
            ans = max(ans, total)
        }
    }

    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}