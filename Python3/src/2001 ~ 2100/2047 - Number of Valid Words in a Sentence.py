# 链接：https://leetcode.com/problems/number-of-valid-words-in-a-sentence/
# 题意：给定一个句子 sentence ，只含有小写字母，数字和 '-', '!', '.', ',', ' ' 。
#       每个句子都可以被一个或多个空格分割成一个或多个 token ，
#       每个满足以下三个条件的 token 是一个合法的单词，
#           1. 只含有小写字母， '-' 以及 '!', '.', ','
#           2. 最多有一个 '-' ，且 '-' 两边都是小写字母
#           3. 最多有一个 '!' 或 '.' 或 ',' ，且其必定在最后
#       求句子中合法的单词数量？

# 数据限制：
#   1 <= sentence.length <= 1000
#   sentence 只含有英文小写字母，数字 和 '-', '!', '.', ',', ' '
#   至少有一个 token

# 输入： sentence = "cat and  dog"
# 输出： 3
# 解释： 合法的单词有： "cat", "and", "dog"

# 输入： sentence = "!this  1-s b8d!"
# 输出： 0
# 解释： 没有合法的单词

# 输入： sentence = "alice and  bob are playing stone-game10"
# 输出： 5
# 解释： 合法的单词有： "alice", "and", "bob", "are", "playing"

# 输入： sentence = "he bought 2 pencils, 3 erasers, and 1  pencil-sharpener."
# 输出： 6
# 解释： 合法的单词有： "he", "bought", "pencils", "erasers", "and", "pencil-sharpener."


# 思路： 模拟
#
#       先按照空格分成不同的 token ，因为有多个连续空格，所以 token 可能会是空串，
#       然后按照题意判断 token 是否满足三个条件即可。
#
#       时间复杂度： O(n)
#       空间复杂度： O(n)


class Solution:
    def countValidWords(self, sentence: str) -> int:
        # 统计合法的单词数
        ans = 0
        # 将句子分成 token
        for word in sentence.split(' '):
            # 如果 token 满足三个条件，则是一个合法的单词
            if self.is_valid(word):
                ans += 1
            
        return ans
    
    def is_valid(self, word):
        # 空串不是合法的 token
        if len(word) == 0:
            return False

        # 记录 '-' 的数量
        hypen_count = 0
        # 记录 '!', ',', '.' 的数量
        punct_count = 0
        for i, ch in enumerate(word):
            # 如果字符是数字，则不满足题意
            if ch.isdigit():
                return False
            # 如果是 '-'
            if ch == '-':
                # 如果不是第一次出现，或者两边有非小写字母，则不满足题意
                if hypen_count == 1 or i == 0 or i == len(word) - 1 or not word[i - 1].isalpha() or not word[i + 1].isalpha():
                    return False
                # 记录 '-' 的数量
                hypen_count += 1
            # 如果是 '!', ',', '.'
            if ch in ('!', ',', '.'):
                # 如果不是第一次出现，或者不是最后一个字符，则不满足题意
                if punct_count == 1 or i != len(word) - 1:
                    return False
                # 记录 '!', ',', '.' 的数量
                punct_count += 1
        return True
