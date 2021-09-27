# 链接：https://leetcode.com/problems/maximum-difference-between-increasing-elements/
# 题意：给定一个长度为 n 的整数数组 nums ，找到满足以下条件时 nums[i] 与 nums[j] 的最大差值，
#       不存在则返回 -1 。
#       条件： 0 <= i < j < n && nums[i] < nums[j]

# 数据限制：
#   n == nums.length
#   2 <= n <= 1000
#   1 <= nums[i] <= 10 ^ 9

# 输入： nums = [7,1,5,4]
# 输出： 4
# 解释：
#   当 i = 1, j = 2 时， nums[1] = 1 和 nums[2] = 5 之间的差值为 4 。

# 输入： nums = [9,4,3,2]
# 输出： -1
# 解释：
#   不存在 i, j 满足条件 0 <= i < j < n && nums[i] < nums[j] 。

# 输入： nums = [1,5,2,10]
# 输出： 9
# 解释：
#   当 i = 0, j = 3 时， nums[0] = 1 和 nums[3] = 10 之间的差值为 9 。


# 思路： 枚举
#
#       我们维护 mn ，表示 nums[:i] 内最小值，初始化为 1e9 + 1 ，
#       再维护 ans 表示最终结果，初始化为 -1 ，表示没有满足题意的可选解，
#       然后遍历数组，每次先更新 ans ，再更新 mn 即可。
#       1. 此时有 num 的下标 大于等于 mn 的下标，
#           如果 num > mn ，则 num - mn 是一个可选解，可以用于更新 ans
#       2. 然后更新 mn 为 nums[:i] 内最小值
#
#
#       时间复杂度： O(n)
#       空间复杂度： O(1)


class Solution:
    def maximumDifference(self, nums: List[int]) -> int:
        # 先初始化 ans 为不存在
        ans = -1
        # mn 维护 nums[:i] 内最小值
        mn = 1000000001
        for num in nums:
            if num > mn:
                # 此时必定有 num 的下标 大于 mn 的下标，且有 num > mn ，
                # 则 num - mn 是满足题意的可选解
                ans = max(ans, num - mn)
            # 最后更新 mn 为 nums[:i] 内的最小值即可
            mn = min(mn, num)
        return ans
