// 链接：https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
// 题意：给定一个没有重复数字的整数升序数组，并将其中后边一块（数量任意）放在最前面，
//		（例如： [1,2,3,4,5] -> [3,4,5,1,2]）求最小值 ？

// 输入： [3,4,5,1,2]
// 输出： 1

// 输入： [4,5,6,7,0,1,2]
// 输出： 0

// 思路： 二分
//
//		0033 的简化版
//		二分 [l, r] 内的数，每次找到中点 mid = (l + r) >> 1
//      1. nums[l] > nums[mid] ：则最小值在 左边区间，r = mid
//      2. nums[mid] > nums[r] ：则最小值在 右边区间，l = mid + 1
//      3. 上面两个都不满足，则有：nums[l] <= nums[mid] <= nums[r] ，
//     		即数组一定是递增的，l 就是最小值， r = l
//
//		先判断 2 就可以把 1 和 3 合并在一起处理
//
//		时间复杂度： O(logn)
//		空间复杂度： O(1)

func findMin(nums []int) int {
	l, r := 0, len(nums) - 1
	// 二分找到最小值（前闭后闭）
	for ; l < r; {
		mid := (l + r) >> 1
		// 如果 nums[mid] > nums[r] ，则最小值在右边区间
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			// 否则，最小值在左边区间
			r = mid
		}
	}
	return nums[l]
}
