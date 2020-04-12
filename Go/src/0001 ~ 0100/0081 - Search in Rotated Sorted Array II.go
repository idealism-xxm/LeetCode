// 链接：https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
// 题意：给定一个升序的有重复数字的整型数组，将后面一的部分（不清楚有多少数）放到前面，
//		判断指定的数是否在数组内？

// 输入：nums = [2,5,6,0,0,1,2], target = 0
// 输出：true

// 输入：nums = [2,5,6,0,0,1,2], target = 3
// 输出：false

// 思路：一次二分
//	大部分情况下和 0033 题是一样的，由于多了重复数字，
//	所以最差情况下需要 O(n) 的时间复杂度，因为无法判断
//  在每次确定起点（最小值）的区间后后，判断 目标数的相应区间
//  1. nums[l] > nums[mid] -> 起点（最小值）在左边区间（不包括起点为区间开始点）：
//      (1) 目标值在右边区间数的范围内：l = mid + 1 （因为右边区间的数是连续的，所以很好比较）
//      (2) 否则认为目标数在左边区间：  r = mid - 1
//  2. nums[mid] > nums[r] -> 起点（最小值）在右边区间：
//      (1) 目标值在左边区间数的范围内：r = mid - 1 （因为左边区间的数是连续的，所以很好比较）
//      (2) 否则认为目标数在左边区间：  l = mid + 1
//	3. 上面两个都不满足，则有：nums[l] <= nums[mid] <= nums[r] ，
//		此时又分为两种种情况：
//		(1) nums[l] == nums[mid] == nums[r] ：
//			则无法判断起点（最小值）在哪个区间，
//			需要遍历左边区间
//			(a) 若存在，则直接返回 true
//			(b) 若不存在，且左边区间的数不全是 nums[l] ：
//				则起点在左边区间，右边区间的数都是 nums[l] ，直接返回 false
//			(c) 若不存在，且左边区间的数都是 nums[l] ：
//				则可认为起点在右边区间，对右边区间继续二分
//		(2) nums[l] < nums[mid] || nums[mid] < nums[r] ：
//			直接判断是否在两边的区间即可
//			(a) 目标值在左边区间数的范围内： r = mid - 1
//			(b) 目标值在右边区间数的范围内： l = mid + 1
//			(c) 目标值均不在两边区间范围内，直接返回 false
//
//  时间复杂度：平均 O(logn) / 最差 O(n) ，空间复杂度：O(1)

func search(nums []int, target int) bool {
	l, r := 0, len(nums) - 1
	for ; l <= r; {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return true
		}
		if nums[l] > nums[mid] {  // 起点（最小值）在左边区间
			if nums[mid] < target && target <= nums[r] {  // 目标数在右边区间
				l = mid + 1
			} else {  // 目标数在左边区间
				r = mid - 1
			}
		} else if nums[mid] > nums[r] {  // 起点（最小值）在右边区间
			if nums[l] <= target && target < nums[mid] {  // 目标数在左边区间
				r = mid - 1
			} else {  // 目标数在右边区间
				l = mid + 1
			}
		} else {  // nums[l] <= nums[mid] <= nums[r]
			if nums[l] == nums[mid] && nums[mid] == nums[r] {
				// 则无法判断起点（最小值）在哪个区间，
				// 遍历左边区间，并记录左边区间的数字是否都一样
				isAllSame := true
				for i := l + 1; i < mid; i++ {
					if nums[i] == target {
						return true
					}
					if nums[i] != nums[l] {
						isAllSame = false
					}
				}
				// 左边区间不存在目标数字
				if isAllSame {  // 左边区间数字都一样，认为起点在右边区间，继续进行二分
					l = mid + 1
				} else {  // 左边区间存在数字不一样，则起点在左边区间，右边区间数字都一样，直接返回 false
					return false
				}
			} else { // nums[l] < nums[mid] || nums[mid] < nums[r]
				if nums[l] <= target && target < nums[mid] {  // 目标数在左边区间
					r = mid - 1
				} else if nums[mid] < target && target <= nums[r] {   // 目标数在右边区间
					l = mid + 1
				} else { // 目标数均不在左右区间，直接返回false
					return false
				}
			}
		}
	}
	return false
}
