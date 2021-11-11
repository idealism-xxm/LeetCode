# 链接：https://leetcode.com/problems/vowels-of-all-substrings/
# 题意：给定一个字符串，求每个子串中元音个数的和？

# 数据限制：
#   1 <= word.length <= 10 ^ 5
#   word 仅含有英文小写字母

# 输入： word = "aba"
# 输出： 6
# 解释： 
#   0 个元音： "b"
#   1 个元音： "a", "ab", "ba", "a"
#   2 个元音： "aba"
#
#   总共 0 + 1 * 4 + 2 = 6

# 输入： word = "abc"
# 输出： 3
# 解释： 
#   0 个元音： "b", "bc", "c"
#   1 个元音： "a", "ab", "abc"
#
#   总共 0 * 3 + 1 * 3 = 3


# 输入： word = "ltcd"
# 输出： 0
# 解释： 
#   没有元音

# 输入： word = "noosabasboosa"
# 输出： 237


# 思路： 计数
#
#       直接按照题意求解比较复杂，时间复杂度也不允许，所以需要思考其他方式。
#
#       我们可以发现计数是加法，而最后还要求这些加法结果的和，
#       那么我们可以转换思路，运用加法交换律和结合律，
#       可以转换为求拥有某一个元音的子串的个数，将所有这些个数加在一起就是最终结果。
#
#       假设 s[i] 是元音，那么包含 s[i] 的子串的开始下标必定在 [0, i] 之间，
#       结束下标必定在 [i, len(s)) 之间，
#       那么可以包含 s[i] 的子串个数 = (i + 1) * (len(s) - i)
#
#       时间复杂度： O(n)
#       空间复杂度： O(1)


VOWELS = set('aeiou')


class Solution:
    def countVowels(self, word: str) -> int:
        n = len(word)
        # 统计最终结果
        ans = 0
        # 枚举每个字母
        for i, ch in enumerate(word):
            # 如果该字母是元音
            if ch in VOWELS:
                # 该元音对答案的贡献 = 含有该元音的子串的个数
                # 如果一个子串开始下标在 [0, i] 内，结束下标在 [i, n) 内，
                # 那这个子串就包含字符 s[i] ，
                # 这样的子串数有 (i + 1) * (n - i) 个
                ans += (i + 1) * (n - i)
        return ans
