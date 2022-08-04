// 链接：https://leetcode.com/problems/find-closest-node-to-given-two-nodes/
// 题意：给定一个 n 个点的有向图，每个点最多只有一条出边。
//      edges[i] 表示点 i 有一条指向 edges[i] 的出边，
//      如果点 i 没有出边，则 edges[i] == -1 。
//      【注意】这个图可能有环
//
//      求满足以下要求的点 node ：
//          1. 能从 node1 和 node2 到达
//          2. max(distance(node1, node), distance(node2, node)) 在所有点中是最小的。
//              如果存在多个最小值，则选择下标最小的点。
//      如果没有满足要求的点，返回 -1 。


// 数据限制：
//  n == edges.length
//  2 <= n <= 10 ^ 5
//  -1 <= edges[i] < n
//  edges[i] != i
//  0 <= node1, node2 < n


// 输入： edges = [2,2,3,-1], node1 = 0, node2 = 1
// 输出： 2
// 解释： 1 → 2 ← 0
//           ↓
//           3 
//       点 0 到点 2 的距离是 1 ，点 1 到点 2 的距离是 1 。
//       两个距离的最大值是 1 ，没有其他点的距离的最大值比 1 小。

// 输入： edges = [1,2,-1], node1 = 0, node2 = 2
// 输出： 2
// 解释： 0 → 1 → 2
//       点 0 到点 2 的距离是 2 ，点 2 到点 2 的距离是 0 。
//       两个距离的最大值是 2 ，没有其他点的距离的最大值比 2 小。


// 思路： BFS
//
//      这个有向图无边权，所以可以使用 BFS 来求单源最短路。
//
//      那么我们可以先用 BFS 分别求出 node1 和 node2 到所有点的距离 dist1 和 dist2 。
//
//      然后枚举终点 i ，找到 max(dist1[i], dist2[i]) 最小那个点即可。
//
//      【注意】可能存在点不可达，我们需要注意这种情况。
//
//      可以将距离数组初始化为无穷大 MAX 表示不可达，
//      并初始化答案的点为 -1 ，对应的距离最大值为 MAX 。
//
//      这样就不需要特殊判断不可达的情况了，直接求最大值，然后比较更新即可。
//
//
//      时间复杂度：O(n)
//          1. 以 node1 为起点， BFS 找到 node1 到全部 O(n) 个点的距离
//          2. 以 node2 为起点， BFS 找到 node2 到全部 O(n) 个点的距离
//          3. 枚举全部 O(n) 个点，找到距离最大值最小的那个点
//      空间复杂度：O(n)
//          1. 需要维护两个 O(n) 的距离数组


use std::collections::VecDeque;


impl Solution {
    pub fn closest_meeting_node(edges: Vec<i32>, node1: i32, node2: i32) -> i32 {
        // 用 bfs 求出 node1 和 node2 到所有点的距离 dist1 和 dist2
        let dist1 = Self::bfs(&edges, node1);
        let dist2 = Self::bfs(&edges, node2);
        // 初始化答案为 -1 ，对应的距离最大值为 MAX
        let mut ans = -1;
        let mut ans_dist = i32::MAX;
        // 枚举全部点，找到距离最大值最小的那个点
        for i in 0..edges.len() {
            // 如果距离最大值更小，则更新 ans 和 ans_dist
            if dist1[i].max(dist2[i]) < ans_dist {
                ans = i as i32;
                ans_dist = dist1[i].max(dist2[i]);
            }
        }

        ans
    }
    
    fn bfs(edges: &Vec<i32>, start: i32) -> Vec<i32> {
        let start = start as usize;
        // 初始化为无穷大 MAX ，表示不可达，方便后续处理边界情况
        let mut dist = vec![i32::MAX; edges.len()];
        // 最开始只有 start 本身可达，所以 dist[start] = 0 ，且 start 在队列中
        dist[start] = 0;
        let mut q = VecDeque::new();
        q.push_back(start);
        // 不断获取队列 q 中下一个可达的点 cur
        while let Some(cur) = q.pop_front() {
            // 获取 cur 指向的点 nxt
            let nxt = edges[cur];
            // 如果 nxt 是合法的点，且未遍历过，则更新距离
            if nxt != -1 && dist[nxt as usize] == i32::MAX {
                dist[nxt as usize] = dist[cur] + 1;
                q.push_back(nxt as usize);
            }
        }

        dist
    }
}
