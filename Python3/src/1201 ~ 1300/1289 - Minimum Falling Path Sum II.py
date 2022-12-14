# 链接：https://leetcode.com/problems/minimum-falling-path-sum-ii/
# 题意：给定一个 n * n 的矩阵 grid ，求其非零偏移下降路径的最小和？
#
#      非零偏移下降路径的长度为 n ，从第一行开始，在最后一行结束。
#      路径中相邻两行的元素所在列需要不同。
#      如果位于 (row, col) 的元素在下降路径中，
#      则该下降路径的下一个元素应不应该位于 (row + 1, col) 。
#      


# 数据限制：
#  n == grid.length == grid[i].length
#  1 <= n <= 200
#  -99 <= grid[i][j] <= 99


# 输入： grid = [[1,2,3],[4,5,6],[7,8,9]]
# 输出： 13
# 解释： 全部非零偏移下降路径如下：
#          [1,5,9], [1,5,7], [1,6,7], [1,6,8],
#          [2,4,8], [2,4,9], [2,6,7], [2,6,8],
#          [3,4,8], [3,4,9], [3,5,7], [3,5,9]
#
#       其中 [1,5,7] 的和最小，为 13 ：
#          (1) 2  3
#           4 (5) 5
#          (7) 8  9

# 输入： grid = [[7]]
# 输出： 7


# 思路： DP + 贪心
#
#      本题是 LeetCode 931 的加强版，每个状态依赖的状态个数从 O(1) 上升为 O(n) 。
#      但依旧可以使用贪心地优化方式，将时间复杂度保持在 O(n ^ 2) 。
#
#
#      设 dp[i][j] 表示从第 0 行开始到 (i, j) 处的所有路径中，路径和的最小值。
#
#      初始化： dp[0][j] = grid[0][j] ，因为第 0 行只能选择自己作为下降路径开始元素
#      状态转移方程： dp[i][j] = grid[i][j] + min(dp[i - 1][k]), j != k
#
#      根据题意可知，如果 grid[i][j] 在下降路径中，那么其上一个元素不应该位于 (i - 1, j) 。
#
#      那么我们可以优先贪心选择 dp[i - 1] 中的最小值。
#      如果 dp[i - 1] 中的最小值恰好是 dp[i - 1][j] ，
#      那么我们必定贪心选择 dp[i - 1] 中的次小值。
#
#      为此我们可以维护另一个数组 min_dp ，其中 min_dp[i] 表示 dp[i] 中最小的两个值，
#      其中 min_dp[i][0] <= min_dp[i][1] 。
#
#      这样在状态转移时，就能根据 min_dp[i][0] 与 dp[i - 1][j] 是否相等，
#      选择 min_dp[i - 1] 二者之一进行转移即可。    
#
#
#      DP 常见的三种优化方式见 LeetCode 583 这题的思路，
#      本题只能采用滚动数组的方式进行优化，将空间复杂度从 O(n ^ 2) 优化为 O(n) 。
#      本实现为了便于理解，不做优化处理。
#
#
#      时间复杂度： O(n ^ 2)
#          1. 需要初始化并遍历 dp[0] 中全部 O(n) 个状态
#          2. 需要遍历计算 dp 中全部 O(n ^ 2) 个状态
#      空间复杂度： O(n ^ 2)
#          1. 需要维护 dp 中全部 O(n ^ 2) 个状态
#          2. 需要维护 min_dp 中全部 O(n) 个状态


class Solution:
    def minFallingPathSum(self, grid: List[List[int]]) -> int:
        n: int = len(grid)
        # dp[i][j] 表示从第 0 行开始到 (i, j) 处的所有路径中，路径和的最小值
        dp: List[List[int]] = [[0] * n for _ in range(n)]
        # 第 0 行只有一种选择， dp[0] 初始化为 grid[0]
        dp[0] = grid[0][:]
        # min_dp[i] 表示 dp[i] 中最小的两个值，其中 min_dp[i][0] <= min_dp[i][1]
        min_dp: List[List[int]] = [[0x3f3f3f3f] * 2 for _ in range(n)]
        # 初始化 min_dp[0]
        for j in range(n):
            if dp[0][j] <= min_dp[0][0]:
                # 如果 dp[0][j] 是新的最小值，则原最小值变为次小值
                min_dp[0][1] = min_dp[0][0]
                min_dp[0][0] = dp[0][j]
            elif dp[0][j] < min_dp[0][1]:
                # 如果 dp[0][j] 是新的次小值，则进行更新
                min_dp[0][1] = dp[0][j]
        
        for i in range(1, n):
            for j in range(n):
                if dp[i - 1][j] == min_dp[i - 1][0]:
                    # 如果 dp[i - 1][j] 是最小值，则选择次小值进行状态转移
                    dp[i][j] = grid[i][j] + min_dp[i - 1][1]
                else:
                    # 如果 dp[i - 1][j] 不是最小值，则选择最小值进行状态转移
                    dp[i][j] = grid[i][j] + min_dp[i - 1][0]

                if dp[i][j] <= min_dp[i][0]:
                    # 如果 dp[i][j] 是新的最小值，则原最小值变为次小值
                    min_dp[i][1] = min_dp[i][0]
                    min_dp[i][0] = dp[i][j]
                elif dp[i][j] < min_dp[i][1]:
                    # 如果 dp[i][j] 是新的次小值，则进行更新
                    min_dp[i][1] = dp[i][j]
        
        # 下降路径的最小和就是 min_dp[n - 1][0]
        return min_dp[n - 1][0]
