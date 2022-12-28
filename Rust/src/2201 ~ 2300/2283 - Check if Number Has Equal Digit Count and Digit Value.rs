// 链接：https://leetcode.com/problems/check-if-number-has-equal-digit-count-and-digit-value/
// 题意：给定含有 n 个数位的字符串 num ，
//      判断对于每一个下标 i ，数位 i 的出现次数是否都是 num[i] ？


// 数据限制：
//  n == num.length
//  1 <= n <= 10
//  num 仅有数位组成


// 输入： num = "1210"
// 输出： true
// 解释： num[0] = '1', 数位 0 在 num 中出现了 1 次；
//       num[1] = '2', 数位 1 在 num 中出现了 2 次；
//       num[2] = '1', 数位 2 在 num 中出现了 1 次；
//       num[3] = '0', 数位 3 在 num 中出现了 0 次；
//       对于每一个下标 i ， num[i] 满足题意，所以返回 true

// 输入： num = "030"
// 输出： false
// 解释： num[0] = '0', 数位 0 应该出现 0 次，但实际上出现了 2 次；
//       num[1] = '3', 数位 1 应该出现 3 次，但实际上没有出现过；
//       num[2] = '0', 数位 2 在 num 中出现了 0 次；
//       对于下标 0 和 1 ， num[i] 不满足题意，所以返回 false


// 思路： Map
//
//      用一个名为 digit_to_cnt 的 map 维护每个数位的出现次数，
//      然后再判断对于每一个下标 i, digit_to_cnt[i] == num[i] 是否都成立。
//
//
//      时间复杂度：O(n)
//          1. 需要 num 中全部 O(n) 个数位两次
//      空间复杂度：O(n)
//          1. 需要维护 digit_to_cnt 中全部 O(n) 数位的出现次数


use std::collections::HashMap;
use std::ops::AddAssign;


impl Solution {
    pub fn digit_count(num: String) -> bool {
        let num = num.as_bytes();
        // 统计 num 中每个数位的出现次数
        let mut digit_to_cnt = HashMap::new();
        for &ch in num {
            digit_to_cnt.entry(ch - b'0').or_insert(0).add_assign(1);
        }

        num
            .iter()
            .enumerate()
            // 判断对于每一个下标 i, digit_to_cnt[i] == num[i] 是否都成立
            .all(|(i, &ch)| *digit_to_cnt.get(&(i as u8)).unwrap_or(&0) == ch - b'0')
    }
}
