# 链接：https://leetcode.com/problems/remove-one-element-to-make-the-array-strictly-increasing/
# 题意：给定一个整型数组，判断至多删除一个数字后，该数组是否严格单调递增？

# 数据限制：
#   2 <= nums.length <= 1000
#   1 <= nums[i] <= 1000

# 输入： nums = [1,2,10,5,7]
# 输出： true
# 解释： [1,2,10,5,7] -> [1,2,5,7]

# 输入： nums = [2,3,1,2]
# 输出： false

# 输入： nums = [1,1,1]
# 输出： false

# 输入： nums = [1,2,3]
# 输出： true
# 解释： [1,2,3] 已经是否严格单调递增

# 思路： 贪心
#
#       我们记录前一个数字 pre 和当前是否删除过数字 has_removed，
#       1. nums[i] > pre ，则令 pre = nums[i] ，然后继续处理下一个，
#       2. nums[i] <= pre ，则需要删除 pre 或者 nums[i]
#           (1) 如果已经删除过数字，则直接返回 False
#           (2) 令 has_removed = True ，表示已删除数字
#               然后贪心判断是否可以删除 pre ，
#               因为 pre >= nums[i] ，留下较小的数字更有利。
#               如果不能删除 pre ，则只能删除 nums[i]
#
#       时间复杂度： O(n)
#       空间复杂度： O(1)

class Solution:
    def canBeIncreasing(self, nums: List[int]) -> bool:
        # 标记是否已经移除过数字
        has_removed = False
        pre = nums[0]
        for i in range(1, len(nums)):
            # 如果严格单调递增，则继续处理下一个数字
            if pre < nums[i]:
                pre = nums[i]
                continue

            # 此时需要移除一个数字，如果已经移除过了，则直接返回 False
            if has_removed:
                return False

            # 标记已移除数字
            has_removed = True
            # 我们先贪心考虑移除 pre 这个数字，因为 pre >= nums[i] ，
            # 这样可以尽可能留下最小的数，留下大数结果不会更优
            if i > 2 and nums[i] > nums[i - 2]:
                pre = nums[i]

            # 如果没法留下 nums[i] ，那么只能移除它， pre 维持不变

        # 成功遍历完，则数组满足题意
        return True
