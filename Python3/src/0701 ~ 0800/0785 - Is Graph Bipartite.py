# 链接：https://leetcode.com/problems/is-graph-bipartite/
# 题意：给定一个无向图，判断其是否是一个二分图？
#
#      二分图能被分成两个独立的点集 A 和 B ，
#      且图中过的每条边都连接了 A 和 B 中的各一个点。


# 数据限制：
#  graph.length == n
#  1 <= n <= 100
#  0 <= graph[u].length < n
#  0 <= graph[u][i] <= n - 1
#  graph[u] 不含 u ，即没有自环
#  graph[u] 中的所有值都不同，即不含重边
#  如果 graph[u] 包含 v ，则 graph[v] 包含 u


# 输入： graph = [[1,2,3],[0,2],[0,1,3],[0,2]]
# 输出： false
# 解释： 没法将图分成两个独立的点集，
#       使得每条边都分别连接了两个点集中的点

# 输入： graph = [[1,3],[0,2],[1,3],[0,2]]
# 输出： true
# 解释： 可以将点分为 {0, 2} 和 {1, 3} 两个集合


# 思路1： DFS
#
#      我们可以使用染色的方法，将图中的点分成两个独立的点集合。
#
#      维护数组 colors ， 其中 colors[i] 表示点 i 的颜色，
#      0 表示点 i 没有被染色，即还不在集合中；
#      1 和 -1 分别表示两个不同集合的颜色。
#
#      那么我们可以枚举所有的点 i ，如果该未染色，
#      则使用 dfs(i, 1) 进行染色，默认染色为 1 。
#      若染色失败，则点 i 所在的联通子图不是二分图，
#      直接返回 false 。
#
#      在 dfs(cur, color) 中，我们将点 cur 染色为 color ，
#      然后遍历 cur 的所有邻接点 nxt ：
#          1. 如果 nxt 没有被染色，则使用 dfs(nxt, -color) 染色，
#              若染色失败，则 cur 所在的联通子图不是二分图，
#              直接返回 false 。
#          2. 如果 nxt 已经被染色且 nxt 的颜色与 color 相同，
#              则 cur 所在的联通子图不是二分图，直接返回 false 。
#
#      遍历完所有邻接点还未返回，
#      则 cur 所在的联通子图是二分图，返回 true 。
#
#
#      设 n 为点的数量， m 为边的数量。
#
#      时间复杂度：O(n + m)
#          1. 需要遍历全部 O(n) 个点
#          2. 需要遍历全部 O(m) 条边
#      空间复杂度：O(n)
#          1. 需要递归的深度就是一个联通子图的大小，
#              最差情况下，这个图是二分图，
#              且全部 O(n) 点在同一个联通子图中


class Solution:
    def isBipartite(self, graph: List[List[int]]) -> bool:
        # colors[i] 表示每个点的颜色
        #  0 表示未染色，即还不在集合中
        #  1 和 -1 分别表示两个不同集合的颜色
        colors: List[int] = [0] * len(graph)
        # 遍历所有点
        for i in range(len(graph)):
            # 如果点 i 未染色 且 染色失败，
            # 则说明点 i 所在的联通子图不是二分图，
            # 直接返回 false
            if colors[i] == 0 and not Solution.dfs(graph, colors, i, 1):
                return False

        # 此时所有点都染色成功，
        # 即是一个二分图，返回 true
        return True

    @staticmethod
    def dfs(graph: List[List[int]], colors: List[int], cur: int, color: int) -> bool:
        if colors[cur] != 0:
            # 如果点 cur 已经被染色，
            # 那么当且仅当点 cur 的颜色是 color 时，
            # 才表明能染色成功
            return colors[cur] == color

        # 将点 cur 染成 color
        colors[cur] = color
        # 遍历点 cur 的所有邻接点
        for nxt in graph[cur]:
            # 如果点 nxt 未染色 且 染色为 -color 失败，
            # 则说明点 nxt 所在的联通图不是二分图，
            # 直接返回 false
            if not Solution.dfs(graph, colors, nxt, -color):
                return False

        # 此时当前联通子图点所有点都染色成功，
        # 即当前联通子图是一个二分图，返回 true
        return True
