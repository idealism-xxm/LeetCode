# 链接：https://leetcode.com/problems/count-square-sum-triples/
# 题意：给定 n ，求满足 a ^ 2 + b ^ 2 = c ^ 2 的三元组个数？
#       0 < a, b, c <= n

# 数据限制：
#   1 <= n <= 250

# 输入： n = 5
# 输出： 2
# 解释： (3,4,5), (4,3,5)

# 输入： n = 10
# 输出： 4
# 解释： (3,4,5), (4,3,5), (6,8,10), (8,6,10)

# 思路： 模拟
#
#       枚举 a 和 b ，然后判断 c 是否在 [1, n] 之间
#
#       时间复杂度： O(n ^ 2)
#       空间复杂度： O(1)


class Solution:
    def countTriples(self, n: int) -> int:
        ans = 0
        for a in range(1, n + 1):
            aa = a * a
            for b in range(a + 1, n + 1):
                cc = aa + b * b
                c = int(math.sqrt(cc) + 0.000001)
                if c <= n and cc == c * c:
                    ans += 1
        return ans * 2
