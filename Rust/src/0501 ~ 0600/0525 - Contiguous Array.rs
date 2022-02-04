// 链接：https://leetcode.com/problems/contiguous-array/
// 题意：给定一个二进制数组 nums ，
//      返回含有相同数量的 0 和 1 的最长连续子数组的长度。


// 数据限制：
//  1 <= nums.length <= 10 ^ 5
//  nums[i] 是 0 或 1


// 输入：nums = [0,1]
// 输出：2
// 解释：[0, 1] 是具有相同数量 0 和 1 的最长连续子数组。

// 输入：nums = [0,1,0]
// 输出：2
// 解释：[0, 1] (或 [1, 0]) 是具有相同数量 0 和 1 的最长连续子数组。


// 思路：前缀和 + Map
//
//      设 diff[i] 表示 nums[..=i] 中 1 比 0 多的数量，
//      那么我们可以在 O(1) 内计算出 nums[j..=i] 中 1 比 0 多的数量为：
//          diff[i] - diff[j - 1] 
//
//      题目需要我们找到 1 与 0 数量相等的子数组，则有 diff[i] == diff[j - 1] ，
//      又想要这个子数组最长，那么 j - 1 应该尽可能小，
//      即 j - 1 应该是 diff[j - 1] 第一次出现的下标。
//
//      我们可以使用一个 map 来维护不同的 diff 第一次出现的下标，
//      这样当我们计算到 nums[..=i] 的 diff 时，只需要查询 diff 是否在 map 中：
//          1. diff 在 map 中：则 nums[map[diff]..=i] 就是就是以 i 为结尾的
//              满足题意的最长子数组
//          2. diff 不在 map 中：则 diff 是第一次出现，记录下标 i 即可
//
//
//      时间复杂度：O(n)
//      空间复杂度：O(n)

use std::collections::HashMap;

impl Solution {
    pub fn find_max_length(nums: Vec<i32>) -> i32 {
        let mut ans = 0;
        // 维护当前 nums[..=i] 中 1 比 0 多的数量
        let mut diff = 0;
        // 维护每个前缀和 diff 第一次出现的下标 index
        // 初始化 1 和 0 相等时的第一次下标为 -1 ，方便后续计算
        let mut diff_to_first_index = HashMap::<i32, i32>::from([(0, -1)]);
        nums
            // 转成迭代器
            .iter()
            // 加上对应的下标
            .enumerate()
            .for_each(
                |(i, num)| {
                    // 将当前数纳入范围考虑
                    diff += match num {
                        // 如果当前数是 0 ，则 diff 需要 -1
                        0 => -1,
                        // 如果当前数是 1 ，则 diff 需要 +1
                        1 => 1,
                        // 不存在这种情况
                        _ => unreachable!(),
                    };
                    // 获取 diff 第一次出现的下标
                    match diff_to_first_index.get(&diff) {
                        // 如果存在，则 nums[index + 1..=i] 就是以 i 为结尾的
                        // 满足题意的最长子数组，可以更新 ans 的最大值
                        Some(&index) => ans = ans.max(i as i32 - index),
                        // 如果不存在，则 diff 是第一次出现，
                        // 记录下标 i 即可
                        None => {
                            diff_to_first_index.insert(diff, i as i32);
                        },
                    }
                },
            );
    
        ans
    }
}
