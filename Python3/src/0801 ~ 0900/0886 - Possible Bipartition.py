# 链接：https://leetcode.com/problems/possible-bipartition/
# 题意：给定一个整数 n 和一个数组 dislikes 。
#      其中， n 表示人数，
#      dislikes[i] = [a_i, b_i] 表示 a_i 不喜欢 b_i 。
#
#      求是否能将这 n 个人分成两组，使得存在不喜欢关系的人不在同一组？


# 数据限制：
#  1 <= n <= 2000
#  0 <= dislikes.length <= 10 ^ 4
#  dislikes[i].length == 2
#  1 <= dislikes[i][j] <= n
#  a_i < b_i
#  dislikes 中所有的元素都各不相同


# 输入： n = 4, dislikes = [[1,2],[1,3],[2,4]]
# 输出： true
# 解释： 第一组： {1, 4}
#       第二组： {2, 3}

# 输入： n = 3, dislikes = [[1,2],[1,3],[2,3]]
# 输出： false
# 解释： 1 不喜欢 2 和 3 ，所以 2 和 3 需要在同一组；
#       但 2 也不喜欢 3 ，所以 2 和 3 需要在不同组。

# 输入： n = 5, dislikes = [[1,2],[2,3],[3,4],[4,5],[1,5]]
# 输出： false


# 思路1： 递归/DFS
#
#      本题是 LeetCode 785 的加强版，给「判断二分图」增加了一个场景，
#      需要自己推导出题目的实际目的：给定一个无向图，判断其是否是一个二分图？
#
#      所以我们可以先建立邻接表，
#      然后调用 LeetCode 785 的「判断二分图」的方法即可。
#
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
#      空间复杂度：O(n + m)
#          1. 需要维护邻接表中全部 O(m) 条边
#          2. 递归深度就是一个联通子图的大小，
#              最差情况下，这个图是二分图，且全部 O(n) 点在同一个联通子图中


class Solution:
    def possibleBipartition(self, n: int, dislikes: List[List[int]]) -> bool:
        # 构建邻接表
        adj: List[List[int]] = [[] for _ in range(n)]
        for a, b in dislikes:
            # 这里将 [1, n] 映射成 [0, n - 1] ，方便直接复用 isBipartite
            a, b = a - 1, b - 1
            adj[a].append(b)
            adj[b].append(a)

        # 判断该无向图是否为二分图
        return self.isBipartite(adj)

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


# 思路2： 并查集
#
#      本题是 LeetCode 785 的加强版，给「判断二分图」增加了一个场景，
#      需要自己推导出题目的实际目的：给定一个无向图，判断其是否是一个二分图？
#
#      所以我们可以先建立邻接表，
#      然后调用 LeetCode 785 的「判断二分图」的方法即可。
#
#
#      二分图需要将图分成两个独立的点集，
#      并且每条边的两个点分别在不同的集合中。
#
#      那么点 i 的所有邻接点必定在同一个集合中，
#      我们可以使用并查集来维护这个关系。
#
#      所以我们可以枚举每个点 i ，
#      然后再枚举其邻接点 j ：
#          1. 如果 i 和 j 在同一个集合中，
#              那么它们所在的联通子图不是二分图，
#              直接返回 false
#          2. 如果 i 和 j 不在同一个集合中，
#              那么将 i 的第 0 个邻接点和 j 合并入一个集合中
#
#      遍历完所邻接点还没有返回，则所有边的两点都不在同一个集合中，
#      即当前图是一个二分图，返回 true
#
#
#      设 n 为点的数量， m 为边的数量。
#
#      时间复杂度：O(n + m * α(n))
#          1. 需要遍历全部 O(n) 个点
#          2. 需要遍历全部 O(m) 条边
#          3. 遍历每条边时，都需要执行常数次并查集的操作，
#              时间复杂度为 O(α(n))
#      空间复杂度：O(n + m)
#          1. 需要维护邻接表中全部 O(m) 条边
#          2. 需要维护一个大小为 O(n) 的并查集


class Solution:
    def possibleBipartition(self, n: int, dislikes: List[List[int]]) -> bool:
        # 构建邻接表
        adj: List[List[int]] = [[] for _ in range(n)]
        for a, b in dislikes:
            # 这里将 [1, n] 映射成 [0, n - 1] ，方便直接复用 isBipartite
            a, b = a - 1, b - 1
            adj[a].append(b)
            adj[b].append(a)

        # 判断该无向图是否为二分图
        return self.isBipartite(adj)

    def isBipartite(self, graph: List[List[int]]) -> bool:
        # 初始化一个大小为 n 的并查集
        union_find: UnionFind = UnionFind(len(graph))
        # 遍历所有的点
        for i, adjacent in enumerate(graph):
            # 遍历所有邻接点
            for j in adjacent:
                # 如果 i 和 j 已经在同一个集合中，
                # 则说明它们所在的联通子图不是二分图，
                # 返回 false
                if union_find.find(i) == union_find.find(j):
                    return False
    
                # 否则将 i 所有的邻接点合并入同一个集合
                union_find.union(adjacent[0], j)

        # 此时，所有边的两点都不在同一个集合中，
        # 则说明是一个二分图，返回 true
        return True


# 并查集
class UnionFind:

    # 初始化一个大小为 n 的并查集
    def __init__(self, n: int):
        # parent[i] 表示第 i 个元素所指向的父元素
        # 初始每个元素的父元素都是自己
        self.parent = list(range(n))
        # rank[i] 表示以第 i 个元素的深度（秩），
        # 当 i 是根元素（即 parent[i] == i ）时有效
        # 初始化深度（秩）都是 1
        self.rank = [1] * n

    # 查找元素 x 所在集合的根元素
    def find(self, x: int) -> int:
        if self.parent[x] == x:
            # 如果 x 的父元素是自己，那么 x 是根元素
            return x

        # 如果 x 的父元素不是自己，那么递归查找其所在集合的根元素。
        # 这里使用路径压缩优化，将路径上所有的元素都直接挂在根元素下
        self.parent[x] = self.find(self.parent[x])
        # 返回 x 所在集合的根元素
        return self.parent[x]

    # 合并元素 x 和 y 所在的集合
    def union(self, x: int, y: int):
        # 找到 x 和 y 所在集合的根元素
        x_root: int = self.find(x)
        y_root: int = self.find(y)
        # 如果 x 和 y 在同一个集合，则不需要合并
        if x_root == y_root:
            return

        if self.rank[x_root] < self.rank[y_root]:
            # 如果 x_root 深度（秩）更小，
            # 则将 y_root 合并入 x_root 中
            self.parent[x_root] = y_root
        elif self.rank[x_root] > self.rank[y_root]:
            # 如果 x_root 深度（秩）更大，
            # 则将 x_root 合并入 y_root 中
            self.parent[y_root] = x_root
        else:
            # 如果 x_root 深度（秩）相等，
            # 则将 y_root 合并入 x_root 中
            self.parent[y_root] = x_root
            # 同时将 x_root 的深度（秩）加 1
            self.rank[x_root] += 1
