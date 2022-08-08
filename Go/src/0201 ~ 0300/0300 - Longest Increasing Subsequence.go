// 链接：https://leetcode.com/problems/longest-increasing-subsequence/
// 题意：给定一个整数数组，求最长严格上升子序列的长度？
//
//      进阶：使用时间复杂度为 O(nlogn) 的解法。


// 数据限制：
//  1 <= nums.length <= 2500
//  -(10 ^ 4) <= nums[i] <= 10 ^ 4


// 输入： nums = [10,9,2,5,3,7,101,18]
// 输出： 4
// 解释： 最长的上升子序列是 [2,3,7,101] ，它的长度是 4


// 输入： nums = [0,1,0,3,2,3]
// 输出： 4
// 解释： 最长的上升子序列是 [0,1,2,3] ，它的长度是 4

// 输入： nums = [7,7,7,7,7,7,7]
// 输出： 1
// 解释： 最长的上升子序列是 [7] ，它的长度是 1


// 思路：DP + 二分
//
//      最简单地求普通 LIS 就是普通 DP ：
//          设 dp[i] 为以第 i 个元素为结尾的 LIS 的长度，
//          那么在更新 dp[i] 时，需要更新 dp[i] = max(dp[j]) + 1 ，
//          其中 j < i，且 nums[j] < nums[i] 。
//
//      时间复杂度为 O(n ^ 2) ，空间复杂度为 O(n)
//
//
//      可以使用二分将时间复杂度优化为 O(nlogn) ，只需要注意到求解过程中的内在约束：
//          设 minNum[k] 表示长度为 k 的 LIS 的最后一个数字的最小值。
//
//          为了方便后续处理，初始化 minNum = [MIN] ，
//          表示长度为 0 的 LIS 的最后一个数字的最小值为 MIN 。
//
//          那么我们在求解过程中维护的 minNum 必定是一个严格递增的数组。
//
//      注意到这个约束后，我们就不需要遍历前面求出的全部状态 dp[j] ，
//      只需要在 minNum 中找到第一个大于等于当前数 num 的下标 k 即可。
//
//      此时 k 就是以 num 为结尾的 LIS 的长度：
//          1. len(minNum) == k: 
//              说明长度为 k 的 LIS 是第一次出现，直接将 num 加入 minNum 中即可
//          2. len(minNum) > k:
//              说明长度为 k 的 LIS 已经出现过了，
//              由于二分找到的是第一个大于等于 num 的下标 k ，
//              所以必定有 min_num[k] >= num ，
//              可以直接更新 min_num[k] 为 num
//
//      最后 LIS 的长度最长为 len(minNum) - 1
//
//
//      本题是普通的最长上升子序列 (LIS) ，
//      加强版的题目可以继续查看 LeetCode 354 这题，
//      需要保证子序列的两个值都是严格递增的。
//
//
//      时间复杂度：O(nlogn)
//          1. 需要遍历 nums 全部 O(n) 个数字
//          2. 遍历每个数字时，都需要进行 O(logn) 的二分
//      空间复杂度：O(n)
//          1. 需要维护一个大小为 O(n) 的数组 minNum


func lengthOfLIS(nums []int) int {
    // minNum[k] 表示 LIS 的长度为 k 时，最后一个数字的最小值。
    // 初始 LIS 的长度为 0 ，最后一个数的最小值为 MIN ，方便后续处理。
    minNum := []int{ math.MinInt32 }
    // 遍历每个数
    for _, num := range nums {
        // 寻找 minNum 中第一个大于等于 num 的下标 k ，
        // 则说明以 num 为 LIS 的最后一个数时，对应的 LIS 的长度最长为 k
        k := sort.Search(len(minNum), func (i int) bool { return minNum[i] >= num })
        if len(minNum) == k {
            // num 是第一个使 LIS 长度达到 k 的数，所以直接放入 minNum
            minNum = append(minNum, num)
        } else {
            // 此时存在长度为 k 的 LIS ，
            // 因为前面二分寻找的是第一个大于等于 num 的下标，
            // 所以 minNum[k] >= num ，可以直接更新为 num
            minNum[k] = num
        }
    }

    // LIS 的长度最长为 len(minNum) - 1
    return len(minNum) - 1
}
