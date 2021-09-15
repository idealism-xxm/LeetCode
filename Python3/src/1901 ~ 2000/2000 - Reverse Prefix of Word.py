# 链接：https://leetcode.com/problems/reverse-prefix-of-word/
# 题意：给定一个字符串 word 和一个字符 ch ，翻转第一个 ch 及其之前的子串。

# 数据限制：
#   1 <= word.length <= 250
#   word 由英文小写字母组成
#   ch 是英文小写字母

# 输入： word = "abcdefd", ch = "d"
# 输出： "dcbaefd"
# 解释：
#   翻转子串 "abcd" 后得到 "dcba" ，新字符串为 "dcbaefd"

# 输入： word = "xyxzxe", ch = "z"
# 输出： "zxyxxe"
# 解释：
#   反转子串 "xyxz" 后得到 "zxyx" ，新字符串为 "zxyxxe"

# 输入： word = "abcd", ch = "z"
# 输出： "abcd"
# 解释：
#   没有字符 "z" ，不需要翻转


# 思路： 模拟
#
#       按照题意找到第一个 ch 后，翻转前缀串即可
#
#       时间复杂度： O(n)
#       空间复杂度： O(1)


class Solution:
    def reversePrefix(self, word: str, ch: str) -> str:
        for i in range(len(word)):
            # 如果 ch 存在，则翻转前缀串
            if word[i] == ch:
                return word[i::-1] + word[i + 1:]
        return word
