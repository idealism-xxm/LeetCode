# 链接：https://leetcode.com/problems/count-number-of-ways-to-place-houses/
# 题意：有一条街道，两边沿街边各 n 个地块，标号为 1 ~ n 。
#      求有多少种放置房子的方案，每种方案中放置的任意两个房子在街道同一边不相邻。
#
#      结果模 10 ^ 9 + 7 。
#      注意：如果街边第 i 个地块放置了一个房子，那么另一边第 i 个地块还可以放置一个房子，
#           街对面不算相邻。


# 数据限制：
#  1 <= n <= 10 ^ 4


# 输入： n = 1
# 输出： 4
# 解释： 总共有 4 种放置方案：
#          1. 所有地块都空
#          2. 一侧地块放置一个房子
#          3. 另一侧地块放置一个房子
#          4. 两侧地块分别放置一个房子

# 输入： n = 2
# 输出： 9
# 解释： 总共有 9 种放置方案（ '|' 表示街道， '_' 表示空地块，
#       '1'/'2' 分别表示当前标号的地块放置了房子）：
#
#       _|_    _|1    _|_
#       _|_    _|_    _|2
#
#       1|_    1|1    1|_
#       _|_    _|_    _|2
#
#       _|_    _|1    _|_
#       2|_    2|_    2|2


# 思路： DP
#
#      本题是 LeetCode 509 的加强版，需要自己推导出状态和转移方程。
#
#      题目说明街道两侧的房子放置方案相互独立，我们可以只统计一侧的放置方案数 dp[n] ，
#      那么总方案数就是 dp[n] ^ 2 。
#
#      设 dp[i] 表示街道一侧前 n 个地块放置房子（任意两个房子不相邻）的方案数。
#
#      初始化：
#          1. dp[0] = 1: 没有地块，只能不放置房子，共 1 种方案
#          2. dp[1] = 2: 有 1 个地块，能选择放置和不放置，共 2 种方案
#
#      状态转移：
#          1. 第 i 个地块不放置房子，那么第 i - 1 个地块是否放置房子无所谓，
#             对应的方案数可由 dp[i - 1] 转移而来
#          2. 第 i 个地块放置房子，那么第 i - 1 个地块必定不能放置房子，
#             对应的方案数可由 dp[i - 2] 转移而来
#
#      综上， dp[i] = dp[i - 1] + dp[i - 2] ，这其实就是斐波那契数列的状态转移方程。
#
#
#      DP 常见的三种优化方式见 LeetCode 583 这题的思路，
#      本题只能采用滚动数组的方式进行优化，将空间复杂度从 O(n) 优化为 O(1) 。
#      本实现为了便于理解，不做优化处理。
#
#
#      如果 n 达到 10 ^ 9 ，则可以使用矩阵快速幂进行转移，时间复杂度可以优化为 O(logn) 。
#
#
#      时间复杂度：O(n)
#          1. 需要遍历 dp 中全部 O(n) 个状态
#      空间复杂度：O(n)
#          1. 需要维护 dp 中全部 O(n) 个状态


MOD: int = 1_000_000_007


class Solution:
    def countHousePlacements(self, n: int) -> int:
        # dp[i] 表示街道一侧前 n 个地块放置房子（任意两个房子不相邻）的方案数
        dp: List[int] = [1] * (n + 1)
        # 初始化：
        #  1. dp[0] = 1: 没有地块，只能不放置房子，共 1 种方案
        #  2. dp[1] = 2: 有 1 个地块，能选择放置和不放置，共 2 种方案
        dp[1] = 2
        for i in range(2, n + 1):
            # 1. 第 i 个地块不放置房子，那么第 i - 1 个地块是否放置房子无所谓，
            #    对应的方案数可由 dp[i - 1] 转移而来
            # 2. 第 i 个地块放置房子，那么第 i - 1 个地块必定不能放置房子，
            #    对应的方案数可由 dp[i - 2] 转移而来
            dp[i] = (dp[i - 1] + dp[i - 2]) % MOD
        
        # 街道两侧的房子放置方案相互独立，总方案数为 dp[n] ^ 2
        return (dp[n] * dp[n]) % MOD
