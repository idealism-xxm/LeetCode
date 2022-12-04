# 链接：https://leetcode.com/problems/sort-characters-by-frequency/
# 题意：给定一个字符串 s ，对其按照字符出现次数降序排序。
#
#      注意：相同字符必须在一起。


# 数据限制：
#  1 <= s.length <= 5 * 10 ^ 5
#  s 仅由英文大小写字母和数字组成


# 输入： s = "tree"
# 输出： "eert"
# 解释： 'e' 出现两次， 't' 和 'r' 各出现一次。
#       所以 'e' 应该在 't' 和 'r' 之前， 
#       "eert" 和 "eetr" 都满足题意。

# 输入： s = "cccaaa"
# 输出： "aaaccc"
# 解释： 'c' 和 'a' 各出现三次，
#       "cccaaa" 和 "aaaccc" 都满足题意。
#
#       注意： "cacaca" 不满足题意，因为相同字符必须在一起。

# 输入： s = "Aabb"
# 输出： "bbAa"
# 解释： 'b' 出现两次， 'A' 和 'a' 各出现一次。
#       "bbAa" 和 "bbaA" 都满足题意。
#
#       注意： 'A' 和 'a' 是两种不同的字符。


# 思路： Map + 排序
#
#      本题是 LeetCode 1636 加强版，将数组换成了字符串，
#      使用相同的思路即可通过。
#
#
#      先用一个 map 统计 s 中每个字符的出现次数。
#
#      然后对 s 中的字符按照出现次数降序排序，
#      出现次数相同时，按字符升序排序（以保证相同字符在一起）。
#
#      最后转成字符串返回即可。
#
#
#      设字符集大小为 C 。
#
#      时间复杂度： O(nlogn)
#          1. 需要遍历 s 中全部 O(n) 个字符
#          2. 需要对 s 中全部 O(n) 个字符排序，时间复杂度为 O(nlogn)
#      空间复杂度： O(n + C)
#          1. 需要维护全部 O(C) 个不同字符的出现次数
#          2. 需要维护字符数组和结果字符串中全部 O(n) 个字符


class Solution:
    def frequencySort(self, s: str) -> str:
        # ch_to_cnt[ch] 表示 s 中 ch 的出现次数
        ch_to_cnt: Counter = Counter(s)

        # 对 s 中的字符按照出现次数降序排序，
        # 出现次数相同时，按字符升序排序（以保证相同字符在一起）
        chs: List[str] = list(s)
        chs.sort(key=lambda ch: (-ch_to_cnt[ch], ch))

        # 转成字符串返回
        return "".join(chs)
