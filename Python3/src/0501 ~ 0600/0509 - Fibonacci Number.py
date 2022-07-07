# 链接：https://leetcode.com/problems/fibonacci-number/
# 题意：求斐波那契数列的第 n 个数？
#
#      斐波那契数列 F(n) 定义如下：
#          1. F(0) = F(1) = 1
#          2. F(n) = F(n - 1) + F(n - 2) , n > 1


# 数据限制：
#  0 <= n <= 30


# 输入： n = 2
# 输出： 1
# 解释： F(2) = F(1) + F(0) = 1 + 0 = 1

# 输入： n = 3
# 输出： 2
# 解释： F(3) = F(2) + F(1) = 1 + 1 = 2

# 输入： n = 4
# 输出： 3
# 解释： F(4) = F(3) + F(2) = 2 + 1 = 3


# 思路： DP
#
#      题目已经将 DP 的状态定义和转移方程给出，直接按照题意初始化和计算即可。
#
#
#      时间复杂度：O(n)
#          1. 需要遍历计算全部 O(n) 个状态
#      空间复杂度：O(n)
#          1. 只需要维护常数个额外变量    


class Solution:
    def fib(self, n: int) -> int:
        # 定义状态，初始化 dp[0] = 0, dp[1] = 1
        dp: List[int] = [1] * (n + 1)
        dp[0] = 0
        # 从 dp[2] 开始进行状态转移
        for i in range(2, n + 1):
            dp[i] = dp[i - 1] + dp[i - 2]

        return dp[n]