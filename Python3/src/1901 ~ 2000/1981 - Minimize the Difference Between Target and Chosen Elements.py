# 链接：https://leetcode.com/problems/minimize-the-difference-between-target-and-chosen-elements/
# 题意：给定一个 m * n 的矩阵和一个数 target ，每一行选择一个数，
#       求所有这些数的和与 target 差的最小值？

# 数据限制：
#   m == mat.length
#   n == mat[i].length
#   1 <= m, n <= 70
#   1 <= mat[i][j] <= 70
#   1 <= target <= 800

# 输入： mat = [[1,2,3],[4,5,6],[7,8,9]], target = 13
# 输出： 0
# 解释： 
#       第一行选择 1
#       第二行选择 4
#       第三行选择 7
#       和为 13 ，与 target 相差 0 。

# 输入： mat = [[1],[2],[3]], target = 100
# 输出： 94
# 解释： 
#       第一行选择 1
#       第二行选择 2
#       第三行选择 3
#       和为 6 ，与 target 相差 94 。

# 输入： mat = [[1,2,9,8,7]], target = 6
# 输出： 1
# 解释： 
#       第一行选择 7
#       和为 7 ，与 target 相差 1 。


# 思路： DP
#
#       比赛的时候用 Golang 用 O(m * n * mx ^ 2) 水过去了，
#
#       如果用上 target <= 800 这个条件，
#       我们可以将必定不合法的情况排除在外，这样能将时间复杂度降低为 O(m * n * target)
#
#       我们从每行中选择最小的那个数，那么就可以得到最小的和 mn_sum ，
#       1. 如果 mn_sum >= target ，那么我们就可以直接返回 mn_sum - target ，
#       2. 如果 mn_sum < target ，那么只有形成的和
#            sum < target - (mn_sum - target) = 2 * target - mn_sum 时，
#           才能得到更小的差值，那么和的最大值就是 2 * target - mn_sum ，
#           我们进行 DP 即可，
#
#       时间复杂度： O(m * n * target)
#       空间复杂度： O(target)

class Solution:
    def minimizeTheDifference(self, mat: List[List[int]], target: int) -> int:
        # 如果最小的和都大于等于 target，那么直接返回 mn_sum - target
        mn_sum = sum(min(row) for row in mat)
        if mn_sum >= target:
            return mn_sum - target
        
        # 最大的和 < 2 * target - mn_sum 时，才会得到更小的值
        mx = 2 * target - mn_sum
        nums = {0}
        # 遍历所有数字
        for row in mat:
            nxt_nums = set()
            for cur in row:
            # 将当前行的数 cur 与 nums 中的数求和，如果小于 mx ，则可以放入结果中
                for num in nums:
                    sm = cur + num
                    if sm < mx:
                        nxt_nums.add(sm)
            # 滚动数组
            nums = nxt_nums

        return min(abs(num - target) for num in nums)
