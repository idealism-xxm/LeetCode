# 链接：https://leetcode.com/problems/maximum-value-after-insertion/
# 题意：给定一个整数（有正负）的字符串，现在加一个指定的数字，求加入后能形成的最大的数？

# 输入： n = "99", x = 9
# 输出： "999"

# 输入： n = "-13", x = 2
# 输出： "-123"
# 解释： 可以形成以下三个数字： -213, -123, -132
#       最大的是 -123

# 思路： 贪心
#
#       如果是负数，则要将该数放在第一个大于它的位置，如果都不大于，则放在最后
#       如果是正数，则要将该数放在第一个小于它的位置，如果都不小于，则放在最后
#
#       时间复杂度： O(n)
#       空间复杂度： O(1)


class Solution:
    def maxValue(self, n: str, x: int) -> str:
        if n[0] == '-':
            for i in range(1, len(n)):
                # 找到第一个大于 x 的数位的位置，在此插入
                if int(n[i]) > x:
                    return n[:i] + str(x) + n[i:]
            # 如果没有，则在最后插入
            return n + str(x)

        for i in range(len(n)):
            # 找到第一个小于 x 的数位的位置，在此插入
            if int(n[i]) < x:
                return n[:i] + str(x) + n[i:]
        # 如果没有，则在最后插入
        return n + str(x)
