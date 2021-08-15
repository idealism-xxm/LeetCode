# 链接：https://leetcode.com/problems/number-of-strings-that-appear-as-substrings-in-word/
# 题意：给定一组模式串 patterns 和一个单词 word ，返回是 word 子串的模式串的个数？

# 数据限制：
#   1 <= patterns.length <= 100
#   1 <= patterns[i].length <= 100
#   1 <= word.length <= 100
#   patterns[i] 和 word 仅由小写字母组成

# 输入： patterns = ["a","abc","bc","d"], word = "abc"
# 输出： 3
# 解释： 
#       "a" 是 "abc" 的子串
#       "abc" 是 "abc" 的子串
#       "bc" 是 "abc" 的子串
#       "d" 不是 "abc" 的子串

# 输入： patterns = ["a","b","c"], word = "aaaaabbbbb"
# 输出： 2
# 解释： 
#       "a" 是 "aaaaabbbbb" 的子串
#       "b" 是 "aaaaabbbbb" 的子串
#       "c" 不是 "aaaaabbbbb" 的子串

# 输入： patterns = ["a","a","a"], word = "ab"
# 输出： 3
# 解释： 
#       "a" 是 "ab" 的子串


# 思路： 模拟
#
#       直接遍历 patterns ，然后判断每个 pattern 是否为 word 的子串即可
#       
#       时间复杂度： O(sum(|pattern|) * |word|)
#       空间复杂度： O(1)

class Solution:
    def numOfStrings(self, patterns: List[str], word: str) -> int:
        return sum(1 if pattern in word else 0 for pattern in patterns)
