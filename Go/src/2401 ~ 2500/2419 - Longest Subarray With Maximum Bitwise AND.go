// 链接：https://leetcode.com/problems/longest-subarray-with-maximum-bitwise-and/
// 题意：给定一个长度为 n 的整型数组 nums ，在其所有的非空子数组中，
//      找到子数组的按位与最大的那些，然后返回这些子数组中最长的子数组的长度。


// 数据限制：
//  1 <= nums.length <= 10 ^ 5
//  1 <= nums[i] <= 10 ^ 6


// 输入： nums = [1,2,3,3,2,2]
// 输出： 2
// 解释： 所有子数组中，按位与的最大值为 3 ，
//       按位与为 3 的子数组有 [3], [3,3], [3] ，最长的子数组长度为 2

// 输入： nums = [1,2,3,4]
// 输出： 1
// 解释： 所有子数组中，按位与的最大值为 4 ，
//       按位与为 4 的子数组有 [4] ，最长的子数组长度为 1


// 思路： 贪心
//
//      对于两个非负数 a, b ，有 a & b <= min(a, b) <= max(a, b) ，
//      即两个非负数按位与后，结果值不会更大。
//
//      那么可以贪心地找到 nums 中的最大值 maxNum ，
//      则 maxNum 必定是 nums 所有的非空子数组的按位与的最大值。
//
//      接下来就只需要找到所有数都为 maxNum 的最长子数组的长度即可。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 nums 中全部 O(n) 个数字两次
//      空间复杂度：O(1)
//          1. 只需要维护常数个额外变量


func longestSubarray(nums []int) int {
    // 找到 nums 中的最大值，该值就是 nums 所有的非空子数组的按位与的最大值
    maxNum := 0
    for _, num := range nums {
        maxNum = max(maxNum, num)
    }
    // cnt 维护以当前数为结尾，所有数都是 maxNum 的子数组的最长长度
    cnt := 0
    // ans 维护所有 cnt 的最大值，即为题目所求
    ans := 0
    for _, num := range nums {
        if num == maxNum {
            cnt += 1
        } else {
            ans = max(ans, cnt)
            cnt = 0
        }
    }

    // 【注意】以最后一个数为结尾的长度没有计入，需要手动取最大值
    return max(ans, cnt)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}