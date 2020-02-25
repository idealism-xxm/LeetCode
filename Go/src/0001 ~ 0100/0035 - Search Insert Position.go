// 链接：https://leetcode.com/problems/search-insert-position/
// 题意：给定一个升序排序后无重复元素的数组，给定目标数，若目标数存在，则返回其下标；否则，返回其插入下标

// 输入：[1,3,5,6], 5
// 输出：2

// 输入：[1,3,5,6], 0
// 输出：0

// 思路：二分
//  很简单的一题，题目意思就是求大于等于目标数的第一个数的下标，即 C++ 中的 lower_bound
//  时间复杂度：O(logn)，空间复杂度：O(1)

func searchInsert(nums []int, target int) int {
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