# 链接：https://leetcode.com/problems/unique-paths/
# 题意：给定一个 m * n 的网格，一个机器人在左上角 (0, 0) 处，
#		它每一次可以向右或者向下移动一格，
#		求移动到右下角 (m - 1, n - 1) 有多少种路径？


# 数据限制：
#  1 <= m, n <= 100


# 输入：m = 7, n = 3
# 输出：28

# 输入：m = 3, n = 2
# 输出：3
# 解释：总共有 3 种路径：
#			1. 向右 -> 向下 -> 向下
#			2. 向下 -> 向下 -> 向右
#			3. 向下 -> 向右 -> 向下


# 思路： DP
#
#		本题是 LeetCode 63 的简化版，可以直接复用其思路，化用对应代码即可。
#
#      设 dp[i][j] 表示从 (0, 0) 到 (i - 1, j - 1) 的不同路径数，
#      这里 i 和 j 都比对应格子的下标大 1 ，
#      是为了方便处理第一行和第一列的情况，避免越界判断。
#
#      初始状态： dp[i][j] = 0; dp[0][1] = 1 。
#          【注意】这里也可以令 dp[1][0] = 1 ，但两者不能同时为 1 ，
#          这样初始化也是为了方便后续处理，并能转移出 dp[1][1] = 1 。
#
#      状态转移方程： dp[i][j] = dp[i - 1][j] + dp[i][j - 1] ，
#			即可以从上边或左边的格子走过来。
#
#      最后， dp[m][n] 就是从 (0, 0) 到 (m - 1, n - 1) 的不同路径数。
#
#
#      DP 常见的三种优化方式见 LeetCode 583 这题的思路，
#      本题可以采用一维数组 + 临时变量的方式进行优化，能将空间复杂度从 O(mn) 优化为 O(n) 。
#
#      因为 dp[i][j] 仅依赖 dp[i - 1][j] 和 dp[i][j - 1] ，
#      所以我们可以优化后使用长度为 n 的一维数组 dp ，
#      其中 dp[j] 表示走到 (i - 1, j - 1) 的不同路径数。
#
#      那么初始状态变为： dp[j] = 0; dp[1] = 1 。
#      状态转移方程变为： dp[j] = dp[j] + dp[j - 1] 。
#
#
#      时间复杂度：O(mn)
#          1. 需要遍历 dp 全部 O(mn) 个状态
#      空间复杂度：O(n)
#          1. 需要维护一个大小为 O(n) 的数组 dp


class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        # dp[j] 表示从 (0, 0) 到 (i - 1, j - 1) 的不同路径数，初始化均为 0 。
        # 这里使用了一维数组 + 临时变量的优化，能将空间复杂度从 O(mn) 降到 O(n) 。
        dp: List[int] = [0] * (n + 1)
        # 为了方便后续获得 dp[1][1] 为 1 ，这里将 dp[0][1] 设置为 1
        dp[1] = 1
        # 遍历当前要走到的格子下标 (i, j)
        for i in range(1, m + 1):
            for j in range(1, n + 1):
                # 可以从上边和左边走过来：
                #  1. dp[j] 表示从上边走过来，代表原来的 dp[i - 1][j]
                #  2. dp[j - 1] 从左边走过来，代表原来的 dp[i][j - 1]
                dp[j] = dp[j] + dp[j - 1]

        # dp[n] 就是从 (0, 0) 到 (m - 1, n - 1) 的不同路径数
        return dp[n]