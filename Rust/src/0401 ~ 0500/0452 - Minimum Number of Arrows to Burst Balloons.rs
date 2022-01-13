// 链接：https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/
// 题意：给定一些气球的水平方向所占范围 points ， 
//      points[i] = [x_start, x_end] 表示气球 i 占据 [x_start, x_end] 内，
//      现在可以在 x 上射箭，只要满足 x_start <= x <= x_end 的气球都将被射爆，
//      求至少要射多少次箭，才能将所有的气球射爆？

// 数据限制：
//  1 <= points.length <= 10 ^ 5
//  points[i].length == 2
//  -(2 ^ 31) <= x_start < x_end <= 2 ^ 31 - 1

// 输入： points = [[10,16],[2,8],[1,6],[7,12]]
// 输出： 2
// 解释： - 在 x = 6 处射一箭，气球 [2,8] 和 [1,6] 被射爆
//       - 在 x = 11 处射一箭，气球 [10,16] 和 [7,12] 被射爆

// 输入： points = [[1,2],[3,4],[5,6],[7,8]]
// 输出： 4
// 解释： 每个气球都必须被单独射一箭

// 输入： points = [[1,2],[2,3],[3,4],[4,5]]
// 输出： 2
// 解释： - 在 x = 2 处射一箭，气球 [1,2] 和 [2,3] 被射爆
//       - 在 x = 4 处射一箭，气球 [3,4] 和 [4,5] 被射爆


// 思路： 贪心
//
//      我们可以对 points 按照 x_start 升序排序，然后遍历处理。
//      同时维护两个值：
//          arrow_shots: 当前需要射箭的数量。
//                      初始化为 1 ，因为最后一部分完好的气球不会在遍历中被射爆
//          min_end: 当前遍历过的完好的气球的 x_end 的最小值
//                      初始化为 i32:MAX ，因为最开始没有气球，可以认为无限大
//
//      每次遍历到当前 point = [x_start, x_end] 时，
//          1. x_start > min_end: 则说明至少需要两箭才能将当前气球和前一部分完好的所有气球射爆，
//                  那我们直接在 min_end 射一箭，然后将前一部分完好的所有气球射爆， arrow_shots += 1
//                  而当前气球可能会和后面的气球一起被射爆，咱不处理，设置 min_end = x_end
//          2. x_start <= min_end: 则一箭即可将包括当前气球的所有完好气球射爆，
//                  因为前面所有的气球的 [xstart, xend] 均满足 xstart <= min_end <= xend 。
//                  此时只用更新 min_end = min(min_end, x_end) 即可
//
//
//      时间复杂度： O(nlogn)
//      空间复杂度： O(1)

impl Solution {
    pub fn find_min_arrow_shots(mut points: Vec<Vec<i32>>) -> i32 {
        // 按照 x_start 升序排序即可， x_end 不影响结果，
        // 因为当 x_start 不变时，后续所有气球的 [xstart, xend] 均满足 x_start = xstart <= xend ，
        // 即使 min_end 更新了，也有 x_start = xstart <= min_end ，所以这些气球还是可以被一起射爆的。
        // 只有当 x_start 增加时，才有可能出现 x_start > min_end ，这是才需要至少两箭射爆这些气球。
        points.sort_by_key(|point| point[0]);
        // 转成迭代器
        points.iter()
            // 将思路中的 arrow_shots 和 min_end 运用流式处理的 fold 进行运算。
            .fold((1, i32::MAX), |(arrow_shots, min_end), point| {
                if point[0] > min_end {
                    // x_start > min_end: 则说明至少需要两箭才能将当前气球和前一部分完好的所有气球射爆，
                    // 那我们直接在 min_end 射一箭，然后将前一部分完好的所有气球射爆， arrow_shots += 1
                    // 而当前气球可能会和后面的气球一起被射爆，咱不处理，设置 min_end = x_end
                    (arrow_shots + 1, point[1])
                } else {
                    // x_start <= min_end: 则一箭即可将包括当前气球的所有完好气球射爆，
                    // 因为前面所有的气球的 [xstart, xend] 均满足 xstart <= min_end <= xend 。
                    // 此时只用更新 min_end = min(min_end, x_end) 即可
                    (arrow_shots, min_end.min(point[1]))
                }
            })
            // 二元组第一个值就是所需的最小箭数
            .0
    }
}
