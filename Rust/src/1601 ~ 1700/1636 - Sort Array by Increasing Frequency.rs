// 链接：https://leetcode.com/problems/sort-array-by-increasing-frequency/
// 题意：给定一个整型数组 nums ，对其按照数字出现次数升序排序，
//      出现次数相同时，按照数字将序排序。


// 数据限制：
//  1 <= nums.length <= 100
//  -100 <= nums[i] <= 100


// 输入： nums = [1,1,2,2,2,3]
// 输出： [3,1,1,2,2,2]
// 解释： 3 出现一次， 1 出现两次， 2 出现三次。

// 输入： nums = [2,3,1,3,2]
// 输出： [1,3,3,2,2]
// 解释： 1 出现一次， 2 和 3 各出现两次。
//       3 比 2 大，所以 3 应该在 2 前面。

// 输入： nums = [-1,1,-6,4,5,-6,1,4,1]
// 输出： [5,-1,4,4,-6,-6,1,1,1]


// 思路： Map + 排序
//
//      先用一个 map 统计 nums 中每个数字的出现次数。
//
//      然后对 nums 中的数字按照出现次数升序排序，
//      出现次数相同时，按数字降序排序。
//
//
//      设不同数字的个数为 C 。
//
//      时间复杂度： O(nlogn)
//          1. 需要遍历 nums 中全部 O(n) 个数字
//          2. 需要对 nums 中全部 O(n) 个数字排序，时间复杂度为 O(nlogn)
//      空间复杂度： O(C)
//          1. 需要维护全部 O(C) 个不同数字的出现次数
//          2. 复用 nums 排序，只需要 O(1) 的额外空间


use std::cmp::Reverse;
use std::collections::HashMap;
use std::ops::AddAssign;


impl Solution {
    pub fn frequency_sort(mut nums: Vec<i32>) -> Vec<i32> {
        // num_to_cnt[ch] 表示 nums 中数字的出现次数
        let mut num_to_cnt = HashMap::new();
        for &num in &nums {
            num_to_cnt.entry(num).or_insert(0).add_assign(1);
        }

        // 对 nums 中的数字按照出现次数升序排序，
        // 出现次数相同时，按数字降序排序。
        nums.sort_by_key(|&num| (num_to_cnt[&num], Reverse(num)));

        nums
    }
}
