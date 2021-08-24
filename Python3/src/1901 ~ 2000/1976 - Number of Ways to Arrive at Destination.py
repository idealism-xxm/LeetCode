# 链接：https://leetcode.com/problems/number-of-ways-to-arrive-at-destination/
# 题意：给定 n 个节点（ 0 ~ n - 1） ，在给定一些有权无向边列表 roads ，
#       roads[i] = [u_i, v_i, time_i] 表示 u_i 和 v_i 之间的一条无向边，权重为 time_i，
#       求 0 到 n - 1 的最短路有多少条？

# 数据限制：
#   1 <= n <= 200
#   n - 1 <= roads.length <= n * (n - 1) / 2
#   roads[i].length == 3
#   0 <= u_i, v_i <= n - 1
#   1 <= time_i <= 10 ^ 9
#   u_i != v_i
#   两个节点之间最多只有一条边
#   整个图是联通的

# 输入： n = 7, roads = [[0,6,7],[0,1,2],[1,2,3],[1,3,3],[6,3,3],[3,5,1],[6,5,1],[2,5,1],[0,4,5],[4,6,2]]
# 输出： 4
# 解释： 最短路为 7
#        0 ➝ 6
#        0 ➝ 4 ➝ 6
#        0 ➝ 1 ➝ 2 ➝ 5 ➝ 6
#        0 ➝ 1 ➝ 3 ➝ 5 ➝ 6

# 输入： n = 2, roads = [[1,0,10]]
# 输出： 1
# 解释： 最短路为 10
#       0 ➝ 1


# 思路1： Floyd
#
#       本题就是在求最短路的同时，再加一个状态，就是记录当前最短路的数量，
#       只要能求最短路就行，看了下数据范围，就直接用了最简单的 Floyd 算法。
#
#       最开始先初始化邻接矩阵，把所有位置的权重初始化为 MX ，把有无向边的位置初始化为权重。
#       然后就可以直接跑 Floyd 算法了，
#           1. 当 (i, j) 经过 k 的路径更短时，那么需要更新最短路，同时重新计算路径数 = cnt[i][k] * cnt[k][j]
#           2. 当 (i, j) 经过 k 的路径也是最段路时，那么需要加上这部分的路径数 cnt[i][k] * cnt[k][j]
#       
#       时间复杂度： O(n ^ 3)
#       空间复杂度： O(n ^ 2)

MX = 1000000000000
MOD = 1000000007

class Solution:
    def __init__(self):
        self.dist = [[0] * 203 for _ in range(203)]
        self.cnt = [[0] * 203 for _ in range(203)]

    def countPaths(self, n: int, roads: List[List[int]]) -> int:
        # 初始化邻接矩阵和最短路径数
        dist, cnt = self.dist, self.cnt
        for i in range(n):
            dist[i][i] = MX
            # 自己到自己的最短路径数为 1
            cnt[i][i] = 1
            for j in range(i):
                dist[i][j] = dist[j][i] = MX
                cnt[i][j] = cnt[j][i] = 0
        # 有无向边的需要将最段路变为权重，并更新最短路径数为 1
        for u, v, t in roads:
            dist[u][v] = dist[v][u] = t
            cnt[u][v] = cnt[v][u] = 1
        
        # 枚举要经过的点 k
        for k in range(n):
            # 枚举起点
            for i in range(n):
                # 枚举终点
                for j in range(i):
                    # 计算 (i, j) 经过 k 时的最短路径 cost
                    cost = dist[i][k] + dist[k][j]
                    # 如果其更短，则需要更新 (i, j) 之间的最短路径，并更新最短路径数
                    if cost < dist[i][j]:
                        dist[i][j] = dist[j][i] = cost
                        cnt[i][j] = cnt[j][i] = cnt[i][k] * cnt[k][j]
                    elif cost == dist[i][j]:
                        # 如果其就是最短路径，需要加上经过 k 时产生的最短路径数
                        cnt[i][j] = cnt[j][i] = (cnt[i][j] + cnt[i][k] * cnt[k][j]) % MOD
        # 最后返回 (0, n - 1) 的最短路径数
        return cnt[0][n - 1] 


# 思路2： Dijkstra
#
#       本题就是在求最短路的同时，再加一个状态，就是记录当前最短路的数量，
#       只要能求最短路就行，可以使用时间复杂度更低的 Dijkstra 算法。
#
#       最开始先初始化邻接表，。
#       然后就可以直接跑 Dijkstra 算法了，如果发现计算的路径已经超过了 n - 1 的路径，那么就可以跳出循环
#           1. 当到达 nxt 的路径更短时，那么需要更新最短路，同时重新计算路径数 = cur_cnt
#           2. 当到达 nxt 的路径也是最段路时，那么需要加上这部分的路径数 cur_cnt
#       
#       时间复杂度： O(ElogE)
#       空间复杂度： O(E + V)

MX = 1000000000000
MOD = 1000000007

class Solution:
    def __init__(self):
        self.dist = [0 for _ in range(203)]
        self.cnt = [0 for _ in range(203)]

    def countPaths(self, n: int, roads: List[List[int]]) -> int:
        # 初始化最短路径和最短路径数
        dist, cnt = self.dist, self.cnt
        for i in range(n):
            dist[i] = MX
            cnt[i] = 0
        # 自己到自己的最短路径数为 1
        cnt[0] = 1
        # 构建邻接表
        edges = defaultdict(list)
        for u, v, t in roads:
            edges[u].append((v, t))
            edges[v].append((u, t))
        
        # 初始放入 0 ，最短路径为 0
        q = [(0, 0)]
        heapq.heapify(q)
        while len(q):
            cur_cost, cur = heapq.heappop(q)
            # 如果到 cur 的最短路已经超过了到 n - 1 的最短路，那么就不用继续更新了
            if cur_cost > dist[n - 1]:
                break
            # 遍历 cur 的邻接点
            for nxt, t in edges[cur]:
                # 计算到该点的最短路
                nxt_cost = cur_cost + t
                # 如果 nxt_cost 更短，则更新最短路，并更新最短路径数
                if nxt_cost < dist[nxt]:
                    dist[nxt] = nxt_cost
                    cnt[nxt] = cnt[cur]
                    heapq.heappush(q, (dist[nxt], nxt))
                elif nxt_cost == dist[nxt]:
                    # 如果 nxt_cost 就是最短路，则加上从 cur 来的最短路径数
                    cnt[nxt] = (cnt[nxt] + cnt[cur]) % MOD

        # 最后返回 n - 1 的最短路径数
        return cnt[n - 1] 
