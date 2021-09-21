# 链接：https://leetcode.com/problems/sum-of-beauty-in-the-array/
# 题意：给定一个整数数组 nums ，对于 1 <= i <= nums.length - 1 的 nums[i] 来说，
#       其美丽值为：
#           2: 如果 nums[j] < nums[i] < nums[k] 
#               对于所有 0 <= j < i 和 i < k <= nums.length - 1 成立
#           1: 如果 nums[i - 1] < nums[i] < nums[i + 1] ，且上面条件不成立
#           0: 前两个条件都不成立

# 数据限制：
#   3 <= nums.length <= 10 ^ 5
#   1 <= nums[i] <= 10 ^ 5

# 输入： nums = [1,2,3]
# 输出： 2
# 解释：
#   nums[1]: 美丽值为 2

# 输入： nums = [2,4,6,4]
# 输出： 1
# 解释：
#   nums[1]: 美丽值为 1
#   nums[2]: 美丽值为 0

# 输入： nums = [3,2,1]
# 输出： 0
# 解释：
#   nums[1]: 美丽值为 0


# 思路： 前缀和
#
#       使用 rmin 先初始化后缀最小值， rmin[i] 表示 nums[i:] 的最小值
#       lmax 可以在遍历的时候计算
#
#       这样我们就可以在 O(1) 内判断是否满足第一个条件
#
#       时间复杂度： O(n)
#       空间复杂度： O(n)


class Solution:
    def sumOfBeauties(self, nums: List[int]) -> int:
        # 先初始化 rmin ， rmin[i] 表示 nums[i:] 中的最小值
        n = len(nums)
        rmin = [nums[n - 1]] * n
        for i in range(n - 2, 0, -1):
            rmin[i] = min(rmin[i + 1], nums[i])
        # 在从左计算答案
        ans = 0
        # 初始时 nums[:1] 中的最大值为 nums[0]
        lmax = nums[0]
        for i in range(1, n - 1):
            # 如果当前数大于前面所有的数，且小于后面所有的数，则美丽值为 2
            if nums[i] > lmax and nums[i] < rmin[i + 1]:
                ans += 2
            elif nums[i - 1] < nums[i] < nums[i + 1]:
                # 如果当前数仅大于其相邻两个数，则美丽值为 1
                ans += 1
            
            # 如果当前值更大，则更新 lmax
            if nums[i] > lmax:
                lmax = nums[i]

        return ans
