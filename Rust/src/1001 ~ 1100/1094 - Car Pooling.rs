// 链接：https://leetcode.com/problems/car-pooling/
// 题意：有一辆汽车从地点 0 不回头地开到地点 1000 ，
//      这辆汽车初始有 capacity 个空座位。
//      给定一个乘客上下车信息的数组 trips ，
//      其中 trips[i] = [numPassengers_i, from_i, to_i] ，
//      表示有 numPassengers_i 个乘客
//      在 from_i 上车，并在 to_i 下车。
//      现在求这辆汽车是否能将所有人都送到目的地？

// 数据限制：
//  1 <= trips.length <= 1000
//  trips[i].length == 3
//  1 <= numPassengers_i <= 100
//  0 <= from_i < to_i <= 1000
//  1 <= capacity <= 10 ^ 5


// 输入：trips = [[2,1,5],[3,3,7]], capacity = 4
// 输出：false
// 解释：最大矩形面积是 10 = 2 * 5

// 输入：trips = [[2,1,5],[3,3,7]], capacity = 5
// 输出：true
// 解释：有两种可能的最大矩形，他们的面积都是 4 = 2 * 2 = 1 * 4


// 思路：模拟
//
//      因为汽车只能不回头地从 0 开到 1000 ，所以我们可以模拟汽车的这个过程，
//      维护汽车在每一点的空座位数量 capacity 。
//
//      1. 如果当前位置上下乘客后，空座位数量小于零，
//          则说明汽车无法将所有乘客送往目的地，直接返回 false
//      2. 如果所有位置上下乘客后，空座位数量都大于等于零，
//          则说明汽车能将所有乘客送往目的地，最后返回 true
//
//      那么如何维护汽车在每一点空座位的数量呢？
//
//      我们可以发现地点的取值范围是 [0, 1000] ，
//      所以我们可以直接开辟一个长度为 1001 的数组 capacity_diff ，
//      capacity_diff[i] 表示空座位数量在地点 i 的变化量。
//
//      然后将 trips 转化成 capacity_diff 表示即可，
//      即对每一个 trip = [num, from, to] 进行如下处理：
//          1. capacity_diff[from] -= num: num 人在 from 处上车，空座位数量减少 num
//          2. capacity_diff[to] += num: nums 人在 to 处下次，空座位数量增加 num
//
//
//      【进阶】如果地点的取值范围是 [0, 10 ^ 9] 或者更大时，应该如何处理呢呢？
//
//      这时候一般有两种方式：
//          1. 收集所有取值，对其排序后离散化，然后用本题的方式处理
//          2. 将每个 trip 都拆分成两个元组 (from, -num) 和 (to, num) ，
//              将这些元素排序后遍历处理
//
//		设 n 表示 trips 的长度， m 表示下车地点的最大值，这里直接取作 1001 
//
//		时间复杂度： O(n + m)
//      空间复杂度： O(n + m)

impl Solution {
    pub fn car_pooling(trips: Vec<Vec<i32>>, mut capacity: i32) -> bool {
        // capacity_diff[i] 表示汽车空座位在当前位置的变化
        //  负数表示当前位置有人上车，空座位减小
        //  整数表示当前位置有人下车，空座位增加
        let mut capacity_diff = [0; 1001];
        for trip in trips {
            // trip[1] 处有人上车，这里减少 trip[0] 个空座位
            capacity_diff[trip[1] as usize] -= trip[0];
            // trip[2] 处有人下车，这里增加 trip[0] 个空座位
            capacity_diff[trip[2] as usize] += trip[0];
        }

        // 遍历 capacity_diff 
        for diff in capacity_diff {
            // 计算当前位置的空座位数量
            capacity += diff;
            // 如果空座位数量小于 0 ，则不满足题意，直接返回 false
            if capacity < 0 {
                return false;
            }
        }

        // 所有点的空座位都满足题意，返回 true
        true
    }
}
