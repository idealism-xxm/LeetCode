// 链接：https://leetcode.com/problems/find-if-path-exists-in-graph/
// 题意：给定一个无自环、无重边的无向图，该图 n 个点的标号为 0 到 n - 1 。
//      edges 表示该图的边， edges[i] = [u_i, v_i] 表示点 u_i 和点 v_i 之间存在一条无向边。
//      求点 source 和点 destination 之间是否存在一条路径？


// 数据限制：
//  1 <= n <= 2 * 10 ^ 5
//  0 <= edges.length <= 2 * 10 ^ 5
//  edges[i].length == 2
//  0 <= u_i, v_i <= n - 1
//  u_i != v_i
//  0 <= source, destination <= n - 1
//  无自环、无重边


// 输入： n = 3, edges = [[0,1],[1,2],[2,0]], source = 0, destination = 2
// 输出： true
// 解释：  0 - 1
//         \ /
//          2
//        
//       有两条从 0 到 2 的路径：
//       · 0 -> 1 -> 2
//       · 0 -> 2

// 输入： n = 6, edges = [[0,1],[0,2],[3,5],[5,4],[4,3]], source = 0, destination = 5
// 输出： false
// 解释：  0 - 1      3 - 4
//         \          \ /
//          2          5
//        
//       从 0 到 5 没有合法路径。


// 思路1： BFS
//
//      直接用 BFS 遍历 source 的联通块即可。
//
//      如果在 BFS 时找到 destination ，则说明存在一条合法路径，直接返回 true 。
//
//      如果 BFS 结束时还未返回，则 source 所在的联通块还没找到 destination ，
//      说明不存在合法的路径
//
//
//      时间复杂度：O(n + m)
//          1. 需要遍历全部 O(n) 个点
//          2. 需要遍历全部 O(m) 条边
//      空间复杂度：O(n + m)
//          1. 需要维护 adj 中全部 O(m) 条边的邻接关系
//          2. 需要维护 visited 中全部 O(n) 个点的访问状态
//          3. 需要维护 q 中全部 O(n) 个点



func validPath(n int, edges [][]int, source int, destination int) bool {
    // 构建邻接表 adj ，adj[u] 表示与 u 相连的点集
    adj := make([][]int, n)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }

    // visited 标记某个点是否已访问
    visited := make(map[int]bool)
    // 队列 q ，用于 BFS 遍历
    var q []int
    // 初始时标记 source 已访问，并放入队列中
    visited[source] = true
    q = append(q, source)
    // 当 q 不为空时，按照以下逻辑循环处理
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        // 如果 cur 就是 destination ，那么存在一条合法路径，直接返回 true
        if cur == destination {
            return true
        }

        // 遍历与 cur 相连的点 nxt
        for _, nxt := range adj[cur] {
            // 如果 nxt 还没遍历过，则标记已遍历，并放入队列中
            if !visited[nxt] {
                visited[nxt] = true
                q = append(q, nxt)
            }
        }
    }

    // 遍历完 source 所在的联通块还没找到 destination ，则说明不存在合法的路径
    return false
}
