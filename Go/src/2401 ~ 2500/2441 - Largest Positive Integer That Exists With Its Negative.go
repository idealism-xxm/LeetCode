// 链接：https://leetcode.com/problems/largest-positive-integer-that-exists-with-its-negative/
// 题意：给定一个不含 0 的整数数组 nums ，求满足以下条件的数 k 的最大值？
//          1. k 在 nums 中
//          2. -k 在 nums 中
//
//      如果没有满足条件的 k ，则返回 -1 。


// 数据限制：
//  1 <= nums.length <= 1000
//  -1000 <= nums[i] <= 1000
//  nums[i] != 0


// 输入： nums = [-1,2,-3,3]
// 输出： 3
// 解释： 3 是唯一一个满足题意的数

// 输入： nums = [-1,10,6,7,-7,1]
// 输出： 7
// 解释： 1 和 7 都满足题意，其中 7 最大

// 输入： nums = [-1,10,6,7,-7,1]
// 输出： -1
// 解释： 不存在满足题意的数


// 思路： Set/Map
//
//      先用一个集合维护 nums 中全部数字，
//      并用 ans 维护答案，初始化为 -1 ，表示暂无满足题意的数。
//
//      然后遍历 nums 中每个数字 num ，如果 -num 在集合中，
//      则满足题意，更新 ans 的最大值即可。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 nums 中全部 O(n) 个数
//      空间复杂度：O(n)
//          1. 需要维护 nums 中全部不同的数，最差情况下有 O(n) 个


func findMaxK(nums []int) int {
    // 用集合 numSet 维护 nums 中全部数字
    numSet := make(map[int]bool)
    for _, num := range nums {
        numSet[num] = true
    }
    // ans 维护答案，初始化为 -1 ，表示暂无满足题意的数
    ans := -1
    for num, _ := range numSet {
        // 如果 -num 在集合中，则满足题意，更新 ans 最大值
        if numSet[-num] {
            ans = max(ans, num)
        }
    }

    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
