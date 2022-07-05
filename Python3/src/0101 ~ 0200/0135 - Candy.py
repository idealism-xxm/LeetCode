# 链接：https://leetcode.com/problems/candy/
# 题意：有 n 个孩子站成一排，第 i 个孩子的评分为 ratings[i] ，
#		现在要给这些孩子发糖果，要满足以下限制：
#			1. 每个孩子至少要有一个糖果
#			2. 若 ratings[i] > ratings[i - 1] ，则第 i 个孩子的糖果数要多余第 i - 1 个孩子的糖果数；
#			   若 ratings[i] > ratings[i + 1] ，则第 i 个孩子的糖果数要多余第 i + 1 个孩子的糖果数；
#		求最少要发多少个糖果？


# 数据限制：
#  n == ratings.length
#  1 <= n <= 2 * 10 ^ 4
#  0 <= ratings[i] <= 2 * 10 ^ 4


# 输入： [1,0,2]
# 输出： 5
# 解释： 分别发 2, 1, 2 个糖果

# 输入： [1,2,2]
# 输出： 4
# 解释： 分别发 1, 2, 1 （第三个孩子发一个糖果，因为满足上述两个限制）


# 思路1：DP + 排序
#
#		设 dp[i] 表示第 i 个孩子最少能发的糖果数，
#      那么 dp[i] 可以通过两边的 dp 值进行状态转移获得。
#
#      但如果按照通常的下标顺序遍历的话，那么总有一边的 dp 值无法确定，
#      状态转移也就无法进行。
#
#      不过本题也加了其他限制，所以我们可以自定顺序，
#      即可以通过 rating 从小到大的顺序处理，来进行状态转移，
#      这种处理方法与 LeetCode 329 一致。
#
#      如果我们按照 rating 从小到大的顺序进行状态转移，
#      所有 ratings[j] 小于 ratings[i] 的 dp[j] 都已确定，
#      所以 dp[i] 就能通过两边的的 dp 值转移得到。
#
#      ratings 不会改变，所以我们最开始就将 ratings 转成单元格数组 cells ，
#      其中 cells[i] = (matrix[i], i) ，然后按照 ratings[i] 升序排序。
#
#      初始化令所有 dp[i] = 1 ，即至少要分一个糖果。
#
#      然后遍历 cells 进行状态转移，此时能保证在处理 i 之前，
#      所有 rating 小于 ratings[i] 的 dp 值都已处理完成。
#
#
#      时间复杂度：O(nlogn)
#          1. 需要遍历收集 ratings 全部 O(n) 个数
#          2. 需要对 cells 进行排序，时间复杂度是 O(nlogn)
#          3. 需要遍历 cells 全部 O(n) 个元素进行状态转移
#      空间复杂度：O(n)
#          1. 需要维护 cells 全部 O(n) 个元素


class Solution:
    def candy(self, ratings: List[int]) -> int:
        n: int = len(ratings)
        # 将 ratings 收集成 (rating, index) 的列表，
        # 然后按照 rating 升序排序，方便后续状态转移
        cells: List[Tuple[int, int]] = [(rating, i)for i, rating in enumerate(ratings)]
        cells.sort()
        # dp[i] 表示第 i 个孩子分到的糖果数，初始化最少每人一个
        dp: List[int] = [1] * n
        # ans 维护所有孩子分到的糖果数之和
        ans: int = 0
        # 按 rating 升序进行状态转移，这样就能保证在更新 dp[i] 时，
        # 其左右两侧的 dp 值均已确定
        for (rating, i) in cells:
            # 如果其评分大于左侧的评分，则 dp[i] 至少为 dp[i - 1] + 1 ，
            # 此时不用更新最大值，因为 dp[i] = 1 < 1 + 1 <= dp[i - 1] + 1 
            if i > 0 and rating > ratings[i - 1]:
                dp[i] = dp[i - 1] + 1
            # 如果其评分大于左侧的评分，则 dp[i] 至少为 dp[i + 1] + 1 ，
            # 注意此时要更新最大值
            if i + 1 < n and rating > ratings[i + 1]:
                dp[i] = max(dp[i], dp[i + 1] + 1)
            # 计入第 i 个孩子分到的糖果
            ans += dp[i]

        return ans
