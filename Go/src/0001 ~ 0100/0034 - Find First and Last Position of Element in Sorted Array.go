// 链接：https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// 题意：给定一个升序排序后的数组，在 O(logn) 内找到指定数的起始下标和终止下标

// 输入：nums = [5,7,7,8,8,10], target = 8
// 输出：[3,4]

// 输入：nums = [5,7,7,8,8,10], target = 6
// 输出：[-1,-1]

// 思路：两次二分
//  很简单的一题，就是实现 C++ 中的 lower_bound 和 upper_bound 即可
//  两个函数的差别只有比较
//  lower_bound ：nums[mid] <  target，则目标值在 右边区间
//  upper_bound ：nums[mid] <= target，则目标值在 右边区间
//  当然也可以只用 lower_bound ，第二次 用 target + 1
//  时间复杂度：O(logn)，空间复杂度：O(1)

func searchRange(nums []int, target int) []int {
    length := len(nums)
    if length == 0 {
        return []int {-1, -1}
    }

    start := lowerBound(nums, target) // 搜索第一个 大于等于 target 的数
    if start == length || nums[start] != target { // 如果 不存在 大于等于 target 的数 或者 大于等于 target 的数不是 targt，则没找到
        return []int {-1, -1}
    }
    end := lowerBound(nums, target + 1) // 搜索第一个 大于 target 的数
    return []int {start, end - 1}
}

// 二分搜索，查找第一个 大于等于 target 的下标（前闭后开）
func lowerBound(nums []int, target int) int {
    l, r := 0, len(nums)
    for l < r {
        mid := (l + r) >> 1
        if nums[mid] < target { // 目标值在右边区间
            l = mid + 1
        } else { // 目标值在左边区间
            r = mid
        }
    }
    return l
}

// 二分搜索，查找第一个 大于 target 的下标（前闭后开）
func upperBound(nums []int, target int) int {
    l, r := 0, len(nums)
    for l < r {
        mid := (l + r) >> 1
        if nums[mid] <= target { // 目标值在右边区间
            l = mid + 1
        } else { // 目标值在左边区间
            r = mid
        }
    }
    return l
}