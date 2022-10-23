// 链接：https://leetcode.com/problems/contains-duplicate-iii/
// 题意：给定一个整数数组 nums 和两个正整数 indexDiff, valueDiff ，
//      判断是否存满足以下条件的二元组 (i, j) ？
//          1. i != j
//          2. abs(i - j) <= indexDiff
//          3. abs(nums[i] - nums[j]) <= valueDiff


//  1 <= nums.length <= 10 ^ 5
//  -(10 ^ 9) <= nums[i] <= 10 ^ 9
//  1 <= indexDiff <= nums.length
//  0 <= valueDiff <= 10 ^ 9


// 输入： nums = [1,2,3,1], indexDiff = 3, valueDiff = 0
// 输出： true
// 解释： 选取二元组 (0, 3) ，则有：
//          1. 0 != 3
//          2. abs(0 - 3) <= 3
//          3. abs(1 - 1) <= 0

// 输入： nums = [1,5,9,1,5,9], indexDiff = 2, valueDiff = 3
// 输出： false
// 解释： 所有的二元组都无法同时满足题目的三个条件


// 思路1： 滑动窗口 + TreeSet/TreeMap
//
//      本题是 LeetCode 219 的加强版，在下标之差不超过 indexDiff 的基础上，
//      允许值之差不超过 valueDiff 。
//
//		我们可以用一个数据维护滑动窗口 [i - indexDiff, i) 内的所有数，
//      保证满足下标之差不超过 indexDiff 这个条件。
//
//      在遍历数组 nums 的每个数 num 时，
//      先获取大于等于 num - valueDiff 的第一个数 target 。
//
//      若 target 存在，且 target <= num + valueDiff ，则满足条件 3 ，
//      即满了所有 3 个条件，直接返回 true 。
//
//      否则将 num 放入滑动窗口中，并从滑动窗口中移除 nums[i - indexDiff] 。
//
//      可以发现我们遍历全部 O(n) 个数时，每次都要执行以下操作一次：
//          1. 找到滑动窗口内大于某个数的第一个数
//          2. 将某个数放入滑动窗口
//          3. 从滑动窗口中移除某个数
//
//      如果这 3 个操作中有一个时间复杂度为 O(n) 都会 TLE ，
//      所以我们需要一个最差能在 O(logn) 内完成以上操作的数据结构。
//
//      TeeeSet/TreeMap 恰好支持以上操作，并都能在 O(logn) 内完成，
//      所以我们使用其维护滑动窗口内的数字即可。
//      
//
//      时间复杂度： O(nlog(indexDiff))
//          1. 需要遍历 nums 中全部 O(n) 个数字，
//              每次都需要执行时间复杂度为 O(log(indexDiff)) 的操作
//      空间复杂度： O(indexDiff)
//          1. 需要维护滑动窗口内全部 O(indexDiff) 个数字


import "github.com/emirpasic/gods/maps/treemap"


func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
    // numToIndex 维护滑动窗口 [i - indexDiff, i) 内的所有数
    numToIndex := treemap.NewWithIntComparator()
    for i, num := range nums {
        // 如果滑动窗口内存在一个数 target 与 num 的差不超过 valueDiff ，
        // 即 num - valueDiff <= target <= num + valueDiff ，则满足题意
        target, _ := numToIndex.Ceiling(num - valueDiff)
        if target, ok := target.(int); ok && target <= num + valueDiff {
            return true
        }

        // 将当前数 num 纳入滑动窗口中
        numToIndex.Put(num, i)
        // 将左边界的数移除滑动窗口
        if i >= indexDiff {
            numToIndex.Remove(nums[i - indexDiff])
        }
    }

    // 在循环中没有返回，则所有数都不满足题意
    return false
}
