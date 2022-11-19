// 链接：https://leetcode.com/problems/destroy-sequential-targets/
// 题意：给定一个数组 nums 和一个整数 space 。
//      可以选择数组中的一个数 nums[i] ，并移除所有等于 nums[i] + c * space 的数，
//      其中 c 是非负整数。
//
//      现在需要移除最多的数，求在移除最多的数的前提下， nums[i] 的最小值？


// 数据限制：
//  1 <= nums.length <= 10 ^ 5
//  1 <= nums[i] <= 10 ^ 9
//  1 <= space <= 10 ^ 9


// 输入： nums = [3,7,8,1,1,5], space = 2
// 输出： 1
// 解释： 选择 nums[3] = 1 ，可以移除 1,1,3,5,7 共 5 个数。

// 输入： nums = [1,3,5,2,4,6], space = 2
// 输出： 1
// 解释： 选择 nums[0] = 1 ，可以移除 1,3,5 共 3 个数。
//       选择 nums[3] = 2 ，可以移除 2,4,6 共 3 个数。
//       nums[0] 更小，所以最终选择 nums[0] = 1 。

// 输入： nums = [6,2,5], space = 100
// 输出： 2
// 解释： 无论选择哪个数，都只能移除本身 1 个数。
//       nums[1] 更小，所以最终选择 nums[1] = 2 。


// 思路： 贪心 + Map
//
//      我们对 nums[i] 按照 remain = nums[i] % space 分组。
//
//      那么每一组中，我们必定贪心的选择其中最小的数，这样才能移除这个分组内的全部数。
//      因为 c 是非负数，所以移除的数不会小于选择的数。
//
//      那么在所有分组中，我们必定选择数字最多的那个分组中最小的数。
//      如果存在多个分组，则选择所有分组的最小的数的最小值。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 nums 中全部 O(n) 个数
//      空间复杂度：O(min(n, space))
//          1. 需要维护 nums 中全部 O(n) 个数模 space 的余数的出现次数，
//              最多有 O(min(n, space)) 不同的余数


func destroyTargets(nums []int, space int) int {
    // remainToInfo[remain] = (min_num, cnt)
    //  remain: nums[i] 模 space 的余数
    //  min_num: nums 中模 space 余 remain 的最小的数
    //  cnt: nums 中模 space 余 remain 的数的个数
    remainToInfo := make(map[int]*Info)
    // ans 维护满足题意的结果
    ans := 0
    // total 维护选择 ans 后可移除的数的个数
    total := 0
    for _, num := range nums {
        // remain = num % space 的出现次数 +1 ，并更新对应的最小的数
        info := remainToInfo[num % space]
        if info == nil {
            info = &Info{ math.MaxInt32, 0 }
            remainToInfo[num % space] = info
        }
        info.cnt += 1
        info.minNum = min(info.minNum, num)
        // 如果 remain 的出现次数更多，或者出现次数相等但数字更小，
        // 则贪心地选择 min_num
        if total < info.cnt || (total == info.cnt && ans > info.minNum) {
            ans = info.minNum
            total = info.cnt
        }
    }

    return ans
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

type Info struct {
    minNum, cnt int
}
