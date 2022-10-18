# 链接：https://leetcode.com/problems/check-if-the-sentence-is-pangram/
# 题意：给定一个字符串 sentence ，判断是否全部英文字母都至少出现一次？


# 数据限制：
#  1 <= sentence.length <= 1000
#  sentence 仅由英文小写字母组成


# 输入： sentence = "thequickbrownfoxjumpsoverthelazydog"
# 输出： true

# 输入： sentence = "leetcode"
# 输出： false


# 思路： Set/Map
#
#      我们用一个集合维护所有出现的字母，最后判断集合的大小是否为 26 即可。
#
#
#      设字符集大小为 C 。
#
#      时间复杂度：O(n)
#          1. 需要遍历 sentence 中全部 O(n) 个字母一次
#      空间复杂度：O(C)
#          1. 只需要维护集合中全部 O(C) 种不同的字符


class Solution:
    def checkIfPangram(self, sentence: str) -> bool:
        # 用一个集合维护所有出现的字母。
        # 集合的大小是 26 时，所有字母至少出现了一次
        return len(set(sentence)) == 26
