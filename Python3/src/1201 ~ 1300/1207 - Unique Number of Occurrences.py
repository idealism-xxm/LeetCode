# 链接：https://leetcode.com/problems/unique-number-of-occurrences/
# 题意：给定一个整型数组 arr ，如果不同数的出现次数也各不相同，
#      则返回 true ；否则返回 false 。


# 数据限制：
#  1 <= arr.length <= 1000
#  -1000 <= arr[i] <= 1000


# 输入： arr = [1,2,2,1,1,3]
# 输出： true
# 解释： 1 出现 3 次；
#       2 出现 2 次；
#       3 出现 1 次。
#
#       不同数的出现次数各不相同。

# 输入： arr = [1,2]
# 输出： false
# 解释： 1 和 2 都出现了 1 次。

# 输入： arr = [-3,0,1,-3,1,1,1,-3,10,0]
# 输出： true
# 解释：  1 出现 4 次；
#       -3 出现 3 次；
#        0 出现 2 次；
#       10 出现 1 次。
#
#       不同数的出现次数各不相同。


# 思路： Map + Set
#
#      先用一个名为 num_to_cnt 的 map 统计每个数字的出现次数。
#
#      再用一个名为 cnt_set 的 set 统计不同的出现次数。
#
#      那么当且仅当 num_to_cnt 和 cnt_set 的大小一样时，
#      不同数的出现次数各不相同。
#
#
#      时间复杂度：O(n)
#          1. 需要遍历全部 O(n) 个数字
#      空间复杂度：O(n)
#          1. 需要维护 num_to_cnt 中全部 O(n) 个数字的出现次数
#          2. 需要维护 cnt_set 中全部 O(sqrt(n)) 个不同的出现次数


class Solution:
    def uniqueOccurrences(self, arr: List[int]) -> bool:
        # 统计每个数字的出现次数
        num_to_cnt: Counter = Counter(arr)

        # 收集每个数字的出现次数到集合中
        cnt_set: Set[int] = set(num_to_cnt.values())
        # 当且仅当 num_to_cnt 和 cnt_set 的大小一样时，不同数的出现次数各不相同
        return len(num_to_cnt) == len(cnt_set)
