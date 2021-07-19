# 链接：https://leetcode.com/problems/minimum-cost-to-reach-destination-in-time/
# 题意：给定一个无向图，每条边有时间，每个顶点有费用，求从 src 到 dest 的所有路径中，
#       总时间在 maxTime 内，且所有定点费用最小的路径的费用？

# 数据限制：
#   1 <= maxTime <= 1000
#   n == passingFees.length
#   2 <= n <= 1000
#   n - 1 <= edges.length <= 1000
#   0 <= x_i, y_i <= n - 1
#   1 <= time_i <= 1000
#   1 <= passingFees[j] <= 1000
#   有重复边
#   不含自环

# 输入： maxTime = 30, edges = [[0,1,10],[1,2,10],[2,5,10],[0,3,1],[3,4,10],[4,5,15]], passingFees = [5,1,2,20,20,3]
# 输出： 11
# 解释： 0 -> 1 -> 2 -> 5
#       耗费 30 分钟，费用为 11

# 输入： maxTime = 29, edges = [[0,1,10],[1,2,10],[2,5,10],[0,3,1],[3,4,10],[4,5,15]], passingFees = [5,1,2,20,20,3]
# 输出： 48
# 解释： 0 -> 3 -> 4 -> 5,
#       耗费 26 分钟，费用为 48

# 输入： maxTime = 25, edges = [[0,1,10],[1,2,10],[2,5,10],[0,3,1],[3,4,10],[4,5,15]], passingFees = [5,1,2,20,20,3]
# 输出： -1
# 解释： 没有能 25 分钟达到的路径

# 思路1： Dijkstra
#
#       重新根据初始边和费用重新构造一个有既有时间又有费用的图，
#       即构造临接表 adj ，adj[s] = List[(fee, time, e)] ，
#       表示从 s 到 e 有一条边，该边有费用 fee ，经过会耗时 time
#
#       然后使用 Dijkstra 变形即可，
#       我们在优先队列中存储的数据如下： (cur_fee, cur_time, s)
#       表示当前已经遍历到 s 处，费用为 cur_fee ，耗时为 cur_time ，
#       则初始有 (passingFees[src], 0, src)
#
#       每次更新时，我们只将时间在 maxTime 内的数据放入，这样就一定能把保证时间合法，
#       有两种情况可以更新并放入队列中：
#           1. cur_fee + fee < dist_fee
#               更新 dist_fee
#           2. cur_time + time < dist_time
#               更新 dist_fee
#       然后将 (cur_fee + fee, cur_time + time, s) 放入队列中
#
#       时间复杂度： O((n + V) * logV) ，其中 n 是顶点数量， V 是边数量
#       空间复杂度： O(n)


INF = 1000001


import queue

class Solution:
    def minCost(self, maxTime: int, edges: List[List[int]], passingFees: List[int]) -> int:
        adj: List[List[Tuple[int, int, int]]] = [None] * len(passingFees)
        # 距离数组 (到当前点的总费用，到当前点的总时间)
        dist: List[Tuple[int, int]] = [None] * len(passingFees)
        for i in range(len(passingFees)):
            adj[i] = []
            dist[i] = INF, maxTime + 1
        # 初始化起点的距离
        dist[0] = passingFees[0], 0

        # 初始化临接表
        for s, e, time in edges:
            adj[s].append((e, passingFees[e], time))
            adj[e].append((s, passingFees[s], time))

        # 使用 Dijkstra 算法
        return self.dijkstra(0, len(passingFees) - 1, maxTime, adj, dist)

    def dijkstra(self, src: int, dest: int, max_time: int, adj: List[Tuple[int, int, int]], dist: List[Tuple[int, int]]) -> int:
        q = queue.PriorityQueue()
        # 放入起点的数据（优先 fee 小的，再）
        q.put((*dist[src], src))
        while not q.empty():
            # 拿到当前队首的数据
            cur_fee, cur_time, cur = q.get()
            # 如果找到了重点，由于是优先 fee 小的，所以第一个必定是满足题意的最小的 fee
            if cur == dest:
                return cur_fee

            for nxt, fee, time in adj[cur]:
                # 如果时间超了，则不能走，直接处理下一个
                if cur_time + time > max_time:
                    continue

                nxt_fee, nxt_time = dist[nxt]
                # 如果 fee 更小或者 time 更小，则需要更新 dist ，
                # 当前有可能可以派生出优的结果
                if cur_fee + fee < nxt_fee or cur_time + time < nxt_time:
                    dist[nxt] = min(nxt_fee, cur_fee + fee), min(nxt_time, cur_time + time)
                    q.put((cur_fee + fee, cur_time + time, nxt))

        # 如果最终没有到达 dest ，则无合法解
        return -1


# 思路2： DP
#
#       dp[i][j] 表示第 j 分钟到达 i 时的最小花费
#       则如果当前能到达 i ，则可花 time 和 fee 走到 e 去
#
#       初始化：dp[0][src] = passingFees[src]
#       状态转移：
#           当前花 time 走到 e 去，则 dp[e][i + time] = min(dp[e][i + time], dp[j][i] + fee)
#
#       最终答案时 dp[maxTime][n - 1]
#
#       时间复杂度： O((n + V) * t) ，t 表示时间， V 表示边数
#       空间复杂度： O(n * t)


INF = 1000001


class Solution:
    def minCost(self, maxTime: int, edges: List[List[int]], passingFees: List[int]) -> int:
        n = len(passingFees)
        dp: List[List[int]] = [None] * n
        for i in range(n):
            dp[i] = [INF] * (maxTime + 1)
        dp[0][0] = passingFees[0]

        # 初始化临接表（只记录费用最小的边）
        adj: List[Dict[int, int]] = [None] * n
        for i in range(n):
            adj[i] = {}
        for s, e, time in edges:
            if adj[s].get(e) is None or adj[s][e] >= time:
                adj[s][e] = time
            if adj[e].get(s) is None or adj[e][s] >= time:
                adj[e][s] = time

        # 使用 dp 算法
        for i in range(maxTime):
            for j in range(n):
                # 还是最大值，则当前不可达
                if dp[j][i] == INF:
                    continue

                # 走一条边
                for e, time in adj[j].items():
                    if i + time <= maxTime:
                        dp[e][i + time] = min(dp[e][i + time], dp[j][i] + passingFees[e])

        # 找到到达 n - 1 时的最小花费
        ans = min(dp[n - 1])
        return ans if ans != INF else -1
