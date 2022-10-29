// 链接：https://leetcode.com/problems/continuous-subarray-sum/
// 题意：给定一个整数数组 nums 和一个正整数 k ，判断是否存在满足以下条件的子数组？
//          1. 子数组长度至少为 2
//          2. 子数组的和模 k 余 0（ 0 % k 必定是 0 ）


// 数据限制：
//  1 <= nums.length <= 10 ^ 5
//  0 <= nums[i] <= 10 ^ 9
//  0 <= sum(nums[i]) <= 2 ^ 31 - 1
//  1 <= k <= 2 ^ 31 - 1


// 输入： nums = [23,2,4,6,7], k = 6
// 输出： true
// 解释： 子数组 [2,4] 长度为 2 ，和为 6

// 输入： nums = [23,2,6,4,7], k = 6
// 输出： true
// 解释： 子数组 [23,2,6,4,7] 长度为 5 ，和为 6

// 输入： nums = [23,2,6,4,7], k = 13
// 输出： false
// 解释： 任何长度大于 2 的子数组都不满足题意


// 思路： 前缀和 + Map
//
//      本题是 LeetCode 560 的加强版，与 LeetCode 525 类似，
//      需要多一步转换，在思考时需要运用模运算的性质，
//
//      如果 x >= y, x % k = a, y % k = a ，
//      那么必定有 (x - y) % k = 0 。
//
//      不妨设 x = i * k + a, y = j * k + a ，
//      那么 x - y = i * k + a - j * k - a = (i - j) * k ，
//      即 (x - y) % k = 0 。
//
//      所以我们可以维护每个前缀和模 k 的余数 rmain 第一次出现的下标 j 。
//
//      如果该余数 remain 后续出现的下标 i 满足 i - j > 1 ，
//      那么子数组 nums[j+1..=i] 的长度至少为 2 且数字和模 k 余 0 ，
//      满足题意，可以直接返回 true 。
//
//      如果最后还未返回，则所有子数组都不满足题意，直接返回 false 。
//
//      注意最开始要初始化前缀和 0 % k 的下标为 -1 ，
//      即假设存在一个和为 0 的空子数组，方便处理整个前缀和满足题意的情况。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历计算全部 O(n) 前缀和
//      空间复杂度：O(n)
//          1. 需要维护全部不同的前缀和第一次出现的下标，最差情况下有 O(n) 个


func checkSubarraySum(nums []int, k int) bool {
    // remainToFirstIndex 维护每一个前缀和 % k 第一次出现的下标
    remainToFirstIndex := make(map[int]int)
    // 初始前缀和 0 % k 的下标为 -1 ，方便处理整个前缀和满足题意的情况
    remainToFirstIndex[0] = -1
    preSum := 0
    for i, num := range nums {
        // 计算 nums[..=i] 的前缀和及模 k 的余数
        preSum += num
        remain := preSum % k
        if j, exists := remainToFirstIndex[remain]; exists {
            // 如果模 k 余 remain 的前缀和已存在，
            // 那么 sum(nums[j+1:=i]) % k 为 0 ，
            // 只要子数组 nums[j+1:=i] 长度大于 1 就满足题意
            if i - j > 1 {
                return true
            }
        } else {
            // 如果 remain 是第一次出现，则记录其下标
            remainToFirstIndex[remain] = i
        }
    }

    // 如果最后还未返回，则所有子数组都不满足题意，直接返回 false
    return false
}
