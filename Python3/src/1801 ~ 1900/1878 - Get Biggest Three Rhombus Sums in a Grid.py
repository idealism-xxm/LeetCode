# 链接：https://leetcode.com/problems/get-biggest-three-rhombus-sums-in-a-grid/
# 题意：给定一个二维整型数组，菱形和是正方形旋转 45 度后边上所有数字的和，求最大的三个不同的菱形和（不足三个，则全部返回）？

# 输入： nums = [3,5,2,3]
# 输出： 7
# 解释： 可以找出 2 个数对和 (3, 3) 和 (5, 2) ，
#       max(3 + 3, 5 + 2) = max(6, 7) = 7

# 输入： nums = [3,5,4,2,4,6]
# 输出： 8
# 解释： 可以找出 3 个数对和 (3, 5), (4, 4) 和 (6, 2) ，
#       max(3 + 5, 4 + 4, 6 + 2) = max(8, 8, 8) = 8

# 思路： 贪心
#       最大的数字必须和最小的数字配对，否则答案不会更优
#
#       设当前剩余为配对的数字的最小值为 mn, 最大值为 mx
#       则任意数对和 (nums[i], nums[j]) 有：
#           nums[i] + nums[j] >= mn + nums[j] >= mn + mx
#       其他任意数对，都不会更优，所以 (最小值, 最大值) 是一个最优的数对和
#
#       时间复杂度： O(n)
#       空间复杂度： O(1)

class Solution:
    def minPairSum(self, nums: List[int]) -> int:
        nums = sorted(nums)
        length = len(nums)
        mn = 0
        for i in range(0, length // 2):
            mn = max(mn, nums[i] + nums[length - 1 - i])
        return mn
