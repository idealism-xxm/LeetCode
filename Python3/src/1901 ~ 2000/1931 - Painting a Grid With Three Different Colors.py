# 链接：https://leetcode.com/problems/painting-a-grid-with-three-different-colors/
# 题意：给定一个整型数组 m * n 的格子，可以用三种颜色涂这些格子，要求相邻格子的颜色不相同，
#       求总共有多少种不同的涂法？

# 数据限制：
#   1 <= m <= 5
#   1 <= n <= 1000

# 输入： m = 1, n = 1
# 输出： 3
# 解释：
#       r
#       g
#       b

# 输入： m = 1, n = 2
# 输出： 6
# 解释：
#       rg
#       rb
#       gr
#       gb
#       br
#       bg

# 输入： m = 5, n = 5
# 输出： 580986

# 思路： DP
#
#       由于 m 很小，所以我们可以先计算一列所有合法的状态，
#       合法的状态数为 3 * (2 ^ (m - 1))
#       这个过程时间复杂度为 O(2 ^ m)
#
#       然后我们可以计算每个合法状态可以转移至其他合法状态的集合，
#           nxt[cur] 表示当前状态为 cur 时，下一列所有可转移出的合法状态列表
#       这个过程时间复杂度为 O((2 ^ m) * (2 ^ m)) = O(2 ^ (2m))
#
#       然后我们令为 dp[i][j] 表示第 i 列的状态为 j 时不同的涂法
#       初始化： dp[0][j] = 1 ，其中 j 表示合法的状态
#       状态转移： dp[i + 1][nxt[j]] = (dp[i + 1][nxt[j]] + dp[i][j]) % MOD
#       这个过程时间复杂度为 O(2 ^ (2m) * n)
#
#       如果 n 达到 10 ^ 9 ，则可以使用矩阵快速幂进行转移
#       构建 (2 ^ m) * (2 ^ m) 状态转移矩阵，矩阵乘法复杂度为 O((2 ^ m) ^ 3)
#       快速幂时间复杂度为 O(logn)
#       整体时间复杂度为 O(((2 ^ m) ^ 3) * logn)
#
#       时间复杂度： O(2 ^ (2m) * n)
#       空间复杂度： O((2 ^ m) * n) 【可以使用滚动数组优化为 O(2 ^ m)】


MOD = 1000000007


class Solution:
    def colorTheGrid(self, m: int, n: int) -> int:
        # 初始化合法状态和状态转移集合
        self.state = set()
        self.init_state(0, m, -1)
        self.init_next(m)
        dp = [None] * n
        for i in range(n):
            dp[i] = defaultdict(int)

        for cur in self.state:
            dp[0][cur] = 1

        # 通过第 i 列的状态更新第 i + 1 列的状态
        for i in range(0, n - 1):
            # 遍历所有可能的合法状态
            for cur in self.state:
                # 遍历可以转移到的合法状态
                for nxt in self.nxt[cur]:
                    # 转移涂法数
                    dp[i + 1][nxt] = (dp[i + 1][nxt] + dp[i][cur]) % MOD

        # 计算最后一列所有可能状态的涂法数之和
        return sum(dp[n - 1].values()) % MOD

    def init_next(self, m: int):
        """初始化状态转移集合"""
        self.nxt = defaultdict(list)
        for cur in self.state:
            for nxt in self.state:
                # 如果 cur 可以转移至 nxt ，则放入集合中
                if self.is_ok(cur, nxt, m):
                    self.nxt[cur].append(nxt)

    def is_ok(self, cur: int, nxt: int, m: int):
        """判断两个合法的状态是否能够互相转移"""
        while m > 0:
            # 如果对应位的颜色一样，则不能转移
            if cur % 3 == nxt % 3:
                return False
            cur //= 3
            nxt //= 3
            m -= 1
        return True

    def init_state(self, state: int, remain: int, pre: int):
        """初始化所有合法的状态"""
        # 当所有格子都涂完后，当前状态 state 是一个合法的状态
        if remain == 0:
            self.state.add(state)
            return

        state *= 3
        remain -= 1
        for i in range(3):
            # 如果当前颜色与前一个颜色不一样，则可以选择下一个
            if i != pre:
                self.init_state(state + i, remain, i)
