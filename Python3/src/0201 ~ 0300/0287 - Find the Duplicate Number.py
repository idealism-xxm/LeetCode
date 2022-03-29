# 链接：https://leetcode.com/problems/find-the-duplicate-number/
# 题意：给定一个长度为 n + 1 的数组，所有的数都在 [1, n] 范围内。
#      其中只有一个数字出现两次或以上，其他数字都只出现一次，
#      返回这个重复出现的数字。
#
#      要求：不修改原数组，只使用 O(1) 的额外空间。


# 数据限制：
#  1 <= n <= 10 ^ 5
#  nums.length == n + 1
#  1 <= nums[i] <= n
#  只有一个数字出现两次或以上，其他数字都只出现一次


# 输入： nums = [1,3,4,2,2]
# 输出： 2

# 输入： nums = [3,1,3,4,2]
# 输出： 3


# 思路： 二分
#
#      根据鸽巢原理可得： n + 1 个范围在 [1, n] 中的数字，
#      其中至少有一个数字出现两次。
#
#      题目中有一个关键信息：重复的数字可能出现不止两次，
#      也就是说我们无法通过计算总和来确定重复的数字。
#
#      题目也保证只有一个数字会重复出现，
#      所以我们可以通过二分找到这个重复的数字。
#
#      我们二分这个重复的数字，初始区间为 [1, n] 。
#
#      循环中每次计算区间中点的数字 mid = (l + r) >> 1 ，
#      然后统计 nums 中小于等于 mid 的数字的个数 count ：
#          1. count > mid: 说明重复的数字在左边区间中，
#              令 r = mid - 1 ，继续二分。
#          2. count <= mid: 说明重复的数字在右边区间中
#              令 l = mid + 1 ，继续二分。
#
#      二分结束后， l 就是所求的重复数字。
#
#
#      时间复杂度：O(nlogn)
#          1. 需要二分区间 [1, n] ，时间复杂度为 O(logn)
#          2. 每次二分时，都需要遍历 nums 中全部 O(n) 个数字
#      空间复杂度：O(1)
#          1. 只需要维护常数个额外变量


class Solution:
    def findDuplicate(self, nums: List[int]) -> int:
        # 二分区间左边界，初始化为 1
        l: int = 1
        # 二分区间右边界，初始化为 n
        r: int = len(nums) - 1
        # 当前区间不为空时，继续二分
        while l <= r:
            # 计算区间中点 mid
            mid: int = (l + r) >> 1
            # 统计 nums 中小于等于 mid 的数字个数
            count = sum(1 for num in nums if num <= mid)

            if count > mid:
                # count > mid, 说明重复的数字在左边区间中
                r = mid - 1
            else:
                # count <= mid, 说明重复的数字在右边区间中
                l = mid + 1

        # l 就是重复的数字
        return l