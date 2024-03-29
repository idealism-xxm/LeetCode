// 链接：https://leetcode.com/problems/missing-number/
// 题意：给定一个长度为 n 的数组 nums ， nums 中的每个数都唯一，
//      且都在 [0, n] 范围内，求 [0, n] 内 不在 nums 中的数字？
//
//      进阶：使用 O(1) 空间复杂度和 O(n) 时间复杂度的解法。


// 数据限制：
//  n == nums.length
//  1 <= n <= 10 ^ 4
//  0 <= nums[i] <= n
//  nums 中的所有数字都是唯一的


// 输入： nums = [3,0,1]
// 输出： 2
// 解释： 范围 [0, 3] 内不在 nums 中的数是 2

// 输入： nums = [0,1]
// 输出： 2
// 解释： 范围 [0, 2] 内不在 nums 中的数是 2

// 输入： nums = [9,6,4,2,3,5,7,0,1]
// 输出： 8
// 解释： 范围 [0, 9] 内不在 nums 中的数是 8


// 思路： 异或（位运算）
//
//      我们可以把 [0, n] 范围内的数和 nums 内的数都考虑进去，
//      那么这道题就转化成了：
//
//      给定 2 * n - 1 个数，其中 1 个数字只出现了一次，
//      其余数字都出现了两次。
//
//      这就是 LeetCode 136 - Single Number 这题，
//      所以我们可以采用相同的异或解法。
//
//      因为 a ^ a = 0 ，所以出现两次的数异或后会相互抵消。
//
//      那么求所有数的异或和即可，最后剩余的就是只出现一次的数。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历处理全部 O(n) 个数
//      空间复杂度：O(1)
//          1. 只需要维护常数个额外变量即可


impl Solution {
    pub fn missing_number(nums: Vec<i32>) -> i32 {
        // ans 初始化为 0 ^ n = n ，因为下标的范围为 [0, n - 1]
        let mut ans = nums.len() as i32;
        // 带下标遍历 nums
        for (i, num) in nums.iter().enumerate() {
            // 异或下标
            ans ^= i as i32;
            // 异或数组内的数
            ans ^= num;
        }

        // 此时 ans 就是 [0, n] 内不在 nums 中的数
        ans
    }
}
