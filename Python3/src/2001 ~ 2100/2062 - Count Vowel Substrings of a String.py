# 链接：https://leetcode.com/problems/count-vowel-substrings-of-a-string/
# 题意：给定一个字符串，求有多少个子串仅含有全部的 5 个元音字母？

# 数据限制：
#   1 <= word.length <= 100
#   word 仅含有英文小写字母

# 输入： word = "aeiouu"
# 输出： 2
# 解释： 
#   - "(aeiou)u"
#   - "(aeiouu)"

# 输入： word = "unicornarihan"
# 输出： 0
# 解释： 
#   没有 5 个元音


# 输入： word = "cuaieuouac"
# 输出： 7
# 解释： 
#   - "c(uaieuo)uac"
#   - "c(uaieuou)ac"
#   - "c(uaieuoua)c"
#   - "cu(aieuo)uac"
#   - "cu(aieuou)ac"
#   - "cu(aieuoua)c"
#   - "cua(ieuoua)c"

# 输入： word = "bbaeixoubb"
# 输出： true
# 解释： 
#   含有全部 5 个元音的子串都有辅音


# 思路1： 枚举
#
#       枚举所有的子串，判断是否仅含有全部 5 个元音即可，
#       如果只含有全部 5 个原因，则对 ans + 1 ，
#       如果遇辅音，则从下一个起始位置开始枚举
#
#       时间复杂度： O(n ^ 2)
#       空间复杂度： O(1)


VOWEL = set('aeiou')


class Solution:
    def countVowelSubstrings(self, word: str) -> int:
        # 满足题意的子串数
        ans = 0
        # 枚举子串起始下标
        for i in range(len(word)):
            # 收集子串中的元音
            cur = set()
            # 枚举子串的结束下标
            for j in range(i, len(word)):
                # 如果不是元音，则开始处理下一个起始位置
                if word[j] not in VOWEL:
                    break
                # 当前元音加入 cur 中
                cur.add(word[j])
                # 如果现在元音有 5 个，则满足题意
                if len(cur) == 5:
                    ans += 1
        return ans
