# 链接：https://leetcode.com/problems/group-anagrams/
# 题意：给定一个小写字符串数组，按照字母异位词分组后返回？
#
#        字母异位词对一个源单词的字母顺序重新排列后得到的单词，
#        所有源单词中的字母恰好只用一次。


# 数据限制：
#  1 <= strs.length <= 10 ^ 4
#  0 <= strs[i].length <= 100
#  strs[i] 仅由英文小写字母组成


# 输入： strs = ["eat","tea","tan","ate","nat","bat"]
# 输出： [["bat"],["nat","tan"],["ate","eat","tea"]]

# 输入： strs = [""]
# 输出： [[""]]

# 输入： strs = ["a"]
# 输出： [["a"]]


# 思路1： Map + 排序
#
#      用一个 map 维护所有相同的字母异位词，
#      key 为对应字母异位词按字典序升序排序得到的字符串，
#      value 为字母异位词列表。
#
#      遍历 strs 中的每个字符串 s ，
#      对 s 按字典序升序排序得到对应的 key ，
#      然后将 s 放入 key 对应的字母异位词列表中即可。
#
#      最后收集所有的列表并返回。
#
#
#      设字符串的长度为 C 。
#
#      时间复杂度： O(n * ClogC)
#          1. 需要遍历 strs 中全部 O(n) 个字符串，
#              每次都要对字符串全部 O(C) 个字符进行排序，
#              排序时间复杂度为 O(ClogC) ，总时间复杂度为 O(n * ClogC)
#      空间复杂度： O(nC)
#          1. 需要维护 map 中全部不同的 key ，
#              最差情况下有 O(n) 个，每个大小为 O(C)
#          2. 需要维护结果中全部 O(n) 个字符串（复用原字符串，字符串不占额外空间）


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        # group 维护所有相同的字母异位词
        group: Dict[int, List[str]] = defaultdict(list)
        for s in strs:
            # 对 s 按照字典序升序排序，获取对应的 key
            sorted_s: str = "".join(sorted(s))
            # 将 s 放入对应的字母异位词列表
            group[sorted_s].append(s)
        
        # 收集所有的列表并返回
        return list(group.values())
