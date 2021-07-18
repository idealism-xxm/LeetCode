# 链接：https://leetcode.com/problems/build-array-from-permutation/
# 题意：给定一个 [0, n) 的排列数组 nums ，返回 ans ，
#       其中 ans[i] = nums[nums[i]]。

# 数据限制：
#   1 <= nums.length <= 1000
#   0 <= nums[i] < nums.length
#   所有数字都不相同

# 输入： nums = [0,2,1,5,3,4]
# 输出： [0,1,2,4,5,3]

# 输入： nums = [5,0,1,2,3,4]
# 输出： [4,5,0,1,2,3]

# 思路： 模拟
#
#       按照题意模拟一遍即可
#
#       时间复杂度： O(n)
#       空间复杂度： O(n)


class Solution:
    def buildArray(self, nums: List[int]) -> List[int]:
        return [
            nums[num]
            for num in nums
        ]
