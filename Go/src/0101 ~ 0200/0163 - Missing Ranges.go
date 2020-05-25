// 链接：https://leetcode.com/problems/missing-ranges/
// 题意：给定一个升序整数数组和一个闭区间 [lower, upper]，
//		求闭区间内不在整数数组的区间列表 ？

// 输入： nums = [0, 1, 3, 50, 75], lower = 0 和 upper = 99
// 输出： ["2", "4->49", "51->74", "76->99"]

// 思路： 模拟
//
//		遍历 nums 数组，若当前数或者前一个数在范围内，则可能需要收集对应的范围，
//		初始 l, r := lower, upper ，表示范围的最小值和最大值，
//		然后通过 num 缩小范围
//			1. 如果 num 不是第一个数，则 l 可以扩大为 nums[i - 1] + 1
//			2. 如果 nums[i - 1] 在范围内，则 r 可以缩小为 num - 1
//		经过以上处理后，如果 l <= r ，则对应的范围可以收集
//
//
//		时间复杂度： O(n)
//		空间复杂度： O(n) （额外空间只需要 O(1)）

import (
	"fmt"
	"strconv"
)

func findMissingRanges(nums []int, lower int, upper int) []string {
	// 将 upper + 1 加入到数组中，方便后续处理
	nums = append(nums, upper+1)
	var result []string
	for i, num := range nums {
		// 如果当前数在范围左边，或者上一个数在范围右边，
		// 则直接处理下一个
		if num < lower || (i > 0 && nums[i-1] >= upper) {
			continue
		}

		// 范围必须在 [lower, upper] 内，然后通过 num 的值缩小范围
		l, r := lower, upper
		// 如果当前数是不是第一个数，且 nums[i - 1] < num - 1 ，
		// 则 l 可以扩大为 nums[i - 1] + 1
		if i > 0 {
			l = nums[i-1] + 1
		}
		// 如果当前数在 [lower, upper] 内，
		// 则 r 可以缩小为 num - 1
		if num <= upper {
			r = num - 1
		}

		// 范围合法，则加入到结果中
		if l <= r {
			result = append(result, getRangeString(l, r))
		}
	}
	return result
}

// 返回整数范围 [l, r] 对应的字符串
func getRangeString(l, r int) string {
	// 如果只有一个数，则直接返回当前数即可
	if l == r {
		return strconv.Itoa(l)
	}
	// 如果有多个数，则需要返回 "{l}->{r}"
	return fmt.Sprintf("%v->%v", l, r)
}
