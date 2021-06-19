# 链接：https://leetcode.com/problems/redistribute-characters-to-make-all-strings-equal/
# 题意：给定 n 个单词，每次可以将其中一个单词的一个字符移动到另一个单词的任意位置，
#       求在一些操作后，是否可以将 n 个单词变成一样的单词？

# 数据限制：
#   1 <= words.length <= 100
#   1 <= words[i].length <= 100
#   words[i] 全部由英文小写字母组成

# 输入： words = ["abc","aabc","bc"]
# 输出： true
# 解释： ["abc","aabc","bc"] -> ["abc","abc","abc"]

# 输入： words = ["ab","a"]
# 输出： false

# 思路： 统计
#
#       统计每个字符出现的次数，如果所有字符的出现次数都整除 n ，
#       那么可以达到目标
#
#       由于只有小写字母，所以最多只会记录 26 个字符的出现次数
#
#       时间复杂度： O(sum(|word|))
#       空间复杂度： O(1)


class Solution:
    def makeEqual(self, words: List[str]) -> bool:
        cnt = defaultdict(int)
        # 统计每个单词的每个字符出现的次数
        for word in words:
            for ch in word:
                cnt[ch] += 1
        # 如果有一个字符不能整除 n ，则不满足题意
        for value in cnt.values():
            if value % len(words) != 0:
                return False
        return True
