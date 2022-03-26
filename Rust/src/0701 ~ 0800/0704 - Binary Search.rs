// 链接：https://leetcode.com/problems/binary-search/
// 题意：给定一个升序排序的整型数组 nums ，
//      返回 target 所在的下标，如果不存在则返回 -1 。


// 数据限制：
//  1 <= nums.length <= 10 ^ 4
//  -(10 ^ 4) < nums[i], target < 10 ^ 4
//  nums 中的所有数字各不相同
//  nums 是升序排序的


// 输入： nums = [-1,0,3,5,9,12], target = 9
// 输出： 4
// 解释： 9 在 nums 中，下标为 4

// 输入： nums = [-1,0,3,5,9,12], target = 2
// 输出： -1
// 解释： 2 不在 nums 中，返回 -1


// 思路： 二分
//
//      因为 nums 已经是升序排序的，且数字各不相同，
//      那么可以直接使用二分查找。
//
//      我们定义左闭右闭的二分区间 [l, r] ，
//      其中 l 为左边界，初始化为 0 ， 
//      r 为右边界，初始化为 nums.length - 1 。
//
//      那么只要这个区间 [l, r] 不为空，我们就可以继续循环处理，
//      循环中先找到区间中点下标 m = (l + r) / 2 ，
//      然后根据 nums[m] 和 target 的大小关系进行处理：
//          1. nums[m] == target ，那么直接返回 m 。
//          2. nums[m] < target ，那么区间 [l, m] 内的数都小于 target ，
//              如果 target 存在的话，只可能在右边区间 [m + 1, r] 中，
//              此时需要将二分区间变为 [m + 1, r]
//          3. nums[m] > target ，那么区间 [m, r] 内的数都大于 target ，
//              如果 target 存在的话，只可能在左边区间 [l, m - 1] 中，
//              此时需要将二分区间变为 [l, m - 1]
//
//      如果最后结束循环还没找到 target ，
//      那么说明 target 不在 nums 中，返回 -1 。
//
//
//      时间复杂度：O(logn)
//          1. 需要对 nums 进行二分，时间复杂度为 O(logn)
//      空间复杂度：O(1)
//          1. 只需要常数个额外变量


impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> i32 {
        // 二分区间的左边界，初始化为 0
        let mut l = 0;
        // 二分区间的右边界，初始化为 nums.len() - 1
        let mut r = nums.len() as i32 - 1;
        // 当区间不为空时，继续二分
        // （注意这里取等号是因为我们的区间是左闭右闭区间）
        while l <= r {
            // 计算区间中点下标
            let m = (l + r) >> 1;
            if nums[m as usize] == target {
                // 如果区间中点的值恰好是 target ，
                // 则直接返回对应下标 m
                return m;
            } else if nums[m as usize] < target {
                // 如果区间中点的值小于 target ，
                // 则 target 如果存在的话，
                // 必定在右边区间 [m + 1, r] 中
                l = m + 1;
            } else {
                // 如果区间中点的值大于 target ，
                // 则 target 如果存在的话，
                // 必定在左边区间 [l, m - 1] 中
                r = m - 1;
            }
        }

        // 二分时没找到 target ，那么 target 必定不存在
        -1
    }
}
