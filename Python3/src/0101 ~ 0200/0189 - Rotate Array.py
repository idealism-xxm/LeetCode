# 链接：https://leetcode.com/problems/rotate-array/
# 题意：给定一个整型数组，将它原地循环右移 k 次。
#      进阶：
#          1. 用至少 3 种不同的方法完成
#          2. 使用空间复杂度为 O(1) 的原地方法完成

# 数据限制：
#  1 <= nums.length <= 10 ^ 5
#  -(2 ^ 31) <= nums[i] <= 2 ^ 31 - 1
#  0 <= k <= 10 ^ 5

# 输入：nums = [1,2,3,4,5,6,7], k = 3
# 输出：[5,6,7,1,2,3,4]
# 解释：循环右移 1 次： [7,1,2,3,4,5,6]
#      循环右移 2 次： [6,7,1,2,3,4,5]
#      循环右移 3 次： [5,6,7,1,2,3,4]

# 输入：nums = [-1,-100,3,99], k = 2
# 输出：[3,99,-1,-100]
# 解释：循环右移 1 次： [99,-1,-100,3]
#      循环右移 2 次： [3,99,-1,-100]


# 思路3：三次翻转
#
#	   可以发现循环右移 k 次后，
#      数组末尾的 k % n 个数字会移动到数组开始，
#      而数组开始的 n - k % n 个数字，则会向右移动 k % n 次。
#
#      如果我们想将数组 nums 整体翻转，
#      则可以使得末尾的 k % n 个数移动至数组开始，
#      不过此时 nums[:k % n] 和 nums[k % n:] 的顺序都是反的，
#      所以还需要分别对 nums[:k % n] 和 nums[k % n:] 再次翻转，
#      这样就能获得循环右移 k 次的结果。
#
#
#	   时间复杂度： O(n)
#          1. 需要遍历数组 nums 中的全部 O(n) 个数字
#	   空间复杂度： O(1)
#          1. 只需要使用常数个额外变量


class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        n: int = len(nums)
        # 计算最后有多少数字会被移动到数组开始
        k %= n
        # 翻转整个数组
        Solution.reverse(nums, 0, n - 1)
        # 翻转前 k 个数字
        Solution.reverse(nums, 0, k - 1)
        # 翻转后 n - k 个数字
        Solution.reverse(nums, k, n - 1)

    @staticmethod
    def reverse(nums: List[int], l: int, r: int) -> None:
        # 双指针翻转 [l, r] 内的数字
        while l < r:
            # 交换 l 和 r 位置的数字
            nums[l], nums[r] = nums[r], nums[l]
            # l 向右移动一位
            l += 1
            # r 向左移动一位
            r -= 1
