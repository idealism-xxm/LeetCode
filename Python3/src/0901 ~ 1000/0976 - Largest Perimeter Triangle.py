# 链接：https://leetcode.com/problems/largest-perimeter-triangle/
# 题意：给定一个整型数组 nums ，选取其中 3 个数组成三角形，
#      求能组成面积不为 0 的三角形的最大周长？
#      如果不能组成面不为 0 的三角形，则返回 0 。


# 数据限制：
#  3 <= nums.length <= 10 ^ 4
#  1 <= nums[i] <= 10 ^ 6


# 输入： nums = [2,1,2]
# 输出： 5

# 输入： nums = [1,2,1]
# 输出： 0
# 解释： 无法组成三角形


# 思路： 贪心 + 排序
#
#      这道题需要运用数学的基础知识，三角形的判定定理：三角形两边之和大于第三边。
#
#      设三角形的边为 a, b, c ，不失一般性，令 a <= b <= c ，
#      那么根据三角形的判定定理有： a + b > c 。
#
#      假设已经选取了 nums 中的一个数作为 c ，
#      那么贪心地选取小于等于 c 的最大的两个数分别作为 a 和 b ，
#      判断是否能组成三角形即可。
#
#      如果选取的数选取的数比 a 和 b 小，那么两边之和也会变小，
#      大于 c 的可能性更低，周长也会更小。
#
#      所以可以先对 nums 排序，从大到小枚举每个数 nums[i] 作为 c ，
#      则有 a = nums[i - 2], b = nums[i - 1] ，
#      找到第一个能使 a + b > c 成立的，此时组成的三角形周长最大。
#
#      
#      时间复杂度：O(nlogn)
#          1. 需要对 nums 中全部 O(n) 个数字进行排序，时间复杂度为 O(nlogn)
#          2. 需要遍历 nums 中全部 O(n) 个数字一次
#      空间复杂度：O(1)
#          1. 只需要使用常数个额外变量即可


class Solution:
    def largestPerimeter(self, nums: List[int]) -> int:
        # 按数字升序排序
        nums.sort()
        # 贪心从最大数字倒序遍历
        for i in range(len(nums) - 1, 1, -1):
            # 如果当前当前最大的三个数满足三角形判定定理，则其组成的三角形周长最长
            if nums[i - 2] + nums[i - 1] > nums[i]:
                return nums[i - 2] + nums[i - 1] + nums[i]

        # 此时无法组成任何三角形，返回 0
        return 0
