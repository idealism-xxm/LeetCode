# 链接：https://leetcode.com/problems/find-the-middle-index-in-array/
# 题意：给定一个整型数组 nums ，返回一个最左侧的下标 index ，
#       使得 sum(nums[:index]) == sum(nums[index + 1:]) ，
#       不存在则返回 -1 。

# 数据限制：
#   1 <= nums.length <= 100
#   -1000 <= nums[i] <= 1000

# 输入： nums = [2,3,-1,8,4]
# 输出： 3
# 解释： 
#       sum(nums[:3]) = nums[0] + nums[1] + nums[2] = 2 + 3 + -1 = 4
#       sum(nums[4:]) = nums[4] = 4

# 输入： nums = [1,-1,4]
# 输出： 2
# 解释： 
#       sum(nums[:2]) = nums[0] + nums[1] = 1 + -1 = 0
#       sum(nums[3:]) = 0

# 输入： nums = [2,5]
# 输出： -1

# 输入： nums = [0]
# 输出： 0
#       sum(nums[:0]) = 0
#       sum(nums[1:]) = 0

# 思路： 枚举
#
#       先求出所有数字的和 sm ，
#       然后从开始枚举下标 index ，并记录不含当前下标的前缀和 prefix_sum ，
#       如果 sm - nums[index] == prefix_sum * 2 ，那么 index 就是我们要求的下标。
#
#       时间复杂度： O(nlogn)
#       空间复杂度： O(1)

class Solution:
    def findMiddleIndex(self, nums: List[int]) -> int:
        # 求数组和
        sm = sum(nums)
        # 不含当前数字的前缀和
        prefix_sum = 0
        for i, num in enumerate(nums):
            # 如果 sm - num 是前缀和的两倍，那么 i 就是我们要求的下标
            if prefix_sum * 2 == sm - num:
                return i
            # 当前数字加入到求前缀和中
            prefix_sum += num
        # 遍历完所有的数字都没有找到，则不存在
        return -1
