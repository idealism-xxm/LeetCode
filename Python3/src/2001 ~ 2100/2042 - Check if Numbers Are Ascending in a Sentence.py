# 链接：https://leetcode.com/problems/check-if-numbers-are-ascending-in-a-sentence/
# 题意：给定一个空格分割的句子，判断其中的数字是否严格单调递增？

# 数据限制：
#   3 <= s.length <= 200
#   s 由英文小写字母、空格、数字组成
#   s 中的数字范围为 [2, 100]
#   s 中的 token 只通过一个空格分隔
#   s 中至少有两个数字
#   s 中的每个数字是正数，且小于 100 ，并且不含有前导零
#   s 不以空格开始，也不以空格结束

# 输入： s = "1 box has 3 blue 4 red 6 green and 12 yellow marbles"
# 输出： true
# 解释： s 中的数字 1, 3, 4, 6, 12 ，是严格单调递增的。

# 输入： s = "hello world 5 x 5"
# 输出： false
# 解释： s 中的数字 5, 5 ，非严格单调递增的。

# 输入： s = "sunset is at 7 51 pm overnight lows will be in the low 50 and 60 s"
# 输出： false
# 解释： s 中的数字 7, 51, 50, 60 ，非严格单调递增的。

# 输入： s = "4 5 11 26"
# 输出： true
# 解释： s 中的数字 4, 5, 11, 26 ，是严格单调递增的。

# 思路： 枚举
#
#       初始化前一个数字 pre = -1
#       按照空格分割，然后判断每个 token 是否为数字，
#       1. 如果是数字，就直接转换成数字 cur ，
#           (1) cur <= pre ，直接返回 false
#           (2) cur > pre ，更新 pre = cur ，继续处理下一个 token
#       2. 如果不是数字，则直接处理下一个
#
#       最后还没有返回，则 s 中的所有数字是严格单调递增的
#
#       时间复杂度： O(n)
#       空间复杂度： O(n)


class Solution:
    def areNumbersAscending(self, s: str) -> bool:
        # 所有数字都是正数，所以初始化 pre = -1
        pre = -1
        # 枚举所有 token
        for part in s.split():
            # 如果是数字，就直接转换成数字 cur
            if part.isnumeric():
                cur = int(part)
                # cur <= pre ，则非严格单调递增，直接返回 false
                if cur <= pre:
                    return False

                pre = cur
        # 所有 token 都是数字，则 s 中的所有数字是严格单调递增的
        return True
