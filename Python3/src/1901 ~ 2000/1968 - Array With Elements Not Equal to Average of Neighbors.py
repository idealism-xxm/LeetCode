# 链接：https://leetcode.com/problems/array-with-elements-not-equal-to-average-of-neighbors/
# 题意：给定一个数字各不相同的整数数组 nums ，重新调整每个数字的位置，
#       使得任意一个数 nums[i] != (nums[i - 1] + nums[i + 1]) / 2
#       (1 <= i < n - 1)

# 数据限制：
#   3 <= nums.length <= 10 ^ 5
#   0 <= nums[i] <= 10 ^ 5

# 输入： nums = [1,2,3,4,5]
# 输出： [1,2,4,5,3]
# 解释： 
#       i = 1: nums[1] = 2, (nums[0] + nums[2]) / 2 = (1 + 4) / 2 = 2.5
#       i = 2: nums[2] = 4, (nums[1] + nums[3]) / 2 = (2 + 5) / 2 = 3.5
#       i = 3: nums[3] = 5, (nums[2] + nums[4]) / 2 = (4 + 3) / 2 = 3.5

# 输入： nums = [6,2,0,9,7]
# 输出： [9,7,6,2,0]
# 解释： 
#       i = 1: nums[1] = 7, (nums[0] + nums[2]) / 2 = (9 + 6) / 2 = 7.5
#       i = 2: nums[2] = 6, (nums[1] + nums[3]) / 2 = (7 + 2) / 2 = 4.5
#       i = 3: nums[3] = 2, (nums[2] + nums[4]) / 2 = (6 + 0) / 2 = 3


# 思路： 排序 + 贪心
#
#       从小到达排序后，偶数位置从数组左边取数放入，奇数位置从数组右边取数放入
#
#       这样偶数位置的数比左右两边的数都小，而奇数位置的书比左右两边的数都大，那么必定满足题意
#       
#       时间复杂度： O(nlogn)
#       空间复杂度： O(n)

class Solution:
    def rearrangeArray(self, nums: List[int]) -> List[int]:
        nums.sort()
        length = len(nums)
        l, r = 0, length - 1
        ans = [None] * length
        i = 0
        while i < length:
            # 偶数位置为放入当前最小的数
            ans[i] = nums[l]
            i += 1
            l += 1
            # 奇数位置放入当前最大的数
            if i < length:
                ans[i] = nums[r]
                i += 1
                r -= 1
        return ans

