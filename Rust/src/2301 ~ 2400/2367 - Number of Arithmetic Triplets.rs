// 链接：https://leetcode.com/problems/number-of-arithmetic-triplets/
// 题意：给定一个严格单调递增的数组 nums 和一个整数 diff ，
//      求满足以下条件的三元组 (i, j, k) 的数量：
//          1. i < j < k
//          2. nums[j] - nums[i] = diff
//          3. nums[k] - nums[j] = diff


// 数据限制：
//  3 <= nums.length <= 200
//  0 <= nums[i] <= 200
//  1 <= diff <= 50
//  nums 是严格单调递增的


// 输入： nums = [0,1,4,6,7,10], diff = 3
// 输出： 2
// 解释： (1, 2, 4) 满足题意，因为 nums[2] - nums[1] = 4 - 1 = 3, 
//          nums[4] - nums[2] = 7 - 4 = 3 ；
//       (2, 4, 5) 满足题意，因为 nums[4] - nums[2] = 7 - 4 = 3, 
//          nums[5] - nums[4] = 10 - 7 = 3 。


// 输入： nums = [4,5,6,7,8,9], diff = 2
// 输出： 2
// 解释： (0, 2, 4) 满足题意，因为 nums[2] - nums[0] = 6 - 4 = 2,
//          nums[4] - nums[2] = 8 - 6 = 2 ；
//       (1, 3, 5) 满足题意，因为 nums[3] - nums[1] = 7 - 5 = 2,
//          nums[5] - nums[3] = 9 - 7 = 3 。


// 思路： Set/Map
//
//      由于 nums 是严格单调递增的，所以 nums 中的数字必定各不相同，
//      且三元组 (i, j, k) 只要满足了后两个条件，那么 i < j < k 就一定成立。
//
//      如果我们枚举其中一个数，那么通过后两个条件，我们就可以知道其他两个数是多少，
//      只需要判断其他两个数字是否在 nums 中即可。
//
//      为了方便处理，我们枚举 nums 中的每个数 num 作为 nums[k] ，那么通过后两个条件可知：
//          1. nums[j] = nums[k] - diff
//          2. nums[i] = nums[k] - 2 * diff
//
//      只要 nums[k] - diff 和 nums[k] - 2 * diff 在 nums 中，
//      那么存在以 num 作为 nums[k] 的满足题意的三元组 (i, j, k) ，计入 ans 。
//
//      为了能在 O(1) 内判断一个数是否在 nums 中，我们维护一个名为 num_set 的集合，
//      每次枚举处理完后就将 num 放入集合中，方便后续判断。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 nums 中全部 O(n) 个数字一次
//      空间复杂度：O(n)
//          1. 需要维护 num_set 中全部 O(n) 个不同的数字


use std::collections::HashSet;


impl Solution {
    pub fn arithmetic_triplets(nums: Vec<i32>, diff: i32) -> i32 {
        let mut ans = 0;
        // num_set 维护 nums 中不同的数字，方便后续在 O(1) 内判断一个数字是否存在。
        // （由于 nums 是严格单调递增的，所以 nums 中的数字必定全部各不相同）
        let mut num_set = HashSet::with_capacity(nums.len());
        // 遍历每个数字，作为 nums[k]
        for num in nums {
            // 如果 nums[j] = nums[k] - diff 和 nums[i] = nums[k] - 2 * diff 均存在，
            // 则存在以 num 作为 nums[k] 的满足题意的三元组 (i, j, k) ，计入 ans
            if num_set.contains(&(num - diff)) && num_set.contains(&(num - 2 * diff)) {
                ans += 1;
            }
            // 将 nums 放入集合中，方便后续判断
            num_set.insert(num);
        }

        ans
    }
}
