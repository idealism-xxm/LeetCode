# 链接：https://leetcode.com/problems/substrings-of-size-three-with-distinct-characters/
# 题意：给定一个字符串 s ，求其中所有长度为 3 且不含重复字符的子串的个数？

# 输入： s = "xyzzaz"
# 输出： 1

# 输入： s = "aababcabc"
# 输出： 4

# 思路： 枚举
#       从开始枚举所有长度为 3 的子串，没有重复字符时就对结果 +1
#
#       如果长度 3 也是一个变量 k ，
#       那么使用滑动窗口可以保持时间复杂度为 O(n) ，空间复杂度为 O(k)
#
#       时间复杂度： O(n)
#       空间复杂度： O(1)

class Solution:
    def countGoodSubstrings(self, s: str) -> int:
        cnt = 0
        for i in range(2, len(s)):
            if s[i] != s[i - 1] and s[i] != s[i - 2] and s[i - 1] != s[i - 2]:
                cnt += 1
        return cnt
