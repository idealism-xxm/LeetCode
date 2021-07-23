# 链接：https://leetcode.com/problems/maximum-number-of-points-with-cost/
# 题意：给定一个 m * n 的二维数组 points ，求从其中选择一些数能获得的最大的和？
#       选择方式：
#           1. 每一行选一个数，选中 (r, c) 则给总和加上 points[r][c]
#           2. 如果相邻两行所选数隔得太远，则会有惩罚，会从总和中减去 abs(c1 - c2) ，
#               其中 c1 和 c2 是相邻两行的列下标


# 数据限制：
#   m == points.length
#   n == points[r].length
#   1 <= m, n <= 10 ^ 5
#   1 <= m * n <= 10 ^ 5
#   0 <= points[r][c] <= 10 ^ 5

# 输入： points = [[1,2,3],[1,5,1],[3,1,1]]
# 输出： 9
# 解释： 选择下标分别为 (0, 2)，(1, 1) 和 (2, 0) 的三个点，
#       总和 = 3 + 5 + 3 - abs(1 - 2) - abs(0 - 1) = 9

# 输入： points = [[1,5],[2,3],[4,2]]
# 输出： 11
# 解释： 选择下标分别为 (0, 1)，(1, 1) 和 (2, 0) 的三个点，
#       总和 = 5 + 3 + 4 - abs(1 - 1) - abs(1 - 0) = 11


# 思路： DP
#
#       比赛时已经想到是 DP 了，但是没有注意到更深层次的特性，
#       所以使用了线段树维护下标差，时间复杂度多一个 logn 。
#
#       结束后看了大家的解法才恍然大悟，
#       其实我是用线段树就只是对前一半 +1 ，对后一半 -1 ，
#       那么可以分别求前一半更新到当前的最大值和后一半更新到当前的最大值，
#
#       设 dp[i][j] 表示第 i 行选择第 j 个数时能获取到的最大值，
#       设 lmax 表示 dp[i - 1][:j] 能获取到的最大值，
#       那么向右移动到 j 时， lmax = max(lmax - 1, dp[i - 1][j])
#
#       设 rmax 表示 dp[i - 1][j:] 能获取到的最大值，
#       那么向左移动到 j 时， rmax = max(rmax - 1, dp[i - 1][j])
#
#       那么 dp[i][j] = max(lmax, rmax)
#       为了方便，我们可以遍历两次，这样就不用额外维护一个数组了
#
#       时间复杂度： O(m * n)
#       空间复杂度： O(n)


class Solution:
    def maxPoints(self, points: List[List[int]]) -> int:
        # 初始化 dp 数组
        m, n = len(points), len(points[0])
        dp = points[0]
        for i in range(1, m):
            nxt_dp = [0] * n
            # lmax 表示从左侧开始到当前位置的数中，能获取到的最大值
            lmax = 0
            for j in range(n):
                # 1. 选左侧的值：每次向右移动到 j 列时，左侧对当前列贡献的最大值都需要 -1
                # 2. 选 dp[i - 1][j] ：列下标差为 0 ，直接为 dp[i - 1][j]
                lmax = max(lmax - 1, dp[j])
                nxt_dp[j] = max(nxt_dp[j], lmax + points[i][j]) 

            # rmax 表示从右侧开始到当前位置的数中，能获取到的最大值
            rmax = 0
            for j in range(n - 1, -1, -1):
                # 1. 选右侧的值：每次向左移动到 j 列时，右侧对当前列贡献的最大值都需要 -1
                # 2. 选 dp[i - 1][j] ：列下标差为 0 ，直接为 dp[i - 1][j]
                rmax = max(rmax - 1, dp[j])
                nxt_dp[j] = max(nxt_dp[j], rmax + points[i][j])
            
            # 滚动数组，直接赋值
            dp = nxt_dp

        # 返回所有列中的最大值
        return max(dp)