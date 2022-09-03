// 链接：https://leetcode.com/problems/pacific-atlantic-water-flow/
// 题意：有一个 m * n 的矩形岛屿，其左边界和上边界与太平洋相连，
//      右边界和下边界与大西洋相连。
//
//      给定一个 m * n 的整型数组 heights ，
//      其中 heights[r][c] 表示岛屿 (r, c) 处高于海平面的高度。
//
//      如果相邻单元格的高度 小于或等于 当前单元格的高度，
//      雨水可以直接向四个方向流向相邻单元格。
//
//      雨水可以从海洋附近的任何单元格流入海洋。
//
//      求满足以下要求的单元格坐标列表？
//      落在该单元格的雨水既能流入太平洋，也能流入大西洋。


// 数据限制：
//  m == heights.length
//  n == heights[r].length
//  1 <= m, n <= 200
//  0 <= heights[r][c] <= 10 ^ 5


// 输入： heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
// 输出： [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
// 解释： 以下 7 个带括号的单元格满足题意。
//        1  2 2 3 (5)
//        3  2 3(4)(4)
//        2  4(5)3  1
//       (6)(7)1 4  5
//       (5) 1 1 2  4

// 输入： heights = [[2,1],[1,2]]
// 输出： [[0,0],[0,1],[1,0],[1,1]]
// 解释： 全部单元格都满足题意。


// 思路： BFS + 状态压缩
//
//      我们可以按照水流的方向倒着处理，这样矩形的四个边界就是起点。
//      然后我们就能通过 BFS 从这些起点找到所有能流入海洋的单元格。
//
//      但我们需要找到能同时流入太平洋和大西洋的单元格，
//      所以我们不能仅用类似 visited 这样布尔类型的数组记录状态。
//
//      可以用整型数组 status 来记录状态，
//      其中 status[r][c] 的最低位表示是否能流入太平洋 (PO = 1) ，
//      次低位表示是否能流入大西洋 (AO = 2)。
//
//      同时，在 BFS 时，也要在队列 q 中记录每个单元格的状态，
//      用于将该单元格最终会流入哪个海洋这个信息传递下去。
//
//      这样我们就知道了 q 的每个元素都要维护单元格的坐标 (r, c) ，
//      以及该单元格的雨水最终会流入的海洋 dest 。
//
//      在 BFS 时，我们从 q 中取出当前单元格的相关信息 r, c 和 dest 。
//      如果其相邻单元格 (rr, cc) 合法，
//      且还不能流入对应的海洋 dest （即 status[rr][cc] & dest == 0 ），
//      且高度不小于当前单元格高度（即 heights[rr][cc] >= heights[r][c] ），
//      则标记其能流入海洋 dest （即 status[rr][cc] |= dest ），
//      并将相关信息放入队列 q 中。
//
//      最终再遍历状态数组 status ，收集状态最低位和次低位都为 1 的单元格坐标即可。
//
//
//		时间复杂度： O(mn)
//          1. 需要遍历全部 O(mn) 个单元格
//		空间复杂度： O(mn)
//          1. 需要维护全部 O(mn) 个单元格的状态
//          2. 需要维护一个 O(mn) 的队列 q


use std::collections::VecDeque;


// 每个方向的位置改变量
//  0: 向上走 1 步
//  1: 向右走 1 步
//  2: 向下走 1 步
//  3: 向左走 1 步
static DIRS: [(i32, i32); 4] = [(-1, 0), (0, 1), (1, 0), (0, -1)];
// 用状态最低位表示能流入太平洋
const PO: i32 = 1;
// 用状态次低位表示能流入大西洋
const AO: i32 = 2;
// 状态最低位和次低位同时为 1 ，表示能同时流入太平洋和大西洋
const ANS: i32 = PO | AO;


impl Solution {
    pub fn pacific_atlantic(heights: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let (m, n) = (heights.len(), heights[0].len());
        // 初始所有单元格的状态都是 0 ，暂不确定能流入哪
        let mut status = vec![vec![0; n]; m];
        // q 用于 BFS ，维护能流入太平洋/大西洋的单元格列表
        let mut q = VecDeque::new();
        // 初始左边界能流入太平洋，右边界能流入大西洋
        for r in 0..m {
            status[r][0] |= PO;
            q.push_back((r, 0, PO));

            status[r][n - 1] |= AO;
            q.push_back((r, n - 1, AO));
        }
        // 初始上边界能流入太平洋，下边界能流入大西洋
        for c in 0..n {
            status[0][c] |= PO;
            q.push_back((0, c, PO));

            status[m - 1][c] |= AO;
            q.push_back((m - 1, c, AO));
        }

        // 还有单元格未处理，则继续 BFS
        while let Some((r, c, dest)) = q.pop_front() {
            // 遍历计算当前单元格相邻的四个单元格
            for &(dr, dc) in DIRS.iter() {
                let rr = (r as i32 + dr) as usize;
                let cc = (c as i32 + dc) as usize;
                // 如果相邻单元格合法，
                //（注意： r, c 的类型为 usize ，所以会下溢，无需判断是否小于 0 ）
                // 且还不能流入对应的海洋 dest ，
                // 且高度不小于当前单元格高度，
                // 则标记其能流入海洋 dest ，并放入队列 q 中
                if rr < m &&
                    cc < n &&
                    status[rr][cc] & dest == 0 && 
                    heights[rr][cc] >= heights[r][c] {
                        status[rr][cc] |= dest;
                        q.push_back((rr, cc, dest));
                    }
            }
        }

        // 遍历所有单元格，收集能同时流入太平洋和大西洋的单元格坐标
        let mut ans = Vec::new();
        for r in 0..m {
            for c in 0..n {
                if status[r][c] == ANS {
                    ans.push(vec![r as i32, c as i32]);
                }
            }
        }

        ans
    }
}