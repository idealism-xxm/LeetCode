# 链接：https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/
# 题意：给定一个整型数组，从中选取 k 个数，求最大值减去最小值的最小是多少？

# 数据限制：
#   1 <= k <= nums.length <= 1000
#   0 <= nums[i] <= 10 ^ 5

# 输入： nums = [90], k = 1
# 输出： 0
# 解释： 
#       只能选 [90] ，最大值为 90, 最小值为 90 ，差为 0

# 输入： nums = [9,4,1,7], k = 2
# 输出： 2
# 解释： 
#       选择 [9,7] ，最大值为 9, 最小值为 7, 差为 2


# 思路： 贪心
#
#       当选择的数是按顺序连续的 k 个数时，才有可能取得最小差值，
#       假设选择了一个非连续的数，那么差值不会变小，结果都不会更优。
#
#       所以排序完成后，没有所有长度为 k 的子数组计算差值的最小值即可。
#
#       时间复杂度： O(nlogn)
#       空间复杂度： O(1)

class Solution:
    def minimumDifference(self, nums: List[int], k: int) -> int:
        # 先排序
        nums.sort()
        # 从所有长度为 k 的子数组中找到最小差值
        return min(nums[i + k - 1] - nums[i] for i in range(len(nums) - k + 1))
