# 链接：https://leetcode.com/problems/remove-all-occurrences-of-a-substring/
# 题意：给定两个字符串 s 和 part ，不断执行以下操作：
#       删除 s 中最左侧的 part 子串，令 s 为最新的串，
#       直至无法执行该操作，求最终字符串？

# 数据限制：
#   1 <= s.length <= 1000
#   1 <= part.length <= 1000
#   s 和 part 均由英文小写字母组成

# 输入： s = "daabcbaabcbc", part = "abc"
# 输出： "dab"
# 解释：
#       "da[abc]baabcbc" -> "dabaabcbc"
#       "daba[abc]bc" -> "dababc"
#       "dab[abc]" -> "dab"

# 输入： s = "axxxxyyyyb", part = "xy"
# 输出： "ab"
# 解释：
#       "axxx[xy]yyyb" -> "axxxyyyb"
#       "axx[xy]yyb" -> "axxyyb"
#       "ax[xy]yb" -> "axyb"
#       "a[xy]b" -> "ab"

# 思路： KMP
#
#       数据限制支持暴力直接模拟，时间复杂度为 O(n * m)
#
#       当可以结合 KMP 将时间复杂度降低为 O(n + m)
#
#       我们先计算 part 的前缀函数，
#       然后使用一个 ans 字符数组记录当前还剩字符串，
#       并用 ans_pi 记录这些位置对应的前缀函数，
#       遍历 s 的每一个字符，将其放入 ans 中，
#       计算其前缀函数，并更新至 ans_pi ，
#
#       如果发现匹配完 part 后，则要移除 ans 中匹配的那些字符，
#       并更新下一个在 part 中匹配的位置
#
#       时间复杂度： O(n + m)
#       空间复杂度： O(n + m)

class Solution:
    def removeOccurrences(self, s: str, part: str) -> str:
        s_len, part_len = len(s), len(part)
        part_pi = self.get_pi(part)
        ans = [None] * s_len
        ans_pi = [0] * s_len
        ans_len = 0
        j = 0
        for ch in s:
            # 放入当前字符，开始进行匹配计算最长匹配长度
            ans[ans_len] = ch
            # 如果当前匹配的长度不为 0 ，且下一个字符不等，
            # 则继续找下一个最大长度
            while j > 0 and ch != part[j]:
                j = part_pi[j - 1]
            # 如果当前字符相等，则匹配长度 +1
            if ch == part[j]:
                j += 1
            # 更新 ans[:ans_len] 最长相等的真前缀和真后缀的长度
            ans_pi[ans_len] = j
            ans_len += 1
            # 如果完全匹配 part ，则需要从 ans 中移除该部分
            if j == part_len:
                ans_len -= part_len
                # 删除了 part 子串，下一次要从 ans_pi[ans_len - 1] 开始匹配
                j = ans_pi[ans_len - 1]

        # 剩余的部分就是没有匹配上的，连起来即可
        return "".join(ans[:ans_len])

    def get_pi(self, pattern: str) -> List[int]:
        """
        求模式串的 pi 数组，
        pi[i] 表示 parrern[:i] 最长相等的真前缀和真后缀的长度，
        即当前字符匹配失败后，下一次需要匹配的字符下标
        """
        pi = [0] * len(pattern)
        j = 0
        for i in range(1, len(pattern)):
            # 如果当前匹配的长度不为 0 ，且下一个字符不等，
            # 则继续找下一个最大长度
            while j > 0 and pattern[i] != pattern[j]:
                j = pi[j - 1]
            # 如果当前字符相等，则匹配长度 +1
            if pattern[i] == pattern[j]:
                j += 1
            # 更新 s[:i] 最长相等的真前缀和真后缀的长度
            pi[i] = j

        return pi
