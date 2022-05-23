# 链接：https://leetcode.com/problems/palindromic-substrings/
# 题意：给定一个字符串 s ，求所有回文子串的数量？


# 数据限制：
#  1 <= s.length <= 1000
#  s 仅有英文小写字母组成


# 输入： s = "abc"
# 输出： 3
# 解释： 有 3 个子回文串： "a", "b", "c"

# 输入： s = "aaa"
# 输出： 6
# 解释： 有 6 个子回文串： "a", "a", "a", "aa", "aa", "aaa"


# 思路1： 枚举
#
#      可以发现如果 s[l:r+1] 是回文子串，
#      只需要判断 s[l-1] 与 s[r+1] 是否相等，
#      就能在 O(1) 内判断 s[l-1:r+2] 是否是回文子串。
#
#      我们可以枚举回文子串的中心点，由于字符串的长度可能是奇数和偶数，
#      而每次往外扩展时，都会新增两个字符，子串的奇偶性不会改变。
#
#      所以我们需要分别枚举中心点为一个字符 s[i:i+1] 和两个字符 s[i:i+2] 的情况，
#      找到这两种情况下的回文子串的数量。
#
#      为了方便处理，我们定义一个函数 count(s, l, r) ，
#      用来统计以 s[l:r+1] 为中心的回文子串数量。
#
#      这样只需要遍历一遍字符串的下标 i ，
#      将 count(s, i, i) 和 count(s, i, i + 1) 加到 ans 中即可。
#
#
#      时间复杂度：O(n ^ 2)
#          1. 需要枚举全部 O(n) 个中心，每次枚举时都需要往外扩展，
#              最差情况下每次都需要遍历全部 O(n) 个字符。
#      空间复杂度：O(1)
#          1. 只需要使用常数个额外变量


class Solution:
    def countSubstrings(self, s: str) -> int:
        # ans 表示 s 中所有回文子串的数量
        ans: int = 0
        # 枚举回文子串的中心
        for i in range(len(s)):
            # 加上以 s[i] 为中心的回文子串数量
            ans += Solution.count(s, i, i)
            # 加上以 s[i:i+2] 为中心的回文子串数量
            ans += Solution.count(s, i, i + 1)

        return ans

    # 统计以 s[l:r+1] 为中心的回文子串数量
    @staticmethod
    def count(s: str, l: int, r: int) -> int:
        # ans 表示以 s[l:r+1] 为中心的回文子串数量
        ans: int = 0
        # 如果 l 和 r 合法，且 s[l] == s[r]，
        # 则 s[l:r+1] 是一个回文子串，令 ans += 1
        while l >= 0 and r < len(s) and s[l] == s[r]:
            ans += 1
            # 左右分别向外扩展一个字符
            l -= 1
            r += 1

        return ans


# 思路2： 区间 DP
#
#      可以发现如果 s[l:r+1] 是回文子串，
#      只需要判断 s[l-1] 与 s[r+1] 是否相等，
#      就能在 O(1) 内判断 s[l-1:r+2] 是否是回文子串。
#
#      那么我们只要保证已经确定所有长度小于 length 的子串是否为回文子串，
#      就可以在 O(1) 内判断长度为 length 的某个子串是否为回文子串。
#
#      设 dp[l][r] 表示 s[l:r+1] 是否为回文子串，初始化为 false。
#
#      所有长度为 1 的子串为回文子串，
#      即 dp[r][r] = true; ans += 1;
#
#      如果相邻的两个字符 s[r-1] 和 s[r] 相等，
#      则 s[r-1:r+1] 是长度为 2 的回文子串，
#      即 dp[r-1][r] = true; ans += 1;
#
#      状态转移：枚举长度为 length 的子串，再枚举子串的结束下标 r ，
#      可计算出子串的开始下标 l = r - length + 1 。
#
#      如果此时 s[l] == s[r] 且 s[l+1:r] 是回文子串，
#      则 s[l:r+1] 是长度为 length 的回文子串，
#      即 dp[l][r] = true; ans += 1;
#
#
#      时间复杂度：O(n ^ 2)
#          1. 需要枚举全部 O(n) 个子串长度，
#              每次枚举时都需要枚举子串的全部 O(n) 个结束下标
#      空间复杂度：O(n ^ 2)
#          1. 需要维护一个 O(n ^ 2) 的二维数组 dp


class Solution:
    def countSubstrings(self, s: str) -> int:
        n: int = len(s)
        # ans 表示 s 中所有回文子串的数量
        ans: int = 0
        # dp[l][r] 表示 s[l:r+1] 是否是回文子串，初始化为 false
        dp: List[List[int]] = [[False] * n for _ in range(n)]
        # s[0:0+1] 是长度为 1 的回文子串
        dp[0][0] = True
        ans += 1
        # 枚举回文子串的中心
        for r in range(1, n):
            # s[r:r+1] 是长度为 1 的回文子串
            dp[r][r] = True
            ans += 1
            # 如果 s[r - 1] == s[r] ，则 s[r-1:r+1] 是长度为 2 的回文子串
            if s[r - 1] == s[r]:
                dp[r - 1][r] = True
                ans += 1

        # 枚举回文子串的长度
        for length in range(3, n + 1):
            # 枚举回文子串的结束下标
            for r in range(length-1, n):
                # 计算对应的开始下标
                l: int = r - length + 1
                # 如果 s[l] == s[r] 且 s[l+1:r] 是回文子串，
                # 则 s[l:r+1] 是长度为 length 的回文子串
                if s[l] == s[r] and dp[l + 1][r - 1]:
                    dp[l][r] = True
                    ans += 1

        return ans
