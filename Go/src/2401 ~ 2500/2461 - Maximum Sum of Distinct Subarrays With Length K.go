// 链接：https://leetcode.com/problems/maximum-sum-of-distinct-subarrays-with-length-k/
// 题意：给定一个数组 nums 和一个正整数 k ，
//      求所有长度为 k 且数字各不相同的子数组中，子数组和的最大值？


// 数据限制：
//  1 <= k <= nums.length <= 10 ^ 5
//  1 <= nums[i] <= 10 ^ 5


// 输入： nums = [1,5,4,2,9,9,9], k = 3
// 输出： 15
// 解释： 长度为 3 的子数组如下：
//       · [1,5,4]: 所有的数各不相同，和为 10
//       · [5,4,2]: 所有的数各不相同，和为 11
//       · [4,2,9]: 所有的数各不相同，和为 15
//       · [2,9,9]: 9 重复出现
//       · [9,9,9]: 9 重复出现
//
//       所有长度为 3 且数字各不相同的子数组中，和最大的值为 15

// 输入： nums = [4,4,4], k = 3
// 输出： 0
// 解释： 长度为 3 的子数组如下：
//       · [4,4,4]: 4 重复出现
//
//       不存在长度为 3 且数字各不相同的子数组，所以返回 0


// 思路： 滑动窗口
//
//      如果一道题目需要在所有满足某种状态的连续子串/连续子数组中，
//      找到满足题意的一个，那么可以考虑滑动窗口。
//
//      本题需要在所有长度为 k 且数字各不相同的子数组中，找到和最大的那个。
//
//      那么我们使用滑动窗口 [l, r] 表示一个长度不大于 k 的连续子数组，
//      初始化为左边界 l = 0 ，右边界 r = -1 ，表示初始窗口为空。
//
//      同时我们用 numToCnt 维护滑动窗口 [l, r] 内每个数字的出现次数。
//      并用 sum 维护滑动窗口 [l, r] 内所有数字之和，初始化为 0。
//
//      然后不断右移右边界 r ，将其纳入到滑动窗口中考虑。
//
//      如果此时 r >= k ，则说明滑动窗口 [l ,r] 的长度不满足题意，
//      需要移除 nums[l] 并右移 l ，以保持滑动窗口长度为 k 。
//
//      如果移除 nums[l] 后，其出现次数为 0 ，
//      那么还需要从 numToCnt 移除 nums[l] ，
//      以保证 numToCnt 中维护的都是滑动窗口内的数字的出现次数。
//
//      此时，如果 numToCnt 的大小为 k ，
//      那么说明滑动窗口 [l, r] 内的全部 k 个数字各不相同，
//      更新 ans 的最大值即可。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 nums 中全部 O(n) 个数字
//      空间复杂度：O(k)
//          1. 需要维护 numToCnt 中全部 O(k) 个数字的出现次数


func maximumSubarraySum(nums []int, k int) int64 {
    // ans 维护所有长度为 k 且数字各不相同的子数组中，子数组和的最大值
    ans := 0
    // sum 维护当前滑动窗口 [l, r] 内的数字和
    sum := 0
    // numToCnt 表示滑动窗口 [l, r] 内每个数字的出现次数
    numToCnt := make(map[int]int)
    // 不断右移滑动窗口右边界 r
    for r := range nums {
        // 将 nums[r] 纳入滑动窗口中考虑
        numToCnt[nums[r]] += 1
        sum += nums[r]
        // 如果 r >= k ，则需要将 nums[r - k] 从滑动窗口中移除，
        // 以保持滑动窗口长度为 k
        if r >= k {
            numToCnt[nums[r - k]] -= 1
            sum -= nums[r - k]
            // 如果移除后， nums[r - k] 的出现次数为 0 ，
            // 还需要从 numToCnt 移除 nums[l] ，
            // 以保证 numToCnt 中维护的都是滑动窗口内的数字的出现次数
            if numToCnt[nums[r - k]] == 0 {
                delete(numToCnt, nums[r - k])
            }
        }

        // 如果 numToCnt 的大小为 k ，
        // 那么说明滑动窗口 [l, r] 内的全部 k 个数字各不相同，
        // 更新 ans 的最大值
        if len(numToCnt) == k {
            ans = max(ans, sum)
        }
    }

    return int64(ans)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
