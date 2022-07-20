# 链接：https://leetcode.com/problems/number-of-matching-subsequences/
# 题意：给定一个字符串 s 和一个单词数组 words ，
#      求 words 中有多少个单词是 s 的子序列？


# 数据限制：
#  1 <= s.length <= 5 * 10 ^ 4
#  1 <= words.length <= 5000
#  1 <= words[i].length <= 50
#  s 和 words[i] 仅含有英文小写字母


# 输入： s = "abcde", words = ["a","bb","acd","ace"]
# 输出： 3
# 解释： 单词 "a", "acd", "ace" 是 s 的子序列

# 输入： s = "dsahjpjauf", words = ["ahjpjau","ja","ahbwzgqnuk","tnmlanowax"]
# 输出： 2
# 解释： 单词 "ahjpjau", "ja" 是 s 的子序列


# 思路： 贪心 + DP
#
#      我们很容易就能想到枚举所有单词 word ，判断其是否为 s 的子序列。
#
#      针对每个单词都枚举其中的字母 ch ，
#      每次都贪心地在 s[idx..] 中找到 ch 第一次出现的下标 next_idx 。
#
#      因为如果找第二次或者后续出现的下标， s 中剩余可用于匹配的字符会变少，答案不会更优。
#
#      针对找到的下标 next_idx ，有以下两种情况：
#          1. next_idx == len(s): 即 ch 在 s[idx..] 中不存在，则 word 不是 s 的子序列，
#              直接处理下一个单词
#          2. next_idx < len(s): 更新 idx = next_idx + 1 ，需要继续循环处理，
#              即需要判断单词剩余的字符是否为 s[next_idx + 1..] 的子序列
#
#      但直接这样暴力枚举的时间复杂度为 O(m * (l + n)) ，在本题的数据范围下无法通过。
#
#      可以注意到刚刚的思路中有一个贪心的逻辑：找到 s[idx..] 中字母 ch 第一次出现的下标。
#
#      如果我们能在 O(1) 内知道 s[idx..] 中字母 ch 第一次出现的下标，
#      那么时间复杂度就能优化为 O(ml) 了。
#
#      这个信息可以通过 DP 处理后得到，
#      设 dp[i][j] 表示 s[i..] 中字母 ch 第一次出现的下标。
#
#      初始化为： dp[i][j] = len(s) ，表示 s[i..] 中不存在字母 j 。
#      状态转移：倒着更新 dp[i] 即可，因为 s[i..] 只比 s[i+1..] 多了一个字母，
#              那么 dp[i] 与 dp[i + 1] 除了 s[j] 的下标不同以外均相同。
#              即先将 dp[i + 1] 的值全部复制到 dp[i] ，然后令 dp[i][s[i]] = i 即可。
#
#      这部分处理的时间复杂度为 O(nC) ，整体时间复杂度为 O(nc + ml) ，
#      同时空间复杂度上升为 O(nC) ，典型的空间换时间的优化。
#
#
#      设 s 的长度为 n ， words 的长度为 m ，单词的最长长度为 l ，
#      字符集大小为 C 。
#
#		时间复杂度： O(nC + ml)
#          1. 需要遍历 s 中全部 O(n) 个字母，每次都要处理 O(C) 个状态
#          2. 需要遍历 words 中全部 O(m) 个单词，
#              每次都要遍历单词中全部 O(l) 个字母
#		空间复杂度： O(nC)
#          1. 需要维护一个大小为 O(nC) 二维数组 dp


class Solution:
    def numMatchingSubseq(self, s: str, words: List[str]) -> int:
        n: int = len(s)
        # dp[i][j] 表示 s[i..] 中，字母 j 第一次出现的下标，
        # 不存在则为 n
        dp: List[List[int]] = [None for _ in range(n + 1)]
        dp[-1] = [n] * 26
        for i in range(n - 1, -1, -1):
            # dp[i] 与 dp[i + 1] 除了 s[j] 的下标不同以外均相同
            dp[i] = dp[i + 1].copy()
            # dp[i] 中只有字母 j 的下标需要更新
            dp[i][(ord(s[i]) - ord('a'))] = i

        ans: int = 0
        # 遍历每个单词，判断是否是 s 的子序列
        for word in words:
            # 判断 word 是否是 s 的子序列，默认为是子序列
            idx: int = 0
            is_subsequence: int = True
            for ch in word:
                # 找到当前字母在 s[idx..] 中第一次出现的下标 next_idx
                next_idx: int = dp[idx][ord(ch) - ord('a')]
                # 如果不存在，则标记为不是子序列，并跳出循环
                if next_idx == n:
                    is_subsequence = False
                    break

                # 如果存在，则下一次需要判断剩余子串是否为 s[next_idx+1..] 的子序列
                idx = next_idx + 1

            # 如果是子序列，则计入到 ans 中
            if is_subsequence:
                ans += 1

        return ans
