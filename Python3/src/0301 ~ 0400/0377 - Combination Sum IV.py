# 链接：https://leetcode.com/problems/combination-sum-iv/
# 题意：给定一个不含重复数字的正整数数组 nums 和一个正整数 target ，
#      求有多少种不同的组合，使得数组中所有数字的和为 target ？
#      每个数字可选多次，两个组合不同当且仅当至少有一个位置的数不同。
#
#      进阶：如果 nums 中含有负数，需要添加什么限制条件来允许负数出现？


# 数据限制：
#  1 <= nums.length <= 200
#  1 <= nums[i] <= 1000
#  nums 中所有的数字各不相同
#  1 <= target <= 1000


# 输入： nums = [1,2,3], target = 4
# 输出： 7
# 解释： 可能的组合方式有：
#       (1, 1, 1, 1)
#       (1, 1, 2)
#       (1, 2, 1)
#       (1, 3)
#       (2, 1, 1)
#       (2, 2)
#       (3, 1)

# 输入： nums = [9], target = 3
# 输出： 0


# 思路： DP
#
#      设 dp[i] 表示和为 i 的不同组合数。
#
#      初始化： dp[i] = 0; dp[0] = 1;
#          即最开始所有数字都无合法的组合，而数字 0 对应 1 种合法的空组合。
#      状态转移： dp[i] = dp[i] + dp[i - nums[j]]
#          枚举组合的最后一个数字为 nums[j] ，
#          那么在 dp[i - nums[j]] 的所有组合中加入 nums[j] 后，
#          这些组合的和都为 i 。
#          即 dp[i] 可由 dp[i - nums[j]] 转移而来。
#
#      最后 dp[target] 就是和为 target 的不同组合数，直接返回即可。
#		    
#
#		时间复杂度： O(target * n)
#          1. 需要遍历全部 O(target) 个状态
#          2. 遍历每个状态时，都需要遍历全部 O(n) 个数字进行状态转移
#		空间复杂度： O(target)
#          1. 需要维护全部 O(taget) 个状态


class Solution:
    def combinationSum4(self, nums: List[int], target: int) -> int:
        # dp[i] 表示和为 i 的不同组合数。
        # 最开始所有数字都无合法的组合，而数字 0 对应 1 种合法的空组合。
        dp: List[int] = [0] * (target + 1)
        dp[0] = 1
        # 遍历每个状态 i ，则 dp[0..i] 都已确定，可以放心转移
        for i in range(1, target + 1):
            # 枚举组合的最后一个数字为 num
            for num in nums:
                # 如果 i >= num ，则 dp[i] 可由 dp[i - num] 转移而来
                if i >= num:
                    dp[i] += dp[i - num]

        return dp[target]
