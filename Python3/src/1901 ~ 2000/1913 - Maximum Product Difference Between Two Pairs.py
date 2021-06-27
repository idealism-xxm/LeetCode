# 链接：https://leetcode.com/problems/maximum-product-difference-between-two-pairs/
# 题意：给定一个整型数组 nums，定义乘积差为 (w * x) - (y * z) ，
#      w, x, y, z 不同下标的四个数，求乘积差的最大值？

# 数据限制：
#   4 <= nums.length <= 10 ^ 4
#   1 <= nums[i] <= 10 ^ 4

# 输入： nums = [5,6,2,7,4]
# 输出： (6 * 7) - (2 * 4) = 34

# 输入： nums = [4,2,5,9,7,4,8]
# 输出： (9 * 8) - (2 * 4) = 64.

# 思路： 贪心
#
#       因为所有的数都是正数，所以可以直接排序后，
#       用最大的两个数的乘积减去最小的两个数的乘积即可
#
#       时间复杂度： O(nlogn)
#       空间复杂度： O(1)


class Solution:
    def maxProductDifference(self, nums: List[int]) -> int:
        nums.sort()
        return nums[-1] * nums[-2] - nums[0] * nums[1]
