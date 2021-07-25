# 链接：https://leetcode.com/problems/check-if-all-characters-have-equal-number-of-occurrences/
# 题意：给定一个字符串，判断是否每个出现过的字符都出现了相同次数？


# 数据限制：
#   1 <= s.length <= 1000
#   s 只由英文小写字母组成

# 输入： s = "abacbc"
# 输出： true
# 解释： 'a', 'b', 'c' 都只出现了 2 次

# 输入： s = "aaabb"
# 输出： false
# 解释： 'a' 出现了 3 次， 'b' 只出现了 2 次


# 思路： 模拟
#
#       按照题意统计每个字符出现的次数，最后判断是否所有出现次数相等即可
#
#       时间复杂度： O(n)
#       空间复杂度： O(n)

class Solution:
    def areOccurrencesEqual(self, s: str) -> bool:
        cnt = defaultdict(int)
        for ch in s:
            cnt[ch] += 1
        return len(set(cnt.values())) == 1
