# 链接：https://leetcode.com/problems/check-if-string-is-a-prefix-of-array/
# 题意：给定一个字符串 s 和 一个单词数组，
#       求单词数组中前 k 个按顺序组成的字符串是否等于 s ？

# 数据限制：
#   1 <= words.length <= 100
#   1 <= words[i].length <= 20
#   1 <= s.length <= 1000
#   words[i] 和 s 只由英文小写字母组成

# 输入： s = "iloveleetcode", words = ["i","love","leetcode","apples"]
# 输出： true
# 解释： 前 3 个单词组成了 s ，所以返回 true 。

# 输入： s = "iloveleetcode", words = ["apples","i","love","leetcode"]
# 输出： false
# 解释： 前 i 个单词组成的字符串都不等于 s ，所以返回 false 。  


# 思路： 模拟
#
#       刚开始理解错题意，导致连 WA 两次，
#       我们不断将新单词拼接起来，然后判断其是否和 s 相等即可
#       
#       时间复杂度： O(len(s))
#       空间复杂度： O(1)

class Solution:
    def isPrefixString(self, s: str, words: List[str]) -> bool:
        i = 0
        # 遍历每个单词
        for word in words:
            # 遍历每个字母
            for ch in word:
                # 如果 s 已经遍历完 或者 当前字符不是 s[i] ，则不满足题意
                if i >= len(s) or s[i] != ch:
                    return False
                # 此时满足题意，继续判断下一个字符
                i += 1
            # 如果当前单词遍历完，正好遍历完 s ，则满足题意
            if i == len(s):
                return True
        # 如果没有遍历完 s ，则不满足题意
        return False
