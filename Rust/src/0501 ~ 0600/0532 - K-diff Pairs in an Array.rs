// 链接：https://leetcode.com/problems/k-diff-pairs-in-an-array/
// 题意：给定一个整数数组 nums 和一个整数 k ，
//      返回数组中不同的 k-diff 数对的个数。
//
//      k-diff 数对为满足以下条件的整数对 (nums[i], nums[j]):
//          1. 0 <= i < j < nums.length
//          2. |nums[i] - nums[j]| == k ，其中 |val| 表示 val 的绝对值


// 数据限制：
//  1 <= nums.length <= 10 ^ 4
//  -(10 ^ 7) <= nums[i] <= 10 ^ 7
//  0 <= k <= 10 ^ 7

// 输入：nums = [3, 1, 4, 1, 5], k = 2
// 输出：2
// 解释：数组中有两个 2-diff 数对， (1, 3) 和 (3, 5) 。
//      尽管数组中有两个 1 ，但我们只应返回不同的数对的数量。

// 输入：nums = [1, 2, 3, 4, 5], k = 1
// 输出：4
// 解释：数组中有四个 1-diff 数对， (1, 2), (2, 3), (3, 4) 和 (4, 5) 。

// 输入：nums = [1, 3, 1, 5, 4], k = 0
// 输出：1
// 解释：数组中只有一个 0-diff 数对， (1, 1) 。


// 思路：Map
//
//      先统计 nums 中每个数字出现的次数到 num_to_cnt 中。
//
//      然后看 k-diff 数对要满足两个条件，直接根据题意判断处理比较麻烦。
//      但题目只要求统计 k-diff 数对的的个数，不需要找到所有的 k-diff 数对，
//      所以我们可以将条件转换为找到这样的数对 (nums[i], nums[j]) ：
//          1. i != j
//          2. nums[j] - nums[i] == k
//
//      那么此时我们就可以枚举 num_to_cnt 中的
//      每一个数字 num 及其出现的次数 cnt ，
//      令 num 为数对的第一个数字，那么可以进行如下判断：
//          1. k == 0: 那么 cnt 至少出现 2 次时，才能形成一个满足题意的数对
//          2. k != 0: 那么 num + k 至少出现 1 次时，才能形成一个满足题意的数对
//
//
//      时间复杂度：O(n)
//      空间复杂度：O(n)

use std::collections::HashMap;

impl Solution {
    pub fn find_pairs(nums: Vec<i32>, k: i32) -> i32 {
        // 初始化 num_to_cnt 为 nums 中每个数字出现的次数
        let mut num_to_cnt = HashMap::new();
        for num in nums {
            // 如果 num 不在 num_to_cnt 中，另其出现次数为 0 ，
            // 然后对 num 的出现次数 +1
            *num_to_cnt.entry(num).or_insert(0) += 1;
        }

        // 统计符合条件的数字对的数量
        let mut ans = 0;
        // 枚举每个数字及其出现的次数
        for (&num, &cnt) in num_to_cnt.iter() {
            if k == 0 {
                // 如果 k 为 0 ，则只有 num 出现至少 2 次时，
                // 才能形成一个满足题意的数字对
                if cnt > 1 {
                    ans += 1;
                }
            } else {
                // 如果 k 不为 0 ，则只要存在 num + k ，
                // 那么就能形成一个满足题意的数字对
                if num_to_cnt.contains_key(&(num + k)) {
                    ans += 1;
                }
            }
        }
        
        ans
    }
}