# 链接：https://leetcode.com/problems/paint-house-iii/
# 题意：给定一个整数 n ，表示有 n 种颜色。
#      同时给定一个整数数组 houses ，其中 houses[i] 的含义如下：
#          1. houses[i] == 0: 第 i 个房子未染色，需要染色
#          2. 1 <= houses[i] <= n: 表示第 i 个房子的颜色
#
#      对于未染色的房子，需要将其染成 [1, n] 内的颜色 j + 1 ，花费为 cost[i][j] 。
#
#      连续相同颜色的房子会组成一个街区，
#      例如： houses = [1,2,2,3,3,2,1,1] 包含 5 个街区，
#      分别为 [{1}, {2,2}, {3,3}, {2}, {1,1}]
#
#      求全部染色后，组成 target 个街区所需的最小花费？
#      如果不能组成 target 个街区，则返回 -1 。


# 数据限制：
#   m == houses.length == cost.length
#   n == cost[i].length
#   1 <= m <= 100
#   1 <= n <= 20
#   1 <= target <= m
#   0 <= houses[i] <= n
#   1 <= cost[i][j] <= 10 ^ 4


# 输入： houses = [0,0,0,0,0], cost = [[1,10],[10,1],[10,1],[1,10],[5,1]], m = 5, n = 2, target = 3
# 输出： 9
# 解释： 将 houses 染色成 [1,2,2,1,1] ，则共有 3 个街区 [{1}, {2,2}, {1,1}] 。
#       总花费为 1 + 1 + 1 + 1 + 5 = 9 。

# 输入： houses = [0,2,1,2,0], cost = [[1,10],[10,1],[10,1],[1,10],[5,1]], m = 5, n = 2, target = 3
# 输出： 11
# 解释： 将 houses 染色城 [2,2,1,2,2] ，则共有 3 个街区 [{2,2}, {1}, {2,2}] 。
#       总花费为 10 + 1 = 11

# 输入： houses = [3,1,2,3], cost = [[1,1,1],[1,1,1],[1,1,1],[1,1,1]], m = 4, n = 3, target = 3
# 输出： -1
# 解释： 所有的房子都已染色，共有 4 个街区 [{3},{1},{2},{3}] ，不满足 3 个街区的条件。


# 思路： DP
#
#      设 dp[i][j][k] 表示前 i 个房子中，共有 j 个街区，
#      且第 i 个房子的颜色为 k 时的最小花费。
#
#      初始化： dp[i][j][k] = MAX ，方便后续计算。
#      同时需要处理第 0 个房子对应的初始值，第 0 个房子必定属于第 1 个街区，
#      我们根据它的颜色进行处理：
#          1. houses[i] == 0: 未被染色，则可以染成 [0, n) 的颜色，
#              共有 1 个街区，花费为 cost[0][k] ，即 dp[i][1][k] = cost[i][k]
#          2. houses[i] != 0: 已被染色，共有一个街区，无需任何花费，
#              dp[i][1][houses[i] - 1] = 0
#
#      那么很容易就能想到状态转移方程，对于状态 dp[i][j][k] 来说，
#      我们需要枚举第 i - 1 个房子的颜色 l ，
#      这时就能根据 k 和 l 的关系确定前 i - 1 个房子所需的街区数，进行不同处理即可：
#          1. k == l: 不会产生新的街区，则状态转移自 dp[i - 1][j][l]
#          2. k != l: 会产生新的街区，则状态转移自 dp[i - 1][j - 1][l]
#
#      最后令 ans = min(dp[m - 1][target]) ，根据 ans 的值返回即可：
#          1. ans == MAX: 则不存在满足题意的染色方案，直接返回 -1
#          2. ans != MAX: ans 即为满足题意的所有染色方案中，花费最小的费用
#
#      DP 常见的三种优化方式见 LeetCode 583 这题的思路，
#      本题可以采用滚动数组的方式进行优化，因为每一行的状态依赖上一行的所有状态，
#      所以无法采用其他两种方式进行优化。
#
#      使用滚动数组的话，能将空间复杂度从 O(m * target * n) 优化为 O(target * n) ，
#      本实现为了便于理解，不做优化处理。
#
#
#      时间复杂度：O(m * target * n ^ 2)
#          1. 需要遍历全部 O(m) 个房子
#          2. 遍历每个房子时，需要遍历全部 O(target) 个街区数
#          3. 遍历每个街区数时，需要遍历当前房子全部 O(n) 种颜色
#          4. 遍历当前房子的每种颜色时，需要遍历前一个房子的全部 O(n) 中颜色
#      空间复杂度：O(m * target * n ^ 2)
#          1. 需要用一个 dp 数组维护全部 O(m * target * n) 个状态


MAX: int = 0x3f3f3f3f


class Solution:
    def minCost(self, houses: List[int], cost: List[List[int]], m: int, n: int, target: int) -> int:
        # dp[i][j][k] 表示前 i 个房子中，共有 j 个街区，
        # 且第 i 个房子的颜色为 k 时的最小花费
        dp: List[List[List[int]]] = [[[MAX] * n for _ in range(target + 1)] for _ in range(m)]
        if houses[0] == 0:
            # 如果第 0 个房子未被染色，则可以染成 [0, n) 的颜色，
            # 共有 1 个街区，花费为 cost[0][k]
            for k in range(n):
                dp[0][1][k] = cost[0][k]
        else:
            # 如果第 0 个房子已被染色，则共有 1 个街区，无需任何花费
            dp[0][1][houses[0] - 1] = 0

        # 枚举后续的房子，进行状态转移
        for i in range(1, m):
            # 枚举前 i 个房子的街区数
            for j in range(1, target + 1):
                # 枚举第 i 个房子的颜色
                for k in range(n):
                    # 如果房子已被染色，且颜色不是 k ，则直接处理下一个
                    if houses[i] != 0 and houses[i] != k + 1:
                        continue

                    # 计算把第 i 个房子染成颜色 k 需要的花费，
                    # 默认已经染过色了，无需染色
                    paint_cost: int = 0
                    if houses[i] == 0:
                        # 还未染色时，需要花费 cost[i][k] 进行染色
                        paint_cost = cost[i][k]

                    # 枚举第 i - 1 个房子的颜色
                    for l in range(n):
                        if k == l:
                            # 如果颜色相同，不会产生新的街区，
                            # 则依赖的状态为前 i - 1 个房子的街区数为 j
                            dp[i][j][k] = min(dp[i][j][k], dp[i - 1][j][l] + paint_cost)
                        else:
                            # 如果颜色不同，会产生新的街区，
                            # 则依赖的状态为前 i - 1 个房子的街区数为 j - 1
                            dp[i][j][k] = min(dp[i][j][k], dp[i - 1][j - 1][l] + paint_cost)

        # 前 m 个房子中，共 target 街区的情况下，找到最小花费
        ans: int = min(dp[m - 1][target])
        if ans == MAX:
            # 如果 ans 仍为 MAX ，则不存在满足题意的染色方案，直接返回 -1
            return -1

        # 此时 ans 即为满足题意的所有染色方案中，花费最小的费用
        return ans
