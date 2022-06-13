# 链接：https://leetcode.com/problems/triangle/
# 题意：给定一个数组 triangle ，求从顶部到底部的路径中数字和的最小值。
#
#      如果当前在第 i 列，那么可以选择到达下一行的第 i 列或第 i + 1 列。
#
#      进阶：使用空间复杂度为 O(n) 的算法求解。


# 数据限制：
#  1 <= triangle.length <= 200
#  triangle[0].length == 1
#  triangle[i].length == triangle[i - 1].length + 1
#  -(10 ^ 4) <= triangle[i][j] <= 10 ^ 4


# 输入： triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
# 输出： 11
# 解释： 数字和最小的路径是 2 -> 3 -> 5 -> 1 ，数字和为 11 。
#       (2)
#       (3) 4
#        6 (5) 7
#        4 (1) 8 3

# 输入： triangle = [[-10]]
# 输出： -10


# 思路1： DP
#
#      DP 入门题，定义 dp[i][j] 表示走到第 i 行第 j 列时数字和的最小值。
#
#      初始化： dp[0][0] = triangle[0][0]
#      状态转移：只有 (i - 1, j - 1) 和 (i - 1, j) 能走到 (i, j) ，
#          所以 dp[i][j] = min(dp[i - 1][j], dp[i - 1][j - 1]) + triangle[i][j]
#
#      那么满足题意的答案就是 min(dp[n - 1]) 。
#
#      这个方法的空间复杂度是 O(n ^ 2) 的，不满足进阶的要求，
#      我们看看未注意的关键点，从这些方面来优化突破：
#          1. 状态转移时，只会通过上一行的状态 dp[i - 1] 来更新当前行的状态 dp[i] ，
#              而最终结果会从最后一行的状态 dp[n - 1] 中得到。
#              那么我们只需要维护当前行和上一行这两行的状态即可，
#              使用滚动数组就能优化为 O(n) 的空间复杂度。
#          2. 状态转移时， dp[i][j] 可由 dp[i - 1][j - 1] 和 dp[i - 1][j] 转移。
#              如果只使用一维数组，那么 dp[j] 的状态可由 dp[j] 和 dp[j - 1] 转移，
#              因为每个位置只能走一次，所以在更新 dp[j] 时，要保证 dp[j - 1] 是上一行的状态，
#              那么只要我们从大到小更新 dp[j] 即可。
#
#      这两个优化都是背包常见的优化方法，最终能将空间复杂度优化为 O(n) ，且只使用一个一维数组。
#
#
#      时间复杂度：O(n ^ 2)
#          1. 需要遍历全部 O(n ^ 2) 个状态
#      空间复杂度：O(n)
#          1. 需用维护 dp 中一行的 O(n) 个状态


class Solution:
    def minimumTotal(self, triangle: List[List[int]]) -> int:
        # 定义 dp 数组，为了方便后续处理，初始化为一个极大值 0x3f3f3f3f
        dp: List[int] = [0x3f3f3f3f] * len(triangle)
        # 初始化第一行的状态
        dp[0] = triangle[0][0]
        # 将第 i - 1 行的状态转移至第 i 行
        for i in range(1, len(triangle)):
            # 每一行从大到小更新 dp[j] ，保证 dp[j - 1] 是上一行的状态
            for j in range(i, 0, -1):
                dp[j] = min(dp[j], dp[j - 1]) + triangle[i][j]
            # dp[0] 只能从 dp[0] 转移而来，所以直接加上 triangle[i][0]
            dp[0] += triangle[i][0]

        # 最后一行的状态的最小值就是答案
        return min(dp)
