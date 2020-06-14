// 链接：https://leetcode.com/problems/contains-duplicate/
// 题意：给定一个数组，判断是否存在重复的数？

// 输入： [1,2,3,1]
// 输出： true

// 输入： [1,2,3,4]
// 输出： false

// 输入： [1,1,1,3,3,4,3,2,4,2]
// 输出： true

// 思路： set
//
//		将数组转成集合，判断长度是否相等即可
//
//      时间复杂度： O(n)
//      空间复杂度： O(n)

use std::collections::HashSet;
use std::iter::FromIterator;

impl Solution {
    pub fn contains_duplicate(nums: Vec<i32>) -> bool {
        let set: HashSet<&i32> = HashSet::from_iter(nums.iter());
        return set.len() != nums.len();
    }
}
