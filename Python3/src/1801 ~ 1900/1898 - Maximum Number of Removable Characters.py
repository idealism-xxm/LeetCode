# 链接：https://leetcode.com/problems/maximum-number-of-removable-characters/
# 题意：给定两个字符串 s 和 p ，且 p 是 s 的一个子序列，
#       在给定一个 s 中可移除的字符下标数组 removable ，
#       求最大的 k 使得移除 s 中 removable[:k] 指定的所有字符后，
#       p 仍是 s 的子序列？

# 数据限制：
#   1 <= p.length <= s.length <= 105
#   0 <= removable.length < s.length
#   0 <= removable[i] < s.length
#   p 是 s 的子序列
#   p 和 s 都由英文小写字母组成
#   removable 中的所有下标都不相同

# 输入： s = "abcacb", p = "ab", removable = [3,1,0]
# 输出： 2
# 解释： "abcacb" -> "a_c_cb"

# 输入： s = "abcbddddd", p = "abcd", removable = [3,2,1,4,5,6]
# 输出： 1
# 解释： "abcbddddd" -> "abc_ddddd"

# 输入： s = "abcab", p = "abc", removable = [0,1,2,3,4]
# 输出： 0
# 解释： "abcab" -> "abcab"

# 思路： 二分 + 双指针
#
#       刚开始没注意到是连续的前缀，没想出太好的方法
#       发现是前缀后，那么就可以使用二分了
#
#       我们二分 k 的初始区间为 [l, r] = [0, len(removable)]，
#       每次使用双指针判断 p 是否还是 s 的子序列即可。
#       如果 p 还是 s 的子序列，那么下次二分区间为 [mid + 1, r]
#       如果 p 不是 s 的子序列，那么下次二分区间为 [l, mid - 1]
#
#       当 l > r 时结束二分，此时有 l - 1 == r，
#       结果就是 l - 1 ，
#       因为最开始 l = 0 是合法的，而每次都不会再考虑全部合法的区间，
#       最终 l 必定在不合法区间的开始位置
#
#       时间复杂度： O((|s| + |p|) * logn)
#       空间复杂度： O(n)


class Solution:
    def maximumRemovals(self, s: str, p: str, removable: List[int]) -> int:
        l, r = 0, len(removable)
        while l <= r:
            mid = (l + r) >> 1
            # 如果 p 还是 s 的子序列，那么下次二分区间为 [mid + 1, r]
            if self.is_subsequence(s, p, set(removable[:mid])):
                l = mid + 1
            else:
                # 如果 p 已不是 s 的子序列，那么下次二分区间为 [l, mid - 1]
                r = mid - 1
        return l - 1

    def is_subsequence(self, s: str, p: str, removable: Set[int]) -> bool:
        i = 0
        l = len(s)
        # 遍历 p 中每一个字符
        for ch in p:
            # 找到 s 中第一个匹配的合法字符
            while i in removable or (i < l and ch != s[i]):
                i += 1
            # 如果没有找到，则返回 False
            if i == l:
                return False
            # s 中当前字符已用，移动到下一个位置
            i += 1
        return True
