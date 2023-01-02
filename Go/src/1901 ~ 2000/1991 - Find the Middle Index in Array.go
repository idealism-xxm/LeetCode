// 链接：https://leetcode.com/problems/find-the-middle-index-in-array/
// 题意：给定一个整型数组 nums ，返回一个最左侧的下标 index ，
//      使得 sum(nums[:index]) == sum(nums[index + 1:]) ，
//      不存在则返回 -1 。


// 数据限制：
//  1 <= nums.length <= 100
//  -1000 <= nums[i] <= 1000


// 输入： nums = [2,3,-1,8,4]
// 输出： 3
// 解释： sum(nums[:3]) = nums[0] + nums[1] + nums[2] = 2 + 3 + -1 = 4
//       sum(nums[4:]) = nums[4] = 4

// 输入： nums = [1,-1,4]
// 输出： 2
// 解释： sum(nums[:2]) = nums[0] + nums[1] = 1 + -1 = 0
//       sum(nums[3:]) = 0

// 输入： nums = [2,5]
// 输出： -1

// 输入： nums = [0]
// 输出： 0
// 解释： sum(nums[:0]) = 0
//       sum(nums[1:]) = 0


// 思路： 前缀和
//
//       先求出所有数字的和 total ，
//       然后从开始枚举下标 index ，并记录不含当前下标的前缀和 prefixSum 。
//
//       如果 total - nums[index] == prefixSum * 2 ，
//       那么 index 就是我们要求的下标，直接返回 index 即可。
//
//       最后还未返回时，则说明没有满足题意的下标，返回 -1 。
//
//
//       时间复杂度： O(n)
//           1. 需要遍历 nums 全部 O(n) 个数字两次
//       空间复杂度： O(1)
//           1. 只需要维护常数个额外变量即可


func findMiddleIndex(nums []int) int {
    // 求数组和
    total := 0
    for _, num := range nums {
        total += num
    }
    // 不含当前数字的前缀和
    prefixSum := 0
    for i, num := range nums {
        // 如果 total - num 是前缀和的两倍，那么 i 就是我们要求的下标
        if prefixSum * 2 == total - num {
            return i
        }
        // 当前数字加入到求前缀和中
        prefixSum += num
    }
    // 遍历完所有的数字都没有找到，则不存在，返回 -1
    return -1
}
