# 链接：https://leetcode.com/problems/delete-and-earn/
# 题意：给定一个整型数组 nums ，需要按照以下操作获得分数，
#      求能获得的最大分数？
#
#      操作：删除 nums[i] ，获得 nums[i] 的分数，
#           同时需要再删除所有等于 nums[i] - 1 
#           和 nums[i] + 1 的数字。


# 数据限制：
#   1 <= nums.length <= 2 * 10 ^ 4
#   1 <= nums[i] <= 10 ^ 4


# 输入： nums = [3,4,2]
# 输出： 6
# 解释： 可以按照以下操作执行：
#          1. 删除数字 4 ，获得 4 分，同时需要删除数字 3 ，
#              现在 nums = [2]
#          2. 删除数字 2 ，获得 2 分，现在 nums = []

# 输入： nums = [2,2,3,3,3,4]
# 输出： 9
# 解释： 可以按照以下操作执行：
#          1. 删除数字 3 ，获得 3 分，同时需要删除数字 2 和 4 ，
#             现在 nums = [3,3]
#          2. 删除数字 3 ，获得 3 分，现在 nums = [3]
#          3. 删除数字 3 ，获得 3 分，现在 nums = []


# 思路： DP + Map
#
#		先统计每个数字出现的次数到 num_to_cnt 中，
#       并找到 nums 中最大的数字 max_num 。
#
#       然后定义一个长度为 max_num + 1 的数组 dp ，
#       dp[i] 表示 [1, i] 内能获取到的最大分数。
#
#       初始化： dp[0] = 0, dp[1] = num_to_cnt[1] * 1
#       状态转移： dp[i] 可以由两种状态转移而来：
#           1. 不选数字 i ，无新增分数，从 dp[i - 1] 转移
#           2. 选择数字 i ，新增分数 num_to_cnt[i] * i ，
#              从 dp[i - 2] 转移
#
#           即 dp[i] = max(dp[i - 1], dp[i - 2] + num_to_cnt[i] * i)
#
#       最后 dp[max_num] 就是能获得的最大分数。
#
#
#       设 k 为 nums 中最大的数字。
#
#       时间复杂度：O(n + k)
#           1. 需要遍历 nums 中全部 O(n) 个数字一次，
#           2. 需要 2 ~ k 一次，进行状态转移
#       空间复杂度：O(n + k)
#           1. 需要维护一个 O(n) 的 Map num_to_cnt ，
#               最差情况下 nums 中的数字各不相同
#           2. 需要维护一个 O(k) 的数组 dp 
#               （由于每次只会用到 dp[i - 1] 和 dp[i - 2] ，
#               所以可以只使用两个变量维护，这部分优化为 O(1) ）


class Solution:
    def deleteAndEarn(self, nums: List[int]) -> int:
        # num_to_cnt 维护 nums 中每个数字的出现次数
        num_to_cnt = Counter(nums)
        # max_num 维护 nums 中最大的数字
        max_num = max(nums)

        # dp[i] 表示 [1, i] 内能获取到的最大分数
        dp =[0] * (max_num + 1)
        # 初始化 dp[0] = 0, dp[1] = num_to_cnt[1] * 1
        dp[1] = num_to_cnt.get(1, 0)
        # 从 2 开始枚举每个数字，直到 max_num
        for i in range(2, max_num + 1):
            # 数字 i 对分数的贡献为 num_to_cnt[i] * i
            sum = num_to_cnt.get(i, 0) * i
            # dp[i] 可以由两种状态转移而来：
            #   1. 不选数字 i ，无新增分数，从 dp[i - 1] 转移
            #   2. 选择数字 i ，新增分数 sum ，从 dp[i - 2] 转移
            dp[i] = max(dp[i - 1], dp[i - 2] + sum)

        # 最后 dp[max_num] 就是能获得最大分数
        return dp[max_num]
