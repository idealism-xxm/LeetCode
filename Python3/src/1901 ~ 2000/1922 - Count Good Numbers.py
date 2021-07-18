# 链接：https://leetcode.com/problems/count-good-numbers/
# 题意：给定一个数字字符串，如果偶数下标（从 0 开始）的数是偶数，
#       奇数下标的数是素数 (2, 3, 5, 7) ，那么这个数是一个好数，
#       求长度为 n 的数字字符串有多少个好数？

# 数据限制：
#   1 <= n <= 10 ^ 15

# 输入： n = 1
# 输出： 5
# 解释： "0", "2", "4", "6", "8"

# 输入： n = 4
# 输出： 400

# 输入： n = 50
# 输出： 564908303

# 思路： 快速幂
#
#       比赛时第一次理解错题意了，还以为是数位 DP ，结果发现大家都很快做出来了，
#       重新仔细读题，发现就是一个排列组合统计。
#
#       给定了长度，我们就知道偶数下标的数量 x 和奇数下标的数量 y ，
#       那么最终的个数就是 5 ^ x * 4 ^ y
#
#       时间复杂度： O(logn)
#       空间复杂度： O(1)


MOD = 1000000007


def quick_pow(a, n):
    ans = 1
    while n > 0:
        if n & 1:
            ans = (ans * a) % MOD
        a = (a * a) % MOD
        n >>= 1
    return ans


class Solution:
    def countGoodNumbers(self, n: int) -> int:
        return (quick_pow(5, (n + 1) >> 1) * quick_pow(4, n >> 1)) % MOD
