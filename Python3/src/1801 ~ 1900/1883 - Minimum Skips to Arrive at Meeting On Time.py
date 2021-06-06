# 链接：https://leetcode.com/problems/minimum-skips-to-arrive-at-meeting-on-time/
# 题意：给 n 条路的长度数组 dist，以及你的速度 speed ，
#       你每次走一条路时，必须等待到整数小时，但你可以选择跳过一次等待，立刻出发，
#       求最少跳过多少次等待时，可以在 hoursBefore 及之前走完全部路？
#       （跳过全部等待也无法完成时，返回 -1）

# 输入： dist = [1,3,2], speed = 4, hoursBefore = 2
# 输出： 1
# 解释： 不跳过等待，总共用时 (1/4 + 3/4) + (3/4 + 1/4) + (2/4) = 2.5 hours
#       跳过第一次等待，总共用时 ((1/4 + 0) + (3/4 + 0)) + (2/4) = 1.5 hours

# 输入： dist = [7,3,5,5], speed = 2, hoursBefore = 10
# 输出： 2
# 解释： 不跳过等待，总共用时 (7/2 + 1/2) + (3/2 + 1/2) + (5/2 + 1/2) + (5/2) = 11.5 hours
#       跳过第一次等待，总共用时 ((7/2 + 0) + (3/2 + 0)) + ((5/2 + 0) + (5/2)) = 10 hours

# 输入： dist = [7,3,5,5], speed = 1, hoursBefore = 10
# 输出： -1
# 解释： 跳过全部等待，都无法完成， (7 + 3 + 5 + 5) / 1 = 20 / 1 = 20

# 思路： DP
#
#       dp[i][j] 表示前 i 条路跳过 j 次等待时所用的最少时间
#       初始化：dp[i][j] = hoursBefore + 1.0, dp[0][0] = 0.0
#       状态转移方程 (0 < j < i) ：
#           dp[i][0] = ceil(dp[i - 1][0]) + dist[i - 1] / speed
#           dp[i][j] = min(ceil(dp[i - 1][j]), dp[i - 1][j - 1]) + dist[i - 1] / speed
#           dp[i][i] = dp[i - 1][j - 1] + dist[i - 1] / speed
#
#       注意向上取整时要减去 1e-9 ，因为速度最大 10^6 ，
#       且最多有 1000 条路，所以 1e-9 刚好能避免精度误差累计
#
#       时间复杂度： O(n ^ 2)
#       空间复杂度： O(n ^ 2) 【当然可以使用一维数组，倒序更新优化到 O(n) ，但没必要】


class Solution:
    def minSkips(self, dist: List[int], speed: int, hoursBefore: int) -> int:
        n = len(dist)
        # 初始化 dp 数组
        dp: List[List[float]] = [None] * (n + 1)
        for i in range(n + 1):
            dp[i] = [hoursBefore + 1.0] * (n + 1)
        dp[0][0] = 0

        # 状态转移
        for i in range(1, n + 1):
            # 注意向上取整时要减去 1e-9 ，因为速度时 10^6 ，
            # 且有 1000 条路，所以 1e-9 刚好能避免精度误差累计
            dp[i][0] = ceil(dp[i - 1][0] - 1e-9) + dist[i - 1] / speed
            for j in range(i):
                # 注意向上取整时要减去 1e-9 ，因为速度最大 10^6 ，
                # 且最多有 1000 条路，所以 1e-9 刚好能避免精度误差累计
                dp[i][j] = min(ceil(dp[i - 1][j] - 1e-9), dp[i - 1][j - 1]) + dist[i - 1] / speed
            dp[i][i] = dp[i - 1][j - 1] + dist[i - 1] / speed

        # 找到最少的一个即可
        for j in range(n + 1):
            if dp[n][j] <= hoursBefore:
                return j
        # 没有满足要求的，返回 -1
        return -1
