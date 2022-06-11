# 链接：https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/
# 题意：给定一个数组 nums 和一个整数 x ，每次可以从 nums 最左侧或最右侧移除一个数字 num ，
#      然后将其从 x 中减去，求最少需要移除多少个数字，使得 x 变为 0 ，
#      如果无法使 x 变为 0 ，返回 -1 。


# 数据限制：
#  1 <= nums.length <= 10 ^ 5
#  1 <= nums[i] <= 10 ^ 4
#  1 <= x <= 10 ^ 9


# 输入： nums = [1,1,4,2,3], x = 5
# 输出： 2
# 解释： 移除后 2 个数字可以让 5 变为 0

# 输入： nums = [5,6,7,8,9], x = 4
# 输出： -1

# 输入： nums = [3,2,20,1,1,3], x = 10
# 输出： 5
# 解释： 移除前 2 个数字和后 3 个数字可以让 10 变为 0


# 思路： 滑动窗口
#
#      如果一道题目需要在所有满足某种状态的连续子串/连续子数组中，
#      找到满足题意的一个，那么可以考虑滑动窗口。
#
#      如果我们将 nums 拼接一次，得到一个长度为 2 * len(nums) 的新数组，
#      那么本题就转化为：在新数组中找到一个和为 x 的最短连续子数组的长度。
#
#      可以使用滑动窗口的方法解决本题，这样的空间复杂度为 O(n) 。
#      不过可以特殊处理右边界 r 的情况，能让空间复杂度优化为 O(1) ，
#      但不便于理解且容易出错。
#
#      此时可以考虑原问题的镜像问题：求和为 sum(nums) - x 的最长连续子数组的长度 ans ，
#      那么 len(nums) - ans 就是原问题的答案。
#
#      设 target = sum(nums) - x ，
#      我们使用滑动窗口 [l, r] 表示一个数字和小于等于 target 的连续子数组，
#      初始化为左边界 l = 0 ，右边界 r = -1 ，表示初始窗口为空，数字和 total = 0 。
#
#      我们不断右移右边界 r ，将其纳入到滑动窗口中考虑， total += nums[r] ，
#      然后不断左移左边界 l ，直至 total 小于等于 target 。
#
#      如果此时 total 等于 target ，则当前子数组是一个可选答案，更新 ans 的最大值。
#
#      最后如果 ans 为 0 ，则说明不存在和为 target 的连续子数组，返回 -1 ；
#      否则返回 len(nums) - ans 。
#
#
#      时间复杂度：O(n)
#          1. 需要遍历 nums 中全部 O(n) 个数字
#      空间复杂度：O(1)
#          1. 只需用维护常数个额外变量


class Solution:
    def minOperations(self, nums: List[int], x: int) -> int:
        n: int = len(nums)
        # 原问题可以转化为求和为 target 的最长连续子数组的长度
        target: int = sum(nums) - x
        # 如果 target 等于 0 ，则空数组是最长连续子数组，原问题的解为 n - 0 = n
        if target == 0:
            return n
        # 如果 target 小于 0 ，则不存在这样的连续子数组，因为所有数字都是正数
        if target < 0:
            return -1

        # 初始化滑动窗口为 [0, -1] ，窗口内数字和为 0 ，最长连续子数组长度为 0
        l, total, ans = 0, 0, 0
        # 不断右移右边界 r
        for r in range(n):
            # 将 nums[r] 纳入到滑动窗口中考虑
            total += nums[r]
            # 不断左移左边界 l ，直至 total 小于等于 target
            while total > target:
                total -= nums[l]
                l += 1

            # 如果此时滑动窗口内的数字和为 target ，则更新最长连续子数组长度的最大值
            if total == target:
                ans = max(ans, r - l + 1)

        if ans == 0:
            # 如果 ans 为 0 ，则不存在这样的连续子数组
            return -1
        # 如果存在这样的连续子数组，则 n - ans 就是原问题的解
        return n  - ans
