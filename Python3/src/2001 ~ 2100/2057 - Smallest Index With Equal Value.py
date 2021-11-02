# 链接：https://leetcode.com/problems/smallest-index-with-equal-value/
# 题意：给定一个整数数组 nums ，求第一个满足 i mod 10 == nums[i] 下标 i ？
#       如果不存在这样的 i ，返回 -1 。

# 数据限制：
#   1 <= nums.length <= 100
#   0 <= nums[i] <= 9

# 输入： nums = [0,1,2]
# 输出： 0
# 解释： 
#   i=0: 0 mod 10 = 0 == nums[0]
#   i=1: 1 mod 10 = 1 == nums[1]
#   i=2: 2 mod 10 = 2 == nums[2]

# 输入： nums = [4,3,2,1]
# 输出： 2
# 解释： 
#   i=0: 0 mod 10 = 0 != nums[0]
#   i=1: 1 mod 10 = 1 != nums[1]
#   i=2: 2 mod 10 = 2 == nums[2]
#   i=3: 3 mod 10 = 3 != nums[3]

# 输入： nums = [2,1,3,5,2]
# 输出： 1
# 解释： 
#   i=0: 0 mod 10 = 0 != nums[0]
#   i=1: 1 mod 10 = 1 == nums[1]
#   i=2: 2 mod 10 = 2 != nums[2]
#   i=3: 3 mod 10 = 3 != nums[3]
#   i=4: 4 mod 10 = 4 != nums[4]


# 思路： 枚举
#
#       从前往后枚举下标 i ，判断下标 i 是否满足 i mod 10 == nums[i] ，
#       如果满足则直接返回 i ，否则继续处理下一个。
#       如果遍历完还没有返回，则直接返回 -1
#
#       时间复杂度： O(n)
#       空间复杂度： O(1)


class Solution:
    def smallestEqual(self, nums: List[int]) -> int:
        # 遍历 nums 数组
        for i, num in enumerate(nums):
            # 如果 i mod 10 == nums[i] ，则直接返回 i
            if i % 10 == num:
                return i
        # 如果遍历完还没有返回，则直接返回 -1
        return -1
