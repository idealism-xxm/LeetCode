// 链接：https://leetcode.com/problems/4sum-ii/
// 题意：给定四个整数数组 nums1, nums2, nums3 和 nums4 ，
//      数组长度都是 n ，返回满足以下条件的元组 (i, j, k, l) 数：
//          1. 0 <= i, j, k, l < n
//          2. nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0


// 数据限制：
//  n == nums1.length
//  n == nums2.length
//  n == nums3.length
//  n == nums4.length
//  1 <= n <= 200
//  -(2 ^ 28) <= nums1[i], nums2[i], nums3[i], nums4[i] <= 2 ^ 28


// 输入：nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
// 输出：2
// 解释：(0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0
//      (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0

// 输入：nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
// 输出：1


// 思路：Map
//
//      先用 map 统计 nums1[i] + nums2[j] 的出现次数到 num_to_cnt 中，
//      然后再枚举 nums3[l] 和 nums4[k] ，
//      这样 nums1[i] + nums2[j] 就必须等于 num = -(nums3[l] + nums4[k]) ，
//      所以 num_to_cnt[num] 就是选定 nums3[l] 和 nums4[k] 时的满足题意的四元组数
//
//
//      时间复杂度：O(n ^ 2)
//      空间复杂度：O(n ^ 2)

use std::collections::HashMap;

impl Solution {
    pub fn four_sum_count(nums1: Vec<i32>, nums2: Vec<i32>, nums3: Vec<i32>, nums4: Vec<i32>) -> i32 {
        // 统计 nums1[i] + nums2[j] 的出现次数
        let mut num_to_cnt = HashMap::<i32, i32>::new();
        for num1 in &nums1 {
            for num2 in &nums2 {
                *num_to_cnt.entry(num1 + num2).or_insert(0) += 1;
            }
        }

        // 统计满足题意的四元组数
        let mut ans = 0;
        for num3 in &nums3 {
            for num4 in &nums4 {
                // 选择 num3 和 num4 时，则 -(nums1[i] + nums2[j]) 出现多少次，
                // 对 ans 的贡献就有多少 
                ans += num_to_cnt.get(&-(num3 + num4)).unwrap_or(&0);
            }
        }

        ans
    }
}
