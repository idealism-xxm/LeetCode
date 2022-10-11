// 链接：https://leetcode.com/problems/increasing-triplet-subsequence/
// 题意：给定一个整数数组 nums ，判断是否存在满足以下条件的三元组 (i, j, k) ：
//          1. i < j < k
//          2. nums[i] < nums[j] < nums[k]
//
//      进阶：使用时间复杂度为 O(n) ，空间复杂度为 O(1) 的算法


// 数据限制：
//  1 <= nums.length <= 5 * 10 ^ 5
//  -(2 ^ 31) <= nums[i] <= 2 ^ 31 - 1


// 输入： nums = [1,2,3,4,5]
// 输出： true
// 解释： 任何一个三元组都满足题意

// 输入： nums = [5,4,3,2,1]
// 输出： false
// 解释： 不存在满足题意的三元组

// 输入： nums = [2,1,5,0,4,6]
// 输出： false
// 解释： 三元组 (3, 4, 5) 满足题意，
//       因为 nums[3] == 0 < nums[4] == 4 < nums[5] == 6


// 思路： 贪心
//
//      我们可以维护 i, j 对应的数 num_i 和 num_j ，
//      都初始化为 MAX ，方便后续处理。
//
//      即使 nums 中最大的数为 MAX 也不影响，
//      因为满足题意的三元组必须满足 num_i < num_j < num_k <= MAX 。
//
//      我们遍历 nums 中的每个数，如果当前数 num > num_j ，
//      则已经找到满足题意的三元组，直接返回 true 。
//
//      否则，贪心地尽可能减小 num_i 和 num_j ，
//      这样能扩大 num_j 和 num_k 的取值范围，更可能找到满足题意的三元组。
//		
//
//		时间复杂度： O(n)
//          1. 需要遍历 nums 中全部 O(n) 个数
//		空间复杂度： O(1)
//          1. 只需要维护常数个额外变量即可


impl Solution {
    pub fn increasing_triplet(nums: Vec<i32>) -> bool {
        // 初始化 i, j 对应的值为 MAX ，方便后续处理
        let mut num_i = i32::MAX;
        let mut num_j = i32::MAX;
        // 遍历 nums 中的每个数，贪心地更新三元组对应的数
        for num in nums {
            // 如果当前数比 num_j 大，则存在满足题意的三元组，
            // 三元组对应的值为 (num_i, num_j, num)
            if num > num_j {
                return true;
            }
            
            if num > num_i && num < num_j {
                // 如果当前数比 num_i 大， 且比 num_j 小，
                // 那么贪心地更新 num_j = num ，扩大 num_k 的取值范围
                num_j = num;
            } else if num < num_i {
                    // 如果当前数比 num_i 小，
                // 那么贪心地更新 num_i = num ，扩大 num_j 的取值范围
                num_i = num;
            }
        }

        // 此时没有找到满足题意的三元组，返回 false
        false
    }
}
