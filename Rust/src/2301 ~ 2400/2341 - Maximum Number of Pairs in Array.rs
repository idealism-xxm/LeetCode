// 链接：https://leetcode.com/problems/maximum-number-of-pairs-in-array/
// 题意：给定一个整型数组 nums ，每次操作可以选择 nums 中相等的两个数移除，
//      求最多能进行多少次操作？
//      返回结果 answer 是一个长度为 2 的数组，
//      answer[0] 表示操作次数，answer[1] 表示 nums 最后的长度。


// 数据限制：
//  1 <= nums.length <= 100
//  0 <= nums[i] <= 100


// 输入： nums = [1,3,2,1,3,2,2]
// 输出： [3,1]
// 解释： 第 1 次操作：移除 nums[0] 和 nums[3] ， nums 变为 [3,2,3,2,2]
//       第 2 次操作：移除 nums[0] 和 nums[2] ， nums 变为 [2,2,2]
//       第 3 次操作：移除 nums[0] 和 nums[1] ， nums 变为 [2]
//
//       总共进行了 3 次操作，最后 nums 中还有 1 个数。

// 输入： nums = [1,1]
// 输出： [1,0]
// 解释： 第 1 次操作：移除 nums[0] 和 nums[1] ， nums 变为 []
//
//       总共进行了 1 次操作，最后 nums 中还有 0 个数。

// 输入： nums = [0]
// 输出： [0,1]
// 解释： 总共进行了 0 次操作，最后 nums 中还有 1 个数。


// 思路： Map
//
//      我们可以发现如果一个数字 nums 的出现次数 cnt 有两种情况：
//          1. cnt 是偶数，则需要进行 cnt / 2 次操作， num 剩余 0 个数
//          2. cnt 是奇数，则需要进行 cnt / 2 次操作， num 剩余 1 个数
//
//      那么我们只需要统计 nums 中每个数的出现次数到 num_to_cnt 中，
//      然后遍历 num_to_cnt 中每个数字的出现次数 cnt ：
//          1. 计入每个数字需要进行的操作数 cnt / 2 到 ans[0] 中，即 ans[0] += cnt / 2
//          2. 计入每个数字剩余的个数 cnt % 2 到 ans[1] 中，即 ans[1] += cnt % 2
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 nums 中全部 O(n) 个数字
//          2. 需要遍历 num_to_cnt 中全部不同的数字的出现次数，
//              最差情况下有 O(n) 个不同的数字
//      空间复杂度：O(n)
//          1. 需要维护 nums 中全部不同的数字及其出现次数，
//              最差情况下有 O(n) 个不同的数字


use std::collections::HashMap;


impl Solution {
    pub fn number_of_pairs(nums: Vec<i32>) -> Vec<i32> {
        // 统计 nums 中每个数字出现的次数
        let mut num_to_cnt = HashMap::new();
        for num in nums {
            *num_to_cnt.entry(num).or_insert(0) += 1;
        }

        let mut ans = vec![0, 0];
        // 我们不关心这个数是什么，所以只遍历数字的出现次数
        for cnt in num_to_cnt.values() {
            // 计入当前数字需要进行的操作数
            ans[0] += cnt >> 1;
            // 计入当前数字剩余的个数
            ans[1] += cnt & 1;
        }

        ans
    }
}
