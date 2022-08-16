# 链接：https://leetcode.com/problems/first-unique-character-in-a-string/
# 题意：给定一个字符串 s ，求其中第一个不重复的字符下标？如果不存在则返回 -1 。


# 数据限制：
#  1 <= s.length <= 10 ^ 5
#  s 只包含英文小写字母


# 输入： s = "leetcode"
# 输出： 0
# 解释： 字母 'l' 是第一个不重复的字符

# 输入： s = "loveleetcode"
# 输出： 2
# 解释： 字母 'v' 是第一个不重复的字符

# 输入： s = "aabb"
# 输出： -1
# 解释： 不存在不重复的字符


# 思路： Map
#
#      我们直接用一个名为 ch_to_cnt 的 map 维护每个字符的出现次数。
#
#      先遍历一遍 s ，统计每个字符的出现次数。
#
#      然后再遍历一遍 s ，遇到第一个出现次数为 1 的字符，直接返回其下标即可。
#
#      如果最后还没有返回，则没有不重复的字符，返回 -1 。
#
#
#      设 C 为字符集大小。
#
#      时间复杂度：O(n)
#          1. 需要遍历 s 中全部 O(n) 个字符两次
#      空间复杂度：O(1)
#          1. 需要维护全部 O(C) 个不同字符的出现次数


class Solution:
    def firstUniqChar(self, s: str) -> int:
        # ch_to_cnt[ch] 表示 ch 出现的次数
        ch_to_cnt: Counter = Counter(s)

        # 再遍历一遍 s，找到第一个出现次数为 1 的字符，并返回其下标
        for i, ch in enumerate(s):
            if ch_to_cnt[ch] == 1:
                return i

        # 此时没有不重复的字符，返回 -1
        return -1
