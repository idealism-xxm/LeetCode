# 链接：https://leetcode.com/problems/percentage-of-letter-in-string/
# 题意：给定一个字符串 s 和一个字符 letter ，
#      求 letter 在 s 中出现次数的百分比（向下取整到最近的整数）？


# 数据限制：
#  1 <= s.length <= 100
#  s consists of lowercase English letters.
#  letter is a lowercase English letter.


# 输入： s = "foobar", letter = "o"
# 输出： 33
# 解释： 字符串 s 共有 6 个字符， 字符 'o' 出现了 3 次，
#       则百分比为 2 / 6 * 100% = 33%

# 输入： s = "jjjj", letter = "k"
# 输出： 0
# 解释： 字符串 s 共有 4 个字符， 字符 'k' 在 s 中没有出现，
#       则百分比为 0%


# 思路： 模拟
#
#      统计字符 letter 在 s 中出现的次数 cnt ，
#      那么百分比为 cnt / len(s) * 100% 。
#
#      由于需要向下取整，所以直接使用整数除法即可，计算公式为：
#      cnt * 100 / len(s) ，其中 / 表示整数除法。
#
#
#      时间复杂度：O(n)
#          1. 需要遍历 s 中全部 O(n) 个字符
#      空间复杂度：O(1)
#          1. 只需要使用常数个额外变量


class Solution:
    def percentageLetter(self, s: str, letter: str) -> int:
        # 统计字符串 s 中字符 letter 的个数
        cnt = sum(1 for ch in s if ch == letter)
        # 计算字符 letter 在字符串 s 中的百分比
        return cnt * 100 // len(s)
