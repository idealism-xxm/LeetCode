# 链接：https://leetcode.com/problems/smallest-even-multiple/
# 题意：给定一个正整数 n ，返回 2 和 n 的最小公倍数。


# 数据限制：
#  1 <= n <= 150


# 输入： n = 5
# 输出： 10
# 解释： 2 和 5 的最小公倍数为 10

# 输入： n = 6
# 输出： 6
# 解释： 2 和 6 的最小公倍数为 6


# 思路： 模拟
#
#      如果 n 是奇数，则 2 和 n 的最小公倍数为 2n 。
#      如果 n 是偶数，则 2 和 n 的最小公倍数为 n
#
#
#      时间复杂度：O(1)
#          1. 只需要常数次运算即可
#      空间复杂度：O(1)
#          1. 只需要使用常数个额外变量即可


class Solution:
    def smallestEvenMultiple(self, n: int) -> int:
        if n & 1:
            # 如果 n 是奇数，则 2 和 n 的最小公倍数为 2n
            return n << 1

        # 如果 n 是偶数，则 2 和 n 的最小公倍数为 n
        return n
