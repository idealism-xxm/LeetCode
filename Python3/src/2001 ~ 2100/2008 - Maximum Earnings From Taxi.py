# 链接：https://leetcode.com/problems/maximum-earnings-from-taxi/
# 题意：在一条路上有 n 个点，用 1~n 表示，一辆出租车会按顺序经过 1~n 个点，
#       给定一些乘客的起止点和小费 rides[i] = [start_i, end_i, tip_i] ，
#       拉第 i 个乘客能赚 end_i - start_i + tip_i 元，
#       出租车同时一刻最多能拉一个人，求出租车最多能赚多少？

# 数据限制：
#   1 <= n <= 10 ^ 5
#   1 <= rides.length <= 3 * 10 ^ 4
#   rides[i].length == 3
#   1 <= start_i < end_i <= n
#   1 <= tip_i <= 10 ^ 5

# 输入： n = 5, rides = [[2,5,4],[1,5,1]]
# 输出： 7
# 解释：
#   拉第 0 个乘客，赚 5 - 2 + 4 = 7 元

# 输入： n = 20, rides = [[1,6,1],[3,10,2],[10,12,3],[11,12,2],[12,15,2],[13,18,1]]
# 输出： 20
# 解释： 
#   拉第 1 个乘客，从  3 到 10 ，小费 2 两元，总共 10 -  3 + 2 = 9 元
#   拉第 2 个乘客，从 10 到 12 ，小费 3 两元，总共 12 - 10 + 3 = 5 元
#   拉第 5 个乘客，从 13 到 18 ，小费 1 两元，总共 18 - 13 + 1 = 6 元
#   总共能赚 9 + 5 + 6 = 20


# 思路1： DP
#
#       我们先用 map 映射每个起点的乘客列表 start_to_rides ，
#       设 dp[i] 表示汽车走到第 i 个点时，能赚的最大钱数。
#
#       初始化： dp[1~n] = 0
#       状态转移：
#           枚举当前点上车的乘客列表 start_to_rides[i] ，
#           假设当前乘客为 [start, end, tip] ，那么将其拉到 end 处会赚钱，更新 dp[end] 即可，
#           dp[end] = max(dp[end, dp[start] + end - start + tip)
#
#       最终 dp[n] 就是出租车能赚的最大钱数
#       
#       时间复杂度： O(n + m)
#       空间复杂度： O(n + m)


class Solution:
    def maxTaxiEarnings(self, n: int, rides: List[List[int]]) -> int:
        # 初始化每个起点的乘客列表
        start_to_rides = defaultdict(list)
        for ride in rides:
            start_to_rides[ride[0]].append(ride)
        
        # 初始化 dp
        dp = [0] * (n + 1)
        # 状态转移
        for i in range(1, n + 1):
            # 当前点最大赚的钱数也可以来源于 dp[i] 空载过来
            dp[i] = max(dp[i], dp[i - 1])
            # 枚举该起点的乘客
            for start, end, tip in start_to_rides[i]:
                # 更新终点 end 处能赚到的最大钱数
                dp[end] = max(dp[end], dp[start] + end - start + tip)
        
        return dp[n]


# 思路2： DP
#
#       看了讨论区后，发现了另一种思路，
#       思路 1 是以距离为状态枚举，而这里是以乘客为状态枚举，
#       这样在距离特别大的情况下也可以很快求出结果。
#
#       首先对 rides 按照起点升序排序，
#       设 dp[i] 表示 rides[i:] 中能赚到的最大钱数，
#
#       初始化： dp[1~m] = 0
#       状态转移：
#           倒序枚举乘客，假设当前第 i 个乘客为 [start, end, tip] ，
#           拉完 rides[i:] 中的某些乘客后能赚到的最大钱数有两种情况：
#               1. 不拉当前乘客，则能赚 dp[i + 1] 元
#               2. 拉当前乘客，走到 end 处，赚 end - start + tip 元，
#                   然后从 end 开始继续拉乘客，赚取可能的最大的钱。
#                   我们注意到 rides 是升序的，所以我们可以二分 rides[i:] 中的乘客，
#                   找到第一个乘客的起点大于等于 end 的乘客下标 k ，
#                   则后续可以赚到的最大钱数为 dp[k] ，
#                   即总共赚到的钱数为 end - start + tip + dp[k]
#           综上： dp[i] = max(dp[i + 1], end - start + tip + dp[k])
#               
#
#       最终 dp[0] 就是出租车能赚的最大钱数
#       
#       时间复杂度： O(mlogm)
#       空间复杂度： O(m)


class Solution:
    def maxTaxiEarnings(self, n: int, rides: List[List[int]]) -> int:
        # 乘客列表按照起点升序排序
        rides.sort()
        
        # 初始化 dp
        m = len(rides)
        dp = [0] * (m + 1)
        # 状态转移，倒序枚举乘客
        for i in range(m - 1, -1, -1):
            start, end, tip = rides[i]
            # 二分找到第一个起点大于等于终点的乘客
            l, r = i + 1, m - 1
            while l <= r:
                mid = (l + r) // 2
                if rides[mid][0] >= end:
                    r = mid - 1
                else:
                    l = mid + 1
            # 第 k 个乘客的起点是第一个大于等于 end 的乘客
            k = l

            # 1. 不拉当前乘客，则能赚 dp[i + 1] 元
            # 2. 拉当前乘客，走到 end 处，能赚 end - start + tip 元，
            #    能从剩余的 rides[k:] 中赚到 dp[k] 元
            dp[i] = max(dp[i + 1], end - start + tip + dp[k])

        return dp[0]
