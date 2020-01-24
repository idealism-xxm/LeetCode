// 链接：https://leetcode.com/problems/permutations-ii/
// 题意：给定一个可能含重复数字的数组，求不重复的所有排列？

// 输入：[1,1,2]
// 输出：
// [
//   [1,1,2],
//   [1,2,1],
//   [2,1,1]
// ]

// 思路：递归模拟即可
//		每次枚举未使用的数字（若当前数字已经枚举过，则跳过），
//		放入排列对应的位置，当所有位置都放入后，即找到一个不重复且合法的排列

import "sort"

func permuteUnique(nums []int) [][]int {
	used := make([]bool, len(nums))  // 表示当前数字是否正在使用
	list := make([]int, len(nums))  // 表示当前结果排列
	sort.Ints(nums)  // 按升序排序，方便跳过重复的数字
	return dfs(nums, used, 0, list)
}

func dfs(nums []int, used []bool, current int, list []int) [][]int {
    if current == len(nums) {  // 若所有位置都放入数字，则当前排列放入结果列表中
    	return [][]int {append(list[:0:0], list...)}
	}

	isFirst := true  // 是否为本次递归第一个选择的数
	preNum := nums[0] // 前一个选择的数字
	var result [][]int  // 收集循环中产生的所有结果
	for i := 0; i < len(nums); i++ {
		if !used[i] {  // 如果当前数字未使用
			if isFirst || nums[i] != preNum {  // 当前数字本次递归第一次选择的数 或者 当前数字不等于前一个选择的数
				isFirst = false
				preNum = nums[i]

				used[i] = true  // 使用当前数字
				list[current] = nums[i]  // current 位放入 nums[i]
				result = append(result, dfs(nums, used, current + 1, list)...)  // 收集所有结果
				used[i] = false  // 不再使用当前数字
			}
		}
	}
	return result
}