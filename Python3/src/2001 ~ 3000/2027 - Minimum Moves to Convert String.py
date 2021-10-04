# 链接：https://leetcode.com/problems/minimum-moves-to-convert-string/
# 题意：给定一个只含有 'X' 和 'O' 的字符串，每次可以将任意连续长度为 3 的子串变为 "OOO" ，
#       求最少多少次可以将字符串变为只含有 'O' ？

# 数据限制：
#   3 <= s.length <= 1000
#   s[i] is either 'X' or 'O'


# 输入： s = "XXX"
# 输出： 1
# 解释： (XXX) -> OOO

# 输入： s = "XXOX"
# 输出： 2
# 解释： (XXO)X -> O(OOX) -> OOOO

# 输入： s = "OOOO"
# 输出： 0


# 思路： 贪心
#
#       我们枚举需要修改的字符串的起始位置 i ，
#       如果 s[i] 是 'X' ，则假装将 s[i:i+3] 变为 "OOO" ，
#       然后更新替换字符串的右边界 r = i + 2 ，
#       后续只有当 i > r 时，才会继续执行替换操作，
#       因为 i <= r 时，所有的字符都是 'O' 。
#
#       时间复杂度： O(n)
#       空间复杂度： O(1)


class Solution:
    def minimumMoves(self, s: str) -> int:
        # 记录操作次数
        ans = 0
        # 记录替换字符串的右边界
        r = -1
        # 枚举替换字符串的起始位置
        for i in range(len(s)):
            # 如果 i 不再替换字符串的右边界内，且 s[i] 是 'X' ，
            # 则假装将 s[i:i+3] 变为 "OOO" ，
            # 然后更新替换字符串的右边界 r = i + 2
            if i > r and s[i] == 'X':
                ans += 1
                r = i + 2
        return ans
