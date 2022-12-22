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



use std::collections::{ HashSet, VecDeque };


impl Solution {
    pub fn valid_path(n: i32, edges: Vec<Vec<i32>>, source: i32, destination: i32) -> bool {
        let (n, source, destination) = (n as usize, source as usize, destination as usize);
        // 构建邻接表 adj ，adj[u] 表示与 u 相连的点集
        let mut adj = vec![vec![]; n];
        for edge in edges {
            let (u, v) = (edge[0] as usize, edge[1] as usize);
            adj[u].push(v);
            adj[v].push(u);
        }

        // visited 标记某个点是否已访问
        let mut visited = HashSet::new();
        // 队列 q ，用于 BFS 遍历
        let mut q = VecDeque::new();
        // 初始时标记 source 已访问，并放入队列中
        visited.insert(source);
        q.push_back(source);
        // 当 q 不为空时，按照以下逻辑循环处理
        while let Some(cur) = q.pop_front() {
            // 如果 cur 就是 destination ，那么存在一条合法路径，直接返回 true
            if cur == destination {
                return true;
            }

            // 遍历与 cur 相连的点 nxt
            for &nxt in adj[cur].iter() {
                // 如果 nxt 还没遍历过，则标记已遍历，并放入队列中
                if !visited.contains(&nxt) {
                    visited.insert(nxt);
                    q.push_back(nxt);
                }
            }
        }

        // 遍历完 source 所在的联通块还没找到 destination ，则说明不存在合法的路径
        false
    }
}


// 思路2： 并查集
//
//      本题其实就是判断 source 和 destination 是否联通，那么可以直接使用并查集进行处理。
//
//      初始化大小为 n 的并查集，并遍历每一条边 edges[i] = [u_i, v_i] ，
//      将并查集中的 u_i 和 v_i 合并。
//
//      最后，如果 source 和 destination 在同一个集合中，则说明他们联通，返回 true ；
//      否则返回 false 。
//
//
//      时间复杂度：O(n + m * α(n))
//          1. 需要初始化并查集中全部 O(n) 个点
//          2. 需要遍历全部 O(m) 条边，每次需要执行时间复杂度为 O(α(n)) 并查集操作
//      空间复杂度：O(n)
//          1. 需要维护并查集中全部 O(n) 个点


impl Solution {
    pub fn valid_path(n: i32, edges: Vec<Vec<i32>>, source: i32, destination: i32) -> bool {
        // 初始化大小为 n 的并查集
        let mut union_find = UnionFind::new(n as usize);
        
        // 合并每一条边的两点
        for edge in edges {
            union_find.union(edge[0] as usize, edge[1] as usize)
        }
        
        // 最后，如果 source 和 destination 在同一个集合中，则说明他们联通
        return union_find.find(source as usize) == union_find.find(destination as usize)
    }
}


// 并查集
struct UnionFind {
    // parent[i] 表示第 i 个元素所指向的父元素
    parent: Vec<usize>,
    // rank[i] 表示以第 i 个元素的深度（秩），
    // 当 i 是根元素（即 parent[i] == i ）时有效
    rank: Vec<usize>,
}

impl UnionFind {
    // 初始化一个大小为 n 的并查集
    fn new(n: usize) -> UnionFind {
        UnionFind {
            // 初始每个元素的父元素都是自己
            parent: (0..n).collect(),
            // 初始化深度（秩）都是 1
            rank: vec![1; n],
        }
    }

    // 查找元素 x 所在集合的根元素
    fn find(&mut self, x: usize) -> usize {
        if self.parent[x] == x {
            // 如果 x 的父元素是自己，那么 x 是根元素
            x
        } else {
            // 如果 x 的父元素不是自己，那么递归查找其所在集合的根元素。
            // 这里使用路径压缩优化，将路径上所有的元素都直接挂在根元素下
            self.parent[x] = self.find(self.parent[x]);
            // 返回 x 所在集合的根元素
            self.parent[x]
        }
    }

    // 合并元素 x 和 y 所在的集合
    fn union(&mut self, x: usize, y: usize) {
        // 找到 x 和 y 所在集合的根元素
        let x_root = self.find(x);
        let y_root = self.find(y);
        // 如果 x 和 y 在同一个集合，则不需要合并
        if x_root == y_root {
            return;
        }

        if self.rank[x_root] < self.rank[y_root] {
            // 如果 x_root 深度（秩）更小，
            // 则将 y_root 合并入 x_root 中
            self.parent[x_root] = y_root;
        } else if self.rank[x_root] > self.rank[y_root] {
            // 如果 x_root 深度（秩）更大，
            // 则将 x_root 合并入 y_root 中
            self.parent[y_root] = x_root;
        } else {
            // 如果 x_root 深度（秩）相等，
            // 则将 y_root 合并入 x_root 中
            self.parent[y_root] = x_root;
            // 同时将 x_root 的深度（秩）加 1
            self.rank[x_root] += 1;
        }
    }
}
