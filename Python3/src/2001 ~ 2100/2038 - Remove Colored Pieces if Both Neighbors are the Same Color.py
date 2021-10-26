# 链接：https://leetcode.com/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color/
# 题意：给定一个 AB 字符串 colors ， 其中 colors[i] 是 'A' 或者 'B' 。
#       现在两个人需要轮流从中按照以下规则删除字符，如果一个人不能删除一个字符，那么他就会输掉游戏。
#       1. 先手只能删除字符 'A' ，而且 'A' 的左右两侧必须都是 'A' 
#       2. 后手只能删除字符 'B' ，而且 'B' 的左右两侧必须都是 'B'
#       3. 先手和后手都不能删除开始和结尾的字符
#       假设先手和后手都采用最优策略，判断先手是否必胜？

# 数据限制：
#   1 <= colors.length <= 105
#   colors 仅有字母 'A' 和 'B' 组成

# 输入： colors = "AAABABB"
# 输出： true
# 解释：
#   先手： A(A)ABABB -> AABABB
#   后手： 不能删除任何一个 'B'

# 输入： colors = "AA"
# 输出： false
# 解释：
#   先手：不能删除任何一个 'A'

# 输入： colors = "ABBBBBBBAAA"
# 输出： false
# 解释：
#   先手： ABBBBBBBA(A)A -> ABBBBBBBAA
#   后手： ABBBBB(B)BAA -> ABBBBBBAA
#   先手：不能删除任何一个 'A'


# 思路： 计数
#
#       由于先手和后手的字符不一样，且判断条件也不相关，所以就直接统计先手和后手各能删除多少次即可。
#       如果先手能删除的次数大于后手的，那么先手必胜，否则后手必胜。
#
#       时间复杂度： O(n)
#       空间复杂度： O(1)


class Solution:
    def winnerOfGame(self, colors: str) -> bool:
        # 统计先手和后手能删除的次数
        a_count, b_count = 0, 0
        for i in range(1, len(colors) - 1):
            # 如果当前字符是 'A' ，且其左右两侧都是 'A' ，那么先手可以删除当前字符
            if colors[i] == 'A' and colors[i - 1] == 'A' and colors[i + 1] == 'A':
                a_count += 1
            elif colors[i] == 'B' and colors[i - 1] == 'B' and colors[i + 1] == 'B':
                # 如果当前字符是 'B' ，且其左右两侧都是 'B' ，那么后手可以删除当前字符
                b_count += 1

        # 如果先手能删除的次数比后手的多，那么先手必胜，否则后手必胜
        return a_count > b_count
