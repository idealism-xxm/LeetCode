# 链接：https://leetcode.com/problems/search-insert-position/
# 题意：给定不含重复数字的升序整数数组 nums 和一个整数 target ，
#      如果 target 在 nums 中，返回对应数字的下标，
#      否则返回 target 应该插入的下标。


# 数据限制：
#  1 <= nums.length <= 10 ^ 4
#  -(10 ^ 4) <= nums[i] <= 10 ^ 4
#  nums 中的数字各不相同且按升序排序
#  -(10 ^ 4) <= target <= 10 ^ 4


# 输入： nums = [1,3,5,6], target = 5
# 输出： 2
# 解释： 5 在 nums 中，下标为 2

# 输入： nums = [1,3,5,6], target = 2
# 输出： 1
# 解释： 2 不在 nums 中，插入下标为 1

# 输入： nums = [1,3,5,6], target = 7
# 输出： 4
# 解释： 7 不在 nums 中，插入下标为 4


# 思路： 二分
#
#      这个题其实就是需要找到 nums 中第一个大于等于 target 的数字的下标，
#      如果 target 比 nums 中的任何一个数字都大，那么返回 nums.length 。
#
#      那么我们直接按照这个思路二分即可，
#      找到 nums 中第一个大于等于 target 的数字。
#
#      我们定义左闭右闭的二分区间 [l, r] ，
#      其中 l 为左边界，初始化为 0 ， 
#      r 为右边界，初始化为 nums.length - 1 。
#
#      那么只要这个区间 [l, r] 不为空，我们就可以继续循环处理，
#      循环中先找到区间中点下标 mid = (l + r) / 2 ，
#      然后根据 nums[mid] 和 target 的大小关系进行处理：
#          1. nums[mid] < target ，那么区间 [l, mid] 内的数都小于 target ，
#              第一个大于等于 target 的元素必定在右边区间 [mid + 1, r] 中
#              此时需要将二分区间变为 [mid + 1, r]
#          2. nums[mid] >= target ，那么区间 [mid, r] 内的数都大于等于 target ，
#              则第一个大于等于 target 的元素必定在左边区间 [l, mid - 1] 中
#              （如果此时 mid 指向的元素，就是 nums 中第一个大于等于 target 的元素，
#                  那么在最后 l == r 时，会更新为 l = l + 1 ，
#                  最终选择到此时的 mid ）
#
#      最终 l 就是 nums 中第一个大于等于 target 的元素的下标
#
#      时间复杂度：O(logn)
#          1. 需要二分区间 [0, n - 1] ，时间复杂度为 O(logn)
#      空间复杂度：O(1)
#          1. 只需要维护常数个额外变量


class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        # 二分区间的左边界，初始化为 0
        l: int = 0
        # 二分区间的右边界，初始化为 len(nums) - 1
        r: int = len(nums) - 1
        # 当区间不为空时，继续二分
        # （注意这里取等号是因为我们的区间是左闭右闭区间，
        # 且收缩 r 时不取到 mid）
        while l <= r:
            # 计算区间中点下标
            mid: int = (l + r) >> 1
            if nums[mid] < target:
                # 如果区间中点的值小于 target ，
                # 则第一个大于等于 target 的元素
                # 必定在右边区间 [mid + 1, r] 中
                l = mid + 1
            else:
                # 如果区间中点的值大于等于 target ，
                # 则第一个大于等于 target 的元素
                # 必定在左边区间 [l, mid - 1] 中
                # （如果此时 mid 指向的元素，
                #   就是 nums 中第一个大于等于 target 的元素，
                #   那么在最后 l == r 时，会更新为 l = l + 1 ，
                #   最终选择到此时的 mid ）
                r = mid - 1

        # 此时 l 指向 nums 中
        # 第一个大于等于 target 的元素的下标
        return l
