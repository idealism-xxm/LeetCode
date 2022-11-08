# 链接：https://leetcode.com/problems/maximum-69-number/
# 题意：给定一个只含有 6 和 9 的整数 num ，
#      求最多修改一位数字（ 6 变 9 ， 9 变 6 ）的情况下，
#      能得到的最大整数是什么？


# 数据限制：
#  1 <= palindrome.length <= 1000
#  palindrome 仅由英文小写字母组成


# 输入： num = 9669
# 输出： 9969
# 解释： 修改第一位数字，结果是 6669 ；
#       修改第二位数字，结果是 9969 ；
#       修改第三位数字，结果是 9699 ；
#       修改第四位数字，结果是 9666 ；
#       其中最大的结果是 9969

# 输入： num = 9996
# 输出： 9999
# 解释： 修改第四位数字，结果最大，为 9999

# 输入： num = 9999
# 输出： 9999
# 解释： 不修改任何位数字，结果最大，为 9999


# 思路： 贪心
#
#      贪心地将从左往右的第一个 6 修改为 9 即可。
#      如果没有数位为 6 ，则 num 本身就已是最大，不做修改。
#
#      实际处理时，我们只能从个位开始判断，所以需要找从右往左的最后一个 6 。
#
#      在处理过程中维护每一位的权重 weight ，
#      并用 ans_weight 维护最后一个 6 的权重。
#
#      这样最后只需要对 num 加上 3 * ans_weight ，
#      就能达到将最后一个 6 替换为 9 的效果。
#
#
#      【进阶】
#
#
#      时间复杂度：O(n)
#          1. 需要遍历 palindrome 中 O(n) 个字母
#      空间复杂度：O(n)
#          1. 需要维护结果中全部 O(n) 个字母


class Solution:
    def maximum69Number (self, num: int) -> int:
        # ans_weight 表示从右到左最后一个 6 的权重。
        # 初始化为 0 ，表示还没有 6
        ans_weight: int = 0
        # weight 表示当前最低位的权重，初始化个位的权重为 1
        weight: int = 1
        # 从个位遍历 cur 的每一位
        cur: int = num
        while cur > 0:
            # 如果该位是 6 ，则更新 ans_weight 为当前位的权重
            if cur % 10 == 6:
                ans_weight = weight

            # 移除最低位，并更新最低位的权重
            cur //= 10
            weight *= 10

        # 将 num 从右往左最后一个 6 替换为 9
        return num + 3 * ans_weight
