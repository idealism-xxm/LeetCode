# 链接：https://leetcode.com/problems/largest-odd-number-in-string/
# 题意：给定一个无前导零的正数字符串，求一个最大的奇数子串（不存在则返回空串）？

# 数据限制：
#   1 <= num.length <= 10 ^ 5
#   num 只含有数字，且无前导零

# 输入： num = "52"
# 输出： "5"

# 输入： num = "4206"
# 输出： ""

# 输入： num = "35427"
# 输出： "35427"

# 思路： 贪心
#
#       找到最后一个出现的奇数数字下标 i ，
#       则 sum[:i + 1] 就是答案
#
#       时间复杂度： O(|num|)
#       空间复杂度： O(1)


class Solution:
    def largestOddNumber(self, num: str) -> str:
        i = len(num) - 1
        while i >= 0:
            # 找到第一个奇数位置，就返回
            if int(num[i]) & 1 == 1:
                return num[:i + 1]
            i -= 1
        # 无奇数，则返回空串
        return ""
