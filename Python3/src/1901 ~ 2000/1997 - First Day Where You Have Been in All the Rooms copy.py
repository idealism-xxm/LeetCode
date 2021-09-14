# 链接：https://leetcode.com/problems/first-day-where-you-have-been-in-all-the-rooms/
# 题意：有 n 个房间需要参观，第 0 天在房间 0 ，给定一个长度为 n 的数组 nextVisit ，
#       参观房间的规则如下：某一天在房间 i ，
#           如果当前是第奇数次参观房间 i ，则下一次要去房间 nextVisit[i] (0 <= nextVisit[i] <= i) ；
#           如果当前时第偶数次参观房间 i ，则下一次要去房间 (i + 1) % n 。
#       求第一次参观完所有的房间是第几天（结果模 1e9 + 7）？

# 数据限制：
#   n == nextVisit.length
#   2 <= n <= 10 ^ 5
#   0 <= nextVisit[i] <= i

# 输入： nextVisit = [0,0]
# 输出： 2
# 解释：
#   - 第 0 天，参观房间 0 ，参观过房间 0 总共 1 次（奇数），下一天参观房间 nextVisit[0] = 0
#   - 第 1 天，参观房间 0 ，参观过房间 0 总共 2 次（偶数），下一天参观房间 0 + 1 = 1
#   - 第 2 天，参观房间 1 ，此时已参观完所有房间

# 输入： nextVisit = [0,0,2]
# 输出： 6
# 解释：
#   参观房间顺序为 [0, 0, 1, 0, 0, 1, 2, ...]

# 输入： nextVisit = [0,1,2,0]
# 输出： 6
# 解释：
#   参观房间顺序为 [0, 0, 1, 1, 2, 2, 3, ...]


# 思路： DP
#
#       我们可以发现要第一次参观第 i 个房间，那么前一天必须是偶数次参观第 i - 1 个房间，
#       我们设：
#           dp[i][1] 表示第一次参观第 i 个房间要用的天数，
#           dp[i][0] 表示第二次参观第 i 个房间要用的天数。
#       初始化：dp[0][1] = 0, dp[0][0] = 1
#       状态转移：
#           # 第一次参观房间 i ，前一天一定是第二次参观房间 i - 1
#           dp[i][1] = dp[i - 1][0]
#           # 第二次参观房间 i ，假设 nextVisit[i] = 0 ，
#           # 那么 dp[i][0] = dp[i][1] + dp[i][1]
#           # 由于 nextVisit[i] 是任意值，
#           # 那么就相当于我们直接跳过了第奇数次访问房间 nextVisit[i] 的天数，
#           # 因此 dp[i][0] = dp[i][1] + dp[i][1] + 1 - nextVisit[i]][1]
#           dp[i][0] = dp[i][1] + dp[i][1] + 1 - nextVisit[i]][1]
#        最终结果为第一次访问房间 i - 1 时的天数： dp[n - 1][1]
#
#       时间复杂度： O(n)
#       空间复杂度： O(n)


MOD = 1000000007


class Solution:
    def firstDayBeenInAllRooms(self, nextVisit: List[int]) -> int:
        dp = [[-1, -1] for _ in range(len(nextVisit))]
        dp[0] = [1, 0]
        for i in range(1, len(nextVisit)):
            # 第一次参观房间 i ，前一天一定是第二次参观房间 i - 1
            dp[i][1] = (dp[i - 1][0] + 1) % MOD
            # 第二次参观房间 i ，假设 nextVisit[i] = 0 ，
            # 那么 dp[i][0] = dp[i][1] + dp[i][1] + 1
            # 由于 nextVisit[i] 是任意值，
            # 那么就相当于我们直接跳过了第奇数次访问房间 nextVisit[i] 的天数，
            # 因此 dp[i][0] = dp[i][1] + dp[i][1] + 1 - nextVisit[i]][1]
            dp[i][0] = (dp[i][1] + dp[i][1] + 1 - dp[nextVisit[i]][1] + MOD) % MOD

        # 最终结果为第一次访问房间 i - 1 时的天数
        return dp[len(nextVisit) - 1][1]
