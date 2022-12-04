# 链接：https://leetcode.com/problems/determine-if-two-strings-are-close/
# 题意：给定两个字符串，判断它们是否接近？
#
#      两个字符串接近，当且仅当其中一个字符串能用以下任意次操作变成另一个字符串：
#          1. 交换任意两个字符的位置。例如： "abcde" -> "aecdb"
#          2. 选择任意两种字符 x 和 y ，将所有的 x 替换为 y ，
#             将所有的 y 替换为 x 。例如： "aacabb" -> "bbcbaa"


# 数据限制：
#  1 <= word1.length, word2.length <= 10 ^ 5
#  word1 和 word2 仅含有英文小写字母


# 输入： word1 = "abc", word2 = "bca"
# 输出： true
# 解释： word1 执行以下 2 次操作后，可以变为 word2 ：
#       · 执行操作 1: "abc" -> "acb"
#       · 执行操作 1: "acb" -> "bca"

# 输入： word1 = "a", word2 = "aa"
# 输出： false
# 解释： word1 和 word2 不能任意次操作相互转换。

# 输入： word1 = "cabbba", word2 = "abbccc"
# 输出： true
# 解释： word1 执行以下 3 次操作后，可以变为 word2 ：
#       · 执行操作 1: "cabbba" -> "caabbb"
#       · 执行操作 2: "caabbb" -> "baaccc"
#       · 执行操作 2: "baaccc" -> "abbccc"


# 思路： Map + Set
#
#      操作 1 和操作 2 都不能新增或减少字符，所以以下两种情况，
#      必定无法通过任意次操作相互转换：
#          (1) word1 和 word2 长度不相等
#          (2) word1 和 word2 字符集不同
#
#      其中 (1) 可以直接比较长度判断， (2) 需要收集字符串的全部字符到集合中，
#      再判断集合是否相等。
#
#      操作 1 的效果是交换字符的位置，所以必须保证相同字符的出现次数都相同时，
#      才能用操作 1 实现两个字符串相互转换。
#
#      操作 2 的效果是交换字符的出现次数，所以必须保证相同出现次数的个数相同时，
#      才能用操作 2 转换为两个字符串相同字符的出现次数都相同。
#
#
#      设字符集大小为 C 。
#
#      时间复杂度：O(n + C)
#          1. 需要遍历 word1 和 word2 全部 O(n) 个字符
#          2. 需要遍历 ch_to_cnt 全部 O(C) 个不同字符
#      空间复杂度：O(C)
#          1. 需要维护 word1 和 word2 全部 O(C) 个不同字符的集合和统计信息


class Solution:
    def closeStrings(self, word1: str, word2: str) -> bool:
        # 操作 1 和操作 2 都不能新增或减少字符，所以 word1 和 word2 长度不相等时，
        # 必定无法通过任意次操作相互转换。
        if len(word1) != len(word2):
            return False
        
        # 获取两个字符串的不同字符的出现次数
        word1_ch_to_cnt: Counter = Counter(word1)
        word2_ch_to_cnt: Counter = Counter(word2)
        # 如果字符集不同，则必定无法相互转换
        if word1_ch_to_cnt.keys() != word2_ch_to_cnt.keys():
            return False

        # 获取两个字符串的统计信息
        word1_ch_cnt_to_cnt: Counter = Counter(word1_ch_to_cnt.values())
        word2_ch_cnt_to_cnt: Counter = Counter(word2_ch_to_cnt.values())
 
        # 当两个字符串统计信息相等时，可以相互转换
        return word1_ch_cnt_to_cnt == word2_ch_cnt_to_cnt
