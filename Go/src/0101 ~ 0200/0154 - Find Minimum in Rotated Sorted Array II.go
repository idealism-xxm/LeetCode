// 链接：https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
// 题意：给定一个没有重复数字的整数升序数组，并将其中后边一块（数量任意）放在最前面，
//		（例如： [1,2,3,4,5] -> [3,4,5,1,2]）求最小值 ？

// 输入： [3,4,5,1,2]
// 输出： 1

// 输入： [4,5,6,7,0,1,2]
// 输出： 0

// 思路： 二分
//
//		0081 的简化版， 0153 的强化版
//		二分 [l, r] 内的数，每次找到中点 mid = (l + r) >> 1
//      1. nums[l] > nums[mid] ：则最小值在 左边区间，r = mid
//      2. nums[mid] > nums[r] ：则最小值在 右边区间，l = mid + 1
//      3. 上面两个都不满足，则有：nums[l] <= nums[mid] <= nums[r] ，
//     		此时又分为两种种情况：
//			(1) nums[l] == nums[mid] == nums[r] ：
//				则无法判断最小值在哪个区间，去除最后一个数字即可， r--
//			(2) nums[l] < nums[mid] || nums[mid] < nums[r] ：
//				则最小值一定是 nums[l] ， r = l
//
//  	时间复杂度：平均 O(logn) / 最差 O(n)
//		空间复杂度： O(1)

func findMin(nums []int) int {
	l, r := 0, len(nums) - 1
	// 二分找到最小值（前闭后闭）
	for ; l < r; {
		mid := (l + r) >> 1
		// 如果 nums[l] > nums[mid] ，则最小值在左边区间
		if nums[l] > nums[mid] {
			r = mid
		} else if nums[mid] > nums[r] {
			// 如果 nums[mid] > nums[r] ，则最小值在右边区间
			l = mid + 1
		} else if nums[l] == nums[mid] && nums[mid] == nums[r] {
			// 此时有 nums[l] == nums[mid] == nums[r] ，无法判断最小值在哪边
			// 由于本题不需要找特定数，所以直接直接将整个区间右边减少一个数即可
			r--
		} else {
			// 此时有 nums[l] < nums[mid] || nums[mid] < nums[r]
			// 则最小值是 nums[l]
			r = l
		}
	}
	return nums[l]
}
