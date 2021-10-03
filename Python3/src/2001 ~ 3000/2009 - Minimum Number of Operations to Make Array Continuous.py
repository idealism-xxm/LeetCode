# 链接：https://leetcode.com/problems/minimum-number-of-operations-to-make-array-continuous/
# 题意：给定一个数组 nums ，每次可以将任意一个位置的数字变成任意数，
#       求最少需要多少操作，使得数组 nums 中的数字是连续的？
#       连续满足以下两个条件：
#           1. nums 中的数字都是唯一的
#           2. max(nums) - min(nums) == nums.length - 1

# 数据限制：
#   1 <= nums.length <= 10 ^ 5
#   1 <= nums[i] <= 10 ^ 9

# 输入： nums = [4,2,5,3]
# 输出： 0
# 解释： nums 已是连续的

# 输入： nums = [1,2,3,5,6]
# 输出： 1
# 解释： [1,2,3,5,(6)] -> [1,2,3,5,4]

# 输入： nums = [1,10,100,1000]
# 输出： 3
# 解释： [1,(10),100,1000] -> [1,2,(100),1000] -> [1,2,3,(1000)] -> [1,2,3,4]


# 思路： 滑动窗口
#
#       我们先对 nums 进行去重排序，操作最少的情况下，
#        nums 中的一个数字 num 必定可以是最终数组的左边界值，
#       （如果 num 是右边界，我们可以找到另一个数作为左边界，使得 num 是范围内的最大数，
#           这样的答案不会更差）
#       那么我们就枚举每个数字作为左边界 left ，然后不断右移右边界 right ，
#       直至 nums[right] > nums[left] ，这时 unique_nums[left:right] 内的数字不用操作，
#       剩余的 n - (right - left) 个数字需要操作，
#
#       
#       时间复杂度： O(nlogn)
#       空间复杂度： O(n)


class Solution:
    def minOperations(self, nums: List[int]) -> int:
        n = len(nums)
        # 去重后按升序排序
        nums = sorted(set(nums))
        unique_n = len(nums)
        # 最多只会操作 n - 1 次
        ans = n - 1
        # 滑动窗口的右边界的右侧第一个位置
        right = 0
        # 枚举滑动窗口的左边界
        for left in range(unique_n):
            # 如果 right 合法，且这个位置的数字还在滑动窗口内，则右移 right ，
            # 直至 right 不合法，或者这个位置的数字不在滑动窗口内
            while right < unique_n and nums[right] <= nums[left] + n - 1:
                right += 1
            # 以 [left, right) 为左右边界的的数字，有 right - left 个不同的数字，
            # 需要对剩余的数字操作 n - right + left 次才能满足题意
            ans = min(ans, n - right + left)
        return ans