// 链接：https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/
// 题意：给定一个整型数组，从中选取 k 个数，求最大值减去最小值的差最小是多少？


// 数据限制：
//   1 <= k <= nums.length <= 1000
//   0 <= nums[i] <= 10 ^ 5

 
// 输入： nums = [90], k = 1
// 输出： 0
// 解释： 
//       只能选 [90] ，最大值为 90, 最小值为 90 ，差为 0

// 输入： nums = [9,4,1,7], k = 2
// 输出： 2
// 解释： 
//       选择 [9,7] ，最大值为 9, 最小值为 7, 差为 2


// 思路： 贪心
//
//       当选择的数是按顺序连续的 k 个数时，才有可能取得最小差值。
//
//       假设选择了一个非连续的数，那么差值不会变小，结果不会更优。
//
//       所以我们可以对 nums 排序，然后枚举所有长度为 k 的子数组计算差值的最小值即可。
//
//
//       时间复杂度： O(nlogn)
//           1. 需要对 nums 中全部 O(n) 个数字排序，排序时间复杂度为 O(nlogn)
//           2. 需要遍历 nums 中全部 O(n - k) 个长度为 k 的子数组
//       空间复杂度： O(1)
//           1. 只需要维护常数个额外变量即可


func minimumDifference(nums []int, k int) int {
    // 先排序
    sort.Ints(nums)
    // 从所有长度为 k 的子数组中找到最小差值
    ans := nums[len(nums) - 1] - nums[0]
    for i := k - 1; i < len(nums); i++ {
        ans = min(ans, nums[i] - nums[i - k + 1])
    }
    return ans
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
