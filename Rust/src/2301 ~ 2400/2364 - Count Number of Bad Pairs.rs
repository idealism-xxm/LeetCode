// 链接：https://leetcode.com/problems/count-number-of-bad-pairs/
// 题意：给定一个整型数组 nums ，求有多少个二元组 (i, j) 满足以下条件：
//          1. i < j
//          2. j - i != nums[j] - nums[i]


// 数据限制：
//  1 <= nums.length <= 10 ^ 5
//  1 <= nums[i] <= 10 ^ 9


// 输入： nums = [4,1,3,3]
// 输出： 5
// 解释： 满足条件的二元组有：
//          (0,1): 1 - 0 != 1 - 4
//          (0,2): 2 - 0 != 3 - 4
//          (0,3): 3 - 0 != 3 - 4
//          (1,2): 2 - 1 != 3 - 1
//          (2,3): 3 - 2 != 3 - 3

// 输入： nums = [1,2,3,4,5]
// 输出： 2
// 解释： 没有满足题意的二元组


// 思路： Map
//
//      第二个条件如果不变形的话很难线性处理，因为等式两边都与二元组的两个数有关。
//
//      我们可以对其进行变形：
//          j - i != nums[j] - nums[i]
//      =>  j - nums[j] != i - nums[i]
//
//      这时等式两边分别只与二元组中的一个数有关，而且形式一致，就可以线性处理了。
//
//      我们用 cnt 维护所有 i - nums[i] 的出现次数，并用 ans 维护满足题意的二元组的个数。
//
//      那么我们可以按顺序遍历 nums 中的每个数 nums[j] ，
//      此时 [0, j) 中共有 j 个数可作为 i ，
//      与 j 组成满足第一个条件的二元组 (i, j) 。
//
//      我们需要从中找到不满足第二个条件的二元组的数量，
//      也就是 i - nums[i] = j - nums[j] 的数量，即 cnt[j - nums[j]] 。
//
//      综上： [0, j) 中共有 j - cnt[j - nums[j]] 个数可作为 i ，
//          与 j 组成满足题意的二元组 (i, j) ，计入 ans 即可。
//
//      然后计入 j - nums[j] 的次数，方便后续统计。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 nums 中全部 O(n) 个数一次
//      空间复杂度：O(n)
//          1. 需要维护 cnt 中全部不同数的出现次数，最差情况下有 O(n) 个


use std::collections::HashMap;


impl Solution {
    pub fn count_bad_pairs(nums: Vec<i32>) -> i64 {
        // ans 维护满足题意的二元组的个数
        let mut ans = 0;
        // cnt 维护所有 i - nums[i] 的出现次数
        let mut cnt = HashMap::new();
        // 按顺序遍历 nums 中的每个数 nums[j]，进行如下处理即可
        for (j, &num) in nums.iter().enumerate() {
            let diff = j as i32 - num;
            // [0, j) 中共有 j 个数，但其中有 cnt[j - nums[j]] 无法与 j 组成满足题意的二元组
            ans += j as i64 - (*cnt.get(&diff).unwrap_or(&0)) as i64;
            // 计入 j - nums[j] 的次数，方便后续统计
            cnt.entry(diff).and_modify(|x| *x += 1).or_insert(1);
        }

        ans
    }
}
