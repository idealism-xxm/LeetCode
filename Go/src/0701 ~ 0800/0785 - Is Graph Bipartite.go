// 链接：https://leetcode.com/problems/is-graph-bipartite/
// 题意：给定一个无向图，判断其是否是一个二分图？
//
//      二分图能被分成两个独立的点集 A 和 B ，
//      且图中过的每条边都连接了 A 和 B 中的各一个点。


// 数据限制：
//  graph.length == n
//  1 <= n <= 100
//  0 <= graph[u].length < n
//  0 <= graph[u][i] <= n - 1
//  graph[u] 不含 u ，即没有自环
//  graph[u] 中的所有值都不同，即不含重边
//  如果 graph[u] 包含 v ，则 graph[v] 包含 u


// 输入： graph = [[1,2,3],[0,2],[0,1,3],[0,2]]
// 输出： false
// 解释： 没法将图分成两个独立的点集，
//       使得每条边都分别连接了两个点集中的点

// 输入： graph = [[1,3],[0,2],[1,3],[0,2]]
// 输出： true
// 解释： 可以将点分为 {0, 2} 和 {1, 3} 两个集合


// 思路1： DFS
//
//      我们可以使用染色的方法，将图中的点分成两个独立的点集合。
//
//      维护数组 colors ， 其中 colors[i] 表示点 i 的颜色，
//      0 表示点 i 没有被染色，即还不在集合中；
//      1 和 -1 分别表示两个不同集合的颜色。
//
//      那么我们可以枚举所有的点 i ，如果该未染色，
//      则使用 dfs(i, 1) 进行染色，默认染色为 1 。
//      若染色失败，则点 i 所在的联通子图不是二分图，
//      直接返回 false 。
//
//      在 dfs(cur, color) 中，我们将点 cur 染色为 color ，
//      然后遍历 cur 的所有邻接点 nxt ：
//          1. 如果 nxt 没有被染色，则使用 dfs(nxt, -color) 染色，
//              若染色失败，则 cur 所在的联通子图不是二分图，
//              直接返回 false 。
//          2. 如果 nxt 已经被染色且 nxt 的颜色与 color 相同，
//              则 cur 所在的联通子图不是二分图，直接返回 false 。
//
//      遍历完所有邻接点还未返回，
//      则 cur 所在的联通子图是二分图，返回 true 。
//
//
//      设 n 为点的数量， m 为边的数量。
//
//      时间复杂度：O(n + m)
//          1. 需要遍历全部 O(n) 个点
//          2. 需要遍历全部 O(m) 条边
//      空间复杂度：O(n)
//          1. 需要递归的深度就是一个联通子图的大小，
//              最差情况下，这个图是二分图，
//              且全部 O(n) 点在同一个联通子图中


func isBipartite(graph [][]int) bool {
    // colors[i] 表示每个点的颜色
    //  0 表示未染色，即还不在集合中
    //  1 和 -1 分别表示两个不同集合的颜色
    colors := make([]int, len(graph))
    // 遍历所有点
    for i := 0; i < len(graph); i++ {
        // 如果点 i 未染色 且 染色失败，
        // 则说明点 i 所在的联通子图不是二分图，
        // 直接返回 false
        if colors[i] == 0 && !dfs(graph, colors, i, 1) {
            return false
        }
    }
    // 此时所有点都染色成功，
    // 即是一个二分图，返回 true
    return true
}

func dfs(graph [][]int, colors []int, cur, color int) bool {
    if colors[cur] != 0 {
        // 如果点 cur 已经被染色，
        // 那么当且仅当点 cur 的颜色是 color 时，
        // 才表明能染色成功
        return colors[cur] == color
    }

    // 将点 cur 染成 color
    colors[cur] = color
    // 遍历点 cur 的所有邻接点
    for _, nxt := range graph[cur] {
        // 如果点 nxt 未染色 且 染色为 -color 失败，
        // 则说明点 nxt 所在的联通图不是二分图，
        // 直接返回 false
        if !dfs(graph, colors, nxt, -color) {
            return false
        }
    }
    // 此时当前联通子图点所有点都染色成功，
    // 即当前联通子图是一个二分图，返回 true
    return true
}



