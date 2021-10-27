# 链接：https://leetcode-cn.com/problems/the-time-when-the-network-becomes-idle/
# 题意：给定一个无向图网络， edges[i] = [u_i, v_i] 表示节点 u_i 和 v_i 之间有一条边权为 1 的无向边，
#       每个节点每秒钟可以将其接收到的所有数据发送给相邻节点，
#       0 节点为主节点，最开始其他节点都会沿最短路向 0 节点发送数据，
#       0 节点接收到后立刻将数据按原路发回给对应的节点，
#       i 节点发出数据后，如果没有接收到返回给自己的数据，则会每隔 patience[i] 秒再发送一次，
#       求整个网络中最早没有任何数据的时间点？

# 数据限制：
#   n == patience.length
#   2 <= n <= 10 ^ 5
#   patience[0] == 0
#   1 <= patience[i] <= 10 ^ 5 for 1 <= i < n
#   1 <= edges.length <= min(10 ^ 5, n * (n - 1) / 2)
#   edges[i].length == 2
#   0 <= u_i, v_i < n
#   u_i != v)i
#   没有重复边
#   整个图是联通的

# 输入： edges = [[0,1],[1,2]], patience = [0,2,1]
# 输出： 8
# 解释：
#   0 秒：
#       1 节点发送他自己的数据给 0 节点，标记数据为 1A
#       2 节点发送他自己的数据给 0 节点，标记数据为 2A
#
#   1 秒：
#       数据 1A 到达 0 节点， 0 节点处理后立刻发送回复给 1 节点
#       1 节点没有收到回复，但当前时刻 1 < patience[1] = 2 ，所以 1 节点还不会重新发送数据
#       2 节点没有收到回复，当前时刻 1 == patience[2] = 1 ，所以 2 节点会再次发送数据，标记为 2B
#
#   2 秒：
#       数据 1A 的回复到达 1 节点， 1 节点不会再发送任何数据
#       数据 2A 到达 0 节点， 0 节点处理后立刻发送回复给 2 节点
#       2 节点再次发送数据，标记为 2C
#
#   ...
#
#   4 秒：
#       数据 2A 的回复到达 2 节点， 2 节点不会在发送任何数据
#
#   ...
#
#   7 秒：
#       数据 2D 的回复到达 2 节点
#
#   从第 8 秒开始，网络上没有任何数据

# 输入： edges = [[0,1],[0,2],[1,2]], patience = [0,10,10]
# 输出： 3
# 解释：
#   1 节点和 2 节点都会在第 2 秒收到他们发出数据的回复
#   从第 3 秒开始，网络上没有任何数据


# 思路： BFS
#
#       我们找到以主节点为源的单源最短路，由于边的权值都是 1 ，
#       所以我们使用 BFS 求出所有节点的最短路
#
#       假设 i 节点的最短路为 dist[i] ，每隔 patience[i] 秒会重新发送一次数据，
#       i 节点第一次收到回到自身的数据是在 dist[i] * 2 秒，
#       那么 i 节点在此期间发送了 send_count = (dist[i] * 2 - 1) // patience[i] 次数据，
#       且最后一次发送数据的时间点是 send_count * patience[nxt] ，
#       该数据在 dist[i] * 2 后会回到 nxt ，此后 i 节点相关的数据不会存在，
#       那么 nxt 节点相关的空闲时间为 send_count * patience[nxt] + nxt_dist * 2
#                               = (dist[i] * 2 - 1) // patience[i] * patience[nxt] + nxt_dist * 2
#
#       统计所有空闲时间的最大值即可
#
#       时间复杂度： O(n + m)
#       空间复杂度： O(n)


import queue



class Solution:
    def networkBecomesIdle(self, edges: List[List[int]], patience: List[int]) -> int:
        # 构建邻接表
        adj = defaultdict(list)
        for u, v in edges:
            adj[u].append(v)
            adj[v].append(u)

        # 记录最晚空闲时间
        ans = 0
        # 初始化每个节点到主节点的距离
        # -1 表示还未计算到
        dist = [-1] * len(patience)
        dist[0] = 0
        # 初始化 BFS 的队列，最开始只有主节点到自己的距离已知
        q = queue.Queue(len(patience))
        q.put_nowait(0)

        # 当队列非空，则继续遍历
        while not q.empty():
            # 当前队首节点 cur 到主节点的距离为 dist[cur]
            cur = q.get_nowait()
            # 与 cur 相邻且未被遍历过的节点到主节点的距离为 dist[cur] + 1
            nxt_dist = dist[cur] + 1
            # 遍历 cur 的相邻节点
            for nxt in adj[cur]:
                # 如果 nxt 未被遍历过，则更新为当前最短距离
                if dist[nxt] == -1:
                    dist[nxt] = nxt_dist
                    q.put_nowait(nxt)
                    # 计算第一次回到自身前发送了多少次
                    send_count = (2 * nxt_dist - 1) // patience[nxt]
                    # nxt 节点最后一次发送数据的时间点是 send_count * patience[nxt] ，
                    # 该数据在 nxt_dist * 2 后会回到 nxt ，此后 nxt 节点相关的数据不会存在，
                    # 那么 nxt 节点相关的空闲时间为 send_count * patience[nxt] + 2 * nxt_dist
                    ans = max(ans, send_count * patience[nxt] + 2 * nxt_dist)

        return ans + 1
