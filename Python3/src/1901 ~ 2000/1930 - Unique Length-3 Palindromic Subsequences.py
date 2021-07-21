# 链接：https://leetcode.com/problems/unique-length-3-palindromic-subsequences/
# 题意：给定一个字符串，求所有长度为 3 的子序列中回文串的个数？

# 数据限制：
#   3 <= s.length <= 105
#   s 只含有英文小写字母

# 输入： s = "aabca"
# 输出： 3
# 解释：
#       "aba" (subsequence of "aabca")
#       "aaa" (subsequence of "aabca")
#       "aca" (subsequence of "aabca")

# 输入： nums = [1,3,2,1]
# 输出： [1,3,2,1,1,3,2,1]

# 思路： 枚举
#
#       由于长度为 3 的回文串首尾字母相同，所以可以枚举首尾字母 ch ，
#       找到该字母在 s 中第一次和最后一次出现的位置 first 和 last ，
#       然后统计 s[first+1:last] 中不同字母的个数，
#       这个个数就是首尾字母为 ch 且长度为 3 的回文串的个数
#
#       时间复杂度： O(n)
#       空间复杂度： O(n)


class Solution:
    def countPalindromicSubsequence(self, s: str) -> int:
        ans = 0
        for ch in string.ascii_lowercase:
            first, last = s.find(ch), s.rfind(ch)
            if first != -1:
                ans += len(set(s[first + 1 : last]))
        return ans
