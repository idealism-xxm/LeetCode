// 链接：https://leetcode.com/problems/wiggle-subsequence/
// 题意：给定一个整数数组  nums ，求最长的摆动子序列的长度？
//
//      进阶：使用时间复杂度为 O(n) 的算法求解。
//
//      如果连续数字之间的差严格地在正负之间交替，那么这个序列就是摆动序列。
//      第一个差值可以是正数也可以是负数，长度为 1 和 2 的序列也是摆动序列。
//
//      例如： [1, 7, 4, 9, 2, 5] 是一个 摆动序列 ，
//              因为差值 (6, -3, 5, -7, 3)  是正负交替出现的。
//      相反， [1, 4, 7, 2, 5] 和 [1, 7, 4, 5, 5] 不是摆动序列，
//              前者是因为它的前两个差值都是正数，后者是因为它的最后一个差值为零。


// 数据限制：
//  1 <= nums.length <= 1000
//  0 <= nums[i] <= 1000


// 输入： nums = [1,7,4,9,2,5]
// 输出： 6
// 解释： 整个序列就是一个摆动序列，差值 (6, -3, 5, -7, 3) 是正负交替出现的

// 输入： nums = [1,17,5,10,13,15,10,5,16,8]
// 输出： 7
// 解释： 子序列 [1, 17, 10, 13, 10, 16, 8] 是摆动序列，
//       差值 (16, -7, 3, -3, 6, -8) 是正负交替出现的

// 输入： nums = [1,2,3,4,5,6,7,8,9]
// 输出： 2


// 思路： 贪心
//
//      本题其实就是找波峰和波谷的总数（第一个数和最后一个数可以被认为是波峰和波谷），
//      选择其他数得到的摆动子序列长度不会更长。
//
//      假设通过选择波峰和波谷得到的摆动子序列为 [nums[i], nums[j], nums[k]] ，
//      其中 i < j < k 且 nums[i] < nums[j] && nums[j] > nums[k] 。
//
//      则 nums[i] 和 nums[k] 为波谷， nums[j] 为波峰。
//
//      如果第二个数不选择波峰 nums[j] ，而选择 (i, j) 内的其他数 nums[l] ，
//      那么必定有 nums[l] <= nums[j] 。
//
//      则 nums[i] < nums[l] && nums[l] > nums[k] 不一定成立。
//
//      又因为 nums[i] 和 nums[k] 为波谷，所以上述条件不成立时，
//      找不到更小的数替代，那么摆动子序列必定更短才行。
//
//      我们使用 pre_status 维护数组前一个位置的状态，最开始的状态为待确定。
//
//      根据当前数组中最后两个数的状态进行处理：
//          1. nums[i] < nums[i - 1]: 当前是上升状态，
//              则 pre_status 为下降或待定时，状态会发生改变。
//              此时 nums[i - 1] 是波峰，可以放入结果摆动序列中。
//          2. nums[i] > nums[i - 1]: 当前是下降状态，
//              则 pre_status 为上升或待定时，状态会发生改变。
//              此时 nums[i - 1] 是波谷，可以放入结果摆动序列中。
//          3. nums[i] == nums[i - 1]: 状态不会改变，直接处理下一个数
//
//      注意最后一个数必定在摆动序列中，但并未在循环中统计，所以要返回 ans + 1
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 nums 全部 O(n) 个数字一次
//      空间复杂度：O(1)
//          1. 只需要维护常数个额外变量即可


// 数组当前位置的状态
#[derive(Debug, PartialEq, Eq)]
enum Status {
    // 待确定
    Pending,
    // 上升
    Ascend,
    // 下降
    Descend,
}


impl Solution {
    pub fn wiggle_max_length(nums: Vec<i32>) -> i32 {
        // 初始无任何数字在摆动序列中
        let mut ans = 0;
        // pre_status 表示数组前一个位置的状态，最开始的状态为待确定
        let mut pre_status = Status::Pending;
        for i in 1..nums.len() {
            // 当前数和前一个数相等，则直接处理下一个数
            if nums[i] == nums[i - 1] {
                continue;
            }
            // 计算最后两个数的状态
            let status = if nums[i] < nums[i - 1] {
                Status::Descend
            } else {
                Status::Ascend
            };
            // 以下两种情况说明状态发生了改变：
            //  1. 变为下降：当前是下降 且 前一个是上升或待定
            //  2. 变为上升：当前是上升 切 前一个是下降或待定
            // 此时 nums[i - 1] 是波峰或波谷，可以将其放入到结果摆动序列中
            if status != pre_status {
                pre_status = status;
                ans += 1;
            }
        }
        
        // 最后一个数字必定在摆动序列中
        ans + 1
    }
}
