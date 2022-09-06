# 链接：https://leetcode.com/problems/climbing-stairs/
# 题意：有一个 n 级楼梯，一个人每次可以上一级或者两级，
#      求有多少种方法可以到顶部？


# 数据限制：
#  1 <= n <= 45


# 输入： 2
# 输出： 2
# 解释： 有两种到达顶部的方法：
#          1. 先上一级，再上一级
#          2. 直接上两级

# 输入： 3
# 输出： 3
# 解释： 有两种到达顶部的方法：
#          1. 先上一级，再上一级，再上一级
#          2. 先上一级，再上两级
#          3. 先上两级，再上一级

# 思路：DP
#
#      DP 入门题，设 dp[i] 表示上到第 i 级楼梯的方法数。
#
#      初始化： dp[0] = dp[1] = 1 ，
#          表示最开始在第 0 级和第 1 级各有一种方案。
#      状态转移方程： dp[i] = dp[i - 1] + dp[i - 2] ，
#          即第 i 级可以从第 i - 1 级上一级到达，
#          也可以从第 i - 2 级上两级到达。
#
#      最后 dp[n] 就是上到顶部的方法数。
#
#
#      DP 常见的三种优化方式见 LeetCode 583 这题的思路，
#      本题只能采用滚动数组的方式进行优化，将空间复杂度从 O(n) 优化为 O(1) 。
#      本实现为了便于理解，不做优化处理。
#
#
#      时间复杂度： O(logn)
#          1. 只需要循环 O(logn) 次
#      空间复杂度： O(1)
#          1. 只需要使用常数个额外变量


class Solution:
    def climbStairs(self, n: int) -> int:
        # dp[i] 表示上到第 i 级楼梯的方法数
        dp: List[int] = [0] * (n + 1)
        # 最开始在第 0 级和第 1 级各有一种方案
        dp[0], dp[1] = 1, 1

        for i in range(2, n + 1):
            # 第 i 级可以从第 i - 1 级上一级到达，
            # 也可以从第 i - 2 级上两级到达
            dp[i] = dp[i - 1] + dp[i - 2]

        # dp[n] 就是上到顶部的方法数
        return dp[n]