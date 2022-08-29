// 链接：https://leetcode.com/problems/two-sum/
// 题意：给定一个数组 nums 和一个目标数 target ，
//      求和为 target 的两个数的下标？


// 数据限制：
//  2 <= nums.length <= 10 ^ 4
//  -(10 ^ 9) <= nums[i] <= 10 ^ 9
//  -(10 ^ 9) <= target <= 10 ^ 9
//  只有一个合法的答案


// 输入： nums = [2,7,11,15], target = 9
// 输出： [0,1]
// 解释： nums[0] + nums[1] == 9 ，
//       所以返回 [0, 1]

// 输入： nums = [3,2,4], target = 6
// 输出： [1,2]
// 解释： nums[1] + nums[2] == 6 ，
//       所以返回 [1, 2]

// 输入： nums = [3,3], target = 6
// 输出： [0,1]
// 解释： nums[0] + nums[1] == 6 ，
//       所以返回 [0, 1]


// 思路： Map
//
//      我们可以用 map 维护每个数最后一次出现的下标。
//
//      遍历数组第 i 个数 num ：
//          1. 如果 target - num 在 map 中，则直接返回对应结果即可
//          2. 如果 target - num 不在 map 中，
//              则将 num 对应的下标设置为 i
//
//
//      时间复杂度： O(n)
//          1. 需要遍历 nums 中全部 O(n) 个数组
//      空间复杂度： O(n)
//          1. 需要维护全部不同数字的下标，
//              最差情况下，有 O(n) 个不同的数字


use std::collections::HashMap;


impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        // 记录每个数最后一次出现的下标
        let mut num_to_index = HashMap::with_capacity(nums.len());
        for (i, &num) in nums.iter().enumerate() {
            // 获取需要的数的下标，如果存在，则直接返回
            if let Some(&j) = num_to_index.get(&(target - num)) {
                return vec![i as i32, j as i32];
            }

            // 如果不存在，则更新 num 的下标为 i
            num_to_index.insert(num, i);
        }

        // 题目保证一定有答案，所以不会走到这
        unreachable!();
    }
}
