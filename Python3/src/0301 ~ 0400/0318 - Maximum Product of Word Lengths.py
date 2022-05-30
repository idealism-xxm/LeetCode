# 链接：https://leetcode.com/problems/maximum-product-of-word-lengths/
# 题意：给定一个字符串数组 words ，返回 len(words[i]) * len(words[j]) 的最大值，
#      其中 words[i] 和 words[j] 不含相同字母。
#      如果不存在这样的单词对，返回 0 。


# 数据限制：
#  2 <= words.length <= 1000
#  1 <= words[i].length <= 1000
#  words[i] 仅含有英文小写字母


# 输入： words = ["abcw","baz","foo","bar","xtfn","abcdef"]
# 输出： 16
# 解释： "abcw" 与 "xtfn" 不含相同字母，乘积为 4 * 4 = 16

# 输入： words = ["a","ab","abc","d","cd","bcd","abcd"]
# 输出： 4
# 解释： "ab" 与 "cd" 不含相同字母，乘积为 2 * 2 = 4

# 输入： words = ["a","aa","aaa","aaaa"]
# 输出： 0
# 解释： 不存在满足题意的单词对


# 思路： 枚举 + 状态压缩
#
#      题意很清晰，很容易就能想到题意枚举的做法，
#      使用 ans 维护最大乘积，初始化为 0 。
#
#      先枚举单词对的第一个单词 words[i] (0 <= i < len(words)) ，
#      然后获取其字母集合 letter_set[i] 。
#
#      再枚举单词对的另一个单词 words[j] (0 <= j < i) ，
#      其字母集合 letter_set[j] 已在前面计算过，可以直接使用。
#
#      如果 letter_set[i] 与 letter_set[j] 的交集为空集，
#      则 words[i] 与 words[j] 不含相同字母，
#      可以直接计算乘积，更新 ans 的最大值。
#
#      这样做的时间复杂度为 O(n ^ 2 * L) ，空间复杂度为 O(n * L) ，
#      其中 L 为单词的最大长度。
#
#      但单词只含有英文小写字母，可以使用状态压缩处理，
#      使用 mask[i] 表示第 i 个单词的字母使用状态，
#      其中第 j 位为 1 表示第 j 个字母出现在第 i 个单词中。
#
#      这样能将时间复杂度优化为 O(n * L + n ^ 2) ，空间复杂度优化为 (n) 。
#
#
#      设 L 为单词的最大长度。
#
#      时间复杂度：O(n * L + n ^ 2)
#          1. 需要遍历全部 O(n) 个单词，并遍历每个单词全部 O(n) 个字母
#          2. 需要遍历全部 O(n ^ 2) 个单词对
#      空间复杂度：O(n)
#          1. 需要维护全部 O(n) 个单词的字母使用状态


class Solution:
    def maxProduct(self, words: List[str]) -> int:
        # ans 维护最大积，初始化为 0
        ans: int = 0
        # mask[i] 表示第 i 个单词的字母使用状态，
        # 其中第 j 位为 1 表示第 j 个字母出现在第 i 个单词中
        mask: List[int] = [0] * len(words)
        for i in range(len(words)):
            # 计算第 i 个单词的字母使用状态
            for ch in words[i]:
                # 第 ch - 'a' 个字母出现在第 i 个单词中
                mask[i] |= 1 << (ord(ch) - ord('a'))

            # 枚举单词对 (i, j)
            for j in range(i):
                # 如果两者不含相同字母，则状态按位与就是 0
                if mask[i] & mask[j] == 0:
                    # 更新最大积
                    ans = max(ans, len(words[i]) * len(words[j]))

        return ans
