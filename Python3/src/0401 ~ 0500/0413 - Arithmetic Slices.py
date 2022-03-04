# 链接：https://leetcode.com/problems/arithmetic-slices/
# 题意：给定一个整型数组 nums ，求其中满足以下两个条件的子数组个数？
#          1. 长度大于等于 3
#          2. 子数组是一个等差数列


# 数据限制：
#   1 <= nums.length <= 5000
#   -1000 <= nums[i] <= 1000


# 输入： nums = [1,2,3,4]
# 输出： 3
# 解释： [1, 2, 3], [2, 3, 4] 和 [1,2,3,4] 满足题意

# 输入： nums = [1]
# 输出： 0
# 解释： 没有满足题意的子数组


# 思路： 滑动窗口
#
#       我们维护一个滑动窗口 [l, r] ，
#       使得 nums[l:r + 1] 内的数是一个差为 diff 的等差数列，
#       然后我们不断循环处理。
#   
#       在循环中，我们先计算当前滑动窗口内最后两个数的差
#       last_diff = nums[r] - nums[r - 1] 。
#   
#       如果 last_diff != diff ，则说明现在 nums[l:r + 1] 已不是等差数列，
#       需要移动左指针 l 到 r - 1 ，使得滑动窗口内的数仍是等差数列，
#       差为 nums[r] - nums[r - 1] 。
#   
#       然后，只要滑动窗口的长度大于等于 3 ，
#       则存在以 nums[r] 为结尾的满足题意的子数组。
#   
#       nums[l:r + 1], nums[l + 1:r + 1], ..., nums[r - 2:r + 1] 都满足题意，
#       这些子数组的个数为 (r - 2) - l + 1 = r - l - 1 ，统计入 ans 即可。
#   
#       最后，右移右指针 r ，扩大滑动窗口，并进入下一个循环中。
#
#
#       时间复杂度： O(n)
#           1. 只需要遍历全部 O(n) 个数字一次
#       空间复杂度： O(1)
#           1. 只需要维护常数个额外的变量


class Solution:
    def numberOfArithmeticSlices(self, nums: List[int]) -> int:
        # 如果长度不足 3 ，则必定不满足题意，直接返回 0
        if len(nums) < 3:
            return 0

        # 定义左指针 l ，指向当前等差数列子数组的起始下标，初始化为 0
        l = 0
        # 定义右指针 r ，指向当前等差数列子数组的结束下标，初始化为 1
        r = 0
        # 定义当前滑动窗口 [0, 1] 内的等差数列的差，初始化为 nums[1] - nums[0]
        diff = nums[1] - nums[0]
        # ans 维护满足题意的子数组个数，初始化为 0
        ans = 0
        # 如果还有数字，则继续处理
        while r < len(nums):
            # 计算此时最后两个数的差
            last_diff = nums[r] - nums[r - 1]
            # 如果差不相等，则需要移动左指针 l ，缩小滑动窗口
            if diff != last_diff:
                # l 移动至 r - 1 ，滑动窗口变为 [r - 1, r]
                l = r - 1
                # 滑动窗口内等差数列的差为 nums[r] - nums[r - 1]
                diff = last_diff

            # 现在 nums[l:r + 1] 是一个等差数列，差为 diff 。
            #
            # 如果 nums[l:r + 1] 的长度大于等于 3 ，
            # 则存在以 nums[r] 为结尾的满足题意的子数组。
            if r - l + 1 >= 3:
                # nums[l:r + 1], nums[l + 1:r + 1], ..., 
                # nums[r - 2:r + 1] 都满足题意，
                # 这些子数组的个数为 (r - 2) - l + 1 = r - l - 1
                ans += r - l - 1

            # 右移右指针 r ，扩大滑动窗口
            r += 1

        return ans
