# 链接：https://leetcode.com/problems/check-if-all-the-integers-in-a-range-are-covered/
# 题意：给定一组范围 ranges ，判断 [left, right] 中的每个数是否至少在其中一个范围中？

# 数据限制：
#   1 <= ranges.length <= 50
#   1 <= starti <= endi <= 50
#   1 <= left <= right <= 50

# 输入： ranges = [[1,2],[3,4],[5,6]], left = 2, right = 5
# 输出： true
# 解释： 2 在 [1,2] 中
#       3, 4 在 [3, 4] 中
#       5 在 [5,6] 中

# 输入： ranges = [[1,10],[10,20]], left = 21, right = 21
# 输出： false

# 思路： 枚举
#
#       没好好理解题意，连 WA 两次
#
#       枚举 [left, right] 中的每一个数，
#       如果其不在任何一个范围内，则直接返回 False
#       全部枚举完成后，返回 True
#
#       时间复杂度： O(m * n)
#       空间复杂度： O(1)


class Solution:
    def isCovered(self, ranges: List[List[int]], left: int, right: int) -> bool:
        for num in range(left, right + 1):
            for rang in ranges:
                if rang[0] <= num <= rang[1]:
                    break
            else:
                # 没有 break 就会走到这里，表明不在任何一个范围内
                return False
        return True
