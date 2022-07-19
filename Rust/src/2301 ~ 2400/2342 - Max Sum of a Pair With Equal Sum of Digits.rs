// 链接：https://leetcode.com/problems/max-sum-of-a-pair-with-equal-sum-of-digits/
// 题意：给定一个正整数数组 nums ，可以选择满足以下条件的两个数对 (i, j) ：
//          1. i != j
//          2. nums[i] 的数位和与 nums[j] 的数位和相等
//
//      求 nums[i] + nums[j] 的最大值？
//      没有满足题意的数对，则返回 -1 。


// 数据限制：
//  1 <= nums.length <= 10 ^ 5
//  1 <= nums[i] <= 10 ^ 9


// 输入： nums = [18,43,36,13,7]
// 输出： 54
// 解释： 满足题意的数对有：
//          (0, 2): 下标对应的数位和都是 9 ，数字和为 18 + 36 = 54
//          (1, 4): 下标对应的数位和都是 7 ，数字和为 43 + 7 = 50
//       综上， nums[i] + nums[j] 的最大值为 54

// 输入： nums = [10,12,19,14]
// 输出： -1
// 解释： 没有满足题意的数对。


// 思路： Map
//
//      我们很容易就能想到维护一个名为 digit_sum_to_nums 的 map ，
//      其中 digit_sum_to_nums[digit_sum] 表示数位和为 digit_sum 的数字列表。
//
//      再用 ans 维护满足题意的数字和的最大值，初始化为 -1 。
//
//      然后枚举所有的数字列表，对其进行降序排序，求最大两个数字的和，并更新 ans 的最大值。
//
//      不过这样在最差情况下的时间复杂度为 O(n * C + nlogn) ，空间复杂度为 O(n) 。
//
//      由于题目只需要相同数字和的所有数中最大的两个数，所以我们无需维护全部的数字列表，
//      在初始化时就维护最大的两个数即可。
//
//      这样后续就无需排序，直接计算数字和更新最大值即可。
//
//      时间复杂度能优化为 O(n * C) ，空间复杂度能优化为 O(C ^ 2) 。
//      
//
//      设数字的最大位数为 C 。
//
//      时间复杂度：O(n * C)
//          1. 需要遍历 nums 中全部 O(n) 个数字，
//              每个数字都需要遍历全部 O(C) 个数位。
//      空间复杂度：O(C ^ 2)
//          1. 需要用 map 维护所有不同数位和的最大的两个数，
//              最差情况下有 O(C ^ 2) 个不同的数位和


use std::collections::HashMap;


impl Solution {
    pub fn maximum_sum(nums: Vec<i32>) -> i32 {
        let mut digit_sum_to_largest_two_nums = HashMap::new();
        for num in nums {
            // 计算 num 的数位和
            let mut digit_sum = 0;
            let mut remain = num;
            while remain > 0 {
                digit_sum += remain % 10;
                remain /= 10;
            }

            // 获取数位和 digit_sum 对应的最大两个数字，如果该数不存在，则为 0
            let mut largest_two_nums = digit_sum_to_largest_two_nums.entry(digit_sum).or_insert((0, 0));
            if num > largest_two_nums.0 {
                // 如果 num 比最大数大，则先更新次大数，再更新最大数为 num
                largest_two_nums.1 = largest_two_nums.0;
                largest_two_nums.0 = num;
            } else if num > largest_two_nums.1 {
                // 如果 num 只比次大数大，则更新次大数为 num
                largest_two_nums.1 = num;
            }
        }

        digit_sum_to_largest_two_nums
            // 我们只关心最大的两个数，不关心对应的数位和
            .values()
            // 过滤掉只有一个数的情况
            .filter(|largest_two_nums| largest_two_nums.1 > 0)
            // 计算数字和
            .map(|largest_two_nums| largest_two_nums.0 + largest_two_nums.1)
            // 求数字和的最大值
            .max()
            // 如果没有满足题意的数对，则返回 -1
            .unwrap_or(-1)
    }
}
