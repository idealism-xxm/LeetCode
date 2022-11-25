# 链接：https://leetcode.com/problems/apply-operations-to-an-array/
# 题意：给定一个长度为 n 的非负整数数组 nums ，需要按顺序对其进行 n - 1 次操作。
#      第 i 次操作逻辑如下：
#          1. nums[i] == nums[i + 1]: 
#              则令 nums[i] = nums[i] * 2; nums[i + 1] = 0
#          2. nums[i] != nums[i + 1]: 跳过本次操作
#
#      执行完全部操作后，将全部的 0 移动到数组末尾，保持其他数相对顺序不变，
#      最后返回结果数组。


# 数据限制：
#  2 <= nums.length <= 2000
#  0 <= nums[i] <= 1000


# 输入： nums = [1,2,2,1,1,0]
# 输出： [1,4,2,0,0,0]
# 解释： 依次执行以下操作：
#       · i = 0: nums[0] != nums[1], 跳过本次操作
#       · i = 1: nums[1] == nums[2], 则令 nums[1] = 2 * 2 = 4, nums[2] = 0 ，
#              nums 变为 [1,4,0,1,1,0]
#       · i = 2: nums[2] != nums[3], 跳过本次操作
#       · i = 3: nums[3] == nums[4], 则令 nums[3] = 1 * 2 = 2, nums[4] = 0 ，
#              nums 变为 [1,4,0,2,0,0]
#       · i = 4: nums[4] != nums[5], 跳过本次操作
#       · i = 5: nums[5] == nums[6], 则令 nums[5] = 0 * 2 = 0, nums[2] = 0 ，
#              nums 变为 [1,4,0,2,0,0]
#
#       最后将全部的 0 移动到数组末尾， nums 变为 [1,4,2,0,0,0]


# 输入： nums = [0,1]
# 输出： [1,0]
# 解释： 依次执行以下操作：
#       · i = 0: nums[0] != nums[1], 跳过本次操作
#
#       最后将全部的 0 移动到数组末尾， nums 变为 [1,0]


# 思路： 双指针
#
#      本题分为两步，第一步按照题意模拟即可，第二步将 0 移动到数组末尾。
#
#      其中第二步就是 LeetCode 283 的逻辑，可以直接复用思路和代码。
#
#      第二步时，我们维护两个指针 l 和 r 。
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
#          1. 需要遍历 nums 中全部 O(n) 个数字两次
#      空间复杂度：O(1)
#          1. 只需要维护常数个额外变量即可


class Solution:
    def applyOperations(self, nums: List[int]) -> List[int]:
        # 按顺序执行全部 n - 1 次操作
        for i in range(len(nums) - 1):
            if nums[i] == nums[i + 1]:
                nums[i] <<= 1
                nums[i + 1] = 0

        # l 表示不为 0 的数字个数，也是下一个可以放入数字的下标，初始化为 0
        l: int = 0
        # 遍历所有的数字
        for r in range(len(nums)):
            # 如果当前数字不等于 0 ，则 nums[r] 不需要移动至数组末尾，
            # 与 nums[l] 交换，并右移 l
            if nums[r] != 0:
                nums[l], nums[r] = nums[r], nums[l]
                l += 1

        return nums
