# 链接：https://leetcode.com/problems/minimum-falling-path-sum/
# 题意：给定一个 n * n 的矩阵 matrix ，求其下降路径的最小和？
#
#      下降路径长度为 n ，从第一行开始，在最后一行结束。
#      如果位于 (row, col) 的元素在下降路径中，
#      则该下降路径的下一个元素应该位于以下三处：
#      (row + 1, col - 1), (row + 1, col), (row + 1, col - 1) 


# 数据限制：
#  n == matrix.length == matrix[i].length
#  1 <= n <= 100
#  -100 <= matrix[i][j] <= 100


# 输入： matrix = [[2,1,3],[6,5,4],[7,8,9]]
# 输出： 13
# 解释： 有两种最小和的下降路径，最小和为 13 ：
#        2 (1) 3     2 (1) 3
#        6 (5) 4     6  5 (4)
#       (7) 8  9     7 (8) 9

# 输入： matrix = [[-19,57],[-40,-5]]
# 输出： -59
# 解释： 有一种最小和的下降路径，最小和为 -59 ：
#        (-19) 57
#        (-40) -5


# 思路： DP
#
#      设 dp[i][j] 表示从第 0 行开始到 (i, j) 处的所有路径中，路径和的最小值。
#
#      初始化： dp[0][j] = matrix[0][j] ，因为第 0 行只能选择自己作为下降路径开始元素
#      状态转移方程： dp[i][j] = matrix[i][j] + min(dp[i - 1][j - 1], dp[i - 1][j], dp[i - 1][j + 1])
#
#      根据题意可知，如果 matrix[i][j] 在下降路径中，那么其上一个元素应该位于以下三处：
#      (i - 1, j - 1), (i - 1, j), (i - 1, j - 1) 。
#
#      即状态 dp[i][j] 可从 dp[i - 1][j - 1], dp[i - 1][j], dp[i - 1][j + 1] 转移而来，
#      我们取三者中的最小值进行转移即可。
#
#      处理时需要注意判断 j - 1 和 j + 1 是否越界。
#
#
#      DP 常见的三种优化方式见 LeetCode 583 这题的思路，
#      本题只能采用滚动数组的方式进行优化，将空间复杂度从 O(n ^ 2) 优化为 O(n) 。
#      本实现为了便于理解，不做优化处理。
#
#
#		时间复杂度： O(n ^ 2)
#          1. 需要初始化 dp[0] 全部 O(n) 个状态
#          2. 需要遍历计算 dp 全部 O(n ^ 2) 个状态
#		空间复杂度： O(n ^ 2)
#          1. 需要维护 dp 中全部 O(n ^ 2) 个状态


class Solution:
    def minFallingPathSum(self, matrix: List[List[int]]) -> int:
        n: int = len(matrix)
        # dp[i][j] 表示从第 0 行开始到 (i, j) 处的所有路径中，路径和的最小值
        dp: List[List[int]] = [[0] * n for _ in range(n)]
        # 第 0 行只有一种选择， dp[0] 初始化为 matrix[0]
        dp[0] = matrix[0][:]
        
        for i in range(1, n):
            for j in range(n):
                # dp[i][j] = matrix[i][j] + min(dp[i - 1][j - 1], dp[i - 1][j], dp[i - 1][j + 1])
                # 注意需要判断 j - 1 和 j + 1 是否越界
                min_pre: int = dp[i - 1][j]
                if j > 0:
                    min_pre = min(min_pre, dp[i - 1][j - 1])
                if j < n - 1:
                    min_pre = min(min_pre, dp[i - 1][j + 1])
                
                dp[i][j] = matrix[i][j] + min_pre
        
        # 下降路径的最小和就是 dp[n - 1] 的最小值
        return min(dp[n - 1])
