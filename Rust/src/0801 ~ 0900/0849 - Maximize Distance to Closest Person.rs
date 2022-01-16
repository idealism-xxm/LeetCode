// 链接：https://leetcode.com/problems/maximize-distance-to-closest-person/
// 题意：给定一个 01 数组表示的座位列表， 0 表示没有人坐， 1 表示有人坐，
//      现在你想坐在一个与其他人的最小距离最大的位置，返回这个最小距离的最大值？

// 数据限制：
//  2 <= seats.length <= 2 * 10 ^ 4
//  seats[i] 是 0 或 1
//  至少有一个空座位和一个有人坐的座位

// 输入： seats = [1,0,0,0,1,0,1]
// 输出： 2
// 解释： 坐在 seats[2] 时，最小距离为 2
//       坐在其他空位时，最小距离为 1

// 输入： seats = [1,0,0,0]
// 输出： 3
// 解释： 坐在 seats[3] 时，最小距离为 3

// 输入： seats = [0,1]
// 输出： 1


// 思路： 模拟
//
//      我们维护两个值 ans 和 pre ，
//          ans 表示当前最小距离的最大值，
//          pre 表示上一个有人坐的位置的下标。
//
//      首先找到第一个有人坐的位置的下标 pre ，并令 ans = pre :
//          1. pre == 0: 前面不可再坐人，此时距离为 0 ，
//          2. pre != 0: 可以坐在 0 处，此时距离为 pre ，
//          综上：初始距离 ans = pre
//
//      然后遍历后续有人坐的位置的下标 cur ，此时更新 ans = max(ans, (cur - pre) >> 1) ：
//          1. cur == pre: 中间不可再坐人，此时距离为 0 ，
//          2. cur != pre: 可以坐在 (cur + pre) >> 1 处，此时距离为 (cur - pre) >> 1
//          综上：这一段的距离为 (cur - pre) >> 1
//
//      最后考虑最后一个座位的情况，此时更新 ans = max(ans, len(seats) - 1 - pre) ：
//          1. pre == seats.len() - 1 时，最后不可再坐人，这一段的距离为 0 ，无需更新
//          2. pre != seats.len() - 1 时，最后可以坐在 seats.len() - 1 处，
//              这一段的距离为 seats.len() - 1 - pre ，有可能更新最小距离的最大值
//          综上：这一段的距离为 seats.len() - 1 - pre
//
//
//      时间复杂度： O(n)
//      空间复杂度： O(1)


impl Solution {
    pub fn max_dist_to_closest(seats: Vec<i32>) -> i32 {
        // 转成迭代器
        let mut iter = seats.iter()
            // 附上每个座位的下标
            .enumerate()
            // 只处理有人坐的情况
            .filter(|&(_, &seat)| seat == 1);
        // 因为至少有一个空座位和一个有人坐的座位，
        // 所以可以直接取第一个有人坐的座位的下标
        let pre = iter.next().unwrap().0;
        // 将思路中的 ans 和 pre 运用 fold 进行运算
        let (ans, pre) = iter.fold(
            // 1. 当 pre == 0 时，前面不可再坐人，此时距离为 0 ，
            // 2. 当 pre != 0 时，可以坐在 0 处，此时距离为 pre ，
            // 综上：初始距离 ans = pre
            (pre, pre), 
            //  每次选择坐在两人中间位置，
            //    更新最小距离的最大值 ans = ans.max((cur - pre) >> 1) ，
            //      1. cur == pre: 中间不可再坐人，此时距离为 0 ，
            //      2. cur != pre: 可以坐在 (cur + pre) >> 1 处，此时距离为 (cur - pre) >> 1
            //      综上：这一段的距离为 (cur - pre) >> 1
            //  同时更新上一次有人坐的下标 pre = cur
            |(ans, pre), (cur, _)| (ans.max((cur - pre) >> 1), cur)
        );
        // 最后二元组的第一个就是最小距离的最大值，
        // 同时还要考虑最后一个座位的情况：
        //  1. pre == seats.len() - 1 时，最后不可再坐人，这一段的距离为 0 ，无需更新
        //  2. pre != seats.len() - 1 时，最后可以坐在 seats.len() - 1 处，
        //      这一段的距离为 seats.len() - 1 - pre ，有可能更新最小距离的最大值
        ans.max(seats.len() - 1 - pre) as i32
    }
}
