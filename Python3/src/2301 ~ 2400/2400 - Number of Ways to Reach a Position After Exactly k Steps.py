# 链接：https://leetcode.com/problems/number-of-ways-to-reach-a-position-after-exactly-k-steps/
# 题意：给定三个整数 startPos, endPos, k ，求恰好用 k 步从 startPos 走到 endPos 的方案数？
#      如果当前在 i ，那么走一步后可以到 i - 1 或 i + 1 。


# 数据限制：
#  1 <= startPos, endPos, k <= 1000


# 输入： startPos = 1, endPos = 2, k = 3
# 输出： 3
# 解释： 有以下三种方式能恰好用 3 步从 1 走到 2
#       - 1 -> 2 -> 3 -> 2
#       - 1 -> 2 -> 1 -> 2
#       - 1 -> 0 -> 1 -> 2

# 输入： startPos = 2, endPos = 5, k = 10
# 输出： 0
# 解释： 无法恰好用 10 步从 2 走到 5


# 思路： DP + Map
#
#      很容易就会想到使用 DP 进行处理，
#      设 dp[i][j] 表示恰好用 i 步从 start_pos 走到 j 的方案数。
#
#      初始化：初始从 start_pos 恰好走 0 步到 start_pos 的方案数为 1 。
#          即 dp[i][j] = 0; dp[i][start_pos] = 1;
#      状态转移：对于 dp[i][j] 来说，恰好走 i + 1 步只能到达 j - 1 或 j + 1 。
#          所以 dp[i + 1][j - 1] += dp[i][j]; dp[i + 1][j + 1] += dp[i][j];
#
#      由于位置 j 可能为负数，且对于第 i 步来说，至少有一半的位置不合法
#      （要么是奇数位置不合法，要么是偶数位置不合法）。
#
#      所以我们可以使用 map 替代，这样只用对合法的状态进行转移，时空上都有常数优化。
#
#      DP 常见的三种优化方式见 LeetCode 583 这题的思路，
#      本题采用滚动数组的方式进行优化，能将空间复杂度从 O(k ^ 2) 优化为 O(k) 。
#
#
#      时间复杂度：O(k ^ 2)
#          1. 需要进行 O(k) 次状态转移，每次都需要对全部 O(k) 个状态进行转移
#      空间复杂度：O(k)
#          1. 需要维护全部 O(k) 个状态


MOD: int = 1_000_000_007


class Solution:
    def numberOfWays(self, start_pos: int, end_pos: int, k: int) -> int:
        # dp[pos] 表示从 start_pos 恰好走 i 步到 pos 的方案数
        dp: Dict[int, int] = defaultdict(int)
        # 初始从 start_pos 恰好走 0 步到 start_pos 的方案数为 1
        dp[start_pos] = 1
        for _ in range(k):
            # 初始化第 i + 1 步的状态
            nxt: Dict[int, int] = defaultdict(int)
            # 将第 i 步的状态转移至第 i + 1 步的状态
            for j, cnt in dp.items():
                # 将状态 dp[i][j] 转移至 dp[i + 1][j - 1] 和 dp[i + 1][j + 1]
                nxt[j - 1] = (nxt[j - 1] + cnt) % MOD
                nxt[j + 1] = (nxt[j + 1] + cnt) % MOD

            dp = nxt

        # 此时 dp[end_pos] 就是从 start_pos 恰好走 k 步到 end_pos 的方案数
        return dp[end_pos]
