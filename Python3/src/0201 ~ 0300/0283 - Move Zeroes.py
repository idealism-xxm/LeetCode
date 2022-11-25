# 链接：https://leetcode.com/problems/move-zeroes/
# 题意：给定一个整数数组 nums ，原地将全部的 0 移动到数组末尾，
#      保持其他数相对顺序不变。


# 数据限制：
#  1 <= nums.length <= 10 ^ 4
#  -(2 ^ 31) <= nums[i] <= 2 ^ 31 - 1


# 输入： nums = [0,1,0,3,12]
# 输出： [1,3,12,0,0]

# 输入： nums = [0]
# 输出： [0]


# 思路： 双指针
#
#      本题是 LeetCode 27 的加强版，将移除变为了移动，
#      所以需要交换数字，而非直接覆盖。
#
#      我们维护两个指针 l 和 r 。
#
#      l 表示不为 0 的数字个数，也是下一个可以放入数字的下标。
#      初始化为 0 ，表示目前还没遇到不为 0 的数字。
#
#      r 表示下一个待考虑的数字的下标，初始化为 0 。
#
#      不断右移 r ，如果 nums[r] != 0 ，
#      则交换 nums[l] 和 nums[r] 的值，以达到移动 0 的效果，然后右移 l 。
#
#      交换 nums[l] 和 nums[r] 时，只存在两种情况：
#          1. nums[l] == 0: 则必定有 l < r ，交换后，相当于将 0 移动到 r 处
#          2. nums[l] != 0: 则必定有 l == r ，交换后， nums 保持不变
#
#
#      时间复杂度：O(n)
#          1. 需要遍历 nums 中全部 O(n) 个数字
#      空间复杂度：O(1)
#          1. 只需要维护常数个额外变量即可


class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        # l 表示不为 0 的数字个数，也是下一个可以放入数字的下标，初始化为 0
        l: int = 0
        # 遍历所有的数字
        for r in range(len(nums)):
            # 如果当前数字不等于 0 ，则 nums[r] 不需要移动至数组末尾，
            # 与 nums[l] 交换，并右移 l
            if nums[r] != 0:
                nums[l], nums[r] = nums[r], nums[l]
                l += 1
