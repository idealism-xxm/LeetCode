# 链接：https://leetcode.com/problems/reduction-operations-to-make-the-array-elements-equal/
# 题意：给定一个整数数组，求使用多少次以下操作把所有的数变成相等的数？
#       操作：找到最大的数（多个选下标最小的），将其变为次大的数。

# 输入： nums = [5,1,3]
# 输出： 3
# 解释： [5,1,3] -> [3,1,3] -> [1,1,3] -> [1,1,1]

# 输入： nums = [1,1,1]
# 输出： 0

# 输入： nums = [1,1,2,2,3]
# 输出： 4
# 解释： [1,1,2,2,3] -> [1,1,2,2,2] -> [1,1,1,2,2] ->
#       [1,1,1,1,2] -> [1,1,1,1,1]

# 思路： 统计
#
#       排序后，从大到小遍历，并一直记录当前最大的数的个数为 cnt ，
#       当每次发现次大的数时，答案 ans += cnt ，
#       即需要有 cnt 次操作才能将当前所有最大数变为次大数
#
#       时间复杂度： O(nlogn)
#       空间复杂度： O(1)


class Solution:
    def reductionOperations(self, nums: List[int]) -> int:
        # 从大到小排序
        nums = sorted(nums, reverse=True)
        ans = 0
        pre = nums[0]
        cnt = 0
        for cur in nums[1:]:
            cnt += 1
            # 发现次大数时，需要将所有的最大数变为次大数
            if cur != pre:
                ans += cnt
                pre = cur
        return ans
