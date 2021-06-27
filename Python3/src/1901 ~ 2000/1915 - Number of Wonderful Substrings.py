# 链接：https://leetcode.com/problems/number-of-wonderful-substrings/
# 题意：给定一个字符串 s ，如果一个子串中出现奇数次的字符的个数不超过一个，
#       那么这个子串是最美子串，求所有的最美子串？

# 数据限制：
#   1 <= word.length <= 10 ^ 5
#   word 中的每个字符都由 'a' ~ 'j' 这 10 个英文小写字母中

# 输入： word = "aba"
# 输出： 4
# 解释： 有 4 个最美子串
#       "[a]ba" -> "a"
#       "a[b]a" -> "b"
#       "ab[a]" -> "a"
#       "aba" -> "[aba]"

# 输入： word = "aabb"
# 输出： 9
# 解释： 有 9 个最美子串
#       "[a]abb" -> "a"
#       "[aa]bb" -> "aa"
#       "[aab]b" -> "aab"
#       "[aabb]" -> "aabb"
#       "a[a]bb" -> "a"
#       "a[abb]" -> "abb"
#       "aa[b]b" -> "b"
#       "aa[bb]" -> "bb"
#       "aab[b]" -> "b"

# 思路： 状态压缩 + 前缀异或和
#
#       比赛只是想到了一种状压 DP ，时间复杂度 O(n * 2 ^ k) ，其中 k = 10
#       状态定义：dp[i][j] 表示以 s[i] 结尾的子串中，不同字符出现奇偶数状态的数量
#       初始化： dp[0][j] = 0
#       状态转移：s[i] 对应字符为 digit ，那么 dp[i][j ^ (1 << digit)] = dp[i][j]
#               然后再加上当前字符形成的子串： dp[i][1 << digit] += 1
#
#       先写了一遍 Python 的，结果没过，但又想不出其他的解法，
#       就写了一个 Golang 的， 3200ms 通过
#
#       比赛后看了大家的解法，才发现使用前缀异或和的方式避免每次更新所有的状态，
#       prefix[j] 表示 s[:1] ~ s[:i] 这 i 个子串中 j 状态出现的次数，
#       当遍历到第 s[i] 时，我们会得到当前 s[:i] 的状态 state ，
#       由于最美子串的的状态共有 0, 1 << 0, 1 << 1, ..., 1 << 9 这 11 个，
#       所以我们只要找到能使当前状态 state 变为这 11 个的状态即可，
#       这 11 个状态就是 state, state ^ (1 << 0), ..., state ^ (1 << 9) ，
#       那么当前就能计算出以 s[i] 为结尾的子串能产生的最美子串的个数。
#
#
#       时间复杂度： O(n * k), 其中 k = 10
#       空间复杂度： O(2 ^ k)


class Solution:
    def wonderfulSubstrings(self, word: str) -> int:
        prefix = [0] * 1024
        # 要计算空串的状态，因为要考虑从第一个字符开始的子串
        prefix[0] = 1
        # 当前 s[:i] 的状态
        state = 0
        ans = 0
        for ch in word:
            # 获取对应的数位
            digit = ord(ch) - ord('a')
            # 更新 s[:i] 的状态
            state ^= 1 << digit
            # 首先记录 state 产生的最美子串数
            ans += prefix[state]
            for j in range(10):
                # 再依次记录 state ^ (1 << j) 产生的最美子串数
                ans += prefix[state ^ (1 << j)]
            # 更新当前状态出现的次数
            prefix[state] += 1
        return ans
