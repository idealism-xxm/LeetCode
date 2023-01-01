// 链接：https://leetcode.com/problems/find-the-kth-largest-integer-in-the-array/
// 题意：给定一个字符串表示的数字数组，返回第 k 大的数？


// 数据限制：
//  1 <= k <= nums.length <= 10 ^ 4
//  1 <= nums[i].length <= 100
//  nums[i] 仅由数字组成
//  nums[i] 不含前导零


// 输入： nums = ["3","6","7","10"], k = 4
// 输出： "3"
// 解释： 升序排序后为 ["3","6","7","10"] ，第 4 大的数为 "3"

// 输入： nums = ["2","21","12","1"], k = 3
// 输出： "2"
// 解释： 升序排序后为 ["1","2","12","21"] ，第 3 大的数为 "2"

// 输入： nums = ["0","0"], k = 2
// 输出： "0"
// 解释： 升序排序后为 ["0","0"] ，第 3 大的数为 "0"


// 思路： 排序
//
//       第 k 大就是第 n - k + 1 小。
//
//       对 nums 按正数升序排序，然后返回 nums[n - k + 1 - 1] 即可。
//
//       比较两个字符串表示对正数 a 和 b 流程如下：
//           1. 如果 a 和 b 长度不等，则长度更长的数更大
//           2. 否则，找到左起第一个不同的字符，对应数字大的数更大
//           3. 如果全部字符都相同，则 a 和 b 相等
//
//
//       设字符串长度为 L 。
//
//       时间复杂度： O(L * nlogn)
//           1. 需要对 nums 中全部 O(n) 个长度为 O(L) 的字符串排序，
//               排序时间复杂度为 O(L * nlogn)
//       空间复杂度： O(1)
//           1. 只需要维护常数个额外变量即可


use std::cmp::Ordering;


impl Solution {
    pub fn kth_largest_number(mut nums: Vec<String>, k: i32) -> String {
        // 第 k 大就是第 n - k + 1 小
        let k = nums.len() - k as usize + 1;
        // 按升序排序后，返回 nums[k - 1] 即可
        nums.sort_by(|a, b| {
            // 如果长度不等，则长度更长的数更大
            if a.len() != b.len() {
                return a.len().cmp(&b.len());
            }
            // 如果长度相等，则比较字符串的大小
            for (ach, bch) in a.chars().zip(b.chars()) {
                // 对应位置字符不等，则数字大的数更大
                if ach != bch {
                    return ach.cmp(&bch);
                }
            }
            // 都相等，则两个字符串数字相同
            Ordering::Equal
        });

        nums[k - 1].clone()
    }
}