// 思路2： 并查集
//
//      二分图需要将图分成两个独立的点集，
//      并且每条边的两个点分别在不同的集合中。
//
//      那么点 i 的所有邻接点必定在同一个集合中，
//      我们可以使用并查集来维护这个关系。
//
//      所以我们可以枚举每个点 i ，
//      然后再枚举其邻接点 j ：
//          1. 如果 i 和 j 在同一个集合中，
//              那么它们所在的联通子图不是二分图，
//              直接返回 false
//          2. 如果 i 和 j 不在同一个集合中，
//              那么将 i 的第 0 个邻接点和 j 合并入一个集合中
//
//      遍历完所邻接点还没有返回，则所有边的两点都不在同一个集合中，
//      即当前图是一个二分图，返回 true
//
//
//      设 n 为点的数量， m 为边的数量。
//
//      时间复杂度：O(n + m * α(n))
//          1. 需要遍历全部 O(n) 个点
//          2. 需要遍历全部 O(m) 条边
//          3. 遍历每条边时，都需要执行常数次并查集的操作，
//              时间复杂度为 O(α(n))
//      空间复杂度：O(n)
//          1. 需要维护一个大小为 O(n) 的并查集


func isBipartite(graph [][]int) bool {
    // 初始化一个大小为 n 的并查集
    unionFind := newUnionFind(len(graph))
    // 遍历所有的点
    for i, adjacent := range graph {
        // 遍历所有邻接点
        for _, j := range adjacent {
            // 如果 i 和 j 已经在同一个集合中，
            // 则说明它们所在的联通子图不是二分图，
            // 返回 false
            if unionFind.find(i) == unionFind.find(j) {
                return false
            }

            // 否则将 i 所有的邻接点合并入同一个集合
            unionFind.union(adjacent[0], j)
        }
    }

    // 此时，所有边的两点都不在同一个集合中，
    // 则说明是一个二分图，返回 true
    return true
}

// 并查集
type UnionFind struct {
    // parent[i] 表示第 i 个元素所指向的父元素
    parent []int
    // rank[i] 表示以第 i 个元素的深度（秩），
    // 当 i 是根元素（即 parent[i] == i ）时有效
    rank []int
}


    // 初始化一个大小为 n 的并查集
func newUnionFind(n int) *UnionFind {
    parent := make([]int, n)
    rank := make([]int, n)
    for i := 0; i < n; i++ {
        // 初始每个元素的父元素都是自己
        parent[i] = i
        // 初始化深度（秩）都是 1
        rank[i] = 1
    }
    return &UnionFind{parent, rank}
}

// 查找元素 x 所在集合的根元素
func (uf *UnionFind) find(x int) int {
    if uf.parent[x] == x {
        // 如果 x 的父元素是自己，那么 x 是根元素
        return x
    }

    // 如果 x 的父元素不是自己，那么递归查找其所在集合的根元素。
    // 这里使用路径压缩优化，将路径上所有的元素都直接挂在根元素下
    uf.parent[x] = uf.find(uf.parent[x])
    // 返回 x 所在集合的根元素
    return uf.parent[x]
}

// 合并元素 x 和 y 所在的集合
func (uf *UnionFind) union(x int, y int) {
    // 找到 x 和 y 所在集合的根元素
    xRoot := uf.find(x);
    yRoot := uf.find(y);
    // 如果 x 和 y 在同一个集合，则不需要合并
    if xRoot == yRoot {
        return
    }

    if uf.rank[xRoot] < uf.rank[yRoot] {
        // 如果 xRoot 深度（秩）更小，
        // 则将 yRoot 合并入 xRoot 中
        uf.parent[xRoot] = yRoot
    } else if uf.rank[xRoot] > uf.rank[yRoot] {
        // 如果 xRoot 深度（秩）更大，
        // 则将 xRoot 合并入 yRoot 中
        uf.parent[yRoot] = xRoot
    } else {
        // 如果 xRoot 深度（秩）相等，
        // 则将 yRoot 合并入 xRoot 中
        uf.parent[yRoot] = xRoot
        // 同时将 xRoot 的深度（秩）加 1
        uf.rank[xRoot] += 1
    }
}
