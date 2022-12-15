# 链接：https://leetcode.com/problems/house-robber/
# 题意：给定一个数组，不能选择相邻的两个数，求选择某些数的和的最大值？


# 数据限制：
#  1 <= nums.length <= 100
#  0 <= nums[i] <= 400


# 输入： [1,2,3,1]
# 输出： 4
# 解释： 选择第 1 个和第 3 个， 1 + 3 = 4

# 输入： [2,7,9,3,1]
# 输出： 12
# 解释： 选择第 1 个、第 3 个和第 5 个， 2 + 9 + 1 = 12


# 思路： DP
#
#      设 dp[i] 表示在 nums[..=i] 中选择的数的最大和。
#
#      初始化： 
#          1. dp[0] = nums[0]: nums[..=0] 中只能选择 nums[0]
#          2. dp[1] = max(nums[0], nums[1]): 
#              nums[..=1] 中不能两个都选，只能二选一，贪心选择最大的
#            
#      状态转移： dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])
#          1. 不选 nums[i] ，则 dp[i] 由 dp[i - 1] 转移而来
#          2. 选择 nums[i] ，则 dp[i] 由 dp[i - 2] + nums[i] 转移而来
#
#
#      DP 常见的三种优化方式见 LeetCode 583 这题的思路，
#      本题只能采用滚动数组的方式进行优化，将空间复杂度从 O(n) 优化为 O(1) 。
#      本实现为了便于理解，不做优化处理。
#
#
#      时间复杂度： O(n)
#          1. 需要遍历 dp 中全部 O(n) 个状态
#      空间复杂度： O(n)
#          2. 需要维护 dp 中全部 O(n) 个状态


class Solution:
    def rob(self, nums: List[int]) -> int:
        n: int = len(nums)
        # 如果只有一个数，那么必定是选择这个数
        if n == 1:
            return nums[0]

        # dp[i] 表示从 nums[..=i] 中选择的数的最大和
        dp: List[int] = [0] * n
        # nums[..=0] 中只能选择 nums[0]
        dp[0] = nums[0]
        # nums[..=1] 中不能两个都选，只能二选一，贪心选择最大的
        dp[1] = max(nums[0], nums[1])
        for i in range(2, n):
            # 1. 不选 nums[i] ，则 dp[i] 由 dp[i - 1] 转移而来
            # 2. 选择 nums[i] ，则 dp[i] 由 dp[i - 2] + nums[i] 转移而来
            dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])

        return dp[n - 1]
