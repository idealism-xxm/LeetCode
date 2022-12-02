# 链接：https://leetcode.com/problems/determine-if-string-halves-are-alike/
# 题意：给定一个偶数长度的字符串 s ，将其分成长度相等的两个子串。
#      判断第一个子串和第二个子串是否相似？
#
#      当且仅当两个字符串含有的元音字母数相同时，它们相似。
#      元音字母有：'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'


# 数据限制：
#  2 <= s.length <= 1000
#  s.length 是偶数
#  s 仅有大小写英文字母组成


# 输入： s = "book"
# 输出： true
# 解释： 两个子串分别为 "bo" 和 "ok" 。
#       前者含有一个元音字母 o ，后者也含有一个元音字母 o 。

# 输入： s = "textbook"
# 输出： false
# 解释： 两个子串分别为 "text" 和 "book" 。
#       前者含有一个元音字母 e ，后者也含有两个元音字母 o 和 o 。


# 思路： Set/Map
#
#      根据题意统计 s 的前一半子串和后一半子串的元音字母数，判断是否相等即可。
#
#      可以用一个元音字母集合判断一个字母是否为元音，
#      也可以硬编码成 switch/match/if 等语句进行判断。
#
#
#      时间复杂度：O(n)
#          1. 需要遍历 s 中全部 O(n) 个字母
#      空间复杂度：O(1)
#          1. 只需要维护常数个变量


# 初始化全部元音字母的集合
VOWELS: Set[str] = set("aeiouAEIOU")


class Solution:
    def halvesAreAlike(self, s: str) -> bool:
        n: int = len(s)
        # 如果 s 前一半子串和后一半子串的元音字母数相同，则满足题意
        return Solution.count_vowels(s[:n>>1]) == Solution.count_vowels(s[n>>1:])

    @staticmethod
    def count_vowels(s: str) -> int:
        # 统计元音字母数
        return sum(
            1
            for ch in s
            # 过滤出元音字母
            if ch in VOWELS
        )
