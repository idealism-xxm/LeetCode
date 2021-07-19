# 链接：https://leetcode.com/problems/sum-game/
# 题意：给定偶数长度的字符串，含有数字和 '?' ，
#       如果还有 '?' ，则当前需要挑一个 '?' 变成一个数字，
#       两个人轮流进行操作，直至没有问号。
#       如果最后 左半边的数字和 = 右半边的数字和，则后手胜，否则先手胜，
#       求是否先手必胜？

# 数据限制：
#   2 <= num.length <= 10 ^ 5
#   num.length 是偶数
#   num 只含有 数字 和 '?'

# 输入： num = "5023"
# 输出： false

# 输入： num = "25??"
# 输出： true
# 解释： 先手填入一个 9 ，则后手无论填什么数字，都必败

# 输入： num = "?3295???"
# 输出： false
# 解释： 后手必胜，一种可能的过程如下
#       先手： "?3295???" -> "93295???"
#       后手： "93295???" -> "932959??"
#       先手： "932959??" -> "9329592?"
#       先手： "9329592?" -> "93295927"
#       后手胜 9 + 3 + 2 + 9 = 5 + 9 + 2 + 7

# 思路： 分类讨论
#
#       统计左右两边数字和 lsum 和 rsum ，
#       然后统计左右两边的问号数 lcnt 和 rcnt
#       1. 如果 lcnt + rcnt 是奇数，则由先手选择最后一个数字，
#           那么总共有 10 种选择，最多只有 1 种选择使得先手会输，
#           而先手必定可以不选择让自己输的，则先手必胜
#       2. 如果 lcnt + rcnt 是偶数，则由后手选择最后一个数字，
#           (1) lsum > rsum && lcnt >= rcnt ，先手必胜
#               先手在左边放 9 、右边放 0 ，
#               后手无论是在右边放 0 、还是在左边放 9 ，
#               lsum - rsum 都不会变得更小
#           (2) lsum < rsum && lcnt <= rcnt ，先手必胜
#               先手在左边放 0 、右边放 9 ，
#               后手无论是在右边放 9 、还是在左边放 0 ，
#               rsum - lsum 都不会变得更小
#           (3) lsum >= rsum && lcnt < rcnt ：
#               先手目的是使 lsum - rsum 尽可能大，则会优先在左边放 9 ，
#               后手目的是使 lsum - rsum 尽可能趋近于 0 ，则当先手在左边放 9 时，
#                   会在右边也放 9
#               这样当左边的问号没有时，此时 lsum - rsum 的大小没有改变，
#                   rcnt 也变为 rcnt - lcnt
#               接下来只能都在右边放数字
#                   ① 若 lsum - rsum = 9 * (rcnt >> 1) ，则后手必胜
#                       因为无论先手填入什么数字 x ，后手都可以填入 9 - x ，
#                       使得每一轮后， lsum - rsum 都减少 9 ，
#                       当进行全部 rcnt >> 1 轮后，lsum - rsum 刚好变为 0 ，
#                   ② 若 lsum - rsum > 9 * (rcnt >> 1) ，则先手必胜
#                       此时先手每次选 0 即可，这样每一轮 lsum - rsum 最多减少 9 ，
#                       最终必定还是正数
#                   ② 若 lsum - rsum < 9 * (rcnt >> 1) ，则先手必胜
#                       此时先手每次选 9 即可，这样每一轮 lsum - rsum 最少减少 9 ，
#                       最终必定是负数
#           (4) lsum <= rsum && lcnt > rcnt:
#               讨论方法如上
#
#       时间复杂度： O(n)
#       空间复杂度： O(1)


class Solution:
    def sumGame(self, num: str) -> bool:
        lcnt, rcnt = 0, 0
        lsum, rsum = 0, 0
        len_half = len(num) // 2

        for i, digit in enumerate(num):
            if digit == '?':
                if i < len_half:
                    lcnt += 1
                else:
                    rcnt += 1
            else:
                if i < len_half:
                    lsum += int(digit)
                else:
                    rsum += int(digit)

        # 奇数个问号，先手必胜
        if (lcnt + rcnt) & 1 == 1:
            return True

        # 如果和大的一边问号多，则先手必胜
        if (lsum > rsum and lcnt >= rcnt) or (lsum < rsum and lcnt <= rcnt):
            return True

        dsum = abs(lsum - rsum)
        # dcnt 必定是偶数
        dcnt = abs(lcnt - rcnt)
        return dsum != 9 * dcnt // 2
