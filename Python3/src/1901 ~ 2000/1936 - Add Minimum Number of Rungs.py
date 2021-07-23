# 链接：https://leetcode.com/problems/add-minimum-number-of-rungs/
# 题意：给定一个严格单调递增的数组 rungs ，表示一个梯子不同层的高度，
#       现在一个人在地面上，准备爬梯子，但每次最多只能登 dist 的高度，
#       求最少加多少层后，这个人可以爬到梯子顶端？

# 数据限制：
#   1 <= rungs.length <= 10 ^ 5
#   1 <= rungs[i] <= 10 ^ 9
#   1 <= dist <= 10 ^ 9
#   rungs 严格单调递增

# 输入： [1,3,5,10], dist = 2
# 输出： 2
# 解释： 添加高度为 7 和 8 的两层后变为 [1,3,5,7,8,10]

# 输入： rungs = [3,6,8,10], dist = 3
# 输出： 0

# 输入： [3,4,6,7], dist = 2
# 输出： 1
# 解释： 添加高度为 1 的一层后变为 [1,3,4,6,7]

# 输入： rungs = [5], dist = 10
# 输出： 0


# 思路： 贪心
#
#       如果两层之间的高度差小于等于 dist ，则不需要添加新层，
#       如果两层之间的高度差每多大于 dist 一次，则需要添加一个新层，
#
#       即 ans = sum(rungs[i] - rungs[i - 1] - 1) // dist
#
#
#       时间复杂度： O(n)
#       空间复杂度： O(1)


class Solution:
    def addRungs(self, rungs: List[int], dist: int) -> int:
        pre = 0
        ans = 0
        for rung in rungs:
            ans += (rung - pre - 1) // dist
            pre = rung

        return ans
