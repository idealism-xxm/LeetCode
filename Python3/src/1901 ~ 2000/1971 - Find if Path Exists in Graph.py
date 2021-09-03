# 链接：https://leetcode.com/problems/find-if-path-exists-in-graph/
# 题意：给定 n 个节点的无向图，判断 start 和 end 是否联通？

# 数据限制：
#   1 <= n <= 2 * 10 ^ 5
#   0 <= edges.length <= 2 * 10 ^ 5
#   edges[i].length == 2
#   1 <= u_i, v_i <= n - 1
#   u_i != v_i
#   1 <= start, end <= n - 1
#   没有重复边
#   没有自环

# 输入： n = 3, edges = [[0,1],[1,2],[2,0]], start = 0, end = 2
# 输出： true
# 解释： 有两条路径能从 0 到 2
#       - 0 → 1 → 2
#       - 0 → 2

# 输入： n = 6, edges = [[0,1],[0,2],[3,5],[5,4],[4,3]], start = 0, end = 5
# 输出： false
# 解释： 
#       没有路径能从 0 到 5


# 思路： 并查集
#
#       第一反应肯定是 BFS ，每个点每条边只会遍历一次，时间复杂度为 O(E)
#
#       看到有人用并查集，立刻反应过来，这颗题目可以使用，因为只需要判断是否联通即可
#
#       时间复杂度： O(E * alpha(E))
#       空间复杂度： O(1)

class Solution:
    def validPath(self, n: int, edges: List[List[int]], start: int, end: int) -> bool:
        parent = [i for i in range(n)]
        
        def find(x: int) -> int:
            if parent[x] == x:
                return x
            parent[x] = find(parent[x])
            return parent[x]
        
        def union(x: int, y: int) -> None:
            x, y = find(x), find(y)
            if x != y:
                parent[x] = y
        
        # 合并边的两点
        for u, v in edges:
            union(u, v)
        
        # 如果最后 start 和 end 在同一个并查集种，则他们联通
        return find(start) == find(end)
