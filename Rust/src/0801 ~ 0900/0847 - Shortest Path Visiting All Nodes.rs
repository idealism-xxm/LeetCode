// 链接：https://leetcode.com/problems/shortest-path-visiting-all-nodes/
// 题意：给定一个 n 个结点的无向联通图，结点值为 [0, n - 1] 。
//      给定邻接表 graph ，graph[i] 表示与 i 相连的结点列表。
//
//      返回遍历完所有结点的最短路径的长度，可以从任意结点开始和结束。


// 数据限制：
//  n == graph.length
//  1 <= n <= 12
//  0 <= graph[i].length < n
//  graph[i] 不含 i
//  如果 graph[a] 包含 b ，那么 graph[b] 包含 a
//  给定的图是联通的


// 输入： graph = [[1,2,3],[0],[0],[0]]
// 输出： 4
// 解释： 一种可能的最短路径是： [1, 0, 2, 0, 3]
//       0
//     ↗ ⇅ ↘
//    1  2  3   

// 输入： graph = [[1],[0,2,4],[1,3,4],[2],[1,2]]
// 输出： 4
// 解释： 一种可能的最短路径是： [0,1,4,2,3]
//       2 - 1 ← 0
//       ↓ ↖ ↓
//       3   4


// 思路： BFS + 状态压缩
//
//      因为边长都是 1 ，所以求最短路径就可以使用 BFS 即可。
//
//      由于要求遍历完所有结点，所以需要维护额外状态 state ，
//      以便识别当前已遍历过的结点。
//
//      同时发现结点数最大为 12 ，一般小于 20 的都可以用状态压缩处理。
//
//      我们初始化一个状态压缩距离数组 dist ，
//      其中 dist[i][j] 表示最后在结点 i ，
//      已遍历过的结点状态为 j 时的最短路径长度，
//      初始化都是 i32::MAX 。
//
//      同时定义一个队列 q ，用于 BFS 。
//
//      因为我们可以从任意结点开始，
//      所以我们要将 dist[i][1 << i] 的最短路径长度设为 0 ，
//      并将 (i, 1 << i) 放入队列 q 中。
//
//      然后我们不断获取队首的结点信息 (cur, state) ，
//      并遍历 cur 相邻的结点 next ，
//      同时计算其状态 next_state = state | (1 << next) 。
//
//      再判断这种状态是否遍历过，
//      如果没有遍历过，即 dist[next][next_state] == i32::MAX 时，可以继续处理。
//
//      如果 next_state 已遍历完所有结点，则直接返回 dist[cur][state] + 1 。
//      否则，更新 dist[next][next_state] 的最短路径长度为 dist[cur][state] + 1 ，
//      并将 (next, next_state) 放入队列 q 中。
//		
//
//		时间复杂度： O(n * 2 ^ n)
//          1. 需要遍历全部 O(n) 结点
//          2. 最差情况下，每个结点都需要遍历全部 O(2 ^ n) 个状态
//		空间复杂度： O(n * 2 ^ n)
//          1. 状态压缩的距离数组 dist 的空间复杂度为 O(n * 2 ^ n)
//          2. 队列 q 的空间复杂度为 O(n * 2 ^ n)

use std::collections::VecDeque;

impl Solution {
    pub fn shortest_path_length(graph: Vec<Vec<i32>>) -> i32 {
        // 如果只有 1 个结点，则直接返回 0
        if graph.len() == 1 {
            return 0;
        }

        // 压缩后状态的最大值，表示所有点都已遍历过
        let mx = (1 << graph.len()) - 1;
        // 初始化状态压缩的距离数组 dist ，
        // 其中 dist[i][j] 表示最后在结点 i ，
        // 已遍历过的结点状态为 j 时的最短路径长度，
        // 初始化都是 i32::MAX
        let mut dist = vec![vec![i32::MAX; mx + 1]; graph.len()];
        // q 用于 BFS ，存储后续可以遍历的结点等信息
        let mut q = VecDeque::new();
        // 初始以每个结点为开始结点
        for i in 0..graph.len() {
            // 初始在结点 i ，已遍历过的结点状态为 1 << j 时，
            // 最短路径长度为 0 ，可以放入队列继续处理
            dist[i][1 << i] = 0;
            q.push_back((i, 1 << i));
        }

        // 图是联通的，所以在 BFS 中第一次遍历完所有结点的时，
        // 就完成了一种合法的最短路径，可以直接返回此时的路径长度
        loop {
            // 获取当前可往外拓展的结点 cur 及其状态 state
            let (cur, state) = q.pop_front().unwrap();
            // 遍历 cur 相邻的所有结点
            for &next in &graph[cur] {
                // 计算此时移动到 next 的对应状态 next_state
                let next_state = state | (1 << next);
                // 如果这种状态目前没有遍历过，则可以进行遍历
                if dist[next as usize][next_state] == i32::MAX {
                    // 如果此时已经遍历完所有的结点，则可以直接返回最短路径长度
                    if next_state == mx {
                        return dist[cur][state] + 1;
                    }

                    // 此时还未遍历完所有结点，需要继续 BFS 处理，
                    // 更新当前状态的最短路径，并放入队列中继续处理
                    dist[next as usize][next_state] = dist[cur][state] + 1;
                    q.push_back((next as usize, next_state));
                }
            }
        }
    }
}
