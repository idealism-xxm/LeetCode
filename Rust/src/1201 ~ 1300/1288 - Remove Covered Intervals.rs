// 链接：https://leetcode.com/problems/remove-covered-intervals/
// 题意：给定一个数组 intervals ，
//      其中 intervals[i] = [l_i, r_i] ，表示区间 [l_i, r_i) 。
//      现在移除所有被另一个区间覆盖的区间，返回最后剩余的区间数。
//
//      区间 [a, b) 被 区间 [c, d) 覆盖，当且仅当 a >= c && b <= d 。


// 数据限制：
//  1 <= intervals.length <= 1000
//  intervals[i].length == 2
//  0 <= l_i <= r_i <= 10 ^ 5
//  所有的区间各不相同


// 输入：intervals = [[1,4],[3,6],[2,8]]
// 输出：2
// 解释：区间 [3,6] 被区间 [2,8] 覆盖，所以它将被移除

// 输入：intervals = [[1,4],[2,3]]
// 输出：1
// 解释：区间 [2,3] 被区间 [1,4] 覆盖，所以它将被移除


// 思路：排序 + 贪心
//
//		被覆盖的区间要同时满足两个条件 a >= c && b <= d ，
//      像这种题一般都需要通过某种方式，让其中一个条件一直满足，
//      然后就只用判断另一个条件是否满足即可。
//
//      比如昨天的 LeetCode 1675 这题，小满就没有想到消除其中一个操作，
//      所以看了题解后才恍然大悟。
//
//      本题我们可以对 intervals 按照左边界 l 升序排序，
//      那么我们在遍历时， a >= c 这个条件肯定一直满足。
//
//      此时我们只要贪心地维护右边界的最大值 max_right 就行了，
//      这样只要当前区间右边界 r <= max_right ，这个区间肯定被覆盖了。
//
//      需要注意的是，左边界 l 相同时，右边界小的区间会被右边界大的区间覆盖，
//      例如： [1, 3], [1, 4] 这两个区间，前者就会被后者覆盖。
//
//      所以为了统计到这种情况，我们在排序时发现两个区间左边界 l 相等时，
//      再按照右边界 r 降序排序，那么范围更大的区间会先被处理。
//
//
//      时间复杂度：O(nlogn) 
//          1. 排序时间复杂度时 O(nlogn)
//      空间复杂度：O(1)
//          1. 只用了常数个额外变量，所以空间复杂度为 O(1)

impl Solution {
    pub fn remove_covered_intervals(mut intervals: Vec<Vec<i32>>) -> i32 {
        // intervals 按照 l 升序排序， l 相同时，按照 r 降序排序
        intervals.sort_by(|a, b| {
            if a[0] == b[0] {
                // 如果 l 相同，按照 r 降序排序
                b[1].cmp(&a[1])
            } else {
                // 如果 l 不同，按照 l 升序排序
                a[0].cmp(&b[0])
            }
        });

        // covered_cnt 维护需要被移除的区间个数
        let mut covered_cnt = 0;
        // max_right 维护当前所以区间中，最大的右边界 r
        let mut max_right = 0;
        // 按顺序遍历所有区间
        for interval in intervals.iter() {
            if interval[1] <= max_right {
                // 如果当前区间 interval[1] <= max_right ，
                // 则说明其必定被一个区间覆盖
                covered_cnt += 1;
            } else {
                // 如果当前区间 interval[1] > max_right ，
                // 则说明其未被任何区间覆盖，
                // 同时需要更新 max_right = interval[1]
                max_right = interval[1];
            }
        }

        // 最后剩余的区间，就是所有区间数 减去 被覆盖的区间数
        intervals.len() as i32 - covered_cnt
    }
}
