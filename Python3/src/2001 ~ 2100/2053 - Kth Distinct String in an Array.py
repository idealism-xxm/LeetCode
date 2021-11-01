# 链接：https://leetcode.com/problems/kth-distinct-string-in-an-array/
# 题意：给定一个字符串数组，求其中第 k 个唯一的字符串，不存在则返回空串。

# 数据限制：
#   1 <= k <= arr.length <= 1000
#   1 <= arr[i].length <= 5
#   arr[i] 只含有英文小写字母

# 输入： arr = ["d","b","c","b","c","a"], k = 2
# 输出： "a"
# 解释： 
#   唯一的字符串只有： "d", "a"
#   第 2 个唯一的字符串是 "a"

# 输入： arr = ["aaa","aa","a"], k = 1
# 输出： "aaa"
# 解释： 
#   唯一的字符串只有： "aaa", "aa", "a"
#   第 1 个唯一的字符串是 "aaa"

# 输入： arr = ["a","b","a"], k = 3
# 输出： ""
# 解释： 
#   唯一的字符串只有： "b"
#   不足 3 个唯一的字符串，所以返回 ""


# 思路： 计数
#
#       按照题意模拟即可，先统计所有字符串出现的次数，
#       然后遍历原数组，找到第 k 个唯一的字符串并返回，如果不存在，则返回 ""
#
#       时间复杂度： O(n)
#       空间复杂度： O(n)


class Solution:
    def kthDistinct(self, arr: List[str], k: int) -> str:
        # 统计每个字符串出现的次数
        cnt = Counter(arr)
        # 计算当前是第几个唯一的字符串
        i = 0
        # 遍历原数组
        for s in arr:
            # 如果当前字符串出现的次数为 1，则说明其是唯一的字符串
            if cnt[s] == 1:
                # 如果是第 k 个唯一的字符串，则直接返回
                i += 1
                if i == k:
                    return s
        # 如果找不到第 k 个唯一的字符串，则返回 ""
        return ""
