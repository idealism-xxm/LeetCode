// 链接：https://leetcode.com/problems/running-sum-of-1d-array/
// 题意：给定一个整数数组 nums ，求其前缀和数组？


// 数据限制：
//  1 <= nums.length <= 1000
//  -(10 ^ 6) <= nums[i] <= 10 ^ 6


// 输入： nums = [1,2,3,4]
// 输出： [1,3,6,10]
// 解释： 前缀和是 [1, 1+2, 1+2+3, 1+2+3+4] 。

// 输入： nums = [1,1,1,1,1]
// 输出： [1,2,3,4,5]
// 解释： 前缀和是 [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1] 。

// 输入： nums = [3,1,2,10,1]
// 输出： [3,4,6,16,17]
// 解释： 前缀和是 [3, 4, 6, 16, 17] 。


// 思路： 前缀和
//
//      直接按照题意求前缀和即可。
//
//      设 prefix[i] 表示前缀和 sum(nums[:i]) ，
//      那么 prefix[i + 1] = sum(nums[:i + 1])
//                        = sum(nums[:i]) + nums[i + 1]
//                        = prefix[i] + nums[i + 1]
//
//      即只要知道 prefix[i] ，我们就能在 O(1) 内求出 prefix[i + 1] 。
//
//      所以初始化 prefix[0] = nums[0] ，
//      然后按照上面的推导公式计算剩余的前缀和即可。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 nums 中全部 O(n) 个数字
//      空间复杂度：O(n)
//          1. 需要用 prefix 维护全部 O(n) 个前缀和


func runningSum(nums []int) []int {
    // prefix[i] 表示 nums[:i] 的前缀和
    prefix := make([]int, len(nums))
    // 初始化 prefix[0] ，然后从 1 开始进行状态转移即可
    prefix[0] = nums[0]
    for i := 1; i < len(nums); i++ {
        prefix[i] = prefix[i - 1] + nums[i]
    }

    return prefix
}
