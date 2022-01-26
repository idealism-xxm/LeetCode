// 链接：https://leetcode.com/problems/valid-mountain-array/
// 题意：给定一个整型数组 arr ，判断 arr 是否是一个合法的山形数组？
//      山形数组必须满足以下两个条件：
//          1. arr.length >= 3
//          2. 存在一个 i (0 < i < arr.length - 1) 满足：
//              arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
//              arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

// 数据限制：
//  1 <= arr.length <= 10 ^ 4
//  0 <= arr[i] <= 10 ^ 4

// 输入： arr = [2,1]
// 输出： false
// 解释： 长度为 2 ，不满足第 1 个条件

// 输入： arr = [3,5,5]
// 输出： false
// 解释： arr[1] == arr[2] ，不满足第 2 个条件

// 输入： arr = [0,3,2,1]
// 输出： true
// 解释： a[0] < a[1] > a[2] > a[3]


// 思路：模拟
//
//      按照题意，山形数组必须先严格单调递增，然后严格单调递减，
//      我们可以直接一次遍历就能判断是否满足这两个条件。
//
//      维护一个 status 表示当前位置的状态，总共有三种状态：
//          Pending: 未开始，即还未开始遍历
//                  （注意不能使用 Ascend ，因为 arr[1] 必须大于 arr[0] ，
//                  使用 Pending 才能判断出这种情况）
//          Ascend:  已开始上升，即当前还处于上升状态
//          Descend: 已开始下降，即当前已经从上升状态变为下降状态
//
//      然后我们开始从 i = 1 开始遍历 arr 数组，每次进行如下判断即可：
//          1. arr[i - 1] < arr[i] && status == Pending: 如果当前是上升趋势，
//              只要现在不是下降状态，就可以维持上升状态
//          2. arr[i - 1] > arr[i] && status == Descend: 如果当前是下降趋势，
//              只有现在不是未开始状态，才可以维持下降状态
//          3. 其他情况：不满足第 2 个条件，直接返回 false
//
//      最后只有 status 是 Descend 时，才返回 true ，
//      即只有数组先上升，再下降，才是一个合法的山形数组
//
//      时间复杂度： O(n)
//      空间复杂度： O(1)

// 数组当前位置的状态
#[derive(Debug, PartialEq, Eq)]
enum Status {
    // 未开始
    Pending,
    // 已开始上升
    Ascend,
    // 已开始下降
    Descend,
}

impl Solution {
    pub fn valid_mountain_array(arr: Vec<i32>) -> bool {
        // 用大小为 2 的滑动窗口遍历
        arr.windows(2)
            // 使用 try_fold 尝试积累最后的状态
            .try_fold(
                // 最开始还未遍历，默认未开始
                // （注意不能使用 Ascend ，因为 arr[1] 必须大于 arr[0] ，
                //  使用 Pending 才能判断出这种情况）
                Status::Pending, 
                // 使用函数根据当前滑动窗口 window 内的两个值更新 status 的状态
                |status, window| {
                    if window[0] < window[1] && status != Status::Descend {
                        // 如果当前是上升趋势，只要现在不是下降状态，就可以维持上升状态
                        Some(Status::Ascend)
                    } else if window[0] > window[1] && status != Status::Pending {
                        // 如果当前是下降趋势，只有现在不是未开始状态，才可以维持下降状态
                        Some(Status::Descend)
                    } else {
                        // 其他情况都不满足第 2 个条件，直接停止积累
                        None
                    }
                }
            )
            // 将最后的状态映射成是否为山形数组的 bool 值
            .map_or(false, |status| status == Status::Descend)
    }
}
