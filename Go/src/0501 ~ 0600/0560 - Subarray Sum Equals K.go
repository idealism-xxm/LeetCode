// 链接：https://leetcode.com/problems/subarray-sum-equals-k/
// 题意：给定一个整数数组 nums 和一个整数 k ，
//      返回该数组中和为 k 的连续子数组的个数。


// 数据限制：
//  1 <= nums.length <= 2 * 10 ^ 4
//  -1000 <= nums[i] <= 1000
//  -(10 ^ 7) <= k <= 10 ^ 7


// 输入：nums = [1,1,1], k = 2
// 输出：2
// 解释：sum(nums[..2]) = sum([1,1]) = 2
//      sum(nums[1..]) = sum([1,1]) = 2

// 输入：nums = [1,2,3], k = 3
// 输出：2
// 解释：sum(nums[..2]) = sum([1,2]) = 3
//      sum(nums[2..]) = sum([3]) = 3


// 思路：前缀和 + Map
//
//      使用前缀和来进行处理，假设 preSum[i] 表示前缀 nums[0..=i] 的和，
//      那么子数组 nums[i..=j] 的和为 preSum[j] - preSum[i - 1] 。
//
//      根据题意，我们需要让这个和为 k ，
//      即 preSum[j] - preSum[i - 1] = k ，
//      变形为 preSum[j] - k = preSum[i - 1] 。
//
//      那么在 j 之前，前缀和为 preSum[j] - k 数量就是
//      以 nums[j] 为结尾的和为 k 的子数组个数。
//
//      注意最开始要初始化前缀和 preSum[-1] = 0 出现过一次，
//      即假设存在一个和为 0 的空子数组，
//      这是为了方便统计和为 k 的子数组 nums[..=j] 的情况​。
//
//
//		时间复杂度： O(n)
//          1. 需要遍历 nums 中全部 O(n) 个数字
//      空间复杂度： O(n)
//          1. 需要用 map 维护全部不同的前缀和，最差情况下有 O(n) 个


func subarraySum(nums []int, k int) int {
    // preSumToCnt[preSum] 表示前缀和 preSum 出现的次数
    preSumToCnt := make(map[int]int)
    // 最开始初始化前缀和 preSum[-1] = 0 出现过一次，
    // 即假设存在一个和为 0 的空子数组，为了方便统计和为 k 的子数组 nums[..=j] 的情况
    preSumToCnt[0] = 1
    // 维护当前的前缀和
    preSum := 0
    // 维护满足题意的子数组个数
    ans := 0
    // 遍历数组
    for _, num := range nums {
        // 前缀和加上当前的数字
        preSum += num
        //  前缀和 preSum - k 出现的次数就是以当前数为结尾的和为 k 的子数组个数。
        //
        //  子数组 nums[i..=j] 的和为 preSum[j] - preSum[i - 1] ，
        //  那么根据题意，我们需要让这个和为 k ，
        //  即 preSum[j] - preSum[i - 1] = k
        //      => preSum[j] - k = preSum[i - 1]
        //  也就是要找到在 j 之前，前缀和为 preSum[j] - k 数量
        ans += preSumToCnt[preSum - k]
        // 当前前缀和 preSum 的出现次数 +1
        preSumToCnt[preSum] += 1
    }

    return ans
}
