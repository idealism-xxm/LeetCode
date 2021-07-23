# 链接：https://leetcode.com/problems/maximum-number-of-words-you-can-type/
# 题意：给定一段文本 text ，现在一直键盘坏掉的字母列表 brokenLetters ，
#       求能完整打出的单词个数？

# 数据限制：
#   1 <= text.length <= 10 ^ 4
#   0 <= brokenLetters.length <= 26
#   text 由任意个单词构成，每两个单词之间有一个空格，首尾不含空格
#   每个单词均由小写字母组成
#   brokenLetters 由不同的小写字母组成

# 输入： text = "hello world", brokenLetters = "ad"
# 输出： 1

# 输入： text = "leet code", brokenLetters = "lt"
# 输出： 1

# 输入： text = "leet code", brokenLetters = "e"
# 输出： 0


# 思路： 模拟
#
#       按照题意模拟即可
#
#       先按空格分成多个单词，然后判断每个单词中是否存在字母坏掉，
#       如果不存在，则对结果 +1
#
#       时间复杂度： O(n)
#       空间复杂度： O(n)


class Solution:
    def canBeTypedWords(self, text: str, brokenLetters: str) -> int:
        brokenLetters = set(brokenLetters)
        ans = 0
        for word in text.split(' '):
            if not (set(word) & brokenLetters):
                ans += 1
        return ans
