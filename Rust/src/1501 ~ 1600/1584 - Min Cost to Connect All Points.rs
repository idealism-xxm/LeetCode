// 链接：https://leetcode.com/problems/min-cost-to-connect-all-points/
// 题意：给定一个数组 points ，其中 points[i] = [x_i, y_i] ，
//      表示第 i 个点在二维平面上的位置。
//
//      连接两个点 (x_i, y_i) 和 (x_j, y_j) 的花费为它们间的曼哈顿距离，
//      即 |x_i - x_j| + |y_i - y_j| ，其中 |val| 是 val 的绝对值。
//
//      求连接所有点的最小花费。


// 数据限制：
//  1 <= points.length <= 1000
//  -（10 ^ 6) <= xi, yi <= 10 ^ 6
//  所有的 (x_i, y_i) 都互不相同


// 输入： points = [[0,0],[2,2],[3,10],[5,2],[7,0]]
// 输出： 20

// 输入： points = [[3,12],[-2,5],[-4,1]]
// 输出： 18


// 思路： Prim 算法
//
//      这个题意其实就相当于给我们一个完全图，求这个完全图的最小生成树。
//
//      那么我们直接可以使用求最小生成树的算法求解，
//      共有两种方法： Prim 算法 和 Kruskal 算法。
//
//      其中 Kruskal 算法主要是通过并查集枚举所有边来处理，
//      时间复杂度是 O(ElogV) ，空间复杂度为 O(E)
//      对于本题的完全图来说，
//      时间复杂度为 O((n ^ 2) * logn) ，空间复杂度为 O(n ^ 2)。
//
//      而朴素的 Prim 算法的时间复杂度是 O(V ^ 2) ，空间复杂度为 O(E) ，
//      但对于本题来说， 时间复杂度为 O(n ^ 2) ，
//      边的权重可以动态计算，所以空间复杂度优化为 O(n) 。
//
//      因此本题选用朴素的 Prim 算法加上一点优化进行求解。
//
//      我们维护两个数组 dist 和 in_mst ，
//      其中 dist[i] 表示点 i 到最小生成树的距离，初始化均为无穷大；
//      in_mst[i] 表示点 i 是否在最小生成树中，初始化均为 false 。
//
//      最开始最小生成树为空，但为了放入第一个结点，方便后续处理
//      我们可以将 dist[0] 的值设置为 0 。
//
//      然后我们开始枚举下一个要加入到最小生成树的点 nxt ，
//      直至所有点都在最小生成树中。
//
//      每次循环时先找到 nxt ，即不在最小生成树中 且 距离最小生成树最近的点。
//
//      然后标记其为在最小生成树中，并将最小生成树的权重加上 dist[nxt] 。
//
//      最后更新剩余不在最小生成树的点，到当前最小生成树的距离，
//      因为只加入了 nxt ，所以只需要求 nxt 和 i 的距离，
//      并更新 dist[i] 即可。
//
//
//      时间复杂度：O(n ^ 2)
//          1. 需要将全部 O(n) 个点加入到最小生成树中
//          2. 每次加入点前，都需要遍历全部 O(n) 个点，找到距离最小的点
//          3. 每次加入点后，都需要更新全部 O(n) 个点到最小生成树的距离
//      空间复杂度：O(n)
//          1. 需要存储全部 O(n) 个点到最小生成树的距离
//          2. 需要存储全部 O(n) 个点是否在最小生成树中


impl Solution {
    pub fn min_cost_connect_points(points: Vec<Vec<i32>>) -> i32 {
        // 获取点的个数
        let n = points.len();
        // dist[i] 表示 i 到当前最小生成树的距离，初始化均为无穷大
        let mut dist = vec![i32::MAX; n];
        // 初始化 0 到最小生成树的距离为 0 ，
        // 方便在空的最小生成树中放入第一个点
        dist[0] = 0;
        // in_mst[i] 表示 i 是否在最小生成树中，
        // 初始化均不在最小生成树中
        let mut in_mst = vec![false; n];
        // 最小生成树中的点数，初始没有任何点
        let mut mst_size = 0;
        // 最小生成树的权重，初始为 0
        let mut mst_weight = 0;

        // 当最小生成树不足 n 个点时，
        // 还要继续贪心放入下一个点
        while mst_size < n {
            // 表示下一个可放入的点
            let mut nxt = 0;
            // 表示 nxt 距离当前最小生成树的距离
            let mut nxt_dist = i32::MAX;
            // 找到当前最小生成树外，距离最小生成树最近的点
            for i in 0..n {
                // 如果点 i 不在最小生成树中，
                // 且距离当前最小生成树的距离更小
                if !in_mst[i] && dist[i] < nxt_dist {
                    // 更新 i 为下一个可放入的点
                    nxt = i;
                    // 更新 i 到当前最小生成树的距离
                    nxt_dist = dist[i];
                }
            }

            // 将 nxt 放入最小生成树中
            in_mst[nxt] = true;
            // 最小生成树中的点数加 1
            mst_size += 1;
            // 最小生成树的权重加上 nxt 到当前最小生成树的距离
            mst_weight += nxt_dist;
            
            // 更新所有点到最小生成树的距离
            for i in 0..n {
                // 计算 i 到 nxt 的距离
                let i_dist = i32::abs(points[nxt][0] - points[i][0]) + 
                             i32::abs(points[nxt][1] - points[i][1]);
                // 如果 i 不在最小生成树中，且到 nxt 的距离更小
                if !in_mst[i] && dist[i] > i_dist {
                    // 那么 i 到最小生成树的距离就是 i 到 nxt 的距离
                    dist[i] = i_dist;
                }
            }
        }

        mst_weight
    }
}