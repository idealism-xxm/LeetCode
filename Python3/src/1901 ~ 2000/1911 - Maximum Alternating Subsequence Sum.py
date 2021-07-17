# 链接：https://leetcode.com/problems/maximum-alternating-subsequence-sum/
# 题意：给定一个整型数组 nums ，求所有子序列形成的数组的交错和最大值？
#      数组交错和 = 偶数下标的数的和 - 奇数下标的数的和

# 数据限制：
#   1 <= nums.length <= 10 ^ 5
#   1 <= nums[i] <= 10 ^ 5

# 输入： s = "daabcbaabcbc", part = "abc"
# 输出： nums = [4,2,5,3]
# 解释：
#       [4,2,5] -> (4 + 5) - 2 = 7

# 输入： nums = [5,6,7,8]
# 输出： 8
# 解释：
#       [8] -> 8 = 8

# 输入： nums = [6,2,1,2,4,5]
# 输出： 10
# 解释：
#       [6,1,5] -> (6 + 5) - 1 = 10

# 思路1： DP
#
#       状态定义（以下范围为前闭后开）：
#           dp[i][0] 表示 nums[:i] 中选取奇数个数能形成的交错和的最大值
#           dp[i][1] 表示 nums[:i] 中选取偶数个数能形成的交错和的最大值
#       初始状态： dp[0][0] = dp[0][1] = 0
#       状态转移方程：
#           # 不选当前数： dp[i - 1][0]
#           # 选当前数：   dp[i - 1][1] + nums[i - 1]
#           dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + nums[i - 1])
#           # 不选当前数： dp[i - 1][1]
#           # 选当前数：   dp[i - 1][0] - nums[i - 1]
#           dp[i][1] = max(dp[i - 1][1], dp[i - 1][0] - nums[i - 1])
#
#       结果： max(dp[len(nums)][0], dp[len(nums)][1])
#       其实最后长度一定是奇数，因为偶数长度的必定比奇数长度的多减去一个数
#
#       时间复杂度： O(n)
#       空间复杂度： O(1) 【使用滚动数组优化了】

class Solution:
    def maxAlternatingSum(self, nums: List[int]) -> int:
        dp = [[0, 0], [0, 0]]
        pre, cur = 0, 1
        for num in nums:
            # 不选当前数： dp[i - 1][0]
            # 选当前数：   dp[i - 1][1] + nums[i - 1]
            dp[cur][0] = max(dp[pre][0], dp[pre][1] + num)
            # 不选当前数： dp[i - 1][1]
            # 选当前数：   dp[i - 1][0] - nums[i - 1]
            dp[cur][1] = max(dp[pre][1], dp[pre][0] - num)
            # 准备切换滚动数组下标
            pre, cur = cur, pre

        # 获取最后一次更新后的最大值
        return max(dp[pre])


# 思路2： 贪心
#
#       和 122. Best Time to Buy and Sell Stock II 题意类似
#       https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
#
#       只不过本题相当于无成本先给了你股票
#       所以我们每次只要当前数比前一个数大，就可以选上这两个数（即相当于一次股票买卖）
#
#       第一个数直接入选，相当于直接卖掉获利，
#       后面只要发现当前数比前一个数大，
#       则表明前一个数可以买入，当前数可以卖出，利润是增加的
#       （如果前一个数已经用过一次了，则表明前一个数是假装卖出买入的，对结果无影响）
#
#       时间复杂度： O(n)
#       空间复杂度： O(1)

class Solution:
    def maxAlternatingSum(self, nums: List[int]) -> int:
        ans = nums[0]
        for i in range(1, len(nums)):
            # 只要当前数比前一个数大，就可以选择
            if nums[i] > nums[i - 1]:
                ans += nums[i] - nums[i - 1]

        return ans
