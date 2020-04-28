// 链接：https://leetcode.com/problems/largest-number/
// 题意：给定一个非负整数数组，求将其拼接起来后能得到对最大整数对应对字符串？

// 输入： [10,2]
// 输出： "210"

// 输入： [3,30,34,5,9]
// 输出： "9534330"

// 思路： 排序
//
//		将每个数字转换成字符串，然后对比即可
//		对于任意两个要排序对字符串 a, b 来说，
//		如果 a + b < b + a ，则 a + b 更小，
//		那么在拼接后的字符串里面必定是 b 能贡献更大的结果，
//		所以要将 b 放在前面
//
//		【注意】如果全都是 0 ，那么需要直接返回 "0"
//
//		时间复杂度： O(nlogn)
//		空间复杂度： O(n)

import (
	"sort"
	"strconv"
)

type StringSlice []string

func largestNumber(nums []int) string {
	// 1. 转换成 []string 以便于进行排序
	stringSlice := make([]string, len(nums))
	for i, num := range nums {
		stringSlice[i] = strconv.Itoa(num)
	}

	// 2. 按照自定义顺序排序即可
	sort.Sort(StringSlice(stringSlice))
	if stringSlice[0] == "0" {
		return "0"
	}

	// 3. 从后往前拼接即可
	result := ""
	for _, str := range stringSlice {
		result += str
	}
	return result
}

func (r StringSlice) Len() int {
	return len(r)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (r StringSlice) Less(i, j int) bool {
	a, b := r[i], r[j]
	return a + b > b + a
}

// Swap swaps the elements with indexes i and j.
func (r StringSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
