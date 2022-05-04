# 链接：https://leetcode.com/problems/shortest-unsorted-continuous-subarray/
# 题意：给定一个整数数组 nums ，对其一个子数组按升序排序，
#      使得 nums 变成升序排序的数组，求子数组的最小长度。


# 数据限制：
#  1 <= nums.length <= 10 ^ 4
#  -(10 ^ 5) <= nums[i] <= 10 ^ 5


# 输入： nums = [2,6,4,8,10,9,15]
# 输出： 5
# 解释： 需要排序子数组 [6, 4, 8, 10, 9] ，
#       使得 nums 变为升序排序的数组。

# 输入： nums = [1,2,3,4]
# 输出： 0
# 解释： nums 本身已是升序排序的数组。

# 输入： nums = [1]
# 输出： 0
# 解释： nums 本身已是升序排序的数组。


# 思路： 一次迭代
#
#      最简单的方法就是先排序，然后对比原数组和排序后的数组，
#      找到第一处和最后一处不同的下标，就是最短的子数组起止下标。
#
#      时间复杂度为 O(nlogn) ，空间复杂度为 O(n) 。
#
#
#      根据这个方法可以进一步挖掘出四次迭代的解法：
#          1. 第一次迭代：从左往右找到第一个下降的数的下标 l ，
#              然后找到 nums[l:] 中的最小值 min
#          2. 第二次迭代：找到 nums[l:] 中第一个大于 min 的下标 start ，
#              那么 start 就是最短的子数组起始下标。
#          3. 第三次迭代：从右往左找到第一个上升的数的下标 r ，
#              然后找到 nums[:r+1] 中的最大值 max
#          4. 第四次迭代：找到 nums[:r+1] 中最后一个小于 max 的下标 end ，
#              那么 end 就是最短的子数组结束下标。
#
#      时间复杂度为 O(n) ，空间复杂度为 O(1) 。
#
#
#      这样时空复杂度已经满足题意，但看了讨论区的解法后，
#      发现还可以继续挖掘内在条件和限制，优化为一次迭代的解法。
#
#      可以发现前两次迭代必须分开，
#      是为了避免上升阶段的最小值小于 min ，
#      因为这会使得第二次迭代找到错误的 start 。
#
#      那么反转一下思路，我们可以从右往左遍历，
#      并维护 nums[r+1:] 的最小值 min ，
#      同时维护大于 min 的最小下标 start ，
#      这样就能避免错误最小值导致找到错误 start 的问题：
#          1. nums[r] > min: 则说明 min 排序后最靠后也要在 r 处，
#              即排序的子数组的起始下标 start 最大为 r
#          2. nums[r] <= min: 更新 min 为 nums[r]
#
#      针对后两次迭代也可以如此优化。
#
#      同时可以注意到优化后的两次迭代相互独立，所以可以继续合并成一次迭代。  
#
#
#      时间复杂度：O(n)
#          1. 需要遍历 nums 中全部 O(n) 个数字
#      空间复杂度：O(1)
#          1. 只需要维护常数个额外变量     


class Solution:
    def findUnsortedSubarray(self, nums: List[int]) -> int:
        # 定义两个指针，分别从 nums 的两端开始遍历
        l, r = 0, len(nums) - 1
        # max 表示 nums[:l] 中的最大值，
        # min 表示 nums[r+1:] 中的最小值。
        max, min = -0x3f3f3f3f, 0x3f3f3f3f
        # start 表示需要排序的子数组的起始下标，
        # end 表示需要排序的子数组的结束下标。
        start, end = -1, -1
        # 当 nums 数组还未全部遍历完时，继续处理
        while r >= 0:
            if nums[l] < max:
                # 如果 nums[l] < max ，
                # 则说明 max 排序后最靠前也要在 l 处，
                # 即排序的子数组的结束下标最小为 l
                end = l
            else:
                # 如果 nums[l] >= max ，
                # 则更新 nums[:l] 的最大值
                max = nums[l]
            
            if nums[r] > min:
                # 如果 nums[r] > min ，
                # 则说明 min 排序后最靠后也要在 r 处，
                # 即排序的子数组的起始下标最大为 r
                start = r
            else:
                # 如果 nums[r] <= min ，
                # 则更新 nums[r+1:] 的最小值
                min = nums[r]

            # l 向右移动一位
            l += 1
            # r 向左移动一位
            r -= 1

        if start == -1:
            # 如果 start 还是 -1 ，则 nums 是升序排序的数组，
            # 无需排序任何子数组，直接返回 0 即可
            return 0
        else:
            # 如果 start 不是 -1 ，
            # 则 end - start + 1 就是需要排序的子数组的最小长度
            return end - start + 1
        # 可以注意到 end 为 -1 时， start 也不会改变，
        # 所以可以初始化 start 为 0 ，
        # 这就最后就不需要判断，直接返回 end - start + 1 即可
