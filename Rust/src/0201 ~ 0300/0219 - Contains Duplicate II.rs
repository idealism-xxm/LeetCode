// 链接：https://leetcode.com/problems/contains-duplicate-ii/
// 题意：给定一个数组和一个整数 k ，判断是否存在一对重复的数的下标差最多为 k ？

// 输入： nums = [1,2,3,1], k = 3
// 输出： true

// 输入： nums = [1,0,1,1], k = 1
// 输出： true

// 输入： nums = [1,2,3,1,2,3], k = 2
// 输出： false

// 思路： map
//
//		遍历数组 nums ，每次先获取 nums[i] 在前面的最右下标 j ，
//      若存在，且 i - j <= k ，则直接返回 true，
//      然后将 nums[i] 对应的最右下标设置为 i 即可
//      在循环中没有返回，则所有数都不满足题意，直接返回 false
//
//      时间复杂度： O(n)
//      空间复杂度： O(n)

use std::collections::HashMap;

impl Solution {
    pub fn contains_nearby_duplicate(nums: Vec<i32>, k: i32) -> bool {
        let k = k as usize;
        let mut map: HashMap<i32, usize> = HashMap::new();
        for (i, num) in nums.into_iter().enumerate() {
            // 如果 num 已存在，且两者下标差小于等于 k ，则直接返回 true
            if let Some(j) = map.get(&num) {
                if i - j <= k {
                    return true;
                }
            }
            // 标记 num 出现的最右下标为 i
            map.insert(num, i);
        }
        // 循环内未返回，则所有数都不满足题意，直接返回 false
        return false;
    }
}
