# 链接：https://leetcode.com/problems/length-of-last-word/
# 题意：给定一个字符串 s ，返回最后一个单词的长度。


# 数据限制：
#  1 <= s.length <= 10 ^ 4
#  s 仅由英文字母和空格组成
#  s 中至少有一个单词


# 输入： s = "Hello World"
# 输出： 5
# 解释： 最后一个单词是 "World" ，长度为 5

# 输入： s = "   fly me   to   the moon  "
# 输出： 4
# 解释： 最后一个单词是 "moon" ，长度为 4

# 输入： s = "luffy is still joyboy"
# 输出： 6
# 解释： 最后一个单词是 "joyboy" ，长度为 6


# 思路： 模拟
#
#      从后往前遍历，过滤掉末尾的空格，统计最后一个单词的长度 ans ，
#      这时遇到第一个空格返回就表明最后一个单词长度统计完成，返回 ans 即可。
#
#      循环中未返回，则说明 s 中只有一个单词 且 s 无前导空格，
#      ans 就是这个单词的长度，直接返回即可。
#
#
#      时间复杂度：O(n)
#          1. 需要遍历完最后一个单词，最差情况下要遍历 s 中全部 O(n) 个字符
#      空间复杂度：O(1)
#          1. 只需要维护常数个额外变量


class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        # ans 表示最后一个单词的长度
        ans: int = 0
        for ch in reversed(s):
            if ch == ' ':
                if ans > 0:
                    # 如果 ch 是空格且遇到过单词，
                    # 则 ans 就是最后一个单词的长度，直接返回即可
                    return ans
                # 此时 ch 是空格，但还未遇到过单词，
                # 即还处于末尾空格中，不做处理
            else:
                # 统计最后一个单词的长度
                ans += 1
        # 此时 s 中只有一个单词 且 s 无前导空格，
        # ans 就是这个单词的长度，直接返回即可
        return ans
