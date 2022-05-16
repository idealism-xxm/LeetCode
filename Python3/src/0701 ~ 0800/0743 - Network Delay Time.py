# 链接：https://leetcode.com/problems/network-delay-time/
# 题意：有一个 n 个结点的网络，结点编号为 1 到 n 。
#      给定一个列表 times ，表示信号经过有向边传递的时间。
#      times[i] = [u_i, v_i, w_i] 表示一个信号从结点 u_i 发出，
#      到达结点 v_i 的时间为 w_i 。
#
#      现在从指定的结点 k 发出一个信号，需要多久才能时所有结点都收到信号？
#      如果无法使所有结点收到信号，返回 -1 。


# 数据限制：
#  1 <= k <= n <= 100
#  1 <= times.length <= 6000
#  times[i].length == 3
#  1 <= u_i, v_i <= n
#  u_i != v_i
#  0 <= w_i <= 100
#  所有的二元组 (u_i, v_i) 都是不同的（即不含重边）


# 输入： times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
# 输出： 2
# 解释： 2
#  (1)↙ ↓(1)
#    1  3
#       ↓(1)
#       4

# 输入： times = [[1,2,1]], n = 2, k = 1
# 输出： 1
# 解释： 1
#       ↓(1)
#       2


# 输入： times = [[1,2,1]], n = 2, k = 2
# 输出： -1
# 解释： 1
#       ↓(1)
#       2


# 思路： Dijkstra
#
#      这题其实就是求有向图的单源最短路，然后找到到每个结点的最短路的最大值即可，
#      求单源最短路就可以用 Dijkstra 算法。
#
#      先根据 times 构建有向图的邻接表 adj 。
#
#      然后定义一个数组 dist ，
#      其中 dist[i] 表示从结点 k 到结点 i 的最短路长度，初始化均为 MAX 。
#
#      再定义一个优先队列 q ，存储二元组 (d, u) ， d 代表结点 k 到结点 u 的距离， 
#      u 代表结点编号，该优先队列按照 d 升序排列。
#
#      最开始只有结点 k 的距离是已知的，并且为 0 ，
#      所以 dist[k] = 0 ，并将 (0, k) 加入优先队列 q 中。
#
#      当 q 不为空时，我们从 q 队首取出一个二元组 (d, u) ，
#          1. d > dist[u]: 则前边必定以 dist[u] 处理过了，
#              此时处理结果只会更差，所以可以直接跳过
#          2. d == dist[u]: 遍历 u 的所有相邻结点 v 及其花费 w ，
#              继续求解经过 u 到达 v 的长度 d + w ，
#              如果这个长度小于 dist[v] ，则更新 dist[v] 为 d + w ，
#              并将 (dist[v], v) 加入优先队列 q 中。
#          3. d < dist[u]: 不存在这种情况
#
#      当优先队列 q 为空时，所有结点的最短路都已求出，
#      找到这些最短路的最大值 max_dist = max(dist[1:]) 。
#
#      如果 max_dist == MAX ，则说明无法到达结点 i ，返回 -1 ；
#      否则 max_dist 就是最终答案。
#
#
#      设 times 的长度为 E 。
#
#      时间复杂度：O(n + Elogn)
#          1. 需要遍历 times 中全部 O(E) 条边，用于构建邻接表 adj
#          2. Dijkstra 算法的时间复杂度为 O(Elogn)
#          3. 需要遍历 dist 数组全部 O(n) 个最短路，求最短路最大值 max_dist
#      空间复杂度：O(n + E)
#          1. 需要维护一个大小为 O(E) 的邻接表 adj
#          2. 需要维护一个大小为 O(n) 的距离数组 dist
#          3. 需要维护一个优先队列 q ，最多存储 O(E) 个元素


MAX = 0x3f3f3f3f


class Solution:
    def networkDelayTime(self, times: List[List[int]], n: int, k: int) -> int:
        # 根据 times 构建邻接表
        adj: List[List[int]] = [[] for _ in range(n + 1)]
        for u, v, w in times:
            adj[u].append((v, w))

        # dist[i] 表示从 k 到 i 的最短路径长度，初始化均为 MAX
        dist: List[int] = [MAX] * (n + 1)
        # k 到 k 的最短路为 0
        dist[k] = 0
        # 定义优先队列，存储二元组 (d, u) ， d 代表结点 k 到结点 u 的距离，
        # u 代表结点编号，该优先队列按照 d 升序排列。
        # 由于 BinaryHeap 是最大堆，所以 d 需要使用 Reverse 包一层。
        q = []
        # 初始只有结点 k 对应的距离已确定，放入优先队列中
        heapq.heappush(q, (0, k))
        # 当优先队列不为空时，可以继续处理
        while q:
            # 获取队首的二元组 (d, u)
            d, u = heapq.heappop(q)
            # 如果 d > dist[u] ，则前边必定以 dist[u] 处理过了，
            # 此时处理结果只会更差，所以可以直接跳过
            if d > dist[u]:
                continue

            # 此时必定有 d == dist[u]

            # 遍历 u 的所有邻接结点
            for v, w in adj[u]:
                # 如果经过 u 到达 v 的路径比之前的路径更短，
                # 则更新 dist[v] ，并将 (dist[v], v) 加入优先队列
                if dist[v] > d + w:
                    dist[v] = d + w
                    heapq.heappush(q, (dist[v], v))

        # 求所有最短路的最大值，注意结点 0 不存在，不要参与计算
        max_dist = max(dist[1:])
        if max_dist == MAX:
            # 如果最大值是 MAX ，则说明存在结点无法到达，返回 -1
            return -1

        # 如果最大值不是 MAX ，则 max_dist 就是最终答案
        return max_dist
