#  链接：https://leetcode.com/problems/detect-capital/
#  题意：给定一个单词 word ，判断是否满足以下三个条件之一：
#           1. 全部字母都是大写，例如 "USA"
#           2. 全部字母都是小写，例如 "leetcode"
#           3. 首字母大写，后续字母都是小写，例如 "Google"

#  数据限制：
#   1 <= word.length <= 100
#   word 仅由英文大小写字母组成

#  输入： word = "USA"
#  输出： true

#  输入： word = "FlaG"
#  输出： false


#  思路：模拟
# 
#       直接按照题意模拟判断即可，
#       可以注意到第 2 个和第 3 个条件可以融合成一种情况：除首字母外全是小写字母。
#       （如果首字母是小写，则是第 2 个条件；如果首字母是大写，则是第 3 个条件。）
# 
#       时间复杂度： O(n)
#       空间复杂度： O(1)


class Solution:
    def detectCapitalUse(self, word: str) -> bool:
        # 1. 全是大写： word.isupper()
        # 2. 首字母大写，其余全是小写
        # 3. 首字母小写，其余全是小写
        # 
        # 可以发现 2 和 3 可以合并，除首字母外都是小写就满足题意。
        # 所以可以用 word[1:].islower() 判断
        # 
        # 综上，最简单解法如下：
        #   return word.isupper() or word[1:].islower()
        #
        # 但由于 str 取切片是拷贝，所以空间复杂度会上升为 O(n)
        # 可以将 word[1:].islower() 
        # 优化为 all(word[i].islower() for i in range(1, len(word)))
        return word.isupper() or all(
            word[i].islower() for i in range(1, len(word))
        )
