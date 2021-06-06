# 链接：https://leetcode.com/problems/check-if-word-equals-summation-of-two-words/
# 题意：给定三个只含有 a~j 的字符串，分别代表数字 0~9 ，
#       求前两个字符串代表的数加起来是否等于第三个字符串代表的数字？

# 输入： firstWord = "acb", secondWord = "cba", targetWord = "cdb"
# 输出： true
# 解释： "acb" -> "021" -> 21
#       "cba" -> "210" -> 210
#       "cdb" -> "231" -> 231
#       21 + 210 == 231

# 输入： firstWord = "aaa", secondWord = "a", targetWord = "aab"
# 输出： false
# 解释： "aaa" -> "000" -> 0
#       "a" -> "0" -> 0
#       "aab" -> "001" -> 1
#       0 + 0 != 1

# 输入： firstWord = "aaa", secondWord = "a", targetWord = "aaaa"
# 输出： false
# 解释： "aaa" -> "000" -> 0
#       "a" -> "0" -> 0
#       "aaaa" -> "0000" -> 0
#       0 + 0 == 0

# 思路： 模拟
#
#       直接按照题意转成十进制数判断即可
#
#       时间复杂度： O(1)
#       空间复杂度： O(1)

mp = {
    "a": 0,
    "b": 1,
    "c": 2,
    "d": 3,
    "e": 4,
    "f": 5,
    "g": 6,
    "h": 7,
    "i": 8,
    "j": 9,
}


class Solution:
    def isSumEqual(self, firstWord: str, secondWord: str, targetWord: str) -> bool:
        return self.get(firstWord) + self.get(secondWord) == self.get(targetWord)

    def get(self, s: str) -> int:
        it = 0
        for ch in s:
            it = it * 10 + mp[ch]
        return it
