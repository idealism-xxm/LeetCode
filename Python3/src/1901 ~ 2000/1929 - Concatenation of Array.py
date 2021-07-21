# 链接：https://leetcode.com/problems/concatenation-of-array/
# 题意：给定一个整型数组 nums ，返回 nums 与自己拼接一次后的数组。

# 数据限制：
#   n == nums.length
#   1 <= n <= 1000
#   1 <= nums[i] <= 1000

# 输入： nums = [1,2,1]
# 输出： [1,2,1,1,2,1]

# 输入： nums = [1,3,2,1]
# 输出： [1,3,2,1,1,3,2,1]

# 思路： 模拟
#
#       Python 中直接拼接即可
#       其他语言要分配一个两倍大小的数组，然后放入对应位置即可
#
#       时间复杂度： O(n)
#       空间复杂度： O(n)


class Solution:
    def getConcatenation(self, nums: List[int]) -> List[int]:
        return nums + nums
