# 链接：https://leetcode.com/problems/minimum-number-of-swaps-to-make-the-binary-string-alternating/
# 题意：给定一个 01 串，你可以交换任意两个字符，
#       求最少多少次可以使其变成 0 和 1 交替的串，不可能时返回 -1 。

# 数据限制：
#   1 <= s.length <= 1000
#   s[i] is either '0' or '1'

# 输入： s = "111000"
# 输出： 1
# 解释： "111000" -> "101010"


# 输入： s = "010"
# 输出： 0

# 输入： s = "1110"
# 输出： -1

# 思路： 贪心
#
#       我们先用 zero_cnt 统计 0 分别在奇数位和偶数位出现的次数：
#           zero_cnt[0] 表示奇数位的 0 的个数
#           zero_cnt[1] 表示偶数位的 0 的个数
#
#       然后计算 0 和 1 的总数为 total_zero 和 total_one ，
#           1. total_zero == total_one:
#               则 0 既可以在奇数位，也可以在偶数位，
#               让 0 在个数较多的那个位置即可，那么较少的个数就是需要交换的次数
#           2. total_zero == total_one + 1:
#               则 0 必须在偶数位，那么奇数位的 0 必须与偶数位的 1 交换
#           3. total_zero + 1 == total_one:
#               则 0 必须在奇数位，那么偶数位的 0 必须与奇数位的 1 交换
#           4. abs(total_zero - total_one) > 1:
#               不可能变成 0 和 1 交替的串
#       
#       时间复杂度： O(n)
#       空间复杂度： O(1)


class Solution:
    def minSwaps(self, s: str) -> int:
        # zero_cnt[0] 表示奇数位的 0 的个数
        # zero_cnt[1] 表示偶数位的 0 的个数
        zero_cnt = [0, 0]
        # 统计 0 在奇偶位的个数
        for i, ch in enumerate(s):
            if ch == '0':
                zero_cnt[i & 1] += 1
        
        # 计算 0 和 1 的个数
        total_zero = sum(zero_cnt)
        total_one = len(s) - total_zero
        # 如果两者个数相等，则 0 既可以在奇数位，也可以在偶数位，
        # 让 0 在个数较多的那个位置即可，那么较少的个数就是需要交换的次数
        if total_zero == total_one:
            return min(zero_cnt)
        # 如果 0 比 1 多一个，则 0 必须在偶数位，那么奇数位的 0 必须与偶数位的 1 交换
        if total_zero == total_one + 1:
            return zero_cnt[1]
        # 如果 0 比 1 少一个，则 0 必须在奇数位，那么偶数位的 0 必须与奇数位的 1 交换
        if total_zero + 1 == total_one:
            return zero_cnt[0]

        # 如果 0 和 1 的个数差大于 1 ，则不可能变成 0 和 1 交替的串
        return -1
