# 链接：https://leetcode.com/problems/merge-triplets-to-form-target-triplet/
# 题意：给定 n 个三元组 triplets ，每次可以选两个三元组 triplets[i] 和 triplets[j] ，
#       然后将 triplets[j][k] 更新为 max(triplets[i][k], triplets[j][k]) ，
#       求在一些操作后，是否能将其中一个三元组变为 target ？

# 数据限制：
#   1 <= triplets.length <= 10 ^ 5
#   triplets[i].length == target.length == 3
#   1 <= ai, bi, ci, x, y, z <= 1000

# 输入： triplets = [[2,5,3],[1,8,4],[1,7,5]], target = [2,7,5]
# 输出： true
# 解释： 选择 [2,5,3] 和 [1,7,5] ，将 [1,7,5] 变为 [2,7,5]

# 输入： triplets = [[1,3,4],[2,5,8]], target = [2,5,8]
# 输出： true
# 解释： 已存在一个三元组是 target

# 输入： triplets = [[2,5,3],[2,3,4],[1,2,5],[5,2,3]], target = [5,5,5]
# 输出： true
# 解释： 选择 [2,5,3] 和 [1,2,5] ，将 [1,2,5] 变为 [2,5,5]
#       选择 [5,2,3] 和 [2,5,5] ，将 [2,5,5] 变为 [5,5,5]

# 输入： triplets = [[3,4,5],[4,5,6]], target = [3,2,5]
# 输出： false
# 解释： 三元组的第二个数没有比 2 小的

# 思路： 贪心
#
#       如果一个三元组存在一个数比 target 的数大，就需要排除，
#       因为如果纳入计算，则必定会超过 target 对应位置的数。
#
#       计算排除后所有三元组每个位置的最大数，
#       如果最后和 target 的对应位置的数相同，则返回 True
#
#       时间复杂度： O(n)
#       空间复杂度： O(1)


class Solution:
    def mergeTriplets(self, triplets: List[List[int]], target: List[int]) -> bool:
        x, y, z = target
        ma, mb, mc = 0, 0, 0
        for a, b, c in triplets:
            # 如果三个数都小于 target 对应位置的数，则更新对应位置的最大数
            if a <= x and b <= y and c <= z:
                ma = max(ma, a)
                mb = max(mb, b)
                mc = max(mc, c)
        # 如果对应位置的最大数和 target 的想等，则返回 True
        return ma == x and mb == y and mc == z
